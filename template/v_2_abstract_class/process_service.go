package v_2_abstract_class

type ProcessorName string

const (
	FIRST_RROCESSOR ProcessorName = "FirstProcessor"
	SECOND_PROCESSOR ProcessorName = "SecondProcessor"
)

type ProcessService struct {
	processors map[ProcessorName]Processor
}

func NewProcessService() *ProcessService {
	// Init Processors
	processorService := &ProcessService{}
	processorService.processors = make(map[ProcessorName]Processor, 0)
	processorService.processors[FIRST_RROCESSOR] = &FirstProcessor{}
	processorService.processors[SECOND_PROCESSOR] = &SecondProcessor{}
	return processorService
}

func (p *ProcessService) doSomething(processorName ProcessorName) {
	processor := p.processors[processorName]

	// Common operation of abstract class
	processor.CommonOperation()

	// Sub-class function
	processor.TopUp()

	// Sub-class function
	processor.Withdraw()
}