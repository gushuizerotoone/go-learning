package v_2_abstract_class

type ProcessService struct {
	processors map[string]Processor
}

func NewProcessService() *ProcessService {
	// Init Processors
	processorService := &ProcessService{}
	processorService.processors = make(map[string]Processor, 0)
	processorService.processors["FirstProcessor"] = &FirstProcessor{}
	processorService.processors["SecondProcessor"] = &SecondProcessor{}
	return processorService
}

func (p *ProcessService) doSomething(processorName string) {
	processor := p.processors[processorName]

	// Common operation of abstract class
	processor.CommonOperation()

	// Sub-class function
	processor.TopUp()

	// Sub-class function
	processor.Withdraw()
}