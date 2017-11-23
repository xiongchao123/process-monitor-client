package xlog

import (
	"os"
	"log"
	"time"
	"fmt"
	"path"
)

const MAX_SIZE = 10485760 //10M

var (
	Logfile   = "./xlog/monitor.log"
	ErrorFile = "./xlog/error.log"
)

func Info(logfile string, v ...interface{}) {
	if (logfile == "") {
		logfile = Logfile
	}
	xlog, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer xlog.Close()
	if err != nil {
		Fatal(Logfile, "Can't open file:", logfile)
	}
	Log := log.New(xlog, "[Info]", log.LstdFlags)
	//fmt.Println(v ...)
	Log.Println(v ...)
	go checkSize(logfile)
}

func Warn(logfile string, v ...interface{}) {
	if (logfile == "") {
		logfile = Logfile
	}
	xlog, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer xlog.Close()
	if err != nil {
		Fatal(Logfile, "Can't open file:", logfile)
	}
	Log := log.New(xlog, "[Warn]", log.LstdFlags)
	fmt.Println(v ...)
	Log.Println(v ...)
	go checkSize(logfile)
}

func Fatal(logfile string, v ...interface{}) {
	if (logfile == "") {
		logfile = Logfile
	}
	xlog, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer xlog.Close()
	if err != nil {
		Fatal(Logfile, "Can't open file:", logfile)
	}
	Log := log.New(xlog, "[Fatal]", log.LstdFlags)
	fmt.Println(v ...)
	Log.Println(v ...)
	go checkSize(logfile)
	os.Exit(1)
}

func checkSize(logfile string) {
	if (getSize(logfile) > MAX_SIZE) {
		dir, _ := path.Split(logfile)
		err := os.Rename(logfile, dir+time.Now().Format("20060102")+path.Base(logfile))
		if err != nil {
			Fatal(Logfile, err)
		}
	}
}

func getSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err != nil {
		Fatal(Logfile, err)
	}
	fileSize := fileInfo.Size() //获取size
	return fileSize
}
