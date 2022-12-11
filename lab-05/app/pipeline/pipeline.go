package pipeline

import (
	"time"
	"../dbscan"
	"../structures"
)


func Pipeline(count int, ch chan int, filename string, minPtx int, eps float64) *structures.Queue {
	first := make(chan *structures.PipeTask, count)
	second := make(chan *structures.PipeTask, count)
	third := make(chan *structures.PipeTask, count)

	line := InitQueue(count)

	get_file := func() {
		for {
			select {
			case pipe_task := <- first:
				pipe_task.Generated = true

				pipe_task.Start_generating = time.Now()

				pipe_task.Source = dbscan.ReadFile(filename)
				
				pipe_task.End_generatig = time.Now()

				second <- pipe_task
			}
		}
	}

	get_dbscan := func() {
		for {
			select {
			case pipe_task := <- second:
				pipe_task.Dbscan_mode = true

				pipe_task.Start_dbscan = time.Now()

				pipe_task.Dbscan = dbscan.DbscanAlgorithm(pipe_task.Source, minPtx, eps)

				pipe_task.End_dbscan = time.Now()

				third <- pipe_task
			}
		}
	}

	match := func ()  {
		for {
			select {
			case pipe_task := <- third:
				pipe_task.Pattern_matched = true

				pipe_task.Start_match = time.Now()
				pipe_task.Pattern_index = dbscan.SaveFile(pipe_task.Dbscan)
				pipe_task.End_match = time.Now()

				line.Enqueue(pipe_task)
				if (pipe_task.Num == count - 1) {
					ch <- 0
				}
			}
		}
	}

	go get_file()
	go get_dbscan()
	go match()

	for i := 0; i < count; i++ {
		pipe_task := new(structures.PipeTask)
		pipe_task.Num = i + 1
		first <- pipe_task
	}
	return line
}