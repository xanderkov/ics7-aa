package levenstein

import "testing"

func BenchmarkRecursiveLen10(b *testing.B) {
	src := "abaoboaobj"
	dest := "da;ldfjalj"

	for i := 0; i < b.N; i++ {
		CountLevNoRec(src, dest)
	}

}

func BenchmarkRecursiveLen40(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;lj"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofj"

	for i := 0; i < b.N; i++ {
		CountLevNoRec(src, dest)
	}

}

func BenchmarkRecursiveLen80(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountLevNoRec(src, dest)
	}

}

func BenchmarkRecursiveLen160(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountLevNoRec(src, dest)
	}

}



func BenchmarkRecursiveLen240(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountLevNoRec(src, dest)
	}

}

func BenchmarkCountDamNoRec5(b *testing.B) {
	src := "abaob"
	dest := "da;ld"

	for i := 0; i < b.N; i++ {
		CountDamNoRec(src, dest)
	}

}

func BenchmarkCountDamNoRec10(b *testing.B) {
	src := "abaoboaobj"
	dest := "da;ldfjalj"

	for i := 0; i < b.N; i++ {
		CountDamNoRec(src, dest)
	}

}

func BenchmarkCountDamNoRec40(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;lj"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofj"

	for i := 0; i < b.N; i++ {
		CountDamNoRec(src, dest)
	}

}

func BenchmarkCountDamNoRec80(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountDamNoRec(src, dest)
	}
}

func BenchmarkCountDamNoRec160(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountDamNoRec(src, dest)
	}

}



func BenchmarkCountDamNoRec240(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountDamNoRec(src, dest)
	}

}

func BenchmarkCountDamRecCacheLen5(b *testing.B) {
	src := "abaob"
	dest := "da;ld"

	for i := 0; i < b.N; i++ {
		CountDamRecCache(src, dest)
	}

}


func BenchmarkCountDamRecCacheLen10(b *testing.B) {
	src := "abaoboaobj"
	dest := "da;ldfjalj"

	for i := 0; i < b.N; i++ {
		CountDamRecCache(src, dest)
	}

}

func BenchmarkCountDamRecCacheLen40(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;lj"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofj"

	for i := 0; i < b.N; i++ {
		CountDamRecCache(src, dest)
	}

}

func BenchmarkCountDamRecCacheLen80(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountDamRecCache(src, dest)
	}

}

func BenchmarkCountDamRecCacheLen160(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountDamRecCache(src, dest)
	}

}



func BenchmarkCountDamRecCacheLen240(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountDamRecCache(src, dest)
	}

}

func BenchmarkCountDamRecNoCacheLen5(b *testing.B) {
	src := "abaob"
	dest := "da;ld"

	for i := 0; i < b.N; i++ {
		CountDamRecNoCache(src, dest)
	}

}

func BenchmarkCountDamRecNoCacheLen10(b *testing.B) {
	src := "abaoboaobj"
	dest := "da;ldfjalj"

	for i := 0; i < b.N; i++ {
		CountDamRecNoCache(src, dest)
	}

}

/*

func BenchmarkCountDamRecNoCacheLen20(b *testing.B) {
	src := "abaoboaobjaobjoajlaf"
	dest := "da;ldfjaljdasl;fjals"

	for i := 0; i < b.N; i++ {
		CountDamRecNoCache(src, dest)
	}

}

func BenchmarkCountDamRecNoCacheLen80(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountDamRecNoCache(src, dest)
	}

}


func BenchmarkCountDamRecNoCacheLen240(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjfabaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdajda;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

	for i := 0; i < b.N; i++ {
		CountDamRecNoCache(src, dest)
	}

}
*/