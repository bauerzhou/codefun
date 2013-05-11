package arr

import "fmt"

var arr[10] int

func main(){
   arr[1] = 10;
   arr[2] = 12;
   arr[3] = 13;
   fmt.Println(arr[1:3])
}

func TestArr(){
   arr[1] = 10;
   arr[2] = 12;
   arr[3] = 13;
   fmt.Println(arr[1:3])
}

func TestSlice(){
   sl := make([]int,10)
   s1 := arr[0:3]
   s2 := arr[0:8]
   fmt.Println(s1,s2,sl)

}
