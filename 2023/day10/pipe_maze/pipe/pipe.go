package pipe

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

type GridIndex struct {
	Row int
	Col int
}

type Exit struct {
	Index     GridIndex
	Direction Direction
}

type Pipe interface {
	Enter(direction Direction) Exit
	GetGridIndex() GridIndex
	ValidEntryDirections() []Direction
}
