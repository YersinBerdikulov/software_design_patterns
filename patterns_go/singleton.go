package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
	host     string
	password string
}

var singleInstance *single

func (s *single) check(host, password string) bool {
	if s.host == host && password == s.password {
		return true
	}
	return false
}

func getInstance(host, password string) *single {
	if !singleInstance.check(host, password) {
		return nil
	}
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
		} else {
			fmt.Println("Connection is already established.")
		}
	} else {
		fmt.Println("Connection is already established.")
	}

	return singleInstance
}

func main() {

	for i := 0; i < 5; i++ {
		go getInstance("192.180.01.10", "12412")
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
	fmt.Println("end")

}
