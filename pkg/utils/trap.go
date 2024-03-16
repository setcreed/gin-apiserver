/*
Copyright 2023 The gin-apiserver Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
