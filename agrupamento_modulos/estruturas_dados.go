package agrupamento_modulos

import (
	"fmt"
	"modulo/array"
	"modulo/heranca"
	"modulo/maps"
	"modulo/ponteiro"
	"modulo/slice"
	"modulo/structs"
)

func ExecutarEstruturasDados() {
	fmt.Println("--------------------------------")
	fmt.Println("STRUCTS")
	fmt.Println("--------------------------------")
	structs.Structs()
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("HERANÇA")
	fmt.Println("--------------------------------")
	cachorro := heranca.Cachorro{}
	cachorro.Cor = "Preto"
	cachorro.Nome = "Rex"
	cachorro.Idade = 5
	cachorro.Peso = 15.5
	cachorro.Raça = "Labrador"
	fmt.Println(cachorro.Cor)
	fmt.Println(cachorro.Nome)
	fmt.Println(cachorro.Idade)
	fmt.Println(cachorro.Peso)
	fmt.Println(cachorro.Raça)
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("ARRAY")
	fmt.Println("--------------------------------")
	array.DeclaracaoEAtribuicaoSeparada()
	array.DeclaracaoEInicializacaoNaMesmaLinha()
	array.InicializacaoComTamanhoInferido()
	array.InicializacaoComIndicesEspecificos()
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("SLICE")
	fmt.Println("--------------------------------")
	slice.Slice()
	slice.Append()
	slice.AppendMultiplos()
	slice.RemoverItemPorIndice()
	slice.AtribuiArrayASlice()
	slice.AtribuiArrayASlicePeloIndice()
	slice.CriarSliceComMake()
	slice.CriarSliceComMakeVazioComCapacidadeInicial()
	slice.CriarSliceComMakePrePreenchidoComZeros()
	slice.CriarSliceComMakeTamanhoECapacidadeDiferentes()
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("PONTEIRO")
	fmt.Println("--------------------------------")
	ponteiro.AtribuiValorParaVariavel()
	ponteiro.DiferencaEntrePonteiroEValor()
	ponteiro.ModificarValorDaVariavelApontadaPorPonteiro()
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("MAPS")
	fmt.Println("--------------------------------")
	maps.Maps()
	maps.MapAninhado()
	maps.DeletarItemDoMap()
	maps.AdicionarItemNoMap()
	maps.AdicionarMapEmOutroMap()
	fmt.Print("\n")
}

