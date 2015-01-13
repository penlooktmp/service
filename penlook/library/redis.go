package library

type Redis struct {
	Name string
}

func (r Redis) Print() string {
	return "Hello"
}
 