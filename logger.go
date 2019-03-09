package libs

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	// Mnt - is custom logger for monitor events
	Mnt = log.New(logWriter{}, "{monitor}", 0)
	// Log - is custom info logger
	Log = log.New(logWriter{}, "{log}", 0)
	// Err is custom error logger
	Err = log.New(logWriter{}, "{ERROR}", 0)
	// Critical - is custom logger for critical errors
	Critical = log.New(logWriter{}, "{CRITICAL}", 0)
	// Stack is custom error logger for show path error
	Stack = log.New(logWriter{}, "-->", 0)
)

type logWriter struct{}

func (f logWriter) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	str := string(p[:])
	var pStr []string
	log.SetFlags(0)

	if strings.Contains(str, "monitor") {
		pStr = strings.SplitAfter(str, "{monitor}")
		log.Printf("\033[2m%s ==> %s", pStr[0], pStr[1])
	}

	if strings.Contains(str, "log") {
		pStr = strings.SplitAfter(str, "{log}")
		log.Printf("\033[0;36m%s ------------------------------------------ \n\033[0;33m%s | %s:%d \033[0m==> %s", pStr[0], filepath.Base(filepath.Dir(file))+"/"+filepath.Base(file), fnName, line, pStr[1])
	}

	if strings.Contains(str, "ERROR") {
		pStr = strings.SplitAfter(str, "{ERROR}")
		log.Printf("\033[0;95m%s ---------------------------------------- \n\033[0;33m%s --> %s:%d \033[0m==> %s", pStr[0], filepath.Base(filepath.Dir(file))+"/"+filepath.Base(file), fnName, line, pStr[1])
	}

	if strings.Contains(str, "-->") {
		pStr = strings.SplitAfter(str, "-->")
		log.Printf("\033[0;31m%s \033[0;33m%s --> %s:%d", pStr[0], filepath.Base(filepath.Dir(file))+"/"+filepath.Base(file), fnName, line)
	}

	if strings.Contains(str, "CRITICAL") {
		pStr = strings.SplitAfter(str, "{CRITICAL}")
		log.Printf("\033[0;31m%s ------------------------------------- \n\033[0;33m%s | %s:%d \033[0;31m===> %s", pStr[0], filepath.Base(filepath.Dir(file))+"/"+filepath.Base(file), fnName, line, pStr[1])
	}

	return len(p), nil
}
