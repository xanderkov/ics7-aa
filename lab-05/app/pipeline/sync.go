package pipeline

import (
	"time"

	"../dbscan"
	"../structures"
)

func gen_string_sync(task *structures.PipeTask) *structures.PipeTask {
	task.Generated = true

	task.Start_generating = time.Now()

	task.Source = dbscan.ReadFile("1.txt")

	task.End_generatig = time.Now()

	return task
}

func get_table_sync(task *structures.PipeTask) *structures.PipeTask {
	task.Dbscan_mode = true

	task.Start_dbscan = time.Now()

	task.Dbscan = dbscan.DbscanAlgorithm(task.Source, 2, 2)
	task.End_dbscan = time.Now()

	return task
}

func match_sync(task *structures.PipeTask) *structures.PipeTask {
	task.Pattern_matched = true

	task.Start_match = time.Now()
	task.Pattern_index = dbscan.SaveFile(task.Dbscan)
	task.End_match = time.Now()

	return task
}


func Sync(count int) *structures.Queue {
	line_first := InitQueue(count)
	line_second := InitQueue(count)
	line_third := InitQueue(count)

	for i := 0; i < count; i++ {
		pipe_task := new(structures.PipeTask)
		pipe_task = gen_string_sync(pipe_task)
		line_first.Enqueue(pipe_task)
		if (len(line_first.Queue) != 0) {
			pipe_task = get_table_sync(line_first.Dequeue())
			line_second.Enqueue(pipe_task)
			if (len(line_second.Queue) != 0) {
				pipe_task = match_sync(line_second.Dequeue())
				line_third.Enqueue(pipe_task)
			}
		}
	}
	return line_third
}