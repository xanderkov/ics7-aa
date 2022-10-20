package main
import "fmt"
import levenstein "lab-01/levenstein"

/*
#include <pthread.h>
#include <time.h>
#include <stdio.h>

static long long getThreadCpuTimeNs() {
    struct timespec t;
    if (clock_gettime(CLOCK_THREAD_CPUTIME_ID, &t)) {
        perror("clock_gettime");
        return 0;
    }
    return t.tv_sec * 1000000000LL + t.tv_nsec;
}
*/
import "C"

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
	fmt.Print("Дaмерау-Левенштейн не рекурсивный: ", distDamNoRec)
	fmt.Println()
	distMat.PrintMatrix()
	fmt.Println()
	
	distDamRecNoCache := levenstein.CountDamRecNoCache(firstWord, secondWord)
	fmt.Print("Дaмерау-Левенштейн рекурсивный без кэша: ", distDamRecNoCache)
	
	fmt.Println()

	distDamRecCache := levenstein.CountDamRecCache(firstWord, secondWord)
	fmt.Print("Дaмерау-Левенштейн рекурсивный c кэшем: ", distDamRecCache)
	fmt.Println()
}
