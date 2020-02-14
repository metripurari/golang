package main

import(
	"fmt"
	"time"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	ch := make(chan int)
	ch1 := make(chan []int)
	go fib(ch, ch1, &wg)
	for i :=0; i<20;i++{
		select {
		case n := <- ch:
			fmt.Println("no recieved", n)
		case <-time.After(1000):
			fmt.Println("timeout 1")
		}
		
		select {
		case n1 := <- ch1:
			fmt.Println("Array recieved", n1)
		case <-time.After(1):
			fmt.Println("timeout 1")
		}
		
		fmt.Println("loop", i)
		time.Sleep(2 * time.Second)
	}
	
	wg.Wait()
}


func fib(ch chan int, ch1 chan []int, wg *sync.WaitGroup){
	i, j := 0, 1
	for {
		ch <- j
		ch1 <- []int{i,j}
		i, j = j, i + j
		wg.Add(1)
		
	}
	wg.Done()
}