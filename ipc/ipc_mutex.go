package ipc

// An inter-process mutex.
type IPCMutex interface {

	// Acquires this semaphore, blocking first if necessary.
	Lock() error

	// Releases this semaphore.
	Unlock() error

	// Attempts to acquire this semaphore, but does not block.
	TryLock() error

	// Closes this semaphore. It will continue to exist and may be opened again.
	// If the semaphore is already closed, do nothing.
	Close() error

	// Deletes the semaphore.
	Destroy() error

}
