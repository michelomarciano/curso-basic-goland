# PONTEIRO

Ponteiros em Go são variáveis que armazenam o endereço de memória de outra variável. Eles permitem trabalhar diretamente com a memória e modificar valores de forma indireta.

**OBSERVAÇÃO**: Quando atribuímos um valor para uma variável, por padrão esse valor é uma **cópia**. Ponteiros permitem trabalhar com a referência original, não uma cópia.

## Atribuição por Cópia vs Referência

### Atribuição por Cópia (Padrão)
Quando você atribui uma variável a outra, Go cria uma cópia do valor. Modificações em uma variável não afetam a outra.

**Exemplo:**
```go
func AtribuiValorParaVariavel() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println("Valor da variavel1:", variavel1, "Valor da variavel2:", variavel2)
	// Saída: Valor da variavel1: 10 Valor da variavel2: 10

	variavel1++
	fmt.Println("Valor da variavel1:", variavel1, "Valor da variavel2:", variavel2)
	// Saída: Valor da variavel1: 11 Valor da variavel2: 10
	// variavel2 não foi afetada porque é uma cópia
}
```

## Operadores de Ponteiro

### 1. Operador `&` (Endereço)
O operador `&` retorna o endereço de memória de uma variável.

**Exemplo:**
```go
valor := 10
ponteiro := &valor  // ponteiro armazena o endereço de memória de valor
fmt.Println(ponteiro) // Mostra o endereço de memória (ex: 0xc0000140a0)
```

### 2. Operador `*` (Dereferência)
O operador `*` acessa ou modifica o valor apontado por um ponteiro.

**Exemplo:**
```go
valor := 10
ponteiro := &valor

fmt.Println(*ponteiro) // 10 - acessa o valor apontado
*ponteiro = 20         // modifica o valor apontado
fmt.Println(valor)     // 20 - valor foi modificado através do ponteiro
```

## Declarando Ponteiros

**Sintaxe:**
```go
var ponteiro *tipo
```

**Exemplos:**
```go
var ptr *int           // ponteiro para int (nil por padrão)
valor := 10
ptr = &valor           // ptr agora aponta para valor

// Ou de forma mais direta:
valor := 10
ptr := &valor          // declaração curta
```

## Passando Ponteiros para Funções

Ponteiros são muito úteis para modificar valores dentro de funções, já que Go passa argumentos por valor (cópia) por padrão.

**Exemplo - Sem ponteiro (não funciona):**
```go
func modificarValor(x int) {
    x = 42  // Modifica apenas a cópia local
}

valor := 10
modificarValor(valor)
fmt.Println(valor)  // 10 - valor original não foi modificado
```

**Exemplo - Com ponteiro (funciona):**
```go
func modificarValor(ptr *int) {
    if ptr != nil {
        *ptr = 42  // Modifica o valor original
    }
}

valor := 10
fmt.Println("Antes da modificação:", valor)  // 10

modificarValor(&valor)  // Passa o endereço de valor

fmt.Println("Depois da modificação:", valor)  // 42
```

## Verificando se Ponteiro é Nil

Ponteiros não inicializados têm valor `nil`. Sempre verifique se um ponteiro não é `nil` antes de dereferenciá-lo.

**Exemplo:**
```go
var ptr *int
fmt.Println(ptr == nil)  // true

if ptr != nil {
    fmt.Println(*ptr)  // Seguro - só executa se ptr não for nil
} else {
    fmt.Println("Ponteiro é nil")
}
```

**IMPORTANTE**: Tentar acessar um ponteiro `nil` causa um panic em tempo de execução:
```go
var ptr *int
fmt.Println(*ptr)  // PANIC: runtime error: invalid memory address or nil pointer dereference
```

## Quando Usar Ponteiros

**Use ponteiros quando:**
- Você precisa modificar o valor de uma variável dentro de uma função
- Você quer evitar copiar grandes estruturas de dados (melhor performance)
- Você trabalha com structs e quer modificar seus campos
- Você precisa representar valores opcionais (ponteiro pode ser `nil`)

**Não use ponteiros quando:**
- Você está trabalhando com tipos primitivos pequenos (int, bool, etc.) - a cópia é mais eficiente
- Você não precisa modificar o valor original
- Você quer garantir imutabilidade

## Resumo

| Operação | Sintaxe | Descrição |
|----------|---------|-----------|
| Obter endereço | `&variavel` | Retorna o endereço de memória |
| Declarar ponteiro | `var ptr *int` | Declara um ponteiro para int |
| Acessar valor | `*ponteiro` | Acessa o valor apontado |
| Modificar valor | `*ponteiro = valor` | Modifica o valor apontado |
| Verificar nil | `ponteiro == nil` | Verifica se ponteiro é nil |

## Casos de Uso
- **Modificar valores em funções**: Passar variáveis para funções que precisam alterar seus valores
- **Performance com structs grandes**: Evitar copiar structs grandes, passando apenas o ponteiro
- **Valores opcionais**: Representar valores que podem não existir (nil)
- **Estruturas de dados mutáveis**: Modificar elementos de structs, slices, maps
- **APIs que retornam erros**: Padrão Go de retornar (resultado, erro) onde resultado pode ser ponteiro

**Exemplo prático:**
```go
// Modificar struct sem copiar
func atualizarUsuario(u *Usuario) {
    u.Nome = "Novo Nome"  // Modifica o original
}

// Valores opcionais
func buscarUsuario(id int) (*Usuario, error) {
    // Se não encontrar, retorna nil
    if id == 0 {
        return nil, errors.New("ID inválido")
    }
    return &Usuario{ID: id}, nil
}

// Evitar cópia de struct grande
type DadosGrandes struct {
    // ... muitos campos
}
func processar(dados *DadosGrandes) {  // Passa ponteiro, não cópia
    // Processa dados
}
```

