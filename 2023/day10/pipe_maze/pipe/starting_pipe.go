package pipe

type StartingPipe struct {
	Pipe      Pipe
	GridIndex GridIndex
}

func (pipe StartingPipe) Enter(direction Direction) Exit {
	return pipe.Pipe.Enter(direction)
}

func (pipe StartingPipe) GetGridIndex() GridIndex {
	return pipe.GridIndex
}

func (pipe StartingPipe) ValidEntryDirections() []Direction {
	return pipe.Pipe.ValidEntryDirections()
}
