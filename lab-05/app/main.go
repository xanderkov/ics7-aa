package main
import (
	"fmt"
	"./pipeline"
)


func main() {

	fmt.Println("Алгоритм DBSCAN")
	var minPtx int
	var eps float64
	var filename string
	minPtx = 2
	eps = 7
	filename = "data/200.txt"
	/*
	fmt.Print("Имя файла: ")
	fmt.Scan(&filename)
	fmt.Print("Радиус: ")
	fmt.Scan(&minPtx)
	fmt.Print("EPS: ")
	fmt.Scan(&eps)
	*/
	ch := make(chan int)
	count := 20
	pipeline_queue := pipeline.Pipeline(count, ch, filename, minPtx, eps)
	_ = pipeline_queue
	<- ch

	pipeline_queue_sync := pipeline.Sync(count)
	_ = pipeline_queue_sync

	pipeline.Log(pipeline_queue)
	fmt.Println("Done")
}
