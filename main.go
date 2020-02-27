package main

import (
	"github.com/json-iterator/go"
	"github.com/oneminuter/commonly/ruslog"
)

func main() {
	ruslog.SetFormatter(&ruslog.JSONFormatter{
		TimestampFormat:  "",
		DisableTimestamp: false,
		DataKey:          "",
		FieldMap:         nil,
		CallerPrettyfier: nil,
		PrettyPrint:      true,
	})

	ruslog.SetReportCaller(true)

	var u = struct {
		Username string
		Sex      int
	}{
		Username: "linty",
		Sex:      1,
	}

	str, _ := jsoniter.MarshalToString(u)

	ruslog.Println(str)
}
