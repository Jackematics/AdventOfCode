package pipe

type VerticalPipe struct {
	GridIndex GridIndex
}

func (pipe VerticalPipe) Enter(direction Direction) Exit {
	if direction == North {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row + 1, Col: pipe.GridIndex.Col}, Direction: North}
	}

	if direction == South {
		return Exit{Index: GridIndex{Row: pipe.GridIndex.Row - 1, Col: pipe.GridIndex.Col}, Direction: South}
	}

	return Exit{Index: GridIndex{Row: -1, Col: -1}, Direction: North}
}

func (pipe VerticalPipe) GetGridIndex() GridIndex {
	return pipe.GridIndex
}

func (pipe VerticalPipe) ValidEntryDirections() []Direction {
	return []Direction{North, South}
}
