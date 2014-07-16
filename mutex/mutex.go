// rename package to make IPC nature clear
package mutex

// rename interface to make IPC nature clear
type Mutex interface {

	// Acquires this semaphore, blocking first if necessary.
	Lock() error

	// Releases this semaphore.
	// Errors: ErrOverflow
	Unlock() error

	// Attempts to acquire this semaphore, but does not block.
	TryLock() error

	// Closes this semaphore. It will continue to exist and may be opened again.
	// If the semaphore is already closed, do nothing.
	Close() error

	// Deletes the semaphore.
	Destroy() error

}
