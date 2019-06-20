package logs

import (
	"testing"
)

func TestPrintln(t *testing.T) {
	Print("this is log test")
	Printf("this is log test")
	Println("this is log test")
	Debugf("this is log test")
	Debug("this is log test")
	Infof("this is log test")
	Info("this is log test")
	Warnf("this is log test")
	Warn("this is log test")
	Errorf("this is log test")
	Error("this is log test")
	Fatal("this is log test")
	Fatalf("this is log test")
	Fatalln("this is log test")
	Panic("this is log test")
	Panicf("this is log test")
	Panicln("this is log test")
	Stack("this is log test")
}
