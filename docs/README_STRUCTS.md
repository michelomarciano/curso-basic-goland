# STRUCTS

Structs (Estrutura) é o mais próximo de uma classe em orientação a objetos. Ele é uma coleção de campos, onde cada campo tem um nome e um tipo. Imagine que você quer salvar os dados de um usuário — não podemos salvar isso em uma string, então usamos o struct para armazenar esse tipo de dados.

**OBSERVAÇÕES:**
- O padrão do Struct é nomear o tipo em **PascalCase** (exportado) ou **camelCase** (não exportado).
- A estrutura do Struct segue uma ordem, o que quer dizer que precisamos passar os valores na mesma ordem que ele foi criado.
- Também podemos usar a sintaxe nomeada para inicializar os campos: `Usuario{Nome: "Mike", Email: "mike@example.com"}`.

**Exemplo:**
```go
type Usuario struct {
    ID    int
    Nome  string
    Email string
}

// Inicialização por ordem
usuario1 := Usuario{1, "João", "joao@example.com"}

// Inicialização nomeada
usuario2 := Usuario{Nome: "Mike", Email: "mike@example.com"}
```

**Exemplo com Structs Aninhados:**
```go
type Endereco struct {
    Rua    string
    Numero int
    Cidade string
    Estado string
    CEP    string
}

type Usuario struct {
    ID       int
    Nome     string
    Email    string
    Senha    string
    Endereco Endereco
}

// Inicialização
usuario := Usuario{
    1, 
    "João", 
    "joao@example.com", 
    "123456", 
    Endereco{
        "Rua das Flores", 
        123, 
        "São Paulo", 
        "SP", 
        "1234567890",
    },
}

// Ou usando sintaxe nomeada
usuario2 := Usuario{
    Nome: "Mike",
    Endereco: Endereco{
        Cidade: "São Paulo",
        Estado: "SP",
    },
}
```

## Casos de Uso
- **Modelagem de dados**: Representar entidades do mundo real (Usuário, Produto, Pedido)
- **Agrupamento de dados relacionados**: Agrupar campos que pertencem juntos
- **APIs e JSON**: Estruturar dados para serialização/deserialização
- **Banco de dados**: Mapear tabelas para structs
- **Validação**: Agrupar dados para validação em conjunto
- **Passagem de parâmetros**: Agrupar múltiplos parâmetros relacionados

**Exemplo prático:**
```go
// Modelo de usuário para API
type Usuario struct {
    ID    int    `json:"id"`
    Nome  string `json:"nome"`
    Email string `json:"email"`
}

// Configuração da aplicação
type Config struct {
    Port     int
    Database string
    Debug    bool
}

// Dados de formulário
type FormularioContato struct {
    Nome    string
    Email   string
    Mensagem string
}
```

