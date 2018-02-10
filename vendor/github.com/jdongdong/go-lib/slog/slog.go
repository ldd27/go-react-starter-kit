package slog

import (
	"runtime"

	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cihub/seelog"
)

var workingDir = "/"

func init() {
	wd, err := os.Getwd()
	if err == nil {
		workingDir = filepath.ToSlash(wd) + "/"
	}

	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		seelog.Critical("log config not exist")
		return
	}
	//logger.SetAdditionalStackDepth(1)
	seelog.ReplaceLogger(logger)
}

func Trace(v ...interface{}) {
	Tracef(0, v...)
}

func Debug(v ...interface{}) {
	Debugf(0, v...)
}

func Info(v ...interface{}) {
	Infof(0, v...)
}

func Warn(v ...interface{}) {
	Warnf(0, v...)
}

func Error(v ...interface{}) {
	Errorf(0, v...)
}

func Critical(v ...interface{}) {
	Criticalf(0, v...)
}

func Tracef(addCallDepth int, v ...interface{}) {
	if addCallDepth < 0 {
		addCallDepth = 0
	}
	_, _, funcName, line, _ := extractCallerInfo(3 + addCallDepth)
	seelog.Trace(funcName, "-", line, v)
}

func Debugf(addCallDepth int, v ...interface{}) {
	if addCallDepth < 0 {
		addCallDepth = 0
	}
	_, _, funcName, line, _ := extractCallerInfo(3 + addCallDepth)
	seelog.Debug(funcName, "-", line, v)
}

func Infof(addCallDepth int, v ...interface{}) {
	if addCallDepth < 0 {
		addCallDepth = 0
	}
	_, _, funcName, line, _ := extractCallerInfo(3 + addCallDepth)
	seelog.Info(funcName, "-", line, v)
}

func Warnf(addCallDepth int, v ...interface{}) {
	if addCallDepth < 0 {
		addCallDepth = 0
	}
	_, _, funcName, line, _ := extractCallerInfo(3 + addCallDepth)
	seelog.Warn(funcName, "-", line, v)
}

func Errorf(addCallDepth int, v ...interface{}) {
	if len(v) == 0 {
		return
	}

	if v[0] == nil {
		return
	}
	if addCallDepth < 0 {
		addCallDepth = 0
	}
	_, _, funcName, line, _ := extractCallerInfo(3 + addCallDepth)
	seelog.Error(funcName, "-", line, v)
}

func Criticalf(addCallDepth int, v ...interface{}) {
	if len(v) == 0 {
		return
	}
	if v[0] == nil {
		return
	}
	if addCallDepth < 0 {
		addCallDepth = 0
	}
	_, _, funcName, line, _ := extractCallerInfo(3 + addCallDepth)
	seelog.Critical(funcName, "-", line, v)
}

func extractCallerInfo(skip int) (fullPath string, shortPath string, funcName string, line int, err error) {
	pc, fp, ln, ok := runtime.Caller(skip)
	if !ok {
		err = fmt.Errorf("error during runtime.Caller")
		return
	}
	line = ln
	fullPath = fp
	if strings.HasPrefix(fp, workingDir) {
		shortPath = fp[len(workingDir):]
	} else {
		shortPath = fp
	}
	funcName = runtime.FuncForPC(pc).Name()
	if strings.HasPrefix(funcName, workingDir) {
		funcName = funcName[len(workingDir):]
	}
	return
}
