# FUNÇÕES

Funções em Go são blocos de código reutilizáveis que podem receber parâmetros e retornar valores. Go suporta múltiplos retornos, o que é muito útil para retornar valores e erros simultaneamente.

## Retorno Nomeado (Named Return Values)

Retorno nomeado é uma funcionalidade do Go que permite nomear os valores de retorno de uma função diretamente na assinatura. Isso torna o código mais legível e permite usar um `return` sem especificar os valores explicitamente.

**Sintaxe:**
```go
func nomeFuncao(parametros) (nomeRetorno1 tipo1, nomeRetorno2 tipo2) {
    // código
    nomeRetorno1 = valor1
    nomeRetorno2 = valor2
    return  // Retorna automaticamente nomeRetorno1 e nomeRetorno2
}
```

**Exemplo Básico:**
```go
func FuncaoRetornoNomeado(n1, n2 int) (soma int, subtracao int) {
    soma = n1 + n2
    subtracao = n1 - n2
    return  // Retorna soma e subtracao automaticamente
}

// Uso:
soma, subtracao := FuncaoRetornoNomeado(10, 5)
fmt.Println(soma)      // 15
fmt.Println(subtracao) // 5
```

**Vantagens:**
- **Legibilidade**: Os nomes dos retornos tornam a função auto-documentada
- **Código mais limpo**: Não precisa especificar valores no `return`
- **Valores zero**: As variáveis de retorno são inicializadas com valores zero automaticamente

**Exemplo com Valores Zero:**
```go
func Dividir(a, b float64) (resultado float64, erro error) {
    if b == 0 {
        erro = errors.New("divisão por zero")
        return  // resultado já é 0.0 (valor zero de float64)
    }
    resultado = a / b
    return
}
```

**Comparação: Retorno Nomeado vs Retorno Anônimo**

**Retorno Anônimo (tradicional):**
```go
func SomaSubtracao(n1, n2 int) (int, int) {
    soma := n1 + n2
    subtracao := n1 - n2
    return soma, subtracao  // Precisa especificar os valores
}
```

**Retorno Nomeado:**
```go
func SomaSubtracao(n1, n2 int) (soma int, subtracao int) {
    soma = n1 + n2
    subtracao = n1 - n2
    return  // Retorna soma e subtracao automaticamente
}
```

**OBSERVAÇÕES IMPORTANTES:**
- As variáveis de retorno nomeadas são inicializadas com seus valores zero
- Você pode usar `return` sem valores, e os valores atuais das variáveis nomeadas serão retornados
- Você também pode especificar valores explicitamente: `return valor1, valor2`
- Retornos nomeados podem ser úteis, mas também podem tornar o código menos claro se usados incorretamente

**Exemplo Avançado - Múltiplos Retornos:**
```go
func ProcessarUsuario(id int) (usuario *Usuario, erro error) {
    if id <= 0 {
        erro = errors.New("ID inválido")
        return  // usuario é nil, erro tem valor
    }
    
    usuario = &Usuario{
        ID:   id,
        Nome: "João",
    }
    return  // Retorna usuario e erro (nil)
}
```

**Exemplo com Defer:**
```go
func LerArquivo(nome string) (conteudo []byte, erro error) {
    arquivo, erro := os.Open(nome)
    if erro != nil {
        return  // Retorna nil e o erro
    }
    defer arquivo.Close()  // Garante fechamento
    
    conteudo, erro = ioutil.ReadAll(arquivo)
    return  // Retorna conteudo e erro (se houver)
}
```

### Casos de Uso - Retorno Nomeado
- **Funções que retornam múltiplos valores**: Quando você precisa retornar resultado e erro
- **Código mais legível**: Quando os nomes dos retornos tornam a função auto-explicativa
- **Valores padrão**: Quando você quer que valores zero sejam retornados automaticamente
- **Defer com retornos**: Útil quando você usa `defer` e precisa modificar valores de retorno

**Exemplo prático:**
```go
// Validar e processar dados
func ValidarProcessar(dados string) (resultado string, valido bool, erro error) {
    if dados == "" {
        erro = errors.New("dados vazios")
        return  // resultado="", valido=false, erro tem valor
    }
    
    resultado = strings.ToUpper(dados)
    valido = true
    return  // Retorna resultado processado, valido=true, erro=nil
}

// Calcular estatísticas
func CalcularEstatisticas(numeros []int) (media float64, soma int, quantidade int) {
    quantidade = len(numeros)
    if quantidade == 0 {
        return  // Todos retornam valores zero
    }
    
    for _, num := range numeros {
        soma += num
    }
    media = float64(soma) / float64(quantidade)
    return
}
```

## Função Variádica (Variadic Functions)

Funções variádicas são funções que podem receber um número variável de argumentos do mesmo tipo. Em Go, isso é feito usando `...` antes do tipo do último parâmetro. Os argumentos são automaticamente convertidos em um slice dentro da função.

**Sintaxe:**
```go
func nomeFuncao(parametrosFixos, valores ...tipo) tipoRetorno {
    // valores é um slice do tipo especificado
}
```

**Exemplo Básico:**
```go
func FuncaoVariatica(numeros ...int) int {
    total := 0
    for _, numero := range numeros {
        total += numero
    }
    return total
}

// Uso:
soma := FuncaoVariatica(1, 2, 3, 4, 5)
fmt.Println(soma) // 15

// Também pode receber zero argumentos:
soma = FuncaoVariatica()
fmt.Println(soma) // 0
```

**Passando um Slice para Função Variádica:**
Você pode passar um slice existente usando o operador `...`:

```go
numeros := []int{1, 2, 3, 4, 5}
soma := FuncaoVariatica(numeros...)  // Desempacota o slice
fmt.Println(soma) // 15
```

**Exemplo com Múltiplos Parâmetros:**
```go
func ImprimirFormato(formato string, valores ...interface{}) {
    fmt.Printf(formato, valores...)
}

// Uso:
ImprimirFormato("Nome: %s, Idade: %d\n", "João", 30)
// Saída: Nome: João, Idade: 30
```

**Exemplo - Função que Aceita Zero ou Mais Argumentos:**
```go
func Max(numeros ...int) int {
    if len(numeros) == 0 {
        return 0  // ou retornar erro
    }
    
    max := numeros[0]
    for _, num := range numeros {
        if num > max {
            max = num
        }
    }
    return max
}

// Uso:
maximo := Max(10, 5, 8, 20, 3)
fmt.Println(maximo) // 20
```

**Exemplo - Concatenar Strings:**
```go
func Concatenar(separador string, strings ...string) string {
    if len(strings) == 0 {
        return ""
    }
    
    resultado := strings[0]
    for _, s := range strings[1:] {
        resultado += separador + s
    }
    return resultado
}

// Uso:
texto := Concatenar(", ", "Go", "Python", "Java")
fmt.Println(texto) // "Go, Python, Java"
```

**OBSERVAÇÕES IMPORTANTES:**
- O parâmetro variádico deve ser o **último parâmetro** da função
- Você só pode ter **um parâmetro variádico** por função
- Dentro da função, o parâmetro variádico é tratado como um **slice**
- Se nenhum argumento for passado, o slice será vazio (não `nil`)
- Você pode passar zero ou mais argumentos

**Exemplo com Parâmetros Fixos e Variádicos:**
```go
func Processar(prefixo string, valores ...int) {
    for _, valor := range valores {
        fmt.Printf("%s: %d\n", prefixo, valor)
    }
}

// Uso:
Processar("Número", 1, 2, 3)
// Saída:
// Número: 1
// Número: 2
// Número: 3
```

**Exemplo - Função que Retorna Múltiplos Valores:**
```go
func Estatisticas(numeros ...float64) (media, soma, max float64) {
    if len(numeros) == 0 {
        return 0, 0, 0
    }
    
    soma = numeros[0]
    max = numeros[0]
    
    for _, num := range numeros[1:] {
        soma += num
        if num > max {
            max = num
        }
    }
    
    media = soma / float64(len(numeros))
    return
}

// Uso:
media, soma, max := Estatisticas(10.5, 20.3, 15.7, 8.2)
fmt.Printf("Média: %.2f, Soma: %.2f, Máximo: %.2f\n", media, soma, max)
```

### Casos de Uso - Funções Variádicas
- **Funções utilitárias**: Criar funções que aceitam um número variável de argumentos (ex: `fmt.Printf`, `append`)
- **Agregações**: Calcular somas, médias, máximos, mínimos de listas variáveis
- **Formatação**: Funções que formatam strings com múltiplos valores
- **Logging**: Funções de log que aceitam múltiplas mensagens
- **APIs flexíveis**: Criar APIs que podem receber diferentes quantidades de parâmetros
- **Concatenação**: Juntar múltiplas strings ou valores

**Exemplo prático:**
```go
// Função de log variádica
func Log(nivel string, mensagens ...string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    for _, msg := range mensagens {
        fmt.Printf("[%s] %s: %s\n", timestamp, nivel, msg)
    }
}

// Uso:
Log("INFO", "Sistema iniciado", "Conexão estabelecida")
Log("ERRO", "Falha na conexão")

// Função para criar queries SQL
func CriarQuery(tabela string, campos ...string) string {
    if len(campos) == 0 {
        return fmt.Sprintf("SELECT * FROM %s", tabela)
    }
    camposStr := strings.Join(campos, ", ")
    return fmt.Sprintf("SELECT %s FROM %s", camposStr, tabela)
}

// Uso:
query1 := CriarQuery("usuarios", "id", "nome", "email")
query2 := CriarQuery("produtos")  // SELECT * FROM produtos
```

## Função Anônima (Anonymous Function / Closure)

Funções anônimas são funções sem nome que podem ser definidas e usadas diretamente no código. Elas são muito úteis para criar closures (funções que capturam variáveis do escopo externo) e para passar funções como argumentos.

**Sintaxe Básica:**
```go
func() {
    // código da função
}()
```

**Exemplo Básico - Função Anônima Executada Imediatamente (IIFE):**
```go
func() {
    fmt.Println("Ola Mundo")
}()
// Saída: Ola Mundo
```

A função é definida e executada imediatamente. Os parênteses `()` no final executam a função assim que ela é definida.

### Atribuindo Função Anônima a uma Variável

Você pode atribuir uma função anônima a uma variável e chamá-la depois:

**Exemplo:**
```go
// Atribuir função anônima a uma variável
somar := func(a, b int) int {
    return a + b
}

// Chamar a função através da variável
resultado := somar(5, 3)
fmt.Println(resultado) // 8
```

**Exemplo com Tipo Explícito:**
```go
// Declarar variável com tipo de função
var multiplicar func(int, int) int

// Atribuir função anônima
multiplicar = func(a, b int) int {
    return a * b
}

// Usar
resultado := multiplicar(4, 5)
fmt.Println(resultado) // 20
```

### Função Anônima com Parâmetros

Funções anônimas podem receber parâmetros normalmente:

**Exemplo:**
```go
// Função anônima com parâmetros
cumprimentar := func(nome string) {
    fmt.Printf("Olá, %s!\n", nome)
}

cumprimentar("João")  // Olá, João!
cumprimentar("Maria") // Olá, Maria!
```

### Função Anônima com Retorno

Funções anônimas também podem retornar valores:

**Exemplo:**
```go
// Função anônima que retorna valor
calcular := func(x, y int) int {
    return x*x + y*y
}

resultado := calcular(3, 4)
fmt.Println(resultado) // 25 (3*3 + 4*4)
```

### Closure (Captura de Variáveis)

Uma das características mais poderosas das funções anônimas é a capacidade de capturar variáveis do escopo externo, criando um **closure**:

**Exemplo - Closure Simples:**
```go
contador := 0

// Função anônima que captura a variável contador
incrementar := func() {
    contador++
    fmt.Println("Contador:", contador)
}

incrementar() // Contador: 1
incrementar() // Contador: 2
incrementar() // Contador: 3
```

**Exemplo - Closure com Parâmetros Externos:**
```go
func criarMultiplicador(fator int) func(int) int {
    // Retorna uma função anônima que captura 'fator'
    return func(numero int) int {
        return numero * fator
    }
}

// Criar funções multiplicadoras
dobrar := criarMultiplicador(2)
triplicar := criarMultiplicador(3)

fmt.Println(dobrar(5))    // 10 (5 * 2)
fmt.Println(triplicar(5))  // 15 (5 * 3)
```

**Explicação:**
- A função `criarMultiplicador` retorna uma função anônima
- A função anônima **captura** a variável `fator` do escopo externo
- Cada chamada a `criarMultiplicador` cria um novo closure com seu próprio `fator`

### Passando Função Anônima como Argumento

Funções anônimas são frequentemente usadas como argumentos para outras funções:

**Exemplo - Função que Aceita Função como Parâmetro:**
```go
// Função que aceita uma função como parâmetro
func processar(numeros []int, operacao func(int) int) []int {
    resultado := make([]int, len(numeros))
    for i, num := range numeros {
        resultado[i] = operacao(num)
    }
    return resultado
}

// Usar função anônima como argumento
numeros := []int{1, 2, 3, 4, 5}

// Passar função anônima que dobra cada número
dobrados := processar(numeros, func(x int) int {
    return x * 2
})
fmt.Println(dobrados) // [2 4 6 8 10]

// Passar função anônima que eleva ao quadrado
quadrados := processar(numeros, func(x int) int {
    return x * x
})
fmt.Println(quadrados) // [1 4 9 16 25]
```

### Função Anônima em Estruturas de Controle

Funções anônimas são muito usadas com `defer`, `goroutines`, e em loops:

**Exemplo com Defer:**
```go
func exemploDefer() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado de:", r)
        }
    }()
    
    // código que pode causar panic
    panic("algo deu errado")
}
```

**Exemplo com Goroutine:**
```go
// Executar função anônima em uma goroutine
go func() {
    fmt.Println("Executando em goroutine")
    time.Sleep(1 * time.Second)
    fmt.Println("Goroutine finalizada")
}()

fmt.Println("Código principal continua")
```

**Exemplo em Loop:**
```go
// Criar múltiplas funções em um loop
funcoes := make([]func(), 3)

for i := 0; i < 3; i++ {
    // Captura o valor atual de i
    i := i  // Importante: criar nova variável para cada iteração
    funcoes[i] = func() {
        fmt.Printf("Função %d\n", i)
    }
}

// Executar todas as funções
for _, f := range funcoes {
    f()
}
// Saída:
// Função 0
// Função 1
// Função 2
```

**IMPORTANTE - Problema de Captura em Loops:**
```go
// ❌ PROBLEMA - Todas as funções capturam o mesmo 'i'
funcoes := make([]func(), 3)
for i := 0; i < 3; i++ {
    funcoes[i] = func() {
        fmt.Println(i)  // Todas imprimem 3!
    }
}

// ✅ SOLUÇÃO - Criar nova variável em cada iteração
funcoes := make([]func(), 3)
for i := 0; i < 3; i++ {
    i := i  // Nova variável para cada iteração
    funcoes[i] = func() {
        fmt.Println(i)  // Cada uma imprime seu próprio valor
    }
}
```

### Comparação: Função Anônima vs Função Nomeada

**Função Nomeada:**
```go
func somar(a, b int) int {
    return a + b
}

resultado := somar(5, 3)
```

**Função Anônima:**
```go
somar := func(a, b int) int {
    return a + b
}

resultado := somar(5, 3)
```

**Quando usar cada uma:**
- **Função nomeada**: Quando a função será reutilizada em múltiplos lugares
- **Função anônima**: Quando a função é usada apenas uma vez ou como argumento para outra função

### Casos de Uso - Funções Anônimas
- **Callbacks**: Passar funções como argumentos para outras funções
- **Closures**: Capturar e modificar variáveis do escopo externo
- **Goroutines**: Executar código de forma assíncrona
- **Defer com recuperação**: Tratamento de erros e panics
- **Transformações de dados**: Aplicar operações a coleções (map, filter, reduce)
- **Event handlers**: Definir comportamentos para eventos
- **Configuração dinâmica**: Criar funções com comportamento configurável

## Closure (Função de Fechamento)

Closure é uma função que captura e "fecha sobre" variáveis do escopo externo, permitindo que a função acesse essas variáveis mesmo depois que o escopo externo tenha terminado. Em Go, closures são criadas usando funções anônimas.

**Definição:**
Um closure é uma função que tem acesso a variáveis de um escopo externo (enclosing scope), mesmo após o escopo externo ter sido encerrado.

### Como Funciona um Closure

Quando uma função anônima captura uma variável do escopo externo, Go cria um closure que mantém uma referência a essa variável. Isso permite que a função acesse e modifique a variável mesmo depois que a função externa retornou.

**Exemplo Básico:**
```go
func FuncaoClosure() func() {
    texto := "Dentro da funcao closure"
    var funcao = func() {
        fmt.Println(texto)  // Captura a variável 'texto'
    }
    return funcao  // Retorna a função que capturou 'texto'
}

// Uso:
novaFuncao := FuncaoClosure()  // 'texto' foi capturado
novaFuncao()  // Imprime: "Dentro da funcao closure"
// Mesmo que FuncaoClosure tenha terminado, 'texto' ainda está acessível
```

**Explicação:**
1. `FuncaoClosure()` cria uma variável local `texto`
2. Cria uma função anônima que referencia `texto`
3. Retorna essa função anônima
4. Quando a função é chamada depois, ela ainda tem acesso a `texto`, mesmo que `FuncaoClosure()` já tenha terminado

### Closure Simples

**Exemplo - Contador:**
```go
contador := 0

// Função anônima que captura a variável contador
incrementar := func() {
    contador++  // Modifica a variável capturada
    fmt.Println("Contador:", contador)
}

incrementar() // Contador: 1
incrementar() // Contador: 2
incrementar() // Contador: 3
```

### Closure com Parâmetros

**Exemplo - Criar Funções com Comportamento Específico:**
```go
func criarMultiplicador(fator int) func(int) int {
    // Retorna uma função anônima que captura 'fator'
    return func(numero int) int {
        return numero * fator  // Usa 'fator' capturado
    }
}

// Criar funções multiplicadoras
dobrar := criarMultiplicador(2)
triplicar := criarMultiplicador(3)
quadruplicar := criarMultiplicador(4)

fmt.Println(dobrar(5))       // 10 (5 * 2)
fmt.Println(triplicar(5))   // 15 (5 * 3)
fmt.Println(quadruplicar(5)) // 20 (5 * 4)
```

**Explicação:**
- Cada chamada a `criarMultiplicador` cria um novo closure
- Cada closure captura seu próprio valor de `fator`
- As funções retornadas são independentes e mantêm seus próprios valores

### Closure com Múltiplas Variáveis

**Exemplo:**
```go
func criarCalculadora(soma, multiplicacao int) func(int) int {
    return func(x int) int {
        return (x + soma) * multiplicacao  // Captura ambas as variáveis
    }
}

calc := criarCalculadora(10, 2)
fmt.Println(calc(5))  // (5 + 10) * 2 = 30
```

### Closure Modificando Variáveis Capturadas

Closures podem modificar variáveis do escopo externo:

**Exemplo:**
```go
func criarContador() func() int {
    contador := 0  // Variável capturada
    
    return func() int {
        contador++  // Modifica a variável capturada
        return contador
    }
}

contador1 := criarContador()
contador2 := criarContador()

fmt.Println(contador1()) // 1
fmt.Println(contador1()) // 2
fmt.Println(contador2()) // 1 (closure independente)
fmt.Println(contador1()) // 3
fmt.Println(contador2()) // 2
```

**Importante:** Cada closure mantém sua própria cópia das variáveis capturadas. `contador1` e `contador2` são independentes.

### Closure em Loops - Problema Comum

**❌ PROBLEMA - Todas as closures capturam o mesmo valor:**
```go
var funcoes []func()

for i := 0; i < 3; i++ {
    funcoes = append(funcoes, func() {
        fmt.Println(i)  // Todas imprimem 3!
    })
}

for _, f := range funcoes {
    f()
}
// Saída:
// 3
// 3
// 3
```

**Por que acontece?**
- Todas as funções anônimas capturam a mesma variável `i`
- Quando as funções são executadas, `i` já vale 3 (valor final do loop)

**✅ SOLUÇÃO 1 - Criar nova variável em cada iteração:**
```go
var funcoes []func()

for i := 0; i < 3; i++ {
    i := i  // Nova variável para cada iteração
    funcoes = append(funcoes, func() {
        fmt.Println(i)  // Cada uma captura seu próprio i
    })
}

for _, f := range funcoes {
    f()
}
// Saída:
// 0
// 1
// 2
```

**✅ SOLUÇÃO 2 - Passar como parâmetro:**
```go
var funcoes []func()

for i := 0; i < 3; i++ {
    func(valor int) {  // Função que recebe valor como parâmetro
        funcoes = append(funcoes, func() {
            fmt.Println(valor)  // Captura 'valor', não 'i'
        })
    }(i)  // Passa i como argumento
}

for _, f := range funcoes {
    f()
}
// Saída:
// 0
// 1
// 2
```

### Closure com Structs

Closures podem capturar structs e modificar seus campos:

**Exemplo:**
```go
type Contador struct {
    valor int
}

func criarContadorStruct() func() int {
    c := Contador{valor: 0}
    
    return func() int {
        c.valor++
        return c.valor
    }
}

contador := criarContadorStruct()
fmt.Println(contador()) // 1
fmt.Println(contador()) // 2
fmt.Println(contador()) // 3
```

### Closure com Slices e Maps

Closures podem capturar slices e maps, e modificações afetam o original:

**Exemplo:**
```go
func criarAdicionador(lista []int) func(int) {
    return func(valor int) {
        lista = append(lista, valor)  // Modifica o slice capturado
        fmt.Println("Lista:", lista)
    }
}

numeros := []int{1, 2, 3}
adicionar := criarAdicionador(numeros)

adicionar(4)  // Lista: [1 2 3 4]
adicionar(5)  // Lista: [1 2 3 4 5]
```

**Nota:** Modificações em slices e maps capturados afetam o original porque Go passa slices e maps por referência.

### Closure com Ponteiros

Closures podem capturar ponteiros:

**Exemplo:**
```go
func criarIncrementador(ptr *int) func() {
    return func() {
        (*ptr)++  // Modifica o valor apontado
        fmt.Println("Valor:", *ptr)
    }
}

valor := 10
incrementar := criarIncrementador(&valor)

incrementar() // Valor: 11
incrementar() // Valor: 12
fmt.Println("Valor original:", valor) // Valor original: 12
```

### Casos de Uso Práticos de Closures

#### 1. **Factory Functions (Funções Fábrica)**

Criar funções com comportamento configurável:

```go
func criarLogger(prefixo string) func(string) {
    return func(mensagem string) {
        fmt.Printf("[%s] %s\n", prefixo, mensagem)
    }
}

infoLog := criarLogger("INFO")
errorLog := criarLogger("ERROR")
debugLog := criarLogger("DEBUG")

infoLog("Sistema iniciado")   // [INFO] Sistema iniciado
errorLog("Erro encontrado")   // [ERROR] Erro encontrado
debugLog("Debug info")        // [DEBUG] Debug info
```

#### 2. **Middleware e Decorators**

Criar middlewares que envolvem funções:

```go
func criarMiddleware(handler func()) func() {
    return func() {
        fmt.Println("Antes do handler")
        handler()
        fmt.Println("Depois do handler")
    }
}

handler := criarMiddleware(func() {
    fmt.Println("Processando...")
})

handler()
// Saída:
// Antes do handler
// Processando...
// Depois do handler
```

#### 3. **Callbacks com Estado**

Manter estado entre chamadas de callback:

```go
func criarCallbackComEstado() func(string) {
    chamadas := 0
    
    return func(mensagem string) {
        chamadas++
        fmt.Printf("Chamada %d: %s\n", chamadas, mensagem)
    }
}

callback := criarCallbackComEstado()
callback("Primeira")  // Chamada 1: Primeira
callback("Segunda")   // Chamada 2: Segunda
callback("Terceira")  // Chamada 3: Terceira
```

#### 4. **Configuração de Funções**

Criar funções com configuração pré-definida:

```go
func criarProcessador(delimiter string, uppercase bool) func(string) string {
    return func(texto string) string {
        resultado := strings.Split(texto, delimiter)
        if uppercase {
            return strings.ToUpper(strings.Join(resultado, " "))
        }
        return strings.Join(resultado, " ")
    }
}

processar1 := criarProcessador(",", false)
processar2 := criarProcessador(" ", true)

fmt.Println(processar1("a,b,c"))     // a b c
fmt.Println(processar2("hello world")) // HELLO WORLD
```

#### 5. **Gerenciamento de Recursos**

Criar funções que gerenciam recursos:

```go
func criarGerenciadorRecurso() (func(), func()) {
    recurso := "recurso adquirido"
    
    usar := func() {
        fmt.Println("Usando:", recurso)
    }
    
    liberar := func() {
        fmt.Println("Liberando:", recurso)
        recurso = "recurso liberado"
    }
    
    return usar, liberar
}

usar, liberar := criarGerenciadorRecurso()
usar()    // Usando: recurso adquirido
liberar() // Liberando: recurso adquirido
```

#### 6. **Cache e Memoização**

Implementar cache usando closures:

```go
func criarCache() func(string) (string, bool) {
    cache := make(map[string]string)
    
    return func(chave string) (string, bool) {
        if valor, existe := cache[chave]; existe {
            return valor, true
        }
        return "", false
    }
}

cache := criarCache()
cache("chave1") // Busca no cache interno
```

### Vantagens dos Closures

1. **Encapsulamento**: Permite encapsular dados e comportamento
2. **Estado Privado**: Variáveis capturadas não são acessíveis externamente
3. **Flexibilidade**: Cria funções com comportamento configurável
4. **Reutilização**: Permite criar múltiplas instâncias independentes

### Observações Importantes

1. **Cada closure é independente**: Cada chamada cria um novo closure com suas próprias variáveis
2. **Variáveis são capturadas por referência**: Modificações afetam o original (para slices, maps, ponteiros)
3. **Cuidado em loops**: Sempre crie nova variável ou passe como parâmetro
4. **Performance**: Closures têm um pequeno overhead, mas geralmente desprezível
5. **Memória**: Variáveis capturadas permanecem na memória enquanto o closure existir

### Resumo

**Closure é:**
- Uma função que captura variáveis do escopo externo
- Permite acesso a essas variáveis mesmo após o escopo externo terminar
- Útil para criar funções com comportamento configurável
- Permite encapsulamento e estado privado

**Quando usar:**
- Quando você precisa de funções com estado
- Para criar factory functions
- Para implementar callbacks com estado
- Para encapsular dados e comportamento
- Para criar funções configuráveis

**Exemplo prático:**
```go
// Sistema de processamento de pagamentos
func processarPagamentos(pagamentos []float64, aplicarDesconto func(float64) float64) []float64 {
    resultado := make([]float64, len(pagamentos))
    for i, valor := range pagamentos {
        resultado[i] = aplicarDesconto(valor)
    }
    return resultado
}

// Usar função anônima para aplicar desconto de 10%
valores := []float64{100.0, 200.0, 300.0}
comDesconto := processarPagamentos(valores, func(valor float64) float64 {
    return valor * 0.9  // 10% de desconto
})
fmt.Println(comDesconto) // [90 180 270]

// Middleware para HTTP (exemplo conceitual)
func criarMiddleware(handler func()) func() {
    return func() {
        fmt.Println("Antes do handler")
        handler()
        fmt.Println("Depois do handler")
    }
}

// Usar
handler := criarMiddleware(func() {
    fmt.Println("Processando requisição")
})

handler()
// Saída:
// Antes do handler
// Processando requisição
// Depois do handler
```

## Defer

A palavra-chave `defer` em Go adia a execução de uma função até que a função que a contém retorne. É muito útil para garantir que recursos sejam liberados, arquivos sejam fechados, ou operações de limpeza sejam executadas, independentemente de como a função termina (normalmente ou com panic).

**Sintaxe:**
```go
defer funcao()
```

**Características:**
- A função `defer` é executada **após** a função que a contém retornar
- Múltiplos `defer` são executados em ordem **LIFO** (Last In, First Out) - o último `defer` é executado primeiro
- Os argumentos da função `defer` são avaliados **imediatamente**, mas a execução é adiada

### Exemplo Básico

**Sem Defer:**
```go
func SemDefer() {
    fmt.Println("Funcao Sem Defer")
}

func exemplo() {
    fmt.Println("Início")
    SemDefer()
    fmt.Println("Fim")
}

// Saída:
// Início
// Funcao Sem Defer
// Fim
```

**Com Defer:**
```go
func Defer() {
    fmt.Println("Funcao com Defer")
}

func exemplo() {
    fmt.Println("Início")
    defer Defer()  // Executado no final
    fmt.Println("Fim")
}

// Saída:
// Início
// Fim
// Funcao com Defer  <- Executado por último
```

### Múltiplos Defer (Ordem LIFO)

Quando há múltiplos `defer`, eles são executados na ordem inversa (o último é executado primeiro):

**Exemplo:**
```go
func exemplo() {
    defer fmt.Println("Primeiro defer")
    defer fmt.Println("Segundo defer")
    defer fmt.Println("Terceiro defer")
    fmt.Println("Código normal")
}

// Saída:
// Código normal
// Terceiro defer    <- Último defer executado primeiro
// Segundo defer
// Primeiro defer    <- Primeiro defer executado por último
```

### Defer com Argumentos

Os argumentos de uma função `defer` são avaliados **imediatamente** quando a linha `defer` é executada, não quando a função `defer` é realmente chamada. Isso é crucial para entender o comportamento do `defer`.

**Exemplo Básico:**
```go
func exemplo() {
    i := 0
    defer fmt.Println("Valor de i:", i)  // i é avaliado AGORA (0)
    i++
    fmt.Println("i incrementado:", i)
}

// Saída:
// i incrementado: 1
// Valor de i: 0  <- Mostra 0, não 1, porque foi avaliado antes
```

**Explicação detalhada:**
1. Quando Go encontra `defer fmt.Println("Valor de i:", i)`, ele **avalia imediatamente** o valor de `i` (que é 0)
2. O valor `0` é armazenado para ser usado quando o `defer` for executado
3. Mesmo que `i` seja modificado depois, o valor já foi capturado

**Para capturar o valor atual (usando função anônima):**
```go
func exemplo() {
    i := 0
    defer func() {
        fmt.Println("Valor de i:", i)  // Captura i do escopo externo
    }()
    i++
    fmt.Println("i incrementado:", i)
}

// Saída:
// i incrementado: 1
// Valor de i: 1  <- Mostra 1, porque a função anônima captura o valor atual
```

**Por que funciona?**
- A função anônima não recebe `i` como argumento
- Ela **captura** a variável `i` do escopo externo (closure)
- Quando a função é executada (no defer), ela lê o valor atual de `i`

**Exemplo com Slice:**
```go
func exemploSlice() {
    numeros := []int{1, 2, 3}
    
    // ❌ PROBLEMA - captura o slice inteiro, mas mostra valores antigos
    defer fmt.Println("Slice:", numeros)  // Avalia a referência do slice
    
    numeros = append(numeros, 4, 5)  // Modifica o slice
    fmt.Println("Slice modificado:", numeros)
}

// Saída:
// Slice modificado: [1 2 3 4 5]
// Slice: [1 2 3 4 5]  <- Mostra valores atualizados (porque é a mesma referência)

// ✅ Para capturar valores específicos:
func exemploSliceCorreto() {
    numeros := []int{1, 2, 3}
    
    defer func() {
        fmt.Println("Tamanho original:", len(numeros))
    }()
    
    numeros = append(numeros, 4, 5)
    fmt.Println("Slice modificado:", numeros)
}
```

**Exemplo com Ponteiros:**
```go
func exemploPonteiro() {
    valor := 10
    ptr := &valor
    
    defer fmt.Println("Valor via ponteiro:", *ptr)  // Avalia o ponteiro AGORA
    
    *ptr = 20  // Modifica o valor apontado
    fmt.Println("Valor modificado:", valor)
}

// Saída:
// Valor modificado: 20
// Valor via ponteiro: 20  <- Mostra 20, porque o ponteiro aponta para o mesmo endereço
```

**Exemplo com Struct:**
```go
type Pessoa struct {
    Nome string
    Idade int
}

func exemploStruct() {
    p := Pessoa{Nome: "João", Idade: 30}
    
    // ❌ Captura a struct por valor (cópia)
    defer fmt.Println("Pessoa:", p)  // p é copiado AGORA
    
    p.Idade = 31  // Modifica a struct original
    fmt.Println("Pessoa modificada:", p)
}

// Saída:
// Pessoa modificada: {João 31}
// Pessoa: {João 30}  <- Mostra valores antigos (cópia)

// ✅ Para capturar valores atualizados:
func exemploStructCorreto() {
    p := Pessoa{Nome: "João", Idade: 30}
    
    defer func() {
        fmt.Println("Pessoa:", p)  // Captura a struct do escopo
    }()
    
    p.Idade = 31
    fmt.Println("Pessoa modificada:", p)
}

// Saída:
// Pessoa modificada: {João 31}
// Pessoa: {João 31}  <- Mostra valores atualizados
```

**Exemplo com Múltiplos Argumentos:**
```go
func exemploMultiplosArgs() {
    a, b := 1, 2
    
    defer fmt.Println("Soma:", a+b)  // a+b é calculado AGORA (3)
    
    a = 10
    b = 20
    fmt.Println("Valores modificados:", a, b)
}

// Saída:
// Valores modificados: 10 20
// Soma: 3  <- Mostra 3 (1+2), não 30 (10+20)
```

**Resumo - Quando valores são avaliados:**

| Tipo | Comportamento |
|------|---------------|
| **Valores primitivos** (int, string, bool) | Avaliados imediatamente, valor é copiado |
| **Ponteiros** | Ponteiro é avaliado, mas aponta para o mesmo endereço |
| **Slices/Maps** | Referência é avaliada, mas aponta para a mesma estrutura |
| **Structs** | Struct é copiada (por valor) |
| **Função anônima sem args** | Captura variáveis do escopo externo (closure) |

### Casos de Uso Comuns

#### 1. Fechar Arquivos

O uso mais comum de `defer` é garantir que arquivos sejam fechados:

```go
func lerArquivo(nome string) (string, error) {
    arquivo, err := os.Open(nome)
    if err != nil {
        return "", err
    }
    defer arquivo.Close()  // Sempre fecha o arquivo, mesmo se houver erro
    
    conteudo, err := ioutil.ReadAll(arquivo)
    if err != nil {
        return "", err
    }
    
    return string(conteudo), nil
}
```

#### 2. Liberar Recursos

```go
func processarRecurso() {
    recurso := adquirirRecurso()
    defer liberarRecurso(recurso)  // Sempre libera o recurso
    
    // Usar o recurso...
    usarRecurso(recurso)
    // Recurso será liberado automaticamente ao final
}
```

#### 3. Recuperação de Panic

`defer` é essencial para recuperar de panics:

```go
func exemploComRecover() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado de panic:", r)
        }
    }()
    
    // Código que pode causar panic
    panic("algo deu errado")
    fmt.Println("Este código não será executado")
}

// Saída: Recuperado de panic: algo deu errado
```

#### 4. Modificar Valores de Retorno

Com retornos nomeados, você pode modificar valores de retorno usando `defer`:

```go
func exemplo() (resultado int) {
    defer func() {
        resultado++  // Modifica o valor de retorno
    }()
    
    return 5
}

fmt.Println(exemplo()) // 6 (5 + 1)
```

#### 5. Medir Tempo de Execução

```go
func processar() {
    inicio := time.Now()
    defer func() {
        duracao := time.Since(inicio)
        fmt.Printf("Processamento levou: %v\n", duracao)
    }()
    
    // Código a ser medido
    time.Sleep(1 * time.Second)
}
```

#### 6. Logging e Debugging

```go
func funcaoComplexa() {
    defer fmt.Println("Função funcaoComplexa finalizada")
    fmt.Println("Iniciando processamento...")
    // Código complexo...
}
```

### Defer em Loops

**Cuidado ao usar `defer` em loops** - todos os `defer` serão acumulados e executados apenas quando a função retornar:

```go
// ❌ PROBLEMA - Todos os defer são acumulados
func exemploProblema() {
    for i := 0; i < 5; i++ {
        defer fmt.Println(i)  // Todos executam no final
    }
    fmt.Println("Fim do loop")
}

// Saída:
// Fim do loop
// 4
// 3
// 2
// 1
// 0
```

**Por que isso acontece?**
- Cada `defer` é adicionado a uma pilha
- Todos os `defer` só são executados quando a função retorna
- Se você tem 1000 iterações, terá 1000 `defer` acumulados

**✅ SOLUÇÃO 1 - Usar função auxiliar:**
```go
func exemploSolucao1() {
    for i := 0; i < 5; i++ {
        func(valor int) {
            defer fmt.Println("Processando:", valor)
            // Processar valor...
        }(i)
    }
    fmt.Println("Fim do loop")
}

// Saída:
// Processando: 0
// Processando: 1
// Processando: 2
// Processando: 3
// Processando: 4
// Fim do loop
```

**✅ SOLUÇÃO 2 - Executar defer imediatamente:**
```go
func exemploSolucao2() {
    for i := 0; i < 5; i++ {
        func() {
            arquivo, _ := os.Open(fmt.Sprintf("arquivo%d.txt", i))
            defer arquivo.Close()  // Fecha imediatamente após cada iteração
            // Processar arquivo...
        }()
    }
}
```

**✅ SOLUÇÃO 3 - Usar slice para armazenar recursos:**
```go
func exemploSolucao3() {
    var arquivos []*os.File
    
    // Abrir todos os arquivos
    for i := 0; i < 5; i++ {
        arquivo, _ := os.Open(fmt.Sprintf("arquivo%d.txt", i))
        arquivos = append(arquivos, arquivo)
    }
    
    // Fechar todos no final
    defer func() {
        for _, arquivo := range arquivos {
            arquivo.Close()
        }
    }()
    
    // Processar arquivos...
}
```

**Exemplo Real - Processar múltiplos arquivos:**
```go
// ❌ PROBLEMA - Acumula muitos defer
func processarArquivosRuim(nomes []string) error {
    for _, nome := range nomes {
        arquivo, err := os.Open(nome)
        if err != nil {
            return err
        }
        defer arquivo.Close()  // Acumula defer para cada arquivo
        
        // Processar arquivo...
    }
    return nil
}

// ✅ SOLUÇÃO - Fechar imediatamente
func processarArquivosBom(nomes []string) error {
    for _, nome := range nomes {
        func() error {
            arquivo, err := os.Open(nome)
            if err != nil {
                return err
            }
            defer arquivo.Close()  // Fecha após cada iteração
            
            // Processar arquivo...
            return nil
        }()
    }
    return nil
}
```

### Defer com Métodos

Você pode usar `defer` com métodos de structs:

```go
type Arquivo struct {
    nome string
    arquivo *os.File
}

func (a *Arquivo) Abrir() error {
    arquivo, err := os.Open(a.nome)
    if err != nil {
        return err
    }
    a.arquivo = arquivo
    return nil
}

func (a *Arquivo) Fechar() error {
    if a.arquivo != nil {
        return a.arquivo.Close()
    }
    return nil
}

func processarArquivo(nome string) error {
    a := &Arquivo{nome: nome}
    if err := a.Abrir(); err != nil {
        return err
    }
    defer a.Fechar()  // Chama o método Fechar
    
    // Processar arquivo...
    return nil
}
```

### Defer e Retornos Nomeados

Com retornos nomeados, você pode modificar valores de retorno no `defer`:

```go
// Exemplo 1 - Incrementar retorno
func exemplo() (resultado int) {
    defer func() {
        resultado++  // Modifica o valor de retorno
    }()
    return 5
}
fmt.Println(exemplo()) // 6

// Exemplo 2 - Adicionar ao retorno
func processar() (total int, erro error) {
    defer func() {
        if erro != nil {
            total = 0  // Zera total em caso de erro
        }
    }()
    
    total = 100
    // Se houver erro, total será 0
    return total, nil
}

// Exemplo 3 - Logging de retorno
func calcular() (resultado float64) {
    defer func() {
        fmt.Printf("Resultado calculado: %.2f\n", resultado)
    }()
    
    resultado = 3.14159
    return
}
```

### Defer e Panic/Recover

`defer` é essencial para recuperação de panics. O `recover()` só funciona dentro de uma função `defer`:

```go
// Exemplo básico
func exemploRecover() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado:", r)
        }
    }()
    
    panic("erro crítico")
}

// Exemplo avançado - Recuperar e continuar
func processarComRecover() (resultado int, erro error) {
    defer func() {
        if r := recover(); r != nil {
            erro = fmt.Errorf("panic recuperado: %v", r)
            resultado = 0
        }
    }()
    
    // Código que pode causar panic
    resultado = 10 / 0  // Causaria panic, mas é recuperado
    return
}

// Exemplo - Múltiplos recover
func exemploMultiplosRecover() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recover 1:", r)
        }
    }()
    
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recover 2:", r)
            panic(r)  // Re-lança o panic
        }
    }()
    
    panic("teste")
}
```

### Defer em Contextos Específicos

#### Defer em HTTP Handlers

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Garantir que o corpo da requisição seja fechado
    defer r.Body.Close()
    
    // Garantir que headers sejam escritos
    defer func() {
        if err := recover(); err != nil {
            http.Error(w, "Erro interno", http.StatusInternalServerError)
        }
    }()
    
    // Processar requisição...
}
```

#### Defer em Database Transactions

```go
func processarTransacao(db *sql.DB) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()  // Rollback em caso de panic
            panic(r)  // Re-lança o panic
        }
    }()
    
    // Executar operações...
    
    if err != nil {
        tx.Rollback()
        return err
    }
    
    defer tx.Commit()  // Commit no final
    return nil
}
```

#### Defer com Mutexes

```go
func exemploMutex() {
    var mu sync.Mutex
    
    mu.Lock()
    defer mu.Unlock()  // Sempre libera o lock
    
    // Código crítico...
    // Mesmo se houver panic, o lock será liberado
}
```

#### Defer com WaitGroups

```go
func exemploWaitGroup() {
    var wg sync.WaitGroup
    
    wg.Add(1)
    defer wg.Done()  // Sempre decrementa o WaitGroup
    
    // Processar...
}
```

### Armadilhas Comuns com Defer

#### 1. Defer em Loop (já mencionado)
```go
// ❌ ERRADO
for i := 0; i < 1000; i++ {
    defer fmt.Println(i)  // Acumula 1000 defer
}

// ✅ CORRETO
for i := 0; i < 1000; i++ {
    func(valor int) {
        defer fmt.Println(valor)
    }(i)
}
```

#### 2. Defer com Nil Check

```go
// ❌ ERRADO - Pode causar panic se arquivo for nil
func exemploErrado(arquivo *os.File) {
    defer arquivo.Close()  // Panic se arquivo for nil
}

// ✅ CORRETO
func exemploCorreto(arquivo *os.File) {
    if arquivo == nil {
        return
    }
    defer arquivo.Close()
}

// ✅ OU
func exemploCorreto2(arquivo *os.File) {
    defer func() {
        if arquivo != nil {
            arquivo.Close()
        }
    }()
}
```

#### 3. Defer e Return Values

```go
// ❌ CUIDADO - Defer pode modificar retornos nomeados
func exemplo() (resultado int) {
    defer func() {
        resultado = 999  // Modifica o retorno!
    }()
    return 5
}
fmt.Println(exemplo()) // 999, não 5!

// ✅ Para evitar modificação acidental:
func exemploSeguro() int {
    resultado := 5
    defer func() {
        // Não pode modificar resultado (não é retorno nomeado)
        fmt.Println("Resultado:", resultado)
    }()
    return resultado
}
```

#### 4. Defer com Erro

```go
// ❌ ERRADO - Ignora erro do Close
func exemploErrado() {
    arquivo, _ := os.Open("arquivo.txt")
    defer arquivo.Close()  // Erro do Close é ignorado
}

// ✅ CORRETO - Tratar erro do Close
func exemploCorreto() error {
    arquivo, err := os.Open("arquivo.txt")
    if err != nil {
        return err
    }
    defer func() {
        if err := arquivo.Close(); err != nil {
            log.Printf("Erro ao fechar arquivo: %v", err)
        }
    }()
    return nil
}
```

### Performance e Defer

**Custo do Defer:**
- `defer` tem um custo pequeno, mas mensurável
- Em loops muito rápidos (milhões de iterações), pode impactar performance
- Para a maioria dos casos, o custo é desprezível

**Benchmark exemplo:**
```go
// Sem defer (mais rápido)
func semDefer() {
    arquivo, _ := os.Open("arquivo.txt")
    // processar...
    arquivo.Close()
}

// Com defer (ligeiramente mais lento, mas mais seguro)
func comDefer() {
    arquivo, _ := os.Open("arquivo.txt")
    defer arquivo.Close()
    // processar...
}
```

**Quando evitar defer:**
- Loops muito rápidos (milhões de iterações)
- Código crítico de performance (hot paths)
- Quando você precisa de controle preciso do momento de execução

**Quando sempre usar defer:**
- Fechar arquivos
- Liberar recursos (locks, conexões)
- Recuperação de panics
- Cleanup de recursos

### Observações Importantes

1. **Performance**: `defer` tem um pequeno custo de performance, mas geralmente é desprezível. Evite em loops muito rápidos.
2. **Ordem de execução**: Sempre LIFO (Last In, First Out) - o último `defer` é executado primeiro.
3. **Argumentos avaliados imediatamente**: Os argumentos são avaliados quando o `defer` é encontrado, não quando é executado.
4. **Retornos nomeados podem ser modificados**: `defer` pode modificar valores de retorno nomeados.
5. **Recover só funciona em defer**: `recover()` só funciona dentro de uma função `defer`.
6. **Defer sempre executa**: Mesmo com `return` antecipado ou `panic`, o `defer` será executado.
7. **Não use em loops críticos**: Se você precisa de performance máxima, evite `defer` em loops muito rápidos.

### Casos de Uso - Defer
- **Gerenciamento de recursos**: Garantir que arquivos, conexões, locks sejam liberados
- **Tratamento de erros**: Recuperar de panics e garantir limpeza mesmo em caso de erro
- **Logging**: Registrar início/fim de operações
- **Medição de performance**: Medir tempo de execução de funções
- **Modificação de retornos**: Ajustar valores de retorno antes de sair da função
- **Cleanup**: Executar código de limpeza independente do caminho de execução

**Exemplo prático:**
```go
// Processar requisição HTTP com múltiplos recursos
func processarRequisicao() error {
    // Abrir conexão com banco
    db, err := abrirConexao()
    if err != nil {
        return err
    }
    defer db.Close()  // Sempre fecha a conexão
    
    // Abrir arquivo de log
    logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        return err
    }
    defer logFile.Close()  // Sempre fecha o arquivo
    
    // Adquirir lock
    lock.Lock()
    defer lock.Unlock()  // Sempre libera o lock
    
    // Processar requisição...
    return processar()
}
```

## Função Recursiva

Uma função recursiva é uma função que chama a si mesma. A recursão é uma técnica poderosa para resolver problemas que podem ser divididos em subproblemas menores e similares.

**Sintaxe:**
```go
func nomeFuncao(parametros) tipoRetorno {
    // Caso base (condição de parada)
    if condicao {
        return valor
    }
    
    // Chamada recursiva
    return nomeFuncao(parametrosModificados)
}
```

**Componentes essenciais:**
1. **Caso base**: Condição que para a recursão (evita loop infinito)
2. **Caso recursivo**: Chamada à própria função com parâmetros modificados

### Exemplo Básico - Sequência de Fibonacci

**Exemplo do código:**
```go
func FuncaoRecursiva(posicao int) int {
    if posicao <= 1 {
        return posicao  // Caso base
    }
    return FuncaoRecursiva(posicao - 2) + FuncaoRecursiva(posicao - 1)  // Recursão
}

// Uso:
resultado := FuncaoRecursiva(15)
fmt.Println(resultado) // 610 (15º número de Fibonacci)
```

**Explicação:**
- **Caso base**: Se `posicao <= 1`, retorna o próprio valor
- **Caso recursivo**: Soma os dois números anteriores da sequência
- A função chama a si mesma com valores menores até chegar ao caso base

### Exemplo - Fatorial

```go
func fatorial(n int) int {
    // Caso base
    if n <= 1 {
        return 1
    }
    
    // Caso recursivo
    return n * fatorial(n-1)
}

// Uso:
fmt.Println(fatorial(5)) // 120 (5 * 4 * 3 * 2 * 1)
```

**Rastreamento:**
```
fatorial(5)
  = 5 * fatorial(4)
  = 5 * (4 * fatorial(3))
  = 5 * (4 * (3 * fatorial(2)))
  = 5 * (4 * (3 * (2 * fatorial(1))))
  = 5 * (4 * (3 * (2 * 1)))
  = 120
```

### Exemplo - Soma de Array

```go
func somarArray(arr []int, indice int) int {
    // Caso base: fim do array
    if indice >= len(arr) {
        return 0
    }
    
    // Caso recursivo: soma atual + resto do array
    return arr[indice] + somarArray(arr, indice+1)
}

// Uso:
numeros := []int{1, 2, 3, 4, 5}
soma := somarArray(numeros, 0)
fmt.Println(soma) // 15
```

### Exemplo - Busca Binária Recursiva

```go
func buscaBinaria(arr []int, alvo, inicio, fim int) int {
    // Caso base: não encontrado
    if inicio > fim {
        return -1
    }
    
    meio := (inicio + fim) / 2
    
    // Caso base: encontrado
    if arr[meio] == alvo {
        return meio
    }
    
    // Caso recursivo: buscar na metade apropriada
    if arr[meio] > alvo {
        return buscaBinaria(arr, alvo, inicio, meio-1)
    }
    return buscaBinaria(arr, alvo, meio+1, fim)
}

// Uso:
arr := []int{1, 3, 5, 7, 9, 11, 13}
indice := buscaBinaria(arr, 7, 0, len(arr)-1)
fmt.Println(indice) // 3
```

### Exemplo - Potência

```go
func potencia(base, expoente int) int {
    // Caso base
    if expoente == 0 {
        return 1
    }
    
    // Caso recursivo
    return base * potencia(base, expoente-1)
}

// Uso:
fmt.Println(potencia(2, 8)) // 256
```

### Recursão vs Iteração

**Recursão:**
```go
func fatorialRecursivo(n int) int {
    if n <= 1 {
        return 1
    }
    return n * fatorialRecursivo(n-1)
}
```

**Iteração (loop):**
```go
func fatorialIterativo(n int) int {
    resultado := 1
    for i := 2; i <= n; i++ {
        resultado *= i
    }
    return resultado
}
```

**Quando usar cada uma:**
- **Recursão**: Quando o problema é naturalmente recursivo (árvores, grafos, divisão e conquista)
- **Iteração**: Quando a solução iterativa é mais simples e eficiente

### Recursão de Cauda (Tail Recursion)

Recursão de cauda é quando a chamada recursiva é a última operação da função. Go não otimiza recursão de cauda automaticamente, mas é uma boa prática:

```go
// Recursão de cauda
func fatorialTail(n, acumulador int) int {
    if n <= 1 {
        return acumulador
    }
    return fatorialTail(n-1, n*acumulador)
}

// Chamada inicial
func fatorial(n int) int {
    return fatorialTail(n, 1)
}
```

### Limitações e Cuidados

1. **Stack Overflow**: Recursão muito profunda pode causar estouro de pilha
2. **Performance**: Recursão pode ser mais lenta que iteração devido ao overhead de chamadas
3. **Memória**: Cada chamada recursiva usa espaço na pilha

**Exemplo de problema:**
```go
// ❌ Pode causar stack overflow para valores grandes
func fibonacciRecursivo(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacciRecursivo(n-1) + fibonacciRecursivo(n-2)
}

// ✅ Solução iterativa (mais eficiente)
func fibonacciIterativo(n int) int {
    if n <= 1 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
```

### Casos de Uso - Funções Recursivas
- **Algoritmos de divisão e conquista**: Merge sort, Quick sort
- **Estruturas de dados recursivas**: Árvores, grafos, listas encadeadas
- **Problemas matemáticos**: Fatorial, Fibonacci, potência, GCD
- **Busca e travessia**: Busca binária, travessia de árvores
- **Parsing**: Analisar estruturas aninhadas (JSON, XML, expressões)
- **Backtracking**: Resolver problemas como N-Queens, Sudoku
- **Problemas de combinatória**: Permutações, combinações

**Exemplo prático:**
```go
// Calcular profundidade de uma árvore binária
type No struct {
    Valor int
    Esquerda *No
    Direita *No
}

func profundidade(no *No) int {
    if no == nil {
        return 0  // Caso base
    }
    
    profEsquerda := profundidade(no.Esquerda)
    profDireita := profundidade(no.Direita)
    
    if profEsquerda > profDireita {
        return profEsquerda + 1
    }
    return profDireita + 1
}

// Contar nós em uma árvore
func contarNos(no *No) int {
    if no == nil {
        return 0  // Caso base
    }
    return 1 + contarNos(no.Esquerda) + contarNos(no.Direita)
}
```

## Panic e Recover

`panic` e `recover` são mecanismos de tratamento de erros excepcionais em Go. Eles são usados para lidar com situações que não podem ser tratadas normalmente e que requerem interrupção imediata do fluxo de execução.

### O que é Panic?

`panic` é uma função built-in que interrompe o fluxo normal de execução do programa. Quando `panic` é chamado, a função atual para de executar, todos os `defer` são executados, e então o `panic` se propaga para a função chamadora até encontrar um `recover` ou até o programa terminar.

**Sintaxe:**
```go
panic(valor interface{})
```

**Exemplo básico:**
```go
func exemploPanic() {
    panic("algo deu muito errado")
    fmt.Println("Este código nunca será executado")
}
```

### O que é Recover?

`recover` é uma função built-in que permite recuperar o controle de um programa que entrou em `panic`. `recover` **só funciona dentro de uma função `defer`**.

**Sintaxe:**
```go
recover() interface{}
```

**Exemplo básico:**
```go
func recuperarExecucao() {
    if r := recover(); r != nil {
        fmt.Println("Recuperado de panic:", r)
    }
}

func FuncaoPanic(n1, n2 int8) {
    defer recuperarExecucao()
    if media := (n1 + n2) / 2; media < 6 {
        panic("Media menor que 6")
    }
    fmt.Println("Media calculada. Resultado será retornado")
}
```

## Quando Usar Panic?

### ✅ Use Panic Quando:

#### 1. **Erros de Programação (Bugs)**
Use `panic` para indicar erros que são bugs no código, não condições de erro esperadas:

```go
// ✅ BOM - Erro de programação
func dividir(a, b int) int {
    if b == 0 {
        panic("divisão por zero: bug no código - deveria ter verificado antes")
    }
    return a / b
}

// ✅ BOM - Condição impossível
func processar(usuario *Usuario) {
    if usuario == nil {
        panic("usuario não pode ser nil - bug no código")
    }
    // processar...
}
```

#### 2. **Erros Irrecuperáveis**
Quando o programa não pode continuar de forma segura:

```go
// ✅ BOM - Erro crítico que impede continuação
func inicializarSistema() {
    db, err := conectarBanco()
    if err != nil {
        panic(fmt.Sprintf("Não é possível continuar sem banco de dados: %v", err))
    }
    // Sistema não pode funcionar sem banco
}
```

#### 3. **Asserções de Invariantes**
Para verificar condições que devem sempre ser verdadeiras:

```go
// ✅ BOM - Verificar invariante
func processarLista(lista []int) {
    if len(lista) == 0 {
        panic("lista não pode estar vazia - invariante violado")
    }
    // processar...
}
```

#### 4. **Erros em Bibliotecas Internas**
Em bibliotecas, quando você quer forçar o usuário a tratar o erro:

```go
// ✅ BOM - Em biblioteca
func ParseConfig(caminho string) *Config {
    arquivo, err := os.Open(caminho)
    if err != nil {
        panic(fmt.Sprintf("config inválido: %v", err))
    }
    defer arquivo.Close()
    // parse...
}
```

### ❌ NÃO Use Panic Quando:

#### 1. **Erros Esperados e Tratáveis**
Use retorno de erro (`error`) em vez de `panic`:

```go
// ❌ ERRADO - Erro esperado
func dividir(a, b int) int {
    if b == 0 {
        panic("divisão por zero")  // Erro esperado, não use panic
    }
    return a / b
}

// ✅ CORRETO - Retornar erro
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}
```

#### 2. **Validação de Entrada do Usuário**
Validações de entrada devem retornar erros, não causar panic:

```go
// ❌ ERRADO
func validarEmail(email string) {
    if !strings.Contains(email, "@") {
        panic("email inválido")  // Entrada do usuário, não use panic
    }
}

// ✅ CORRETO
func validarEmail(email string) error {
    if !strings.Contains(email, "@") {
        return fmt.Errorf("email inválido")
    }
    return nil
}
```

#### 3. **Erros de I/O**
Erros de arquivo, rede, etc. devem ser retornados como `error`:

```go
// ❌ ERRADO
func lerArquivo(nome string) string {
    conteudo, err := ioutil.ReadFile(nome)
    if err != nil {
        panic(err)  // Erro de I/O, não use panic
    }
    return string(conteudo)
}

// ✅ CORRETO
func lerArquivo(nome string) (string, error) {
    conteudo, err := ioutil.ReadFile(nome)
    if err != nil {
        return "", err
    }
    return string(conteudo), nil
}
```

#### 4. **Erros de Negócio**
Regras de negócio devem retornar erros:

```go
// ❌ ERRADO
func sacar(saldo, valor float64) float64 {
    if valor > saldo {
        panic("saldo insuficiente")  // Regra de negócio, não use panic
    }
    return saldo - valor
}

// ✅ CORRETO
func sacar(saldo, valor float64) (float64, error) {
    if valor > saldo {
        return saldo, fmt.Errorf("saldo insuficiente")
    }
    return saldo - valor, nil
}
```

## Quando Usar Recover?

### ✅ Use Recover Quando:

#### 1. **Recuperar de Panics de Bibliotecas de Terceiros**
Quando você usa bibliotecas que podem causar panic:

```go
func processarComSeguranca() (resultado string, erro error) {
    defer func() {
        if r := recover(); r != nil {
            erro = fmt.Errorf("erro na biblioteca: %v", r)
            resultado = ""
        }
    }()
    
    // Usar biblioteca que pode causar panic
    resultado = bibliotecaTerceiros.Processar()
    return
}
```

#### 2. **Garantir Limpeza de Recursos**
Garantir que recursos sejam liberados mesmo em caso de panic:

```go
func processarComLimpeza() {
    recurso := adquirirRecurso()
    defer func() {
        liberarRecurso(recurso)
        if r := recover(); r != nil {
            log.Printf("Panic durante processamento: %v", r)
            // Recurso já foi liberado
        }
    }()
    
    // Processar (pode causar panic)
    processarRecurso(recurso)
}
```

#### 3. **Em Servidores e Aplicações Longas**
Para evitar que um panic encerre todo o servidor:

```go
func handlerHTTP(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
            log.Printf("Panic no handler: %v", r)
        }
    }()
    
    // Processar requisição (pode causar panic)
    processarRequisicao(w, r)
}
```

#### 4. **Em Goroutines**
Para evitar que um panic em uma goroutine encerre todo o programa:

```go
func processarEmGoroutine() {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Panic na goroutine: %v", r)
                // Goroutine termina, mas programa continua
            }
        }()
        
        // Código que pode causar panic
        processar()
    }()
}
```

#### 5. **Em Testes**
Para testar comportamento em caso de panic:

```go
func TestPanic(t *testing.T) {
    defer func() {
        if r := recover(); r != nil {
            // Teste passou - panic foi causado como esperado
            if r != "erro esperado" {
                t.Errorf("panic inesperado: %v", r)
            }
        } else {
            t.Error("esperava panic, mas não ocorreu")
        }
    }()
    
    funcaoQueDeveCausarPanic()
}
```

### ❌ NÃO Use Recover Quando:

#### 1. **Para Mascarar Erros**
Não use `recover` para esconder problemas:

```go
// ❌ ERRADO - Mascara o problema
func processar() {
    defer func() {
        if r := recover(); r != nil {
            // Ignora o erro silenciosamente
        }
    }()
    // Código problemático...
}

// ✅ CORRETO - Trata o erro apropriadamente
func processar() error {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Erro recuperado: %v", r)
            // Toma ação apropriada
        }
    }()
    // Código...
}
```

#### 2. **Como Substituição de Tratamento de Erros**
Não use `recover` como alternativa a retornar erros:

```go
// ❌ ERRADO
func dividir(a, b int) int {
    defer func() {
        recover()  // Não faça isso!
    }()
    return a / b  // Deveria retornar erro se b == 0
}

// ✅ CORRETO
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}
```

## Como Usar Recover Corretamente

### Padrão Básico

```go
func exemploBasico() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado:", r)
        }
    }()
    
    panic("erro")
}
```

### Recover e Retorno de Erro

```go
func processar() (resultado string, erro error) {
    defer func() {
        if r := recover(); r != nil {
            erro = fmt.Errorf("panic recuperado: %v", r)
            resultado = ""
        }
    }()
    
    // Código que pode causar panic
    resultado = processarDados()
    return
}
```

### Recover e Logging

```go
func processarComLog() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Panic recuperado: %v\nStack trace: %s", r, debug.Stack())
            // Toma ação apropriada
        }
    }()
    
    // Código...
}
```

### Recover e Re-lançar Panic

Às vezes você quer fazer cleanup e depois re-lançar o panic:

```go
func processarComCleanup() {
    defer func() {
        if r := recover(); r != nil {
            // Fazer cleanup
            limparRecursos()
            // Re-lançar panic para que chamador saiba
            panic(r)
        }
    }()
    
    // Código...
}
```

### Recover em Múltiplos Níveis

```go
func nivel1() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recover nível 1:", r)
            // Não re-lança, trata aqui
        }
    }()
    
    nivel2()
}

func nivel2() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recover nível 2:", r)
            panic(r)  // Re-lança para nível 1
        }
    }()
    
    panic("erro")
}
```

## Exemplos Práticos

### Exemplo 1: Validação com Panic (Biblioteca Interna)

```go
// Em uma biblioteca interna, você pode usar panic para erros de programação
func ParseConfig(caminho string) *Config {
    arquivo, err := os.Open(caminho)
    if err != nil {
        panic(fmt.Sprintf("config inválido (bug no código): %v", err))
    }
    defer arquivo.Close()
    
    // Parse...
    return config
}

// Usuário da biblioteca usa recover
func inicializar() (config *Config, erro error) {
    defer func() {
        if r := recover(); r != nil {
            erro = fmt.Errorf("erro ao carregar config: %v", r)
        }
    }()
    
    config = ParseConfig("config.json")
    return
}
```

### Exemplo 2: Servidor HTTP

```go
func handler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            http.Error(w, "Erro interno", http.StatusInternalServerError)
            log.Printf("Panic no handler %s: %v\nStack: %s", 
                r.URL.Path, r, debug.Stack())
        }
    }()
    
    // Processar requisição
    processarRequisicao(w, r)
}
```

### Exemplo 3: Worker Pool

```go
func worker(jobs <-chan Job, results chan<- Result) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Panic no worker: %v", r)
            // Worker termina, mas pool continua
        }
    }()
    
    for job := range jobs {
        result := processarJob(job)
        results <- result
    }
}
```

### Exemplo 4: Transação de Banco de Dados

```go
func executarTransacao(db *sql.DB, operacoes []Operacao) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
            log.Printf("Panic na transação: %v", r)
            panic(r)  // Re-lança após rollback
        }
    }()
    
    for _, op := range operacoes {
        if err := executarOperacao(tx, op); err != nil {
            tx.Rollback()
            return err
        }
    }
    
    return tx.Commit()
}
```

## Boas Práticas

### 1. **Panic é para Erros de Programação, não para Erros Esperados**

```go
// ✅ BOM
func processar(usuario *Usuario) {
    if usuario == nil {
        panic("usuario nil - bug no código")
    }
}

// ❌ RUIM
func validarEmail(email string) {
    if !strings.Contains(email, "@") {
        panic("email inválido")  // Erro esperado, use return error
    }
}
```

### 2. **Sempre Use Defer com Recover**

```go
// ✅ BOM
defer func() {
    if r := recover(); r != nil {
        // tratar
    }
}()

// ❌ RUIM - recover não funciona fora de defer
if r := recover(); r != nil {  // Não funciona!
    // tratar
}
```

### 3. **Documente Quando Funções Podem Causar Panic**

```go
// Processar pode causar panic se dados estiverem em formato inválido
// (erro de programação, não erro de I/O)
func Processar(dados []byte) Result {
    // ...
}
```

### 4. **Use Recover Apenas Onde Faz Sentido**

Não coloque `recover` em todo lugar. Use apenas onde:
- Você precisa garantir limpeza de recursos
- Você está em um ponto de entrada (handler HTTP, main, goroutine)
- Você está lidando com código de terceiros que pode causar panic

### 5. **Não Mascare Panics**

```go
// ❌ RUIM
defer func() {
    recover()  // Ignora silenciosamente
}()

// ✅ BOM
defer func() {
    if r := recover(); r != nil {
        log.Printf("Erro: %v", r)
        // Toma ação apropriada
    }
}()
```

## Resumo: Quando Usar Cada Um

| Situação | Use |
|----------|-----|
| Erro de programação (bug) | `panic` |
| Erro irrecuperável | `panic` |
| Erro esperado e tratável | `return error` |
| Validação de entrada | `return error` |
| Erro de I/O | `return error` |
| Regra de negócio | `return error` |
| Recuperar de panic de terceiros | `recover` |
| Garantir limpeza de recursos | `recover` |
| Em servidores/aplicações | `recover` |
| Em goroutines | `recover` |
| Mascarar erros | ❌ Não use `recover` |
| Substituir tratamento de erros | ❌ Não use `recover` |

**Regra de Ouro:**
- **Panic**: Use para erros de **programação** que indicam bugs
- **Error**: Use para erros **esperados** que podem ser tratados
- **Recover**: Use para **recuperar** de panics e garantir **limpeza** de recursos

## Função Receiver (Métodos)

Em Go, uma função receiver é uma função especial que pertence a um tipo específico (geralmente uma struct). Essas funções são chamadas de **métodos** e permitem associar comportamentos a tipos, criando uma forma de programação orientada a objetos.

**Sintaxe:**
```go
func (receiver Tipo) NomeDaFuncao(parametros) tipoRetorno {
    // corpo da função
}
```

**Componentes:**
- `receiver` - A instância do tipo que recebe a função (pode ser por valor ou ponteiro)
- `Tipo` - O tipo ao qual a função pertence (geralmente uma struct)
- `NomeDaFuncao` - Nome da função
- `parametros` - Parâmetros de entrada (opcional)
- `tipoRetorno` - Tipo de retorno (opcional)

### Diferença entre Função Normal e Função Receiver

**Função Normal (independente):**
```go
func Somar(a, b int) int {
    return a + b
}

// Chamada:
resultado := Somar(5, 3)
```

**Função Receiver (método):**
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

### Value Receiver vs Pointer Receiver

Existem dois tipos de receivers em Go: **value receiver** e **pointer receiver**. A escolha entre eles afeta o comportamento da função.

#### Value Receiver (Recebedor por Valor)

Um value receiver recebe uma **cópia** da instância. Modificações feitas na função não afetam a instância original.

**Sintaxe:**
```go
func (u Usuario) NomeDaFuncao() {
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

#### Pointer Receiver (Recebedor por Ponteiro)

Um pointer receiver recebe um **ponteiro** para a instância. Modificações feitas na função afetam a instância original.

**Sintaxe:**
```go
func (u *Usuario) NomeDaFuncao() {
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

### Comparação Prática

**Exemplo completo mostrando a diferença:**
```go
type Usuario struct {
    Nome  string
    Idade int
}

// Value receiver - NÃO modifica a instância original
func (u Usuario) IncrementarIdade() {
    u.Idade++  // Modifica apenas a cópia
    fmt.Println("Idade dentro da função:", u.Idade)
}

// Pointer receiver - MODIFICA a instância original
func (u *Usuario) AtualizarIdade() {
    u.Idade++  // Modifica a instância original
    fmt.Println("Idade dentro da função:", u.Idade)
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
Idade dentro da função: 21
Depois de IncrementarIdade: 20
Idade dentro da função: 21
Depois de AtualizarIdade: 21
```

### Quando Usar Cada Tipo de Receiver

#### Use Value Receiver quando:
- ✅ A função **não precisa modificar** a instância
- ✅ A função apenas **lê dados** ou **calcula valores**
- ✅ A struct é **pequena** (poucos campos)
- ✅ Você quer **garantir imutabilidade**
- ✅ A função é **puramente funcional** (sem efeitos colaterais)

**Exemplos:**
```go
// Funções de leitura/consulta
func (u Usuario) ObterNome() string {
    return u.Nome
}

// Funções de cálculo
func (c Calculadora) Somar(numero int) int {
    return c.Valor + numero
}

// Funções de formatação
func (u Usuario) String() string {
    return fmt.Sprintf("Usuario: %s", u.Nome)
}
```

#### Use Pointer Receiver quando:
- ✅ A função **precisa modificar** a instância
- ✅ A struct é **grande** (muitos campos ou campos grandes)
- ✅ Você quer **evitar cópias** desnecessárias (performance)
- ✅ A função altera o **estado interno** do objeto
- ✅ Você quer **consistência** (se uma função usa pointer, todas devem usar)

**Exemplos:**
```go
// Funções que modificam estado
func (u *Usuario) AtualizarIdade() {
    u.Idade++
}

// Funções de atualização
func (u *Usuario) AlterarEmail(novoEmail string) {
    u.Email = novoEmail
}

// Funções de inicialização
func (u *Usuario) Inicializar(nome, email string) {
    u.Nome = nome
    u.Email = email
}
```

### Convenções e Boas Práticas

#### 1. Consistência
Se você tem funções que modificam a struct, use pointer receiver para **todas** as funções daquele tipo, mesmo as que não modificam. Isso mantém a API consistente.

```go
// ✅ BOM - Consistente
func (u *Usuario) AtualizarIdade() { ... }
func (u *Usuario) ObterNome() string { ... }  // Mesmo sendo leitura, usa pointer

// ❌ EVITAR - Inconsistente
func (u Usuario) ObterNome() string { ... }      // Value receiver
func (u *Usuario) AtualizarIdade() { ... }      // Pointer receiver
```

#### 2. Funções Receiver em Tipos Primitivos
Você pode definir funções receiver em tipos primitivos também, não apenas em structs.

```go
type MeuInt int

func (m MeuInt) Dobrar() MeuInt {
    return m * 2
}

// Uso:
numero := MeuInt(5)
dobro := numero.Dobrar()  // 10
```

#### 3. Funções Receiver com Múltiplos Parâmetros e Retornos
Funções receiver podem ter múltiplos parâmetros e retornos, assim como funções normais.

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

### Funções Receiver vs Funções Normais

**Quando usar funções receiver:**
- Quando a operação está **intimamente relacionada** ao tipo
- Quando você quer **encapsular** comportamento com dados
- Quando você quer criar uma **API orientada a objetos**

**Quando usar funções normais:**
- Quando a operação é **genérica** e não pertence a um tipo específico
- Quando você precisa de **flexibilidade** (funções podem ser passadas como valores)
- Quando a operação não precisa de **estado** de um objeto

**Exemplo comparativo:**
```go
// Função receiver - pertence ao tipo Usuario
func (u Usuario) ValidarEmail() bool {
    return strings.Contains(u.Email, "@")
}

// Função normal - independente, pode validar qualquer string
func ValidarEmail(email string) bool {
    return strings.Contains(email, "@")
}

// Uso:
usuario := Usuario{Email: "teste@example.com"}
usuario.ValidarEmail()  // Função receiver

ValidarEmail("teste@example.com")  // Função normal
```

### Exemplo Completo

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

### Casos de Uso - Funções Receiver
- **Encapsulamento**: Agrupar dados e comportamentos relacionados
- **APIs orientadas a objetos**: Criar interfaces mais intuitivas e fáceis de usar
- **Modificação de estado**: Atualizar campos de structs de forma controlada
- **Cálculos baseados em dados**: Funções que usam os dados da struct para calcular valores
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

// Funções receiver de consulta (value receiver)
func (p Produto) EstaDisponivel() bool {
    return p.Estoque > 0
}

func (p Produto) CalcularDesconto(percentual float64) float64 {
    return p.Preco * (percentual / 100)
}

// Funções receiver de modificação (pointer receiver)
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

### Observações Importantes

1. **Go automaticamente converte entre value e pointer**: Se você tem um value receiver, pode chamar com um ponteiro, e vice-versa. Go faz a conversão automaticamente.

```go
type Usuario struct {
    Nome string
}

func (u Usuario) ObterNome() string {
    return u.Nome
}

// Ambos funcionam:
usuario := Usuario{Nome: "Mike"}
usuario.ObterNome()  // ✅ Funciona

ptr := &usuario
ptr.ObterNome()      // ✅ Também funciona (Go converte automaticamente)
```

2. **Escolha baseada em comportamento, não apenas performance**: Embora pointer receivers sejam mais eficientes para structs grandes, a escolha principal deve ser baseada em se você precisa modificar o estado ou não.

3. **Métodos são funções especiais**: Em Go, métodos são tecnicamente funções com um receiver. A sintaxe especial `(receiver Tipo)` é apenas uma forma conveniente de associar funções a tipos.

## Função Init

A função `init` é uma função especial em Go que é executada automaticamente antes da função `main()`. Cada pacote pode ter uma ou múltiplas funções `init`, e elas são executadas automaticamente quando o pacote é importado.

**Sintaxe:**
```go
func init() {
    // código de inicialização
}
```

**Características:**
- Não recebe parâmetros
- Não retorna valores
- É executada automaticamente (não pode ser chamada manualmente)
- Pode haver múltiplas funções `init` no mesmo pacote
- Ordem de execução: variáveis globais → `init()` → `main()`

### Exemplo Básico

```go
package main

import "fmt"

var global = inicializarGlobal()

func inicializarGlobal() string {
    fmt.Println("1. Inicializando variável global")
    return "valor global"
}

func init() {
    fmt.Println("2. Função init executada")
}

func main() {
    fmt.Println("3. Função main executada")
    fmt.Println("Global:", global)
}

// Saída:
// 1. Inicializando variável global
// 2. Função init executada
// 3. Função main executada
// Global: valor global
```

### Múltiplas Funções Init

Você pode ter múltiplas funções `init` no mesmo pacote. Elas são executadas na ordem em que aparecem no código:

```go
package main

import "fmt"

func init() {
    fmt.Println("Init 1")
}

func init() {
    fmt.Println("Init 2")
}

func init() {
    fmt.Println("Init 3")
}

func main() {
    fmt.Println("Main")
}

// Saída:
// Init 1
// Init 2
// Init 3
// Main
```

### Ordem de Execução

A ordem de execução em Go é:

1. **Variáveis globais** (inicializadas)
2. **Funções `init()`** (todas, na ordem)
3. **Função `main()`**

```go
package main

import "fmt"

var a = func() int {
    fmt.Println("1. Variável a")
    return 1
}()

var b = func() int {
    fmt.Println("2. Variável b")
    return 2
}()

func init() {
    fmt.Println("3. Init 1")
}

func init() {
    fmt.Println("4. Init 2")
}

func main() {
    fmt.Println("5. Main")
    fmt.Println("a =", a, "b =", b)
}

// Saída:
// 1. Variável a
// 2. Variável b
// 3. Init 1
// 4. Init 2
// 5. Main
// a = 1 b = 2
```

## Quando Usar Init?

### ✅ Use Init Quando:

#### 1. **Inicialização de Variáveis Globais Complexas**

Quando você precisa de lógica complexa para inicializar variáveis globais:

```go
package config

var DatabaseURL string
var APIKey string

func init() {
    // Lógica complexa de inicialização
    env := os.Getenv("ENV")
    if env == "production" {
        DatabaseURL = os.Getenv("PROD_DB_URL")
        APIKey = os.Getenv("PROD_API_KEY")
    } else {
        DatabaseURL = os.Getenv("DEV_DB_URL")
        APIKey = os.Getenv("DEV_API_KEY")
    }
}
```

#### 2. **Registro de Componentes**

Registrar handlers, plugins, ou componentes:

```go
package handlers

var handlers = make(map[string]Handler)

func init() {
    handlers["/home"] = HomeHandler
    handlers["/about"] = AboutHandler
    handlers["/contact"] = ContactHandler
}

func GetHandler(path string) Handler {
    return handlers[path]
}
```

#### 3. **Validação de Configuração**

Validar configurações ao carregar o pacote:

```go
package config

var Port int
var Host string

func init() {
    Port = 8080
    Host = "localhost"
    
    // Validar configuração
    if Port < 1 || Port > 65535 {
        panic("porta inválida")
    }
    if Host == "" {
        panic("host não pode ser vazio")
    }
}
```

#### 4. **Inicialização de Bibliotecas de Terceiros**

Configurar bibliotecas que precisam de inicialização:

```go
package database

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func init() {
    var err error
    DB, err = gorm.Open("postgres", "connection string")
    if err != nil {
        panic("falha ao conectar ao banco: " + err.Error())
    }
}
```

#### 5. **Registro de Rotas ou Handlers**

Em frameworks web, registrar rotas:

```go
package routes

import "net/http"

func init() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)
}
```

#### 6. **Inicialização de Caches ou Pools**

Criar pools de conexões, caches, etc:

```go
package cache

var cache = make(map[string]interface{})

func init() {
    // Pré-carregar cache
    cache["config"] = loadConfig()
    cache["metadata"] = loadMetadata()
}
```

#### 7. **Registro de Drivers**

Registrar drivers de banco de dados:

```go
package main

import (
    _ "github.com/lib/pq"  // Driver PostgreSQL
    _ "github.com/go-sql-driver/mysql"  // Driver MySQL
)

// Os drivers se registram automaticamente via init()
```

### ❌ NÃO Use Init Quando:

#### 1. **Lógica que Deve Ser Testável**

Se você precisa testar a inicialização, use uma função normal:

```go
// ❌ ERRADO - Difícil de testar
var db *sql.DB

func init() {
    db, _ = sql.Open("postgres", "connection")
}

// ✅ CORRETO - Testável
func InicializarDB(connectionString string) (*sql.DB, error) {
    return sql.Open("postgres", connectionString)
}
```

#### 2. **Inicialização que Pode Falhar**

Se a inicialização pode falhar, use uma função que retorna erro:

```go
// ❌ ERRADO - Panic em caso de erro
func init() {
    db, err := sql.Open("postgres", "connection")
    if err != nil {
        panic(err)  // Não é ideal
    }
}

// ✅ CORRETO - Retorna erro
func InicializarDB() (*sql.DB, error) {
    return sql.Open("postgres", "connection")
}
```

#### 3. **Inicialização que Precisa de Parâmetros**

Se você precisa passar parâmetros, use uma função normal:

```go
// ❌ ERRADO - Não pode passar parâmetros
func init() {
    // Como passar configuração?
}

// ✅ CORRETO - Aceita parâmetros
func Inicializar(config Config) error {
    // Usar config
    return nil
}
```

#### 4. **Lógica de Negócio**

Não coloque lógica de negócio em `init`:

```go
// ❌ ERRADO
func init() {
    processarPedidos()  // Lógica de negócio não deve estar em init
}

// ✅ CORRETO
func main() {
    processarPedidos()  // Lógica de negócio em main ou função específica
}
```

## Init em Múltiplos Pacotes

Quando você importa múltiplos pacotes, as funções `init` de cada pacote são executadas na ordem de importação:

```go
package main

import (
    "fmt"
    "pacote1"  // init de pacote1 executa primeiro
    "pacote2"  // init de pacote2 executa depois
)

func init() {
    fmt.Println("Init do main")
}

func main() {
    fmt.Println("Main")
}

// Ordem de execução:
// 1. init de pacote1
// 2. init de pacote2
// 3. init do main
// 4. main
```

**Exemplo prático:**
```go
// pacote1/init.go
package pacote1

import "fmt"

func init() {
    fmt.Println("Pacote1: init")
}

// pacote2/init.go
package pacote2

import "fmt"

func init() {
    fmt.Println("Pacote2: init")
}

// main.go
package main

import (
    "fmt"
    _ "pacote1"
    _ "pacote2"
)

func init() {
    fmt.Println("Main: init")
}

func main() {
    fmt.Println("Main: main")
}

// Saída:
// Pacote1: init
// Pacote2: init
// Main: init
// Main: main
```

## Init e Imports com Efeitos Colaterais

Às vezes você importa um pacote apenas para executar seu `init` (efeito colateral). Use `_` para indicar isso:

```go
package main

import (
    "database/sql"
    _ "github.com/lib/pq"  // Import apenas para executar init do driver
)

func main() {
    // Driver já foi registrado via init
    db, _ := sql.Open("postgres", "connection")
    // ...
}
```

## Exemplos Práticos

### Exemplo 1: Configuração de Aplicação

```go
package config

import (
    "os"
    "strconv"
)

var (
    Port     int
    Host     string
    Debug    bool
    Database string
)

func init() {
    // Carregar configurações do ambiente
    Port = getEnvInt("PORT", 8080)
    Host = getEnv("HOST", "localhost")
    Debug = getEnvBool("DEBUG", false)
    Database = getEnv("DATABASE", "postgres://localhost/db")
    
    // Validar
    if Port < 1 {
        panic("PORT deve ser maior que 0")
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
    if value := os.Getenv(key); value != "" {
        if boolValue, err := strconv.ParseBool(value); err == nil {
            return boolValue
        }
    }
    return defaultValue
}
```

### Exemplo 2: Registro de Handlers

```go
package handlers

import "net/http"

type Handler func(http.ResponseWriter, *http.Request)

var routes = make(map[string]Handler)

func init() {
    // Registrar todas as rotas
    routes["/"] = homeHandler
    routes["/about"] = aboutHandler
    routes["/contact"] = contactHandler
    routes["/api/users"] = usersHandler
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Home"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("About"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Contact"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Users"))
}

func GetHandler(path string) Handler {
    return routes[path]
}
```

### Exemplo 3: Inicialização de Logger

```go
package logger

import (
    "log"
    "os"
)

var Logger *log.Logger

func init() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Falha ao abrir arquivo de log:", err)
    }
    
    Logger = log.New(file, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
}
```

### Exemplo 4: Pré-carregamento de Dados

```go
package cache

var (
    ConfigCache    map[string]interface{}
    MetadataCache  map[string]interface{}
)

func init() {
    // Pré-carregar caches
    ConfigCache = make(map[string]interface{})
    MetadataCache = make(map[string]interface{})
    
    // Carregar dados iniciais
    loadConfig()
    loadMetadata()
}

func loadConfig() {
    // Carregar configurações...
    ConfigCache["app_name"] = "Minha App"
    ConfigCache["version"] = "1.0.0"
}

func loadMetadata() {
    // Carregar metadados...
    MetadataCache["author"] = "João"
    MetadataCache["license"] = "MIT"
}
```

## Boas Práticas

### 1. **Mantenha Init Simples**

`init` deve ser usado apenas para inicialização simples. Lógica complexa deve estar em funções normais:

```go
// ✅ BOM
func init() {
    config = loadConfig()
}

// ❌ EVITAR - Muito complexo
func init() {
    // 100 linhas de código complexo...
}
```

### 2. **Evite Panic em Init**

Se possível, evite `panic` em `init`. Se necessário, documente claramente:

```go
// ✅ BOM - Documentado
// Init pode causar panic se configuração estiver inválida
func init() {
    if !validarConfig() {
        panic("configuração inválida")
    }
}
```

### 3. **Use Init para Efeitos Colaterais Esperados**

Quando o propósito do import é executar `init`, use `_`:

```go
import (
    _ "github.com/lib/pq"  // Import apenas para init
)
```

### 4. **Não Dependa da Ordem de Init**

A ordem de execução de `init` em diferentes pacotes pode mudar. Evite dependências entre `init` de pacotes diferentes:

```go
// ❌ EVITAR - Dependência entre pacotes
// pacote1/init.go
func init() {
    pacote2.Variavel = "valor"  // Depende de pacote2
}

// ✅ MELHOR - Inicialização independente
func init() {
    // Inicialização independente
}
```

### 5. **Teste Inicializações**

Se possível, torne a inicialização testável:

```go
// ✅ BOM - Testável
var db *sql.DB

func InicializarDB(connectionString string) error {
    var err error
    db, err = sql.Open("postgres", connectionString)
    return err
}

func init() {
    if err := InicializarDB(getConnectionString()); err != nil {
        panic(err)
    }
}
```

## Resumo: Quando Usar Init

| Situação | Use Init? |
|----------|-----------|
| Inicialização de variáveis globais complexas | ✅ Sim |
| Registro de componentes/handlers | ✅ Sim |
| Validação de configuração ao carregar | ✅ Sim |
| Inicialização de bibliotecas de terceiros | ✅ Sim |
| Registro de rotas | ✅ Sim |
| Inicialização de caches/pools | ✅ Sim |
| Lógica que precisa ser testável | ❌ Não |
| Inicialização que pode falhar | ⚠️ Cuidado |
| Inicialização que precisa de parâmetros | ❌ Não |
| Lógica de negócio | ❌ Não |

**Regra de Ouro:**
- **Init**: Use para inicialização **automática** e **simples** que deve acontecer quando o pacote é importado
- **Funções normais**: Use para inicialização que precisa de **parâmetros**, pode **falhar**, ou precisa ser **testável**

