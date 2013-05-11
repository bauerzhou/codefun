package main

import "fmt"
import "even"
import "arr"
import "x"

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)
func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	k := 7;
	fmt.Println("k is even ",even.Even(k))
	
	arr.TestArr()
	arr.TestSlice()
	x.ReadX()
	fmt.Println("x.ReadBuf...................")
	x.ReadBuf("g://jobs.txt")
// ouwtput: From the string we read: 56.12 5212 Go
}