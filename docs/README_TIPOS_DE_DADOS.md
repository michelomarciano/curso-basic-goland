# TIPOS DE DADOS

**INT**: Caso use somente `int` sem especificar o tamanho, o compilador usará a arquitetura do computador como base. 

**Exemplo:** Em uma arquitetura de 64 bits, `int` será equivalente a `int64`.

**UINT**: É um tipo inteiro sem sinal (unsigned integer). Isso significa que ele não representa números negativos — apenas 0 e números positivos.

**CHAR**: Não existe tipo `char` no Go. Quando você usa um caractere, ele sempre retornará um `int` representando o código ASCII/Unicode do caractere.

**Exemplo:** `'B'` retorna `66` (valor da tabela ASCII).

| Categoria | Tipo | Tamanho | Descrição | Exemplo |
|-----------|------|---------|-----------|---------|
| **Inteiros com Sinal** | `int8` | 8 bits | Inteiro de -128 a 127 | `var x int8 = 127` |
| | `int16` | 16 bits | Inteiro de -32.768 a 32.767 | `var x int16 = 32767` |
| | `int32` | 32 bits | Inteiro de -2.147.483.648 a 2.147.483.647 | `var x int32 = 2147483647` |
| | `int64` | 64 bits | Inteiro de -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807 | `var x int64 = 9223372036854775807` |
| | `int` | 32 ou 64 bits | Tamanho depende da arquitetura (32 ou 64 bits) | `var x int = 100` |
| **Inteiros sem Sinal** | `uint8` (byte) | 8 bits | Inteiro de 0 a 255 | `var x uint8 = 255` |
| | `uint16` | 16 bits | Inteiro de 0 a 65.535 | `var x uint16 = 65535` |
| | `uint32` | 32 bits | Inteiro de 0 a 4.294.967.295 | `var x uint32 = 4294967295` |
| | `uint64` | 64 bits | Inteiro de 0 a 18.446.744.073.709.551.615 | `var x uint64 = 18446744073709551615` |
| | `uint` | 32 ou 64 bits | Tamanho depende da arquitetura | `var x uint = 100` |
| | `uintptr` | - | Inteiro sem sinal grande o suficiente para armazenar um ponteiro | `var x uintptr` |
| **Ponto Flutuante** | `float32` | 32 bits | Número de ponto flutuante de precisão simples | `var x float32 = 3.14` |
| | `float64` | 64 bits | Número de ponto flutuante de precisão dupla | `var x float64 = 3.14159265359` |
| **Complexos** | `complex64` | 64 bits | Número complexo (float32 real + float32 imaginário) | `var x complex64 = 1 + 2i` |
| | `complex128` | 128 bits | Número complexo (float64 real + float64 imaginário) | `var x complex128 = 1 + 2i` |
| **Texto** | `string` | - | Sequência de caracteres UTF-8 | `var x string = "Hello"` |
| | `byte` | 8 bits | Alias para uint8 | `var x byte = 'A'` |
| | `rune` | 32 bits | Alias para int32, representa um código Unicode | `var x rune = '中'` |
| **Booleano** | `bool` | 1 bit | Valor verdadeiro ou falso | `var x bool = true` |
| **Coleções** | `array` | - | Array de tamanho fixo | `var x [5]int = [5]int{1, 2, 3, 4, 5}` |
| | `slice` | - | Array dinâmico (referência a um array) | `var x []int = []int{1, 2, 3}` |
| | `map` | - | Mapa (chave-valor) | `var x map[string]int = make(map[string]int)` |
| **Estruturas** | `struct` | - | Estrutura de dados personalizada | `type Pessoa struct { Nome string }` |
| **Ponteiros** | `*T` | - | Ponteiro para um tipo T | `var x *int` |
| **Funções** | `func` | - | Tipo de função | `var f func(int) int` |
| **Canais** | `chan` | - | Canal para comunicação entre goroutines | `var c chan int = make(chan int)` |
| **Interfaces** | `interface{}` | - | Interface vazia (aceita qualquer tipo) | `var x interface{} = 42` |

## Casos de Uso
- **Escolha do tipo correto**: Selecionar tipos apropriados para cada situação
  - `int` para contadores, índices
  - `float64` para cálculos científicos, financeiros
  - `string` para texto, nomes, descrições
  - `bool` para flags, estados on/off
  - `byte` para dados binários, imagens
  - `rune` para caracteres Unicode
- **Otimização de memória**: Usar tipos menores quando possível (int8 vs int64)
- **Precisão**: Escolher float32 ou float64 baseado na precisão necessária
- **Compatibilidade**: Usar tipos compatíveis com APIs externas

**Exemplo prático:**
```go
// Contador simples
var contador int = 0

// Cálculo financeiro (precisa de precisão)
var saldo float64 = 1234.56

// Dados binários
var imagem []byte

// Texto Unicode
var nome rune = '中'

// Estado booleano
var estaAtivo bool = true
```

