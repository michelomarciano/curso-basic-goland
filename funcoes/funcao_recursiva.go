package funcoes

func FuncaoRecursiva(posicao int) int {
if posicao <= 1 {
	return posicao
}
	return FuncaoRecursiva(posicao - 2) + FuncaoRecursiva(posicao - 1)

}