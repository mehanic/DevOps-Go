package namedtype

type VertexID int
type EdgeID int

func NewVertexIDGenerator() func() VertexID {
	var next VertexID
	return func() VertexID {
		next++
		return next
	}
}

func NewEdgeIDGenerator() func() EdgeID {
	var next EdgeID
	return func() EdgeID {
		next++
		return next
	}
}

func NewEdge(eid EdgeID) EdgeID {
	return eid
}
