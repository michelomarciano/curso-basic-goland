# MODIFICADOR DE ACESSO

Em Go, a visibilidade de funções, variáveis e tipos é determinada pela primeira letra do nome:

**MAIÚSCULA**: Quando uma função, variável ou tipo está escrito com a primeira letra em MAIÚSCULA, estamos dizendo que é **pública** e pode ser importada por outros pacotes.

**Exemplo:** `func FuncaoPublica() { ... }`

**MINÚSCULA**: Quando uma função, variável ou tipo está escrito com a primeira letra em minúscula, ela só está visível dentro do próprio pacote (privada).

**Exemplo:** `func funcaoPrivada() { ... }`

## Casos de Uso
- **API pública vs implementação interna**: Exportar apenas o que é necessário para outros pacotes usar
- **Encapsulamento**: Esconder detalhes de implementação, expondo apenas a interface pública
- **Controle de acesso**: Prevenir uso indevido de funções auxiliares ou variáveis internas
- **Manutenibilidade**: Facilitar refatoração interna sem quebrar código externo
- **Bibliotecas e pacotes**: Criar APIs limpas e bem definidas para consumo externo

**Exemplo prático:**
```go
// Pacote: calculadora
// Público - pode ser usado por outros pacotes
func Soma(a, b int) int {
    return somar(a, b)  // chama função privada
}

// Privado - apenas para uso interno
func somar(a, b int) int {
    return a + b
}
```

