package levenstein

import "testing"

func BenchmarkRecursiveLen80(b *testing.B) {
	src := "abaoboaobjaobjoajlafjgdsljfl;ajflasjd;ljdasf;ladsj;ldfsajkdasf;ldsfal;dfsjasdkjf"
	dest := "da;ldfjaljdasl;fjalsdjflkadsjflajfhdiofjoijdsodfajodasjodsafjsadoiasdjdzpoijsdaj"

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

