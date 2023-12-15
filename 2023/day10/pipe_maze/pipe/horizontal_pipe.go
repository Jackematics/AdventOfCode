package pipe

type HorizontalPipe struct {
	GridIndex GridIndex
}

func (pipe HorizontalPipe) Enter(direction Direction) Exit {
	if direction == East {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row, Col: pipe.GridIndex.Col - 1}, Direction: East}
	}

	if direction == West {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row, Col: pipe.GridIndex.Col + 1}, Direction: West}
	}

	return Exit{Index: GridIndex{Row: -1, Col: -1}, Direction: North}
}

func (pipe HorizontalPipe) GetGridIndex() GridIndex {
	return pipe.GridIndex
}

func (pipe HorizontalPipe) ValidEntryDirections() []Direction {
	return []Direction{East, West}
}
