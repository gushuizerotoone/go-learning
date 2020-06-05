package v_2_abstract_class

import "fmt"

// Interface
type Processor interface {
	// common function that should be implement by abstract class
	CommonOperation()

	// Detail function that should be implement by sub-class
	TopUp()
	Withdraw()
}

// Abstract class of implement interface
type processor struct {
	// Use composite of abstract class to inherit common functions
	Processor
}

// Abstract class implement common function
func (p *processor) CommonOperation() {
	fmt.Println("This is common operation")
}

// First implement type
type FirstProcessor struct {
	// Use composite of abstract class to inherit common functions
	// common function is not needed to implement in Sub-class
	*processor
}

func (f *FirstProcessor) TopUp() {
	fmt.Println("FirstProcessor TopUp")
}

func (f *FirstProcessor) Withdraw() {
	fmt.Println("FirstProcessor Withdraw")
}

// Second implement type
type SecondProcessor struct {
	// Use composite of abstract class to inherit common functions
	// common function is not needed to implement in Sub-class
	*processor
}

func (s *SecondProcessor) TopUp() {
	fmt.Println("SecondProcessor TopUp")
}

func (s *SecondProcessor) Withdraw() {
	fmt.Println("SecondProcessor Withdraw")
}






