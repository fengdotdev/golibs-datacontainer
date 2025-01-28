package golibsdatacontainer

type StringData struct {
	data string
}

func NewStringData(data string) DataContainer {
	return &StringData{data}
}

func (d *StringData) Get() interface{} {
	return d.data
}

func (d *StringData) Kind() DataKind {
	return STRING
}
