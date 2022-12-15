package pipeline

import (
	"time"
	"../structures"
	"fmt"
	"strings"
	"github.com/logrusorgru/aurora"
)

// Log used to log conveyor tasks time.
func Log(q *structures.Queue) {
	var (
		fdtime time.Duration
		sdtime time.Duration
		tdtime time.Duration
	)

	l := q.Queue
	start := l[0].Start_generating

	fmt.Printf("%38v\n", aurora.BgRed("STARTING TIME"))
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	fmt.Printf(
		"|%3v|%10v|%10v|%10v|\n",
		aurora.Green("N"), aurora.Green("Чтение"), aurora.Green("DBSCAN"), aurora.Green("Запись"),
	)
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	for i := 0; i < len(l); i++ {
		if l[i] != nil {
			fmt.Printf(
				"|%3v|%10v|%10v|%10v|\n",
				i, l[i].Start_generating.Sub(start), l[i].Start_dbscan.Sub(start), l[i].Start_match.Sub(start),
			)
		}
	}
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")

	fmt.Printf("%38v\n", aurora.BgRed("FINISHING TIME"))
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	fmt.Printf(
		"|%3v|%10v|%10v|%10v|\n",
		aurora.Green("N"), aurora.Green("Чтение"), aurora.Green("DBSCAN"), aurora.Green("Запись"),
	)
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	for i := 0; i < len(l); i++ {
		if l[i] != nil {
			fmt.Printf(
				"|%3v|%10v|%10v|%10v|\n",
				i, l[i].End_generatig.Sub(start), l[i].End_dbscan.Sub(start), l[i].End_match.Sub(start),
			)
		}
	}
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")

	fmt.Printf("%38v\n", aurora.BgRed("Время простоя"))
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	fmt.Printf("|%3v|%32v|\n", aurora.Green("N"), aurora.Green("Время простоя"))
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
	for i := 0; i < len(l)-2; i++ {
		fdtime += l[i+1].Start_generating.Sub(start) - l[i].End_generatig.Sub(start)
		sdtime += l[i+1].Start_dbscan.Sub(start) - l[i].End_dbscan.Sub(start)
		tdtime += l[i+1].Start_match.Sub(start) - l[i].End_match.Sub(start)
	}
	ldtime := []time.Duration{fdtime, sdtime, tdtime}
	for i := 0; i < 3; i++ {
		fmt.Printf("|%3v|%32v|\n", i+1, ldtime[i])
	}
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 36), "+")
}