package x

import (
	"os"
	"bufio"
	"fmt"
)

func ReadBuf(path string){
	buf := make([]byte, 1024) 
	f, _ := os.OpenFile(path, os.O_RDONLY, 0666) 
	defer f.Close() 
	r := bufio.NewReader(f)
	w:= bufio.NewWriter(os.Stdout) 
	outFile, er := os.OpenFile("g://job1.txt",os.O_WRONLY,0777)
	if er!= nil {
	   fmt.Println("open file fail",er)
	}
	 
	defer w.Flush() 
	for { 
		n, _ := r.Read(buf) 
		if n == 0 { break } 
		w.Write(buf[0:n])
		outFile.Write(buf[0:n])
	}
}