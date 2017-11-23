package util

import (
	"os"
	"syscall"
	"fmt"
)

//文件锁
type FileLock struct {
	dir string
	f   *os.File
}

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func New(dir string) *FileLock {
	return &FileLock{
		dir: dir,
	}
}

//加锁 in linux 在windows下请注释这段
func (l *FileLock) Lock() error {
	f, err := os.Open(l.dir)
	if err != nil {
		return err
	}
	l.f = f
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}

//加锁in windows  ps:在linux下请注释这段
/*func (l *FileLock)Lock() error {
	f, err := os.Open(l.dir)
	if err != nil {
		return err
	}
	l.f = f
	h, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return err
	}
	defer syscall.FreeLibrary(h)

	addr, err := syscall.GetProcAddress(h, "LockFile")
	if err != nil {
		return err
	}
	for {
		r0, _, _ := syscall.Syscall6(addr, 5, f.Fd(), 0, 0, 0, 1, 0)
		if 0 != int(r0) {
			break
		}
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
		//time.Sleep(100 * time.Millisecond)
	}
	return nil
}*/


//释放锁  linux
/*func (l *FileLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN)
}*/
