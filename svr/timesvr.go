package svr

import (
	"../conf"
	"../comm"
	"../xlog"
	"time"
	"fmt"
)

func AllProcs(cmd string) {
	switch  cmd {
	case comm.START:
		for service, config := range conf.Conf {
			if(config.Autorestart){
				Procs(comm.RESTART, service)
			}else if(config.Autostart){
				Procs(comm.START, service)
			}
		}
		StartCheck(conf.CheckCommand)
	case comm.STOP:
		/*for service, _ := range conf.Conf {
			Procs(cmd, service)
		}*/
		StopCheck(conf.CheckCommand)
	case comm.STATUS:
		for service, _ := range conf.Conf {
			Procs(cmd, service)
		}
		if ok := CheckProc(conf.CheckCommand); ok {
			fmt.Print("Process: monitor ")
			fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, "[is starting]", 0x1B)
		} else {
			fmt.Print("Process: monitor ")
			fmt.Printf("%c[1;40;31%s%c[0m\n", 0x1B, "[is stop]", 0x1B)
		}
	}
}

func Procs(cmd string, service string) {
	defer func() {
		if err := recover(); err != nil {
			xlog.Warn(conf.Conf[service].Logfile, err)
		}
	}()
	switch cmd {
	case comm.START:
		StartProc(conf.Conf[service])
	case comm.STOP:
		StopProc(conf.Conf[service])
	case comm.RESTART:
		RestartProc(conf.Conf[service])
	case comm.STATUS:
		GetProc(conf.Conf[service])
	}
}

func CheckProcs() {
	//进入时间服务
	tick1 := time.Tick(time.Millisecond * 500)
	//tick2 := time.Tick(time.Minute)
	//tick3 := time.Tick(time.Hour * 24)
	for {
		select {
		case <-tick1:
			for _, config := range conf.Conf {
				if ok := CheckProc(config.Command); !ok {
					xlog.Warn(config.Logfile, "Process:", config.Process_name, "进程中断")
					//start the process
					StartProc(config)
				}
			}
			time.Sleep(time.Second * 5)
		}
	}
}

