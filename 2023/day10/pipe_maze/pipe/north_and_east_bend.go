package pipe

type NorthAndEastBend struct {
	GridIndex GridIndex
}

func (pipe NorthAndEastBend) Enter(direction Direction) Exit {
	if direction == North {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row, Col: pipe.GridIndex.Col + 1}, Direction: West}
	}

	if direction == East {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row - 1, Col: pipe.GridIndex.Col}, Direction: South}
	}

	return Exit{Index: GridIndex{Row: -1, Col: -1}, Direction: North}
}

func (pipe NorthAndEastBend) GetGridIndex() GridIndex {
	return pipe.GridIndex
}

func (pipe NorthAndEastBend) ValidEntryDirections() []Direction {
	return []Direction{North, East}
}
