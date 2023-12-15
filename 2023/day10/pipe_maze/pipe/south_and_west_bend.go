package pipe

type SouthAndWestBend struct {
	GridIndex GridIndex
}

func (pipe SouthAndWestBend) Enter(direction Direction) Exit {
	if direction == South {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row, Col: pipe.GridIndex.Col - 1}, Direction: East}
	}

	if direction == West {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row + 1, Col: pipe.GridIndex.Col}, Direction: North}
	}

	return Exit{Index: GridIndex{Row: -1, Col: -1}, Direction: East}
}

func (pipe SouthAndWestBend) GetGridIndex() GridIndex {
	return pipe.GridIndex
}

func (pipe SouthAndWestBend) ValidEntryDirections() []Direction {
	return []Direction{South, West}
}
