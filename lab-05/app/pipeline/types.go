package pipeline

import "time"

const STRING_SIZE = 10
const PATTERN_SIZE = 5


type PipeTask struct {
	// helpers
	num  			int
	generated		bool
	dbscan_mode     bool
	pattern_matched	bool


	// time labels
	start_generating time.Time
	end_generatig    time.Time
	start_dbscan	 time.Time
	end_dbscan		 time.Time
	start_match		 time.Time
	end_match		 time.Time

	// data
	source        []rune
	dbscan        map[rune]int
	pattern_index int
}

type Queue struct {
	queue [](*PipeTask)
	size  int
}