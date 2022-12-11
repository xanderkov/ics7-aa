package main
import (
	"fmt"
	"./pipeline"
)


func main() {

	fmt.Println("Алгоритм DBSCAN")

	ch := make(chan int)
	count := 20
	pipeline_queue := pipeline.Pipeline(count, ch)
	_ = pipeline_queue
	<- ch

	pipeline_queue_sync := pipeline.Sync(count)
	_ = pipeline_queue_sync

	// pipeline.PerfLog(pipeline_queue, pipeline_queue_sync)

}
