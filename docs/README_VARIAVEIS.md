# VARIÁVEIS

**EXPLÍCITO**: Declarar uma variável de forma explícita ajuda a identificar melhor o tipo da variável.

**Exemplo:** `var nome string = "Mike"`

**IMPLÍCITO**: Declarar de forma implícita deixa o código mais enxuto. O compilador consegue inferir o tipo usando como base o valor atribuído à variável.

**Exemplo:** `nome2 := "Mike"`

## Casos de Uso
- **Declaração explícita**: Quando você quer deixar o tipo claro ou quando o tipo não pode ser inferido
- **Declaração implícita**: Para código mais conciso quando o tipo é óbvio
- **Variáveis de escopo**: Declarar variáveis no escopo apropriado
- **Variáveis não inicializadas**: Declarar variáveis que serão atribuídas depois
- **Múltiplas variáveis**: Declarar várias variáveis de uma vez

**Exemplo prático:**
```go
// Explícito - útil quando tipo não é óbvio
var resultado float64
var erro error

// Implícito - código mais limpo
nome := "João"
idade := 30

// Múltiplas variáveis
var x, y int = 1, 2
var a, b = "hello", true

// Declaração curta em if
if idade := calcularIdade(); idade >= 18 {
    fmt.Println("Maior de idade")
}
```

