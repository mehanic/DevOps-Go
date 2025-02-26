package abpattern

type VertexID int

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func NewVertexIDGenerator() func() VertexID {
	gen := NewIntGenerator()
	return func() VertexID {
		return VertexID(gen())
	}

}
