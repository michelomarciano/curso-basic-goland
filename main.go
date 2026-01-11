package main

import (
	"fmt"
	agrupamento_modulos "modulo/agrupamento_modulos"

)

func main() {
	// Módulos
	agrupamento_modulos.ExecutarModulos()

	// Tipos Básicos
	agrupamento_modulos.ExecutarTiposBasicos()

	// Operadores e Controle de Fluxo
	agrupamento_modulos.ExecutarOperadores()
	agrupamento_modulos.ExecutarControleFluxo()

	// Estruturas de Dados
	agrupamento_modulos.ExecutarEstruturasDados()

	// Funções Avançadas
	agrupamento_modulos.ExecutarFuncoesAvancadas()

	// Métodos e JSON
	agrupamento_modulos.ExecutarMetodosEJSON()

	//INTERFACES
	agrupamento_modulos.ExecutarInterface()
	agrupamento_modulos.ExecutarInerfaceGenerica()


	fmt.Println("\n✅ Programa executado com sucesso!")

}
