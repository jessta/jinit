package sshd

import (
	"./network"
	"./files"
)
var Up chan bool
func init(){
	Up = make(chan bool)
	go sshdup(Up)
}

func sshdup(ch chan bool){
	for net:=<-network.Up; net == false; {}
	for need:=<-file.Require("/etc/sshd.conf"); need == false; {}
	for {
		ch <- true
	}
}
