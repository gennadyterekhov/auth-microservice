package serializer

type Interface interface {
	Serialize(data interface{}) ([]byte, error)
}
