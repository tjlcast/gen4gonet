package main

// @SERVER
type HelloWorldService struct {
	Name string
	Age int
	Score float64
}

func NewHelloWorldService(name string, age int, score float64) *HelloWorldService {
	return &HelloWorldService{Name: name, Age: age, Score: score}
}

// @RPC
func (srv *HelloWorldService) Hello() string {
	return "This is " + srv.Name
}

// @RPC
func (srv *HelloWorldService) Hello1(name string, age int) string {
	return "Hello: " + name
}

// @RPC
func (srv *HelloWorldService) Hello2(car Car) string {
	return "Hello: " + car.PlateNo
}