// Package wlog implements a simple logging package for internal WASM workers use.
package wlog

import (
	"log"
)

var (
	// stdLog is standard wasm workers logger.
	stdLog = log.New(StdLog, "", log.LstdFlags|log.Lshortfile)

	// debugLog is a separate logger for debug messages.
	debugLog = log.New(StdDebug, "[DEBUG] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)

	debugLogEnabled = checkDebugLogParam()
)

const callerSkip = 2

func Debugf(format string, v ...any) { _ = "STUB: not implemented"; return }

func Debugln(v ...any) { _ = "STUB: not implemented"; return }

func Debug(v ...any) { _ = "STUB: not implemented"; return }

func Printf(format string, v ...any) { _ = "STUB: not implemented"; return }

func Println(v ...any) { _ = "STUB: not implemented"; return }

func Print(v ...any) { _ = "STUB: not implemented"; return }

func Fatal(v ...any) { _ = "STUB: not implemented"; return }

func Fatalln(v ...any) { _ = "STUB: not implemented"; return }

func Fatalf(format string, v ...any) { _ = "STUB: not implemented"; return }

func Panic(v ...any) { _ = "STUB: not implemented"; return }

func Panicln(v ...any) { _ = "STUB: not implemented"; return }

func Panicf(format string, v ...any) { _ = "STUB: not implemented"; return }

func checkDebugLogParam() bool { _ = "STUB: not implemented"; return false }
