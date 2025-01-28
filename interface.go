package golibsdatacontainer

type DataContainer interface {
	Kind() DataKind
	Get() interface{}
}
