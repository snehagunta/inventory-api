package model

// AcquireLockResult contains the outcome of a distributed lock acquisition.
//domain:constructors NewAcquireLockResult
type AcquireLockResult struct {
	lockToken LockToken
}

func NewAcquireLockResult(lockToken LockToken) AcquireLockResult {
	return AcquireLockResult{lockToken: lockToken}
}

func (r AcquireLockResult) LockToken() LockToken { return r.lockToken }
