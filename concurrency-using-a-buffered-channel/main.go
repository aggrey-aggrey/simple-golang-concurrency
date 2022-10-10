package main


import (
	"fmt"
)

func main (){


	//make a buffered channel by giving a capacity
	// so it won't block till the channel is full
	c := make (chan string,2)

	//send message to a channel
	c <- "hello"
	c <- "world"

	//receive message from a channel
	msg := <- c
	fmt.Println(msg)

	msg = <- c
	fmt.Println(msg)
}