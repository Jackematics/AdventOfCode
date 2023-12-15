package pipe

type Ground struct {
	GridIndex GridIndex
}

func (pipe Ground) Enter(direction Direction) Exit {
	return Exit{Index: GridIndex{Row: -1, Col: -1}, Direction: North}
}

func (pipe Ground) GetGridIndex() GridIndex {
	return pipe.GridIndex
}

func (pipe Ground) ValidEntryDirections() []Direction {
	return []Direction{}
}
