//go:build darwin
// +build darwin

package thread

import "golang.org/x/sys/unix"

func GetTid() uint64 {
	var tid uint64
	unix.PthreadThreadid_np(0, &tid)
	return tid
}
