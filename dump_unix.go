// +build !windows

package luddite

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func dumpGoroutineStacks() {
	sigs := make(chan os.Signal, 1)
	go func() {
		for {
			<-sigs
			buf := make([]byte, maxStackSize)
			size := runtime.Stack(buf, true)
			fmt.Fprintf(os.Stderr, "*** goroutine dump ***\n%s\n", buf[:size])
		}
	}()
	signal.Notify(sigs, syscall.SIGUSR1)
}
