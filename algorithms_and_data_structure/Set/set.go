package set

// Set is a collection, where each element occurs exactly one time
type Set map[int]bool

// New creates a new Set with intially contains the passed elements
func New(ns ...int) Set {
	//TODO: Implementation
	return nil
}

// Insert inserts the given elements into the set
func (set Set) Insert(ns ...int) {
	//TODO: Implementation
}

// Delete deletes the given elements from the set
func (set Set) Delete(ns ...int) {
	//TODO: Implementation
}

// Contains tests, if the given elements are in the set or not
func (set Set) Contains(ns ...int) bool {
	//TODO: Implementation
	return false
}

// ContainsAny tests, if any (at least one) of the given element is in the set or not
func (set Set) ContainsAny(ns ...int) bool {
	//TODO: Implementation
	return false
}

// Values returns the set elements as a slice
func (set Set) Values() []int {
	//TODO: Implementation
	return nil
}

// Union returns the union set of s1 and s2
func Union(s1, s2 Set) Set {
	//TODO: Implementation
	return nil
}

// Union returns the intersection set of s1 and s2
func Intersection(s1, s2 Set) Set {
	//TODO: Implementation
	return nil
}
