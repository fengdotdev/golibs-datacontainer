package golibsdatacontainer

func NewDataContainer(kind DataKind, data interface{}) DataContainer {
	switch kind {
	case BINARY:
		return &BinaryData{data.([]byte)}
	case STRING:
		return &StringData{data.(string)}
	case NUMBER:
		panic("unimplemented")
		return &NumberData{data.(int)}
	}
	panic("Unknown data kind")
	return nil
}
