package ipc

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

// Creates a semaphore with the given name and initialises it to the value 1. This means the
// semaphore can be acquired by at most one user concurrently.
//
// If the named semaphore already exists, returns it without modification.
func GetSemaphoreMutex(semName string) (Mutex, error) {
	n := C.CString(semName)
	defer C.free(unsafe.Pointer(n))

	sem_t, err := C._sem_open(n, C.O_CREAT, C.S_IRWXU, 1)
	if sem_t == C.SEM_FAILED_ {
		return nil, err
	} else {
		return &semaphore{semName, wrapSem(sem_t)}, nil
	}
}

type semaphore struct {
	name string
	sem_t _sem_t
}

func (sem *semaphore) Lock() error {
	panic("unimplemented")
}

func (sem *semaphore) Unlock() error {
	panic("unimplemented")
}

func (sem *semaphore) TryLock() error {
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

func wrapSem(s *_Ctype_sem_t) _sem_t {
	return _sem_t(unsafe.Pointer(s))
}

func unwrapSem(s _sem_t) *_Ctype_sem_t {
	return (*_Ctype_sem_t)(unsafe.Pointer(s))
}
