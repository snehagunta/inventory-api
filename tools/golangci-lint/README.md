# Inventory API Go lint tooling

This module builds the repository's `golangci-lint` executable with the
domain-modeling analyzers from [`go-uber-linters`](../go-uber-linters). The
dependency is pinned in `go.mod` and locally replaced with the sibling module,
so the analyzer library can be developed and reviewed alongside its consumer.

Run the configured linters from the repository root with `make lint`. The
`domainconstruct`, `domainzero`, and `domainserialization` analyzers are
available through `make lint-domain`; they apply to domain types annotated
with `domain:*` directives in their Go doc comments. The domain checks remain
opt-in while existing model findings are assessed and migrated.
