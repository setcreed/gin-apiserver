package service

type helloService struct {
}

type HelloService interface {
	SayHello() interface{}
}

func NewHelloService() HelloService {
	return &helloService{}
}

func (h *helloService) SayHello() interface{} {
	return "Hello, World!"
}
