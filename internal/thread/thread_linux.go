//go:build linux
// +build linux

package thread

import "golang.org/x/sys/unix"

func GetTid() uint64 {
	return uint64(unix.Gettid())
}
