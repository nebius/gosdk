package platform

import (
	"os"
	"runtime"
	"strings"
)

func IsWSL() bool {
	if runtime.GOOS != "linux" {
		return false
	}
	read, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}

	// WSL1 marker is "Microsoft"; WSL2 marker is "microsoft".
	return strings.Contains(string(read), "Microsoft") || strings.Contains(string(read), "microsoft")
}
