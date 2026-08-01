package main

import "fmt"

// func Email(SendEmail <-chan string, flag chan<- bool )// we can reduce the scope of channel
//that in a func we can receive/send channel
func Email(SendEmail chan string, flag chan bool ){
	defer func(){flag <- true}()

    for item := range SendEmail{
		fmt.Println(item)
	}
}

// Channel is basically the producer consumer concept of OS 
// producer is one thread and consumer in another thread otherwise causes blocking

func main() {

//unbuffered channel : Producer and consumer run together (send ↔ receive happen one at a time).
	Channel := make(chan string)     

	go func(){Channel <- "Hello"}()           // send to
	fmt.Println(<-Channel)                    //receiving

// buffered channel : Producer can get ahead until the buffer is full, then it must wait for the consumer.
// Implementation shows the queuing mechanism
	SendEmail := make(chan string,100) 
	flag := make(chan bool)
    
	go Email(SendEmail, flag)

	for i:= 0; i < 10; i++ {
		SendEmail <- fmt.Sprintf("%d@gmail.com",i)
	}
    close(SendEmail)

	<- flag 
	// Send email is a channel it use go routines and send the 100 email 
	// once after the channel completely sends data Email func prints 
	// the  defer keywords ensures in the func that whole func must be completed 
	// in the main <-flag is doing blocking 



	sendInt := make(chan int)
	sendStr := make(chan string)

	go func(){sendInt <- 2}()
	go func(){sendStr <- "hi"}()

//select is used when you want to wait on multiple channel operations at the same time
//if we not use it may have happened the sendStr has to unncessary wait for sendInt 
	for i:=0; i<2; i++{
		select{
		case receiveInt := <- sendInt:
			fmt.Println(receiveInt)
		case receiveStr := <- sendStr:
			fmt.Println(receiveStr)
		}
	}

	close(sendInt)
	close(sendStr)
}

