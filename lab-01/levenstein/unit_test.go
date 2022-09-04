package levenstein

import "testing"

func TestEmptyString(t *testing.T) {
	eth := 0
	w1 := ""
	w2 := ""
	
	res, mat := CountLevNoRec(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res, mat = CountDamNoRec(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
		mat.PrintMatrix()
	}

	res = CountDamRecNoCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = CountDamRecCache(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

}