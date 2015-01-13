package library

type Doc struct {
	Name string
}

func (d Doc) Print() string {
	return "Hello"
}
