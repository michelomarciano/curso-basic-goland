package slice

import "fmt"
import "reflect" // para saber o tipo de um slice

func Slice() {
	slice := []int{ 1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println("Slice:", slice)

	}

	func Append() {
		slice := []int{ 1, 2, 3, 4, 5}
		slice = append(slice, 6)
		fmt.Println("Append:", slice)
	}

	func AppendMultiplos() {
		slice := []int{ 1, 2, 3, 4, 5}
		slice = append(slice, 6, 7, 8, 9, 10)
		fmt.Println("Append multiplos:", slice)
	}

	func AtribuiArrayASlice() {
		array := [5]int{ 1, 2, 3, 4, 5}
		slice := array[:]
		fmt.Println("Atribuindo array a slice:", slice)
	}

	func AtribuiArrayASlicePeloIndice() {
		array := [5]int{ 1, 2, 3, 4, 5}
		slice := array[1:3]
		fmt.Println("Atribuindo array a slice pelo indice 1 até 3:", slice)
	}


	func RemoverItemPorIndice() {
		slice := []int{ 1, 2, 3, 4, 5}
		slice = append(slice[:2], slice[3:]...)
		fmt.Println("Removendo item por indice:", slice, "Tamanho do slice:", len(slice), "Capacidade do slice:", cap(slice))
	}

	func CriarSliceComMake() {
		slice := make([]int, 5)
		fmt.Println("Criando slice com make:", slice, "Tamanho do slice:", len(slice), "Capacidade do slice:", cap(slice))
	}

	func CriarSliceComMakeVazioComCapacidadeInicial() {
		slice := make([]int, 0, 10)
		fmt.Println("Criando slice com make vazio com capacidade inicial:", slice, "Tamanho do slice:", len(slice), "Capacidade do slice:", cap(slice))
	}

	func CriarSliceComMakePrePreenchidoComZeros() {
		slice := make([]int, 5)
		fmt.Println("Criando slice com make preenchido com zeros:", slice, "Tamanho do slice:", len(slice), "Capacidade do slice:", cap(slice))
	}
	func CriarSliceComMakeTamanhoECapacidadeDiferentes() {
		slice := make([]int, 3, 10)
		fmt.Println("Criando slice com make tamanho e capacidade diferentes:", slice, "Tamanho do slice:", len(slice), "Capacidade do slice:", cap(slice))
	}




