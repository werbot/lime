package version

import (
	"fmt"
	"runtime"
)

// Version is a ...
const Version = "0.1.0"

var (
	// GitCommit is a ...
	GitCommit string

	// BuildDate is a ...
	BuildDate = ""

	// GoVersion is a ...
	GoVersion = runtime.Version()

	// OsArch is a ...
	OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
)
