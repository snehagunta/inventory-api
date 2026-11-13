// Command golangci-lint runs upstream linters with Inventory's domain analyzers.
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/golangci/golangci-lint/v2/pkg/commands"
	"github.com/golangci/golangci-lint/v2/pkg/exitcodes"

	_ "github.com/project-kessel/inventory-api/tools/golangci-lint/domainconstruct/plugin"
)

func main() {
	goVersion := runtime.Version()
	if err := os.Setenv("GOTOOLCHAIN", goVersion); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitcodes.Failure)
	}
	version := ""
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dependency := range info.Deps {
			if dependency.Path == "github.com/golangci/golangci-lint/v2" {
				version = dependency.Version
				break
			}
		}
	}
	identity, err := executableHash()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitcodes.Failure)
	}
	version += "-inventory." + identity
	if err := commands.Execute(commands.BuildInfo{Version: version, GoVersion: goVersion}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitcodes.Failure)
	}
}

func executableHash() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}
	binary, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() { _ = binary.Close() }()
	hash := sha256.New()
	if _, err := io.Copy(hash, binary); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
