package golibsdatacontainer

type NumberData struct {
	data int
}

func NewNumberData(data int) DataContainer {
	return &NumberData{data}
}

func (d *NumberData) Get() interface{} {
	return d.data
}

func (d *NumberData) Kind() DataKind {
	return NUMBER
}
