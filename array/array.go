package array

import "fmt"

func DeclaracaoEAtribuicaoSeparada() {
	var array [5]string
	array[0] = "Mike"
	array[1] = "Sophia"
	array[2] = "Luisa"
	array[3] = "Eduardo"
	array[4] = "Familia"
	fmt.Println(array)

}

func DeclaracaoEInicializacaoNaMesmaLinha() {	
	array2 := [5]string{"Mike", "Sophia", "Luisa", "Eduardo", "Familia"}
	fmt.Println(array2)

}

func InicializacaoComTamanhoInferido() { 
	array3 := [...]string{"Mike", "Sophia", "Luisa", "Eduardo", "Familia"}
	fmt.Println(array3)
}

func InicializacaoComIndicesEspecificos() { 
	array4 := [5]string{1: "Sophia", 3: "Eduardo"}
	fmt.Println(array4)
}
