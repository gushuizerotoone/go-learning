package v_1_observers

import (
	"testing"
)

func TestObservers(t *testing.T) {
	// we can register handlers in construct func of higher level type
	serviceConfig := NewServiceConfig()
	serviceConfig.registerHandler(&FirstHandler{})
	serviceConfig.registerHandler(&SecondHandler{})

	service := &Service{A: 3}
	// Change service's value
	service.A = 5

	for _, h := range serviceConfig.ServiceHandlers {
		h.OnServiceAdd(service)
	}
}
