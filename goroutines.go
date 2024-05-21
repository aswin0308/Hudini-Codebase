package main

// import (
// 	"fmt"
// 	"time"
// 	"os"
// )
// type Waitgroup struct {
// 	count int

// }
// func (c *Waitgroup) Add(num int){
// 	c.count += num
// }
// func (c *Waitgroup) Done(){
// 	if(c.count!=0){
// 	c.count = c.count-1
// 	}
// }
// func (c *Waitgroup) Wait(){
// 	for{
// 		if(c.count==0){
// 		os.Exit(0)
// 	}
// }
// }
// func main(){
// 	wg := Waitgroup{}
// 	for i:=0; i<3; i++{
// 	wg.Add(1)
// 	go execute("routine", &wg)
// 	}
// 	wg.Wait()

// }
// func execute(name string, wg *Waitgroup){
// 	for i := 1; i <= 5; i++ {
// 				time.Sleep(time.Second)
// 				fmt.Println(name, i)
// 			}
// 			wg.Done()
		
// 		}


// func display(messages chan string){
// 	messages <- "Hello, World!"
// }

// func main() {

//     messages := make(chan string)

//     go display(messages)

//     fmt.Println(<-messages)
// }

// Objective:
// Learn how to send and receive values using channels.

// Task:

// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
// The main function should receive the message from the channel and print it.
// Hints:

// Use an unbuffered channel for simplicity.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to create and use goroutines.

// Task:

// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:

// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.

// import (
// 	"fmt"
// 	"time"
// 	"os"
// 	"sync"
// )

// func execute(name string, wg *sync.WaitGroup) {
// 	for i := 1; i <= 5; i++ {
// 		time.Sleep(time.Second)
// 		fmt.Println(name, i)
// 	}
// 	wg.Done()

// }

// func main() {
// 	// Launching both the runners
// 	//var timer sync.WaitGroup
// 	wg := new(sync.WaitGroup)
// 	wg.Add(3)
// 	go execute("routine1", wg)
// 	go execute("routine2", wg)
// 	go execute("routine3", wg)
// 	wg.Wait()

// }

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to handle multiple senders with a single receiver using channels.

// Task:

// Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// The main function should receive both messages from the channel and print them.

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to use channels for communication between goroutines.

// import "fmt"

// func display(messages chan string){
// 	messages <- "Hello,"
// }

// func display2(messages chan string){
// 	messages <- "World!"
// }

// func main() {

//     messages := make(chan string)

//     go display(messages)
// 	go display2(messages)

//     fmt.Println(<-messages)
// 	fmt.Println(<-messages)
// }

// Task:

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:

// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.


// func main(){
// 	ch := make(chan int)

// 	go generatenum(ch)

// 	for i:=1; i<=10; i++ {
// 		fmt.Println(<-ch)
// 	}
// }
// func generatenum(ch chan int){
// 	for i:=1; i<=10; i++ {
// 		ch <- i
// 	}
// 	close(ch)

// }

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use a buffered channel.

// Task:

// Write a program that creates a buffered channel with a capacity of 3.
// The main function should send three integers (1, 2, 3) to the buffered channel without blocking.
// Then, receive and print the integers from the channel.

// func main(){
// 	mych := make(chan int,3)
// 	mych <- 1
// 	mych <- 2
// 	mych <- 3

// 	fmt.Println(<-mych)
// 	fmt.Println(<-mych)
// 	fmt.Println(<-mych)
// }





// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// The main function should use a select statement to receive messages from both channels and print them.
// Hints:
// Use two separate channels.
// Use the select statement inside a loop to receive from the channels.

// import ("fmt"
//         "time")


// func main(){
// 	chan1 := make(chan string)
// 	chan2 := make(chan string)
// 	var chan1closed bool
// 	var chan2closed bool


// 	go message1(chan1)
// 	go message2(chan2)
// 	for {
//         select {
//         case msg, ok := <-chan1:
//             if ok {
//                 fmt.Println("Received from sender1:", msg)
//             }else{
// 				chan1closed = true
// 			}
//         case msg, ok := <-chan2:
//             if ok {
//                 fmt.Println("Received from sender2:", msg)
//             }else{
// 				chan2closed = true
// 			}
//         }
// 		if chan1closed && chan2closed == true {
// 			break
// 		}
// 		break
		
//     }
// }
// func message1(ch chan string){
// 	for i:=1; i<=5; i++ {
// 		ch<- "sending message 1"
// 		time.Sleep(time.Second)
// 	}
// 	close(ch)
// }

// func message2(ch chan string){
// 		for i:=1; i<=5; i++ {
// 			ch<- "sending message 2"
// 			time.Sleep(2 *time.Second)
			
// 		}
// 		close(ch)
// 	}

// -------------------------------------------------------------------------------------

// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.

// Task:

// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.
import ("fmt"
	"sync")


func main(){
	 chan1 := make(chan string)
	wg := new(sync.WaitGroup)

    wg.Add(3)
	 go printmsg("Hai Python",chan1,wg)
	 go printmsg("Hai Java",chan1,wg)
	 go printmsg("Hai Golang",chan1,wg)
	 wg.Wait()
	 for i:=1; i<=3; i++ {
	 fmt.Println(<-chan1)
	}
}
func printmsg(msg string,chan1 chan string,wg *sync.WaitGroup){
	chan1 <- msg 
	wg.Done()
}



// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.

// import "fmt"

// func main()
