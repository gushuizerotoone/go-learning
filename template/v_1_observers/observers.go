package v_1_observers

import "fmt"

// This is the object to be observed
type Service struct {
	// define some fields here...
	A int
}

// Interface to implement for observe handlers
type IServiceHandler interface {
	OnServiceAdd(s *Service)
	OnServiceDelete(s *Service)
	OnServiceUpdate(s *Service)
}

// Config to hold observers
type ServiceConfig struct {
	ServiceHandlers []IServiceHandler
}

func NewServiceConfig() *ServiceConfig {
	var serviceHandlers []IServiceHandler
	return &ServiceConfig{ServiceHandlers: serviceHandlers}
}

// register observe handlers
func (s *ServiceConfig) registerHandler(handler IServiceHandler) {
	s.ServiceHandlers = append(s.ServiceHandlers, handler)
}

// First Handler
type FirstHandler struct {
}

func (f *FirstHandler) OnServiceAdd(service *Service) {
	fmt.Printf("FirstHandler OnServiceAdd: %q \n", service.A)
}

func (f *FirstHandler) OnServiceDelete(service *Service) {
	fmt.Printf("FirstHandler OnServiceDelete: %q \n", service.A)
}

func (f *FirstHandler) OnServiceUpdate(service *Service) {
	fmt.Printf("FirstHandler OnServiceUpdate: %q \n", service.A)
}

// Second Handler
type SecondHandler struct {
}

func (s *SecondHandler) OnServiceAdd(service *Service) {
	fmt.Printf("SecondHandler OnServiceAdd: %q \n", service.A)
}

func (s *SecondHandler) OnServiceDelete(service *Service) {
	fmt.Printf("SecondHandler OnServiceDelete: %q \n", service.A)
}

func (s *SecondHandler) OnServiceUpdate(service *Service) {
	fmt.Printf("SecondHandler OnServiceUpdate: %q \n", service.A)
}

