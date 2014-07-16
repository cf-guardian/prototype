package mutex

type Mutex interface {

	// Acquires this semaphore, blocking first if necessary.
	Wait() error

	// Releases this semaphore.
	// Errors: ErrOverflow
	Post() error

	// Attempts to acquire this semaphore, but does not block.
	TryWait() error

	// Closes this semaphore. It will continue to exist and may be opened again.
	// If the semaphore is already closed, do nothing.
	Close() error

	// Deletes the semaphore.
	Destroy() error

}
