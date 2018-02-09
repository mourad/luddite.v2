// +build windows

package luddite

import (
	"fmt"
	"os"
)

func dumpGoroutineStacks() {
	fmt.Fprintln(os.Stderr, "*** goroutine dump is not available on Windows ***")
}
