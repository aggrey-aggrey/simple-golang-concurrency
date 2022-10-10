
package main

import (
	"fmt"
	"time"
)

func main (){
	//make a channel of type string
	c:= make(chan string)

	go count("sheep",c)
	//go count("fish", c)

	//receive a message from the channel
	msg := <- c
	fmt.Println(msg)

}

func count (thing string, c chan string){

	for i:= 1; i<=5 ; i ++{

		//send the value of thing into the channel
		c <- thing
		time.Sleep(time.Millisecond * 500)

	}

}