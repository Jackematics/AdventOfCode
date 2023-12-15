package pipe

type NorthAndWestBend struct {
	GridIndex GridIndex
}

func (pipe NorthAndWestBend) Enter(direction Direction) Exit {
	if direction == North {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row, Col: pipe.GridIndex.Col - 1}, Direction: East}
	}

	if direction == West {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row - 1, Col: pipe.GridIndex.Col}, Direction: South}
	}

	return Exit{Index: GridIndex{Row: -1, Col: -1}, Direction: East}
}

func (pipe NorthAndWestBend) GetGridIndex() GridIndex {
	return pipe.GridIndex
}

func (pipe NorthAndWestBend) ValidEntryDirections() []Direction {
	return []Direction{North, West}
}
