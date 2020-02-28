package main

import (
	"os"

	"github.com/oneminuter/commonly/ruslog"
)

func main() {
	//ruslog.SetFormatter(&ruslog.JSONFormatter{
	//	TimestampFormat:  "",
	//	DisableTimestamp: false,
	//	DataKey:          "",
	//	FieldMap:         nil,
	//	CallerPrettyfier: nil,
	//	PrettyPrint:      true,
	//})
	//
	//ruslog.SetReportCaller(true)
	//
	//var u = struct {
	//	Username string
	//	Sex      int
	//}{
	//	Username: "linty",
	//	Sex:      1,
	//}
	//
	//str, _ := jsoniter.MarshalToString(u)
	//
	//ruslog.Println(str)

	ruslog.New()

	var log = ruslog.Logger{
		Out:   os.Stderr,
		Hooks: nil,
		Formatter: &ruslog.TextFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
		},
		ReportCaller: true,
		Level:        ruslog.DebugLevel,
		ExitFunc:     nil,
	}
	ruslog.NewEntry(&log).Println("dsfsd")
}
