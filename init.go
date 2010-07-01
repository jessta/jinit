package main
/*
8g sshd.go
8g network.go
8g init.go
8l -o goinit init.8
./goinit


So, this is an init.
It's really static. It knows about a set number of boot up 
processes and their dependancies.
Dependancies are managed by startup groutines waiting on channels


Services can be fine gained about what they need and when they actually need it.

*/

//import what the services you want to run.
import (
"time"
"os"
"./sshd"
"./network"
//"jinit/syslog"
)

func SpawnMonitor(argv0 string, argv []string, envv []string, dir string, fd []*File) os.Error {
	pid,err: = os.ForkExec(argv0, argv, envv, dir, fd) (pid int, err Error)
	if err != nil {
		return err
	}
	for ;os.Kill(pid, 0) == 0;{
		_=time.Sleep(10000)
	}
}
func main(){
	
	println(<-sshd.Up)
	println(<-network.Up)
	println("done")
}

