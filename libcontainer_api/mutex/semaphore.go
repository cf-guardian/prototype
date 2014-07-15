package mutex

// #include <fcntl.h>
// #include <semaphore.h>
// #include <stdlib.h>
// #include <sys/stat.h>
// sem_t* _sem_open(const char * name, int oflag, mode_t mode, unsigned int value) {
//     return sem_open(name, oflag, mode, value);
// }
// sem_t* SEM_FAILED_ = SEM_FAILED;
import "C"

import (
	"unsafe"
)

type Semaphore interface {

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

// Creates a semaphore with the given name and initialises it to the value 1. This means the
// semaphore can be acquired by at most one user concurrently.
//
// If the named semaphore already exists, returns it without modification.
func SemOpen(semName string) (Semaphore, error) {
	n := C.CString(semName)
	defer C.free(unsafe.Pointer(n))

	sem_t, err := C._sem_open(n, C.O_CREAT, C.S_IRWXU, 1)
	if sem_t == C.SEM_FAILED_ {
		return nil, err
	} else {
		return &semaphore{semName, wrap(sem_t)}, nil
	}
}

type semaphore struct {
	name string
	sem_t _sem_t
}

func (sem *semaphore) Wait() error {
	panic("unimplemented")
}

func (sem *semaphore) Post() error {
	panic("unimplemented")
}

func (sem *semaphore) TryWait() error {
	panic("unimplemented")
}

func (sem *semaphore) Close() error {
	panic("unimplemented")
}

func (sem *semaphore) Destroy() error {
	n := C.CString(sem.name)
	defer C.free(unsafe.Pointer(n))

	rc, err := C.sem_unlink(n)
	if rc == 0 {
		return nil
	} else {
		return err
	}
}

type _sem_t unsafe.Pointer

func wrap(s *_Ctype_sem_t) _sem_t {
	return _sem_t(unsafe.Pointer(s))
}

func unwrap(s _sem_t) *_Ctype_sem_t {
	return (*_Ctype_sem_t)(unsafe.Pointer(s))
}
