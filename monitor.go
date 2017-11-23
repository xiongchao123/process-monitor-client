package main

import (
	"fmt"
	"flag"
	"./conf"
	"./comm"
	"./xlog"
	"./svr"
	"./util"
	"os"
)

func main() {
	switch {
	case conf.List:
		flag.Usage()
	case conf.Start != "":
		handle_procs(comm.START, conf.Start)
	case conf.Stop != "":
		handle_procs(comm.STOP, conf.Stop)
	case conf.Restart != "":
		handle_procs(comm.RESTART, conf.Restart)
	case conf.Status != "":
		handle_procs(comm.STATUS, conf.Status)
	case conf.Input == "start":
		//check is or not start
		if ok := checkLock(); !ok {
			fmt.Print("Process: monitor ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[is already start]", 0x1B)
			os.Exit(1)
		} else {
			if ok := svr.CheckProc(conf.CheckCommand); ok {
				fmt.Print("Process: monitor ")
				fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[is already start]", 0x1B)
				os.Exit(1)
			} else {
				//start monitor process and start check process for daemon
				svr.AllProcs(comm.START)
			}
		}
	case conf.Input == "stop":
		//check the monitor is or not stop
		if ok := svr.CheckProc(conf.CheckCommand); !ok {
			fmt.Print("Process: monitor ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[is already stop]", 0x1B)
			os.Exit(1)
		}
		//stop monitor process and stop check process for daemon
		svr.AllProcs(comm.STOP)
	case conf.Input == "restart":
		//check the monitor is or not stop
		if ok := svr.CheckProc(conf.CheckCommand); !ok {
			fmt.Print("Process: monitor ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[is already stop]", 0x1B)
		} else {
			//stop monitor process and stop check process
			svr.AllProcs(comm.STOP)
		}
		checkLock()
		//start monitor process and start check process for daemon
		svr.AllProcs(comm.START)
	case conf.Input == "status":
		//get  check process status
		svr.AllProcs(comm.STATUS)
	case conf.Input == "check":
		//check the process
		svr.CheckProcs()
	case conf.Input != "":
		fmt.Print("undefined command: ")
		fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, os.Args[1], 0x1B)
		os.Exit(1)
	default:
		//start self process
		if ok := checkLock(); !ok {
			fmt.Print("Process: monitor ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[is already start]", 0x1B)
			os.Exit(1)
		} else {
			if ok := svr.CheckProc(conf.CheckCommand); ok {
				fmt.Print("Process: monitor ")
				fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[is already start]", 0x1B)
				os.Exit(1)
			} else {
				//start conf process
				for service, _ := range conf.Conf {
					svr.Procs(comm.START, service)
				}
				//start monitor process
				svr.CheckProcs()
			}
		}
	}
}

//根据输入命令对监控进程执行相关操作
func handle_procs(cmd string, service string) {
	//检查检测的服务是否违法
	if _, ok := conf.Conf[service]; ok {
		svr.Procs(cmd, service)
		// ...
	} else {
		xlog.Fatal("Invalid parameters!")
	}
}

//检查文件锁
func checkLock() bool {
	if ok := util.PathExist(conf.LockFile); !ok {
		lock, _ := os.OpenFile(conf.LockFile, os.O_RDWR|os.O_CREATE, 0644)
		lock.Close()
	}
	flock := util.New(conf.LockFile)
	err := flock.Lock()
	if err != nil {
		return false
	}
	return true
}
