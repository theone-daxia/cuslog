package main

import "github.com/theone-daxia/cuslog"

func main() {
	cuslog.Info("std log")
	cuslog.SetOptions(cuslog.WithLevel(cuslog.DebugLevel))
	cuslog.Debug("change std log to debug level")
	cuslog.SetOptions(cuslog.WithFormatter(&cuslog.JsonFormatter{IgnoreBasicFields: false}))
	cuslog.Debug("log in json format")
	cuslog.Info("another log in json format")
}