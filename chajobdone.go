package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
				if(j%5==0){
					fmt.Println("Hello")
				}else{
					fmt.Println("received job", j)
				}
            } else {
                fmt.Println("received all jobs")
                done <- false
                return
            }
        }
    }()

    for j := 1; j <= 10; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    
    fmt.Println("sent all jobs")

    <-done
}