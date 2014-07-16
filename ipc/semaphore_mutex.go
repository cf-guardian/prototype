package ipc

// #include <fcntl.h>
// #include <semaphore.h>
// #include <stdlib.h>
// #include <sys/stat.h>
// sem_t* _sem_open(const char * name, int oflag, mode_t mode, unsigned int value) {
//     return sem_open(name, oflag, mode, value);
// }
// sem_t* _SEM_FAILED = SEM_FAILED;
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

	if sem_t, err := C._sem_open(n, C.O_CREAT, C.S_IRWXU, 1); sem_t == C._SEM_FAILED {
		return nil, err
	} else {
		return newSemaphore(semName, sem_t), nil
	}
}

type semaphore struct {
	name string
	posix_sem _sem_t
}

func (sem *semaphore) Lock() error {
	if rc, err := C.sem_wait(sem.sem_t()); rc == 0 {
		return nil
	} else {
		return err
	}
}

func (sem *semaphore) Unlock() error {
	if rc, err := C.sem_post(sem.sem_t()); rc == 0 {
		return nil
	} else {
		return err
	}
}

func (sem *semaphore) TryLock() error {
	if rc, err := C.sem_trywait(sem.sem_t()); rc == 0 {
		return nil
	} else {
		return err
	}
}

func (sem *semaphore) Close() error {
	if rc, err := C.sem_close(sem.sem_t()); rc == 0 {
		return nil
	} else {
		return err
	}
}

func (sem *semaphore) Destroy() error {
	n := C.CString(sem.name)
	defer C.free(unsafe.Pointer(n))

	if rc, err := C.sem_unlink(n); rc == 0 {
		return nil
	} else {
		return err
	}
}

type _sem_t unsafe.Pointer

func newSemaphore(semName string, sem_t *_Ctype_sem_t) *semaphore {
	return &semaphore{semName, _sem_t(unsafe.Pointer(sem_t))}
}

func (sem *semaphore) sem_t() *_Ctype_sem_t {
	return (*_Ctype_sem_t)(unsafe.Pointer(sem.posix_sem))
}
