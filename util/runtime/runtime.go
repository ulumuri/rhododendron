package runtime

import (
	"fmt"
	"runtime"
)

// Source: https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
func Trace() string {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s:%d %s\n", file, line, f.Name())
}
