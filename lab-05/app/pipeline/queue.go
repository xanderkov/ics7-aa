package pipeline
import "../structures"

func InitQueue(size int) *structures.Queue {
	qu := new(structures.Queue)
	qu.Queue = make([](*structures.PipeTask), size)
	qu.Size = -1

	return qu
}

