package agrupamento_modulos

import (
	"fmt"
	"modulo/funcoes"
)

func ExecutarFuncoesAvancadas() {
	fmt.Println("--------------------------------")
	fmt.Println("FUNÇÕES")
	fmt.Println("--------------------------------")
	fmt.Println("FUNÇÃO COM RETORNO:", funcoes.FuncaoComRetorno(1, 2))
	fmt.Println("REUPERANDO VALOR DA FUNÇÃO COM VARIAVEL:", funcoes.RecuperandoValorDaFuncaoComVariavel(1, 2))
	fmt.Print("FUNÇÃO SEM RETORNO: ")
	funcoes.FuncaoSemRetorno()
	fmt.Println("PASSANDO FUNÇÃO PARA VARIAVEL:", funcoes.PassandoFuncaoParaVariavel(1, 2))
	soma, subtracao := funcoes.FuncaoComMaisDeUmRetorno(1, 2)
	fmt.Println("FUNÇÃO COM MAIS DE UM RETORNO:", soma, subtracao)
	soma, _ = funcoes.FuncaoComMaisDeUmRetorno(1, 2)
	fmt.Println("FUNÇÃO COM MAIS DE UM RETORNO - IGNORANDO SEGUNDO RETORNO:", soma)
	fmt.Print("\n")

	fmt.Println("--------------------------------")
	fmt.Println("FUNÇÕES AVANÇADAS")
	fmt.Println("--------------------------------")
	somaNomeado, subtracaoNomeado := funcoes.FuncaoRetornoNomeado(10, 5)
	fmt.Println("FUNÇÃO COM RETORNO NOMEADO - Soma:", somaNomeado, "Subtração:", subtracaoNomeado)
	funcoes.FuncaoVariaticaComMaisDeUmParametro(1, 2, 3, 4, 10)
	funcoes.FuncaoVariaticaComMaisDeUmParametroComRetorno("Ola Mundo", 1, 2, 3, 4, 10)
	soma, subtracao = funcoes.FuncaoComMaisDeUmRetorno(1, 2)
	fmt.Println("FUNÇÃO COM MAIS DE UM RETORNO:", soma, subtracao)
	fmt.Println("FUNÇÃO RECURSIVA:", funcoes.FuncaoRecursiva(15))
	defer funcoes.Defer()
	funcoes.SemDefer()
	fmt.Println("Aluno aprovado:", funcoes.AlunoAprovado(7, 8))
	funcoes.FuncaoPanic(5, 4)
	texto := "Dentro da main"
	fmt.Println(texto)
	novaFuncao := funcoes.FuncaoClosure()
	novaFuncao()
	numero := 10
	funcoes.FuncaoPonteiro(&numero)
	fmt.Println("Numero:", numero)
	fmt.Print("\n")
}

