[中文 README](#中文)

# [process-monitor](https://github.com/xiongchao123/process-monitor)
## Functions
* Used to monitor and manage processes

## Installation
> go get -u github.com/simplejia/cmonitor

## Implementation
* Check each process by sending signal0 per 500ms, each monitored process is monitored in an independent goroutine.
* When cmonitor is started, it will receive management commands（start|stop|restart|status|...）

## Playing with cmonitor
* Configuration files: [conf.ini](https://github.com/xiongchao123/process-monitor/blob/master/conf/conf.ini) (ini format, supporting annotations)

```
[Demo]  ;监控模块名称
process_name = Demo  ;监控进程名称
command=php 监控程序目录/test/test.php  ;监控进程启动命令
autostart=true       ;是否自启动
autorestart=true     ;是否随进程启动而重启
logfile=/var/log/monitor.log  ;日志文件存储，可以为空，默认为xlog目录下monitor.log
[Test]
process_name = Test
command=php 监控程序目录/test/test2.php
autostart=true
autorestart=true
logfile=/var/log/monitor.log

```
* Run：./monitor [start|stop|restart|status|check]
* Management：./monitor -[list|status|start|stop|restart]

## Notice
* The conf of monitor are reported via ini, they also can be recorded in local ini.  And it is recommended to report the log by [ini](github.com/ini).
* When cmonitor starts its monitor processes, the console logs (monitor.log) of the monitored processes will be output to corresponding process directory which will be saved up to 30 days.，the history logs is like monitor.{day}.log
* When cmonitor is starting, all monitored processes will be started according to its conf.ini configuration. When the monitored processes have already been started and meet the configuration requirements, monitor will automatically add them to the monitor list.
* cmonitor will periodically check the process status. If there is abnormal process exit, it will repeatedly try to restart it and record the error log.

## demo
```
$ ./monitor start

Process:demo is already stop...
Process:demo start [success]
Process:test is already stop...
Process:test start [success]

```

---
中文
===

# [cmonitor](http://github.com/simplejia/cmonitor)
## 功能
* 用于进程监控，管理

## 安装
> go get -u github.com/simplejia/cmonitor

## 实现
* 被监控进程启动后，按每500ms执行一次状态检测（通过发signal0信号检测），每个被监控进程在一个独立的协程里被监测。
* monitor启动后接收管理命令（start|stop|restart|status|...）

## 使用方法
* 配置文件：[conf.ini](https://github.com/xiongchao123/process-monitor/blob/master/conf/conf.ini) (ini格式，支持注释)

```
    [Demo]  ;监控模块名称
    process_name = Demo  ;监控进程名称
    command=php 监控程序目录/test/test.php  ;监控进程启动命令
    autostart=true       ;是否自启动
    autorestart=true     ;是否随进程启动而重启
    logfile=/var/log/monitor.log  ;日志文件存储，可以为空，默认为xlog目录下monitor.log
    [Test]
    process_name = Test
    command=php 监控程序目录/test/test2.php
    autostart=true
    autorestart=true
    logfile=/var/log/monitor.log
```
* 运行方法：./monitor [start|stop|restart|status|check]
* 进程管理：./monitor -[list|status|start|stop|restart]

## 注意
* monitor的配置文件解析使用ini，需要先引入ini包 [ini](github.com/ini)
* 当cmonitor启动时，会根据conf.ini配置启动所有被监控进程，当被监控进程已经启动过，并且符合配置要求时，cmonitor会自动将其加入监控列表
* monitor会定期检查进程运行状态，如果进程异常退出，monitor会反复重试拉起，并且记录日志
* 当被监控进程为多进程运行模式，monitor只监控管理父进程(子进程应实现检测父进程运行状态，并随父进程退出而退出）
* 被监控进程以nohup方式启动，所以你的程序就不要自己设定daemon运行了
* 每30秒通过ps方式检测一次进程状态，如果出现任何异常，比如有多份进程启动等，记日志

## demo
```
$ ./monitor start

Process:demo is already stop...
Process:demo start [success]
Process:test is already stop...
Process:test start [success]

```
