package agrupamento_modulos

import (
	"modulo/interfaces"
)

func ExecutarInterface(){
	r := interfaces.Retangulo{Altura: 10, Largura: 30}
	interfaces.EscreverArea(r)

	c := interfaces.Circulo{Raio: 10}
	interfaces.EscreverArea(c)
}

func ExecutarInerfaceGenerica(){
	interfaces.Generica("Ola Mundo")
	interfaces.Generica(1.0)
}