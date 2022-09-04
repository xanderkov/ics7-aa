package main
import "fmt"
import levenstein "lab-01/levenstein"

/*
	1. Посчитать расстояние Левенштейна не рекурсивно
	2. Домерау-Левенштейн не рекурсивно
	3. Домерау-Левенштейн рек. без кэша
	4. Домерау-Левенштейн рек. с кэшем
*/

func main() {
	fmt.Println("Расстояние Левенштейна")

	fmt.Println()	

	fmt.Printf("Введте первое слово: ")
	firstWord := levenstein.ReadWord()
	fmt.Print("Введите второе слово: ")
	secondWord := levenstein.ReadWord()

	fmt.Println()

	distLevNoRec, distMat := levenstein.CountLevNoRec(firstWord, secondWord)
	fmt.Print("Левейнштейн не рекурсивный: ", distLevNoRec)
	fmt.Println()
	distMat.PrintMatrix()
	fmt.Println()	

	distDamNoRec, distMat := levenstein.CountDamNoRec(firstWord, secondWord)
	fmt.Print("Домерау-Левенштейн не рекурсивный: ", distDamNoRec)
	fmt.Println()
	distMat.PrintMatrix()
	fmt.Println()
	
	distDamRecNoCache := levenstein.CountDamRecNoCache(firstWord, secondWord)
	fmt.Print("Домерау-Левенштейн рекурсивный без кэша: ", distDamRecNoCache)
	
	fmt.Println()

	distDamRecCache := levenstein.CountDamRecCache(firstWord, secondWord)
	fmt.Print("Домерау-Левенштейн рекурсивный c кэшем: ", distDamRecCache)
	fmt.Println()
}