package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func SetupSigusr1Trap() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			dumpStacks()
		}
	}()
	return c
}

func dumpStacks() {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	numBytes := runtime.Stack(buf.Bytes(), true)
	if numBytes > buf.Cap() {
		fmt.Println("WARNING: Stack dump truncated. Consider increasing buffer size.")
	}
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n", buf.String())
}
