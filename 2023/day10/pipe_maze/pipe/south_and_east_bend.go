package pipe

type SouthAndEastBend struct {
	GridIndex GridIndex
}

func (pipe SouthAndEastBend) Enter(direction Direction) Exit {
	if direction == South {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row, Col: pipe.GridIndex.Col + 1}, Direction: West}
	}

	if direction == East {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row + 1, Col: pipe.GridIndex.Col}, Direction: North}
	}

	return Exit{Index: GridIndex{Row: -1, Col: -1}, Direction: North}
}

func (pipe SouthAndEastBend) GetGridIndex() GridIndex {
	return pipe.GridIndex
}

func (pipe SouthAndEastBend) ValidEntryDirections() []Direction {
	return []Direction{South, East}
}
