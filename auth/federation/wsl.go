package federation

import (
	"os"
	"runtime"
	"strings"
)

func IsWSL() bool {
	if runtime.GOOS == "linux" {
		read, err := os.ReadFile("/proc/version")
		if err != nil {
			return false
		}

		// WLS1: marker - Microsoft, WSL2: marker - microsoft
		return strings.Contains(string(read), "Microsoft") || strings.Contains(string(read), "microsoft")
	}

	return false
}
