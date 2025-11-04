package prototype

type Prototype interface {
	Clone() Prototype
}
type Student struct {
	Name string
	Age  int
}

func (s *Student) Clone() Prototype {
	return &Student{
		Name: s.Name,
		Age:  s.Age,
	}
}
