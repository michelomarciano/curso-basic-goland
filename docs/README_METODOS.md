# MÉTODOS

Métodos em Go são funções que pertencem a um tipo específico (geralmente structs). Eles permitem associar comportamentos a tipos, criando uma forma de programação orientada a objetos em Go.

**OBSERVAÇÃO**: Em Go, métodos são diferentes de funções. Funções são independentes, enquanto métodos pertencem a um tipo e são chamados através de uma instância desse tipo.

## Diferença entre Métodos e Funções

**Função (independente):**
```go
func Somar(a, b int) int {
    return a + b
}

// Chamada:
resultado := Somar(5, 3)
```

**Método (pertence a um tipo):**
```go
type Calculadora struct {
    Valor int
}

func (c Calculadora) Somar(numero int) int {
    return c.Valor + numero
}

// Chamada:
calc := Calculadora{Valor: 5}
resultado := calc.Somar(3)  // Método chamado através da instância
```

## Sintaxe de Métodos

**Formato geral:**
```go
func (receiver Tipo) NomeDoMetodo(parametros) tipoRetorno {
    // corpo do método
}
```

**Componentes:**
- `receiver` - A instância do tipo que recebe o método
- `Tipo` - O tipo ao qual o método pertence (geralmente uma struct)
- `NomeDoMetodo` - Nome do método
- `parametros` - Parâmetros de entrada (opcional)
- `tipoRetorno` - Tipo de retorno (opcional)

## Value Receiver vs Pointer Receiver

Existem dois tipos de receivers em Go: **value receiver** e **pointer receiver**. A escolha entre eles é importante e afeta o comportamento do método.

### Value Receiver (Recebedor por Valor)

Um value receiver recebe uma **cópia** da instância. Modificações feitas no método não afetam a instância original.

**Sintaxe:**
```go
func (u Usuario) NomeDoMetodo() {
    // u é uma cópia da instância original
}
```

**Exemplo:**
```go
type Usuario struct {
    Nome  string
    Email string
    Senha string
}

// Value receiver - recebe uma cópia
func (u Usuario) Salvar() {
    fmt.Println("Salvando usuario:", u.Nome)
    fmt.Println("Email:", u.Email)
    fmt.Println("Senha:", u.Senha)
    // Modificações em u não afetam a instância original
}

// Uso:
usuario := Usuario{Nome: "Mike", Email: "mike@example.com", Senha: "123456"}
usuario.Salvar()
```

**Características:**
- Recebe uma cópia da instância
- Modificações não afetam a instância original
- Mais seguro (não pode modificar acidentalmente)
- Pode ser menos eficiente para structs grandes (copia toda a struct)

### Pointer Receiver (Recebedor por Ponteiro)

Um pointer receiver recebe um **ponteiro** para a instância. Modificações feitas no método afetam a instância original.

**Sintaxe:**
```go
func (u *Usuario) NomeDoMetodo() {
    // u é um ponteiro para a instância original
}
```

**Exemplo:**
```go
type Usuario struct {
    Nome  string
    Email string
    Idade int
}

// Pointer receiver - recebe um ponteiro
func (u *Usuario) AtualizarIdade() {
    u.Idade++  // Modifica a instância original
}

// Uso:
usuario := Usuario{Nome: "Mike", Idade: 20}
usuario.AtualizarIdade()  // Idade agora é 21
fmt.Println(usuario.Idade) // 21
```

**Características:**
- Recebe um ponteiro para a instância
- Modificações afetam a instância original
- Mais eficiente para structs grandes (não copia, apenas passa o ponteiro)
- Permite modificar o estado do objeto

## Comparação Prática

**Exemplo completo mostrando a diferença:**
```go
type Usuario struct {
    Nome  string
    Idade int
}

// Value receiver - NÃO modifica a instância original
func (u Usuario) IncrementarIdade() {
    u.Idade++  // Modifica apenas a cópia
    fmt.Println("Idade dentro do método:", u.Idade)
}

// Pointer receiver - MODIFICA a instância original
func (u *Usuario) AtualizarIdade() {
    u.Idade++  // Modifica a instância original
    fmt.Println("Idade dentro do método:", u.Idade)
}

// Uso:
usuario := Usuario{Nome: "Mike", Idade: 20}

fmt.Println("Antes - Idade:", usuario.Idade) // 20

usuario.IncrementarIdade()  // Value receiver
fmt.Println("Depois de IncrementarIdade:", usuario.Idade) // 20 (não mudou!)

usuario.AtualizarIdade()  // Pointer receiver
fmt.Println("Depois de AtualizarIdade:", usuario.Idade) // 21 (mudou!)
```

**Saída:**
```
Antes - Idade: 20
Idade dentro do método: 21
Depois de IncrementarIdade: 20
Idade dentro do método: 21
Depois de AtualizarIdade: 21
```

## Quando Usar Cada Tipo de Receiver

### Use Value Receiver quando:
- ✅ O método **não precisa modificar** a instância
- ✅ O método apenas **lê dados** ou **calcula valores**
- ✅ A struct é **pequena** (poucos campos)
- ✅ Você quer **garantir imutabilidade**
- ✅ O método é **puramente funcional** (sem efeitos colaterais)

**Exemplos:**
```go
// Métodos de leitura/consulta
func (u Usuario) ObterNome() string {
    return u.Nome
}

// Métodos de cálculo
func (c Calculadora) Somar(numero int) int {
    return c.Valor + numero
}

// Métodos de formatação
func (u Usuario) String() string {
    return fmt.Sprintf("Usuario: %s", u.Nome)
}
```

### Use Pointer Receiver quando:
- ✅ O método **precisa modificar** a instância
- ✅ A struct é **grande** (muitos campos ou campos grandes)
- ✅ Você quer **evitar cópias** desnecessárias (performance)
- ✅ O método altera o **estado interno** do objeto
- ✅ Você quer **consistência** (se um método usa pointer, todos devem usar)

**Exemplos:**
```go
// Métodos que modificam estado
func (u *Usuario) AtualizarIdade() {
    u.Idade++
}

// Métodos de atualização
func (u *Usuario) AlterarEmail(novoEmail string) {
    u.Email = novoEmail
}

// Métodos de inicialização
func (u *Usuario) Inicializar(nome, email string) {
    u.Nome = nome
    u.Email = email
}
```

## Convenções e Boas Práticas

### 1. Consistência
Se você tem métodos que modificam a struct, use pointer receiver para **todos** os métodos daquele tipo, mesmo os que não modificam. Isso mantém a API consistente.

```go
// ✅ BOM - Consistente
func (u *Usuario) AtualizarIdade() { ... }
func (u *Usuario) ObterNome() string { ... }  // Mesmo sendo leitura, usa pointer

// ❌ EVITAR - Inconsistente
func (u Usuario) ObterNome() string { ... }      // Value receiver
func (u *Usuario) AtualizarIdade() { ... }      // Pointer receiver
```

### 2. Métodos em Tipos Primitivos
Você pode definir métodos em tipos primitivos também, não apenas em structs.

```go
type MeuInt int

func (m MeuInt) Dobrar() MeuInt {
    return m * 2
}

// Uso:
numero := MeuInt(5)
dobro := numero.Dobrar()  // 10
```

### 3. Métodos com Múltiplos Parâmetros e Retornos
Métodos podem ter múltiplos parâmetros e retornos, assim como funções.

```go
func (u *Usuario) Atualizar(nome, email string) (bool, error) {
    if nome == "" {
        return false, errors.New("nome não pode ser vazio")
    }
    u.Nome = nome
    u.Email = email
    return true, nil
}
```

## Métodos vs Funções

**Quando usar métodos:**
- Quando a operação está **intimamente relacionada** ao tipo
- Quando você quer **encapsular** comportamento com dados
- Quando você quer criar uma **API orientada a objetos**

**Quando usar funções:**
- Quando a operação é **genérica** e não pertence a um tipo específico
- Quando você precisa de **flexibilidade** (funções podem ser passadas como valores)
- Quando a operação não precisa de **estado** de um objeto

**Exemplo comparativo:**
```go
// Método - pertence ao tipo Usuario
func (u Usuario) ValidarEmail() bool {
    return strings.Contains(u.Email, "@")
}

// Função - independente, pode validar qualquer string
func ValidarEmail(email string) bool {
    return strings.Contains(email, "@")
}

// Uso:
usuario := Usuario{Email: "teste@example.com"}
usuario.ValidarEmail()  // Método

ValidarEmail("teste@example.com")  // Função
```

## Exemplo Completo

```go
package metodos

import "fmt"

type Usuario struct {
    Nome  string
    Email string
    Senha string
    Idade int
}

// Value receiver - apenas leitura/formatação
func (u Usuario) Salvar() {
    fmt.Println("Salvando usuario:", u.Nome)
    fmt.Println("Email:", u.Email)
    fmt.Println("Senha:", u.Senha)
}

// Pointer receiver - modifica a instância
func (u *Usuario) AtualizarIdade() {
    u.Idade++
}

// Value receiver - retorna informação
func (u Usuario) ObterNome() string {
    return u.Nome
}

// Pointer receiver - modifica múltiplos campos
func (u *Usuario) AtualizarPerfil(nome, email string) {
    u.Nome = nome
    u.Email = email
}

// Uso:
func Exemplo() {
    usuario := Usuario{
        Nome:  "Mike",
        Email: "mike@example.com",
        Senha: "123456",
        Idade: 20,
    }
    
    usuario.Salvar()              // Value receiver
    fmt.Println(usuario.ObterNome()) // Value receiver
    
    usuario.AtualizarIdade()      // Pointer receiver - modifica
    fmt.Println("Idade:", usuario.Idade) // 21
    
    usuario.AtualizarPerfil("João", "joao@example.com") // Pointer receiver
    fmt.Println(usuario.Nome)     // "João"
}
```

## Casos de Uso
- **Encapsulamento**: Agrupar dados e comportamentos relacionados
- **APIs orientadas a objetos**: Criar interfaces mais intuitivas e fáceis de usar
- **Modificação de estado**: Atualizar campos de structs de forma controlada
- **Cálculos baseados em dados**: Métodos que usam os dados da struct para calcular valores
- **Validação**: Validar dados da struct antes de salvar ou processar
- **Formatação**: Converter structs em strings ou outros formatos
- **Operações CRUD**: Criar, ler, atualizar e deletar operações em entidades

**Exemplo prático:**
```go
// Sistema de e-commerce
type Produto struct {
    Nome     string
    Preco    float64
    Estoque  int
}

// Métodos de consulta (value receiver)
func (p Produto) EstaDisponivel() bool {
    return p.Estoque > 0
}

func (p Produto) CalcularDesconto(percentual float64) float64 {
    return p.Preco * (percentual / 100)
}

// Métodos de modificação (pointer receiver)
func (p *Produto) AtualizarPreco(novoPreco float64) {
    p.Preco = novoPreco
}

func (p *Produto) Vender(quantidade int) error {
    if p.Estoque < quantidade {
        return errors.New("estoque insuficiente")
    }
    p.Estoque -= quantidade
    return nil
}

// Uso:
produto := Produto{Nome: "Notebook", Preco: 1500.0, Estoque: 10}
if produto.EstaDisponivel() {
    desconto := produto.CalcularDesconto(10) // 150.0
    produto.Vender(1) // Reduz estoque
}
```

