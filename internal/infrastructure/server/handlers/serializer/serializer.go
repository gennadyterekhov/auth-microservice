package serializer

type Serializer interface {
	Serialize(data interface{}) ([]byte, error)
	SerializeOne(dto interface{}) ([]byte, error)
}

type Pack struct {
	Users      Serializer
	Orders     Serializer
	Register   Serializer
	Login      Serializer
	Categories Serializer
}
