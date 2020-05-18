package main

import(
	"log"
	"fmt"
	"time"
)


func double(x int) (result int) {
	defer func() { 
		fmt.Printf("double(%d) = %d\n", x, result) 
	}()
	
	return x+x
}


func main(){
	double(3)
}

func bigSlowOperation(){
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)	
}


func trace(msg string) func(){
	start := time.Now()
	log.Printf("enter %s", msg)
	return func(){  log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
