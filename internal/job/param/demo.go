package param

type DemoJobParams struct {
	*abstractJobParam
	// add other fields
}

func NewDemoJobParams() *DemoJobParams {
	return &DemoJobParams{
		abstractJobParam: newAbstractJobParam(""),
	}
}
