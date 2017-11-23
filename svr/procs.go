package svr

import (
	"../conf"
	"../xlog"
	"fmt"
	"os/exec"
)

func StartCheck(command string) {
	cmdStr := fmt.Sprintf(
		"nohup %s 1>/dev/null 2>"+xlog.ErrorFile+"&",
		command,
	)
	cmd := exec.Command("sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		xlog.Fatal(xlog.ErrorFile, "err", err)
	}
}

//停止自动检测进程
func StopCheck(command string) {
	cmdStr := fmt.Sprintf(
		"ps -ef|grep -v grep|grep \"%s\" |cut -c 9-15|xargs kill -9",
		command,
	)
	err := exec.Command("sh", "-c", cmdStr).Run()
	if err != nil {
		xlog.Fatal(xlog.ErrorFile, "err", err)
	}
	fmt.Print("Process: monitor", " stop ")
	fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, "[success]", 0x1B)
}

//启动进程
func StartProc(conf *conf.Config) {
	config := *conf
	command := config.Command
	logfile := config.Logfile
	cmdStr := fmt.Sprintf(
		"nohup %s 1>/dev/null 2>"+logfile+"&",
		command,
	)
	if ok := CheckProc(command); !ok {
		err := exec.Command("sh", "-c", cmdStr).Run()
		if err != nil {
			fmt.Print("Process:", config.Process_name, " start ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[fail]", 0x1B)
			xlog.Fatal(logfile, err)
		}
		fmt.Print("Process:", config.Process_name, " start ")
		fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, "[success]", 0x1B)
	} else {
		fmt.Print("Process:", config.Process_name)
		fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, " is already start...", 0x1B)
	}
}

//停止进程
func StopProc(conf *conf.Config) {
	config := *conf
	command := config.Command
	logfile := config.Logfile
	if ok := CheckProc(command); !ok {
		fmt.Print("Process:", config.Process_name)
		fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, " is already stop...", 0x1B)
	} else {
		cmdStr := fmt.Sprintf(
			"ps -ef| grep -v grep|grep \"%s\"|cut -c 9-15|xargs kill -9",
			command,
		)
		err := exec.Command("sh", "-c", cmdStr).Run()
		if err != nil {
			fmt.Print("Process:", config.Process_name, " stop ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[fail]", 0x1B)
			xlog.Fatal(logfile, err)
		}
		fmt.Print("Process:", config.Process_name, " stop ")
		fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, "[success]", 0x1B)
	}
}

//获取当前进程信息
func GetProc(conf *conf.Config) {
	config := *conf
	command := config.Command
	logfile := config.Logfile
	if ok := CheckProc(command); !ok {
		fmt.Println("Process:", config.Process_name, " is already stop...")
	} else {
		cmdStr := fmt.Sprintf(
			"ps -ef| grep -v grep|grep \"%s\"",
			command,
		)
		cmd := exec.Command("sh", "-c", cmdStr)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Print("Process:", config.Process_name, " Status ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[fail]", 0x1B)
			xlog.Fatal(logfile, err)
		}
		fmt.Println("Process:", config.Process_name, " Status:")
		fmt.Println(string(out))
	}
}

//重启进程
func RestartProc(conf *conf.Config) {
	//stop proc
	StopProc(conf)
	//start proc
	StartProc(conf)
}

//检查进程是否已经启动
func CheckProc(command string) (ok bool) {
	cmdStr := fmt.Sprintf(
		"ps aux| grep -v grep|grep \"%s\" |awk '{print $2}'",
		command,
	)
	cmd := exec.Command("sh", "-c", cmdStr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err:", err)
	}
	if (string(out) == "") {
		return false
	}
	return true
}
