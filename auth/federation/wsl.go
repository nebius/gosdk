package federation

import "github.com/nebius/gosdk/platform"

// IsWSL reports whether the current process runs under Windows Subsystem for Linux.
//
// Deprecated: use platform.IsWSL.
func IsWSL() bool {
	return platform.IsWSL()
}
