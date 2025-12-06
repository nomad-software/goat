//go:build windows
// +build windows

package thread

import "golang.org/x/sys/windows"

func GetTid() uint64 {
	return uint64(windows.GetCurrentThreadId())
}
