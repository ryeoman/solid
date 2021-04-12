package q2

// Reference: https://medium.com/movile-tech/single-responsibility-principle-in-swift-61ee11ee81b5
// GeometricGraphics class/struct case

// Question:
// 1. Is below code conform SRP?
// 2. If not, what part is violating SRP?
// 3. please build a solution that conform SRP

// Some context:
// we have a main class that responsible to draw geometrical Graphics

// Persistence handle Data Structure, store/fetch
type Persistence struct {
}

func NewPersistence() *Persistence {
	return &Persistence{}
}

func (p *Persistence) Save(square Square) {
	// Logic to save square to persistence
}

type Square struct {
	edge        int
	persistence Persistence
}

func NewSquare(edge int, persistence Persistence) *Square {
	return &Square{
		edge:        edge,
		persistence: persistence,
	}
}

// First Responsibility
func (s *Square) Area() int {
	return s.edge * s.edge
}

// Second Responsibility
func (s *Square) Save() {
	s.persistence.Save(*s)
}

// Coupled Responsibilities
func (s *Square) SaveAccordingToArea() {
	area := s.Area()
	if area > 20 {
		s.persistence.Save(*s)
	}
}
