package golibsdatacontainer

type BinaryData struct {
	data []byte
}

func NewBinaryData(data []byte) DataContainer {
	return &BinaryData{data}
}

func (d *BinaryData) Kind() DataKind {
	return BINARY
}
func (d *BinaryData) Get() interface{} {
	return d.data
}
