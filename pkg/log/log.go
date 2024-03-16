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

package log

import (
	"os"
	"runtime"

	"github.com/setcreed/gin-apiserver/pkg/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// set log level
	logLevel := config.GetString(config.FLAG_KEY_LOG_LEVEL)
	if l, err := log.ParseLevel(logLevel); err != nil {
		panic(err)
	} else {
		log.SetLevel(l)
	}
}

const (
	KEY_MODULE_NAME   = "module_name"
	KEY_FUNCTION_NAME = "function_name"
	KEY_LINE_NUM      = "line_num"
)

func getLogMetadata() (string, string, int) {
	var (
		module, method string
		lineNum        int
	)
	if pc, file, line, ok := runtime.Caller(2); ok {
		module = file
		method = runtime.FuncForPC(pc).Name()
		lineNum = line
	}

	return module, method, lineNum
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Info(args)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Infoln(args)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Infof(format, args...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warning(args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Warning(args)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warningln(args ...interface{}) {
	module, mehtod, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: mehtod,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Warningln(args)
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warningf(format string, args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Warningf(format, args...)
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Error(args)
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Errorln(args)
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Errorf(format, args...)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Fatal(args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Fatal(args)
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Fatalln(args)
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	module, method, lineNum := getLogMetadata()
	fields := log.Fields{
		KEY_MODULE_NAME:   module,
		KEY_FUNCTION_NAME: method,
		KEY_LINE_NUM:      lineNum,
	}
	log.WithFields(fields).Fatalf(format, args...)
}
