package ipc

// An inter-process mutex.
type Mutex interface {

	// Locks this mutex, blocking first if necessary.
	// Note: blocks if the current goroutine has already locked the mutex.
	Lock() error

	// Unlocks this mutex.
	// Unlocking a mutex which is not locked is undefined.
	Unlock() error

	// Attempts to lock this mutex, but does not block.
	TryLock() error

	// Closes this mutex, freeing up resources used by the current process.
	Close() error

	// Deletes this mutex.
	Destroy() error
}
