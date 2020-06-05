package v_2_abstract_class

import "testing"

func TestAbstractClass(t *testing.T) {
	service := NewProcessService()
	service.doSomething(FIRST_RROCESSOR)
	service.doSomething(SECOND_PROCESSOR)
}
