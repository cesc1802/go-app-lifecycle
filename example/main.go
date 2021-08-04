package main

import (
	"fmt"
	lifecycle "github.com/cesc1802/go-app-lifecycle"
)

type HelloWorldHttpService struct {

}

func (service *HelloWorldHttpService) Start() error {
	fmt.Print("start helloworld service")
	return nil
}

func (service *HelloWorldHttpService) Stop() error {
	fmt.Print("stop helloworld service")
	return nil
}

type OrderdHttpService struct {

}

func (service *OrderdHttpService) Start() error {
	fmt.Print("start order service")
	return nil
}

func (service *OrderdHttpService) Stop() error {
	fmt.Print("stop order service")
	return nil
}

func main() {
	var service HelloWorldHttpService
	var orderService OrderdHttpService

	app := lifecycle.New(lifecycle.Services(&service, &orderService), lifecycle.Name("hello service"))

	if err := app.Run(); err != nil {
		fmt.Printf("Finish with error: %+v\n", err)
		return
	}
	fmt.Printf("Done!")
}



