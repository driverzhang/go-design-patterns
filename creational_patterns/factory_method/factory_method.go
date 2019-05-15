package factory_method

type Data interface {
	GetData(id string) string
}

func Get(name string) Data {
	switch name {
	case "logs":
		return &Logs{}
	case "simBase":
		return &SimBase{}
	default:
		return nil
	}
}

type Logs struct {
}

func (l Logs) GetData(id string) string { return "this is logs id:" + id }

type SimBase struct{}

func (s SimBase) GetData(id string) string { return "this is simbase ids:" + id }
