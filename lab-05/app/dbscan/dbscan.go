package dbscan

import (

	"../structures"
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

func ReadFile(filename string) structures.Matrix  {
	file, _ := os.Open(filename)
    var lines []string
	var points structures.Matrix
	
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text()) 
    }

	points = structures.Allocate(len(lines), len(lines))

	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")
		for j := 0; j < len(line); j++ {
			points.Values[i][j], _ = strconv.Atoi(line[j])
		}
	}
    return points
}


func DbscanAlgorithm(points structures.Matrix, minPtx int, eps float64) structures.Matrix {
	return points
}

func SaveFile(points structures.Matrix) int {
	if points.Rows > 0 {
		file, err := os.Create("data/Answer.txt")
		if err != nil {
			return 0
		}
		defer file.Close()

		for _, line := range points.Values {
			fmt.Fprintln(file, line)
		}
	}
	return 0
}