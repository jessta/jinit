package network

var Up chan bool
func init(){
	Up=make(chan bool)
	go networkup(Up)
}

func networkup(ch chan bool){
	for {
		ch <- true
	}
}
