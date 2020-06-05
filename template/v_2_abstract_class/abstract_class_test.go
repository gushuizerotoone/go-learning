package v_2_abstract_class

import "testing"

func TestAbstractClass(t *testing.T) {
	service := NewProcessService()
	service.doSomething("FirstProcessor")
	service.doSomething("SecondProcessor")
}
