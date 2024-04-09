package myPackageShape

// type Shaper interface {
// 	GetName() string
// 	GetSize() int
// }

type Shape struct {
	name string
	size int
}

// Constructor
func NewShape(name string, size int) *Shape {
	return &Shape{
		name: name,
		size: size,
	}
}

func (s *Shape) GetName() string {
	return s.name
}

func (s *Shape) GetSize() int {
	return s.size
}
