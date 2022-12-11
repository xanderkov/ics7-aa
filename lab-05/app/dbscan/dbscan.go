package dbscan

import (

	"../structures"
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
	"math"
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

func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}


func DbscanAlgorithm(points structures.Matrix, minPtx int, eps float64) structures.Matrix {
	clusterCount := 0
	currentPoint := points
	img := structures.Allocate(points.Rows, points.Cols)
	color := 0
	for i := 0; i < points.Rows; i++ {
		for j:=0; j < points.Cols; j++ {
			if currentPoint.Values[i][j] == 1 {
				var v1 []int
				var v2 []int

				v1 = append(v1, i)
				v2 = append(v2, j)
				neighborCountCheck := 0
				for k := 0; k < len(v1); k++ {
					if currentPoint.Values[v1[k]][v2[k]] == 1 {
						currentPoint.Values[v1[k]][v2[k]] = 0
						neighborCount := 0
						_, startX := MinMax([]int{0, v1[k] - minPtx})
						_, startY := MinMax([]int{0, v2[k] - minPtx})
						endX, _ := MinMax([]int{points.Rows, v1[k] + minPtx + 1})
						endY, _ := MinMax([]int{points.Cols, v2[k] + minPtx + 1})

						for x := startX; x < endX; x++ {
							for y := startY; y < endY; y++ {
								distance := math.Sqrt(math.Pow(float64(startX - endX), 2) + math.Pow(float64(startY - endY), 2))
								
								if distance <= eps && currentPoint.Values[x][y] == 1 {
									neighborCount++
									
									v1 = append(v1, x)
									v2 = append(v2, y)
								}
							}
						}

						if neighborCount >= minPtx {
							neighborCountCheck++
							img.Values[v1[k]][v2[k]] = color
						}
					}
					
				}
				if neighborCountCheck > 0 {
					clusterCount++
					color++
				}
				currentPoint.Values[i][j] = 0
			}
		}
	}
	return img
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