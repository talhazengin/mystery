package astrarium

// Defines two way connection between two nodes.
type Link struct {
	Id        int
	NextLinks []*Link
}
