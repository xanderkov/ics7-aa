package structures

import "time"

const STRING_SIZE = 10
const PATTERN_SIZE = 5


type Matrix struct {
	Rows int
	Cols int
	Values [][]int
}

type PipeTask struct {
	// helpers
	Num  			int
	Generated		bool
	Dbscan_mode     bool
	Pattern_matched	bool


	// time labels
	Start_generating time.Time
	End_generatig    time.Time
	Start_dbscan	 time.Time
	End_dbscan		 time.Time
	Start_match		 time.Time
	End_match		 time.Time

	// data
	Source        Matrix
	Dbscan        Matrix
	Pattern_index int
}

type Queue struct {
	Queue [](*PipeTask)
	Size  int
}

func (qu *Queue) Enqueue (t *PipeTask) {
	if (qu.Size != len(qu.Queue) - 1) {
		qu.Queue[qu.Size + 1] = t
		qu.Size++
	}
}

func (qu *Queue) Dequeue () *PipeTask {
	t := qu.Queue[0]
	qu.Queue = qu.Queue[1:]
	qu.Size--
	return t
}



func Allocate(rows, cols int) Matrix {
	var dest Matrix

	dest.Rows = rows
	dest.Cols = cols
	dest.Values = make([][]int, dest.Rows)
	for i := range dest.Values {
		dest.Values[i] = make([]int, dest.Cols)
	}

	return dest
}