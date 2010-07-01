package files

const (
	ROOT="/"
	HOME="/home/"
	
)
var fstab map[string] chan bool
func init(){
	fstab = make(map[string]chan bool)	
}

/*TODO: fix smaller prefix problem*/
func Requires(filename string) (chan bool){ 
	for path,channel := range fstab {
		if strings.HasPrefix(filename,path){
			return channel
		}
	}
}