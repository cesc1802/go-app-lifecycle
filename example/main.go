package main

import (
	"fmt"

	lifecycle "github.com/cesc1802/go-app-lifecycle"
)

type Service1 struct {
}

func (service *Service1) Start() error {
	fmt.Println("start service 1")
	return nil
}

func (service *Service1) Stop() error {
	fmt.Println("stop service 1")
	return nil
}

type Service2 struct {
}

func (service *Service2) Start() error {
	fmt.Println("start service 2")
	return nil
}

func (service *Service2) Stop() error {
	fmt.Println("stop service 2")
	return nil
}

func main() {
	var service1 Service1
	var service2 Service2

	appGroup := lifecycle.NewAppGroup(lifecycle.AppGroupOption{
		Name:     "app group",
		Services: []lifecycle.Service{&service1, &service2},
	})

	if err := appGroup.Run(); err != nil {
		fmt.Printf("Finish with error: %+v\n", err)
		return
	}

	fmt.Printf("Done!")
}
