package conf

import (
	"fmt"
	"os"
	"../comm"
	"flag"
	"bytes"
	"github.com/ini"
	"../xlog"
	"path"
)

type Config struct {
	Process_name string
	Command      string
	Autostart    bool
	Autorestart  bool
	Logfile      string
}

var (
	List                                bool
	Start, Stop, Restart, Status, Input string
	CheckCommand                        string
	LockFile                            string
	Conf                                = make(map[string]*Config)
)

func init() {
	Dir, _ := os.Getwd()
	LockFile = Dir + "/monitor.lock"
	var buffer bytes.Buffer
	buffer.WriteString(Dir)
	buffer.WriteString("/")
	buffer.WriteString(path.Base(os.Args[0]))
	buffer.WriteString(" check")
	CheckCommand = buffer.String()
	loadConf()
	count := len(os.Args)
	if (count > 2) {
		var buf bytes.Buffer
		for i := 1; i < count; i++ {
			buf.WriteString(os.Args[i])
			buf.WriteString(" ")
		}
		err_msg := buf.String()
		fmt.Print("undefined command: ")
		fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, err_msg, 0x1B)
		os.Exit(1)
	}
	if (count == 2) {
		Input = os.Args[1]
		/*switch os.Args[1] {
		case "start", "stop", "restart","status","check":
			Input = os.Args[1]
		default:
			fmt.Print("undefined command: ")
			fmt.Printf( "%c[1;40;31m%s%c[0m\n", 0x1B,os.Args[1],0x1B)
			os.Exit(1)
		}*/
	}
	flag.StringVar(&Start, comm.START, "", "start a svr")
	flag.StringVar(&Stop, comm.STOP, "", "stop a svr")
	flag.StringVar(&Restart, comm.RESTART, "", "restart a svr")
	flag.StringVar(&Status, comm.STATUS, "", "status a svr")
	flag.BoolVar(&List, "list", false, "Lists commands")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Another process monitor\n")
		fmt.Fprintf(os.Stderr, "version: 1.0, Created by simplejia [11/2017]\n\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  start\n")
		fmt.Fprintf(os.Stderr, "        Start the monitor process\n")
		fmt.Fprintf(os.Stderr, "  stop\n")
		fmt.Fprintf(os.Stderr, "        Stop the monitor process\n")
		fmt.Fprintf(os.Stderr, "  restart\n")
		fmt.Fprintf(os.Stderr, "        Restart the monitor process\n")
		fmt.Fprintf(os.Stderr, "  status\n")
		fmt.Fprintf(os.Stderr, "        show the monitor process status\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func loadConf() {
	//读取配置文件信息
	// _, err := ini.Load(conf.RootPath+"/conf/conf.ini")
	cfg, err := ini.Load("./conf/conf.ini")
	if (err != nil) {
		xlog.Fatal("", err)
	} else {
		names := cfg.SectionStrings()
		if (names[0] == "DEFAULT") {
			names = remove(names, 0)
		}
		for _, v := range names {
			conf := new(Config)
			isset_process_name := cfg.Section(v).HasKey("process_name")
			isset_command := cfg.Section(v).HasKey("command")
			isset_autostart := cfg.Section(v).HasKey("autostart")
			isset_autorestart := cfg.Section(v).HasKey("autorestart")
			isset_logfile := cfg.Section(v).HasKey("logfile")
			if (!isset_process_name || !isset_command || !isset_autostart || !isset_autorestart) {
				xlog.Fatal("", "process_name|command|autostart|autorestart 必须填写")
			}
			conf.Process_name = cfg.Section(v).Key("process_name").String()
			conf.Command = cfg.Section(v).Key("command").String()
			conf.Autostart, err = cfg.Section(v).Key("autostart").Bool()
			if (err != nil) {
				xlog.Fatal("", "autostart shoule be bool")
			}
			conf.Autorestart, err = cfg.Section(v).Key("autorestart").Bool()
			if (err != nil) {
				xlog.Fatal("", "autorestart shoule be bool")
			}
			if (isset_logfile) {
				conf.Logfile = cfg.Section(v).Key("logfile").String()
				xlog.Info(conf.Logfile, "Process:", conf.Process_name, "is loading")
			} else {
				conf.Logfile = xlog.Logfile
				xlog.Info(conf.Logfile, "Process:", conf.Process_name, "is loading")
			}
			Conf[v] = conf
		}
	}

}

//删除函数
func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
