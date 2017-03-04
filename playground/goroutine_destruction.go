package main

import "fmt"

func mainGoroutine() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
