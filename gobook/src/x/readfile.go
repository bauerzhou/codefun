package x

import (
	"os"
)
func ReadX(){
	buf := make([]byte, 1024) 
	
	f, _ := os.OpenFile("g://jobs.txt", os.O_RDONLY,0666) 
	
	defer f.Close() 
	
	for { 
		n, _ := f.Read(buf) 
		if n == 0 { break } 
		os.Stdout.Write(buf[0:n]) 
	}
}

