package pipeline

import (
	"time"
	"../dbscan"
)

func Pipeline(count int, chan int) *Queue {
	first := make(chan *PipeTask, count)
	second := make(chan *PipeTask, count)
	third := make(chan *PipeTask, count)

	line := initQueue(count)

	get_file := func() {
		for {
			select {
			case pipe_task := <- first:
				pipe_task.generated = true

				pipe_task.start_generating = time.Now()

				pipe_task.source = dbscan.ReadFile("1.txt")
				
				pipe_task.end_generating = time.Now()

				second <- pipe_task
			}
		}
	}

	get_dbscan = func() {
		for {
			select {
			case pipe_task := <- second:
				pipe_task.dbscan_mode = true

				pipe_task.start_dbscan = time.Now()

				pipe_task.dbscan = dbscan.DbscanAlgorithm(pipe_task.source)

				pipe_task.end_dbscan = time.Now()

				third <- pipe_task
			}
		}
	}

	match := func ()  {
		for {
			select {
			case pipe_task := <- third:
				pipe_task.pattern_matched = true

				pipe_task.start_match = time.Now()
				pipe_task.pattern_index = dbscan.SaveFile(pipe_task.dbscan)
				pipe_task.end_match = time.Now()

				line.enqueue(pipe_task)
				if (pipe_task.num == count - 1) {
					ch <- 0
				}
			}
		}
	}

	go get_file()
	go get_dbscan()
	go match()

	for i := 0; i < count; i++ {
		pipe_task := new(PipeTask)
		pipe_task.num = i + 1
		first <- pipe_task
	}
	return line
}