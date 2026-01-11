# JSON

JSON (JavaScript Object Notation) é um formato de dados leve e legível usado para troca de informações. Em Go, o pacote `encoding/json` fornece funcionalidades para serializar (Marshal) e deserializar (Unmarshal) dados JSON.

## Import

```go
import "encoding/json"
```

## Marshal - Convertendo Struct para JSON

O `json.Marshal()` converte uma struct (ou qualquer valor) em um array de bytes JSON.

**Sintaxe:**
```go
jsonBytes, err := json.Marshal(valor)
```

**Exemplo Básico:**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Usuario struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
    Email string `json:"email"`
}

func main() {
    usuario := Usuario{
        Nome:  "Mike",
        Idade: 20,
        Email: "mike@example.com",
    }
    
    jsonBytes, err := json.Marshal(usuario)
    if err != nil {
        fmt.Println("Erro ao fazer Marshal:", err)
        return
    }
    
    fmt.Println("JSON:", string(jsonBytes))
    // Output: JSON: {"nome":"Mike","idade":20,"email":"mike@example.com"}
}
```

## Entendendo `jsonBytes, err := json.Marshal(jsonExemplo)`

Esta linha é fundamental para trabalhar com JSON em Go. Vamos quebrá-la em partes para entender completamente:

### O que faz

A função `json.Marshal()` converte uma struct Go em um array de bytes contendo o JSON correspondente.

### Quebra da linha

```go
jsonBytes, err := json.Marshal(jsonExemplo)
```

**Componentes:**

1. **`json.Marshal()`** - Função do pacote `encoding/json` que serializa (converte) um valor Go em JSON
2. **`jsonExemplo`** - O valor que você quer converter (geralmente uma struct)
3. **Retorno múltiplo** - Go permite funções retornarem múltiplos valores:
   - **`jsonBytes`** - `[]byte` contendo o JSON gerado
   - **`err`** - `error` que será `nil` se tudo deu certo, ou conterá informações sobre o erro se algo falhou

### Como funciona passo a passo

```go
// 1. Você tem uma struct
jsonExemplo := jsons.Usuario{
    Nome:  "Mike",
    Idade: 20,
    Email: "mike@example.com",
}

// 2. json.Marshal converte para bytes JSON
jsonBytes, err := json.Marshal(jsonExemplo)

// 3. jsonBytes contém: []byte(`{"nome":"Mike","idade":20,"email":"mike@example.com"}`)
// 4. err será nil se tudo deu certo, ou terá um erro se algo falhou
```

### Exemplo visual

```go
// ANTES (Struct Go):
jsonExemplo := jsons.Usuario{
    Nome:  "Mike",
    Idade: 20,
    Email: "mike@example.com",
}

// DEPOIS (JSON em bytes):
jsonBytes = []byte(`{"nome":"Mike","idade":20,"email":"mike@example.com"}`)
```

### Retorno múltiplo em Go

Go permite que funções retornem múltiplos valores. A função `json.Marshal` retorna dois valores:

1. **`[]byte`** - Os dados JSON em formato de bytes
2. **`error`** - O erro (se houver)

```go
// Sintaxe de atribuição múltipla
jsonBytes, err := json.Marshal(jsonExemplo)
//     ↑         ↑
//   bytes    erro
```

### Convertendo bytes para string

O `jsonBytes` é do tipo `[]byte` (array de bytes). Para imprimir como texto legível, você precisa converter para `string`:

```go
jsonBytes, err := json.Marshal(jsonExemplo)
if err != nil {
    fmt.Println("Erro:", err)
} else {
    // Converte []byte para string para poder imprimir
    fmt.Println("JSON:", string(jsonBytes))
    // Output: JSON: {"nome":"Mike","idade":20,"email":"mike@example.com"}
}
```

### Tratamento de erro

Sempre verifique o erro antes de usar os dados:

```go
jsonBytes, err := json.Marshal(jsonExemplo)
if err != nil {
    // Algo deu errado na conversão
    fmt.Println("Erro ao fazer Marshal:", err)
    return // ou tratar o erro de outra forma
}
// Se chegou aqui, jsonBytes contém o JSON válido
fmt.Println("JSON:", string(jsonBytes))
```

### Exemplo completo comentado

```go
// 1. Criar uma struct
jsonExemplo := jsons.Usuario{
    Nome:  "Mike",
    Idade: 20,
    Email: "mike@example.com",
}

// 2. Converter struct para JSON (bytes)
//    json.Marshal retorna 2 valores:
//    - jsonBytes: []byte com o JSON
//    - err: error (nil se sucesso, ou erro se falhou)
jsonBytes, err := json.Marshal(jsonExemplo)

// 3. Verificar se houve erro
if err != nil {
    fmt.Println("Erro ao fazer Marshal:", err)
} else {
    // 4. Converter bytes para string e imprimir
    fmt.Println("JSON:", string(jsonBytes))
    // Output: JSON: {"nome":"Mike","idade":20,"email":"mike@example.com"}
}
```

### Resumo

- **`json.Marshal()`** converte struct → JSON (bytes)
- **Retorna dois valores**: `[]byte` e `error`
- **`jsonBytes`** contém o JSON em formato de bytes
- **`err`** será `nil` se tudo deu certo, ou conterá informações sobre o erro
- **Use `string(jsonBytes)`** para converter bytes em texto legível
- **Sempre verifique `err`** antes de usar `jsonBytes`

Esta é a forma padrão em Go de converter dados para JSON, com tratamento de erro integrado.

## Unmarshal - Convertendo JSON para Struct

O `json.Unmarshal()` converte um array de bytes JSON em uma struct.

**Sintaxe:**
```go
err := json.Unmarshal(jsonBytes, &struct)
```

**Exemplo Básico:**
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Usuario struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
    Email string `json:"email"`
}

func main() {
    jsonString := `{"nome":"Mike","idade":20,"email":"mike@example.com"}`
    jsonBytes := []byte(jsonString)
    
    var usuario Usuario
    err := json.Unmarshal(jsonBytes, &usuario)
    if err != nil {
        fmt.Println("Erro ao fazer Unmarshal:", err)
        return
    }
    
    fmt.Printf("Nome: %s, Idade: %d, Email: %s\n", usuario.Nome, usuario.Idade, usuario.Email)
    // Output: Nome: Mike, Idade: 20, Email: mike@example.com
}
```

## Tags JSON

As tags JSON são anotações que controlam como os campos são serializados/deserializados. Elas são definidas usando backticks (`` ` ``) após o tipo do campo.

**Sintaxe:**
```go
type Struct struct {
    Campo string `json:"nome_no_json"`
}
```

### Tags Comuns

**1. Renomear Campo:**
```go
type Usuario struct {
    NomeCompleto string `json:"nome"`  // Campo "NomeCompleto" vira "nome" no JSON
}
```

**2. Omitir Campo Vazio:**
```go
type Usuario struct {
    Nome  string `json:"nome,omitempty"`  // Não inclui se estiver vazio
    Email string `json:"email,omitempty"`
}
```

**3. Ignorar Campo:**
```go
type Usuario struct {
    Nome  string `json:"nome"`
    Senha string `json:"-"`  // Campo nunca aparece no JSON
}
```

**4. Converter para String:**
```go
type Produto struct {
    ID   int    `json:"id,string"`  // Converte int para string no JSON
    Nome string `json:"nome"`
}
```

**Exemplo Completo com Tags:**
```go
type Usuario struct {
    ID          int    `json:"id"`
    Nome        string `json:"nome"`
    Email       string `json:"email,omitempty"`
    Senha       string `json:"-"`  // Nunca serializa
    Ativo       bool   `json:"ativo"`
    CriadoEm    string `json:"criado_em,omitempty"`
}

func main() {
    usuario := Usuario{
        ID:    1,
        Nome:  "Mike",
        Email: "",  // Será omitido por causa do omitempty
        Senha: "123456",  // Nunca será serializado
        Ativo: true,
    }
    
    jsonBytes, _ := json.Marshal(usuario)
    fmt.Println(string(jsonBytes))
    // Output: {"id":1,"nome":"Mike","ativo":true}
}
```

## MarshalIndent - JSON Formatado

O `json.MarshalIndent()` produz JSON formatado (com indentação) para melhor legibilidade.

**Sintaxe:**
```go
jsonBytes, err := json.MarshalIndent(valor, prefixo, indentacao)
```

**Exemplo:**
```go
usuario := Usuario{Nome: "Mike", Idade: 20, Email: "mike@example.com"}

jsonBytes, err := json.MarshalIndent(usuario, "", "  ")
if err != nil {
    fmt.Println("Erro:", err)
    return
}

fmt.Println(string(jsonBytes))
// Output:
// {
//   "nome": "Mike",
//   "idade": 20,
//   "email": "mike@example.com"
// }
```

## Structs Aninhadas

JSON suporta estruturas aninhadas, que são representadas como objetos JSON aninhados.

**Exemplo:**
```go
type Endereco struct {
    Rua    string `json:"rua"`
    Numero int    `json:"numero"`
    Cidade string `json:"cidade"`
}

type Usuario struct {
    Nome     string   `json:"nome"`
    Email    string   `json:"email"`
    Endereco Endereco `json:"endereco"`
}

func main() {
    usuario := Usuario{
        Nome:  "Mike",
        Email: "mike@example.com",
        Endereco: Endereco{
            Rua:    "Rua das Flores",
            Numero: 123,
            Cidade: "São Paulo",
        },
    }
    
    jsonBytes, _ := json.Marshal(usuario)
    fmt.Println(string(jsonBytes))
    // Output: {"nome":"Mike","email":"mike@example.com","endereco":{"rua":"Rua das Flores","numero":123,"cidade":"São Paulo"}}
}
```

## Arrays e Slices

JSON suporta arrays, que em Go são representados como slices.

**Exemplo:**
```go
type Usuario struct {
    Nome  string   `json:"nome"`
    Hobbies []string `json:"hobbies"`
}

func main() {
    usuario := Usuario{
        Nome:    "Mike",
        Hobbies: []string{"programação", "leitura", "música"},
    }
    
    jsonBytes, _ := json.Marshal(usuario)
    fmt.Println(string(jsonBytes))
    // Output: {"nome":"Mike","hobbies":["programação","leitura","música"]}
}
```

## Maps

Maps em Go são convertidos para objetos JSON.

**Exemplo:**
```go
func main() {
    dados := map[string]interface{}{
        "nome":  "Mike",
        "idade": 20,
        "ativo": true,
    }
    
    jsonBytes, _ := json.Marshal(dados)
    fmt.Println(string(jsonBytes))
    // Output: {"ativo":true,"idade":20,"nome":"Mike"}
}
```

## Tratamento de Erros

Sempre trate erros ao usar Marshal e Unmarshal.

**Exemplo com Tratamento de Erro:**
```go
func main() {
    usuario := Usuario{Nome: "Mike", Idade: 20}
    
    // Marshal
    jsonBytes, err := json.Marshal(usuario)
    if err != nil {
        fmt.Printf("Erro ao serializar: %v\n", err)
        return
    }
    fmt.Println("JSON:", string(jsonBytes))
    
    // Unmarshal
    var novoUsuario Usuario
    err = json.Unmarshal(jsonBytes, &novoUsuario)
    if err != nil {
        fmt.Printf("Erro ao deserializar: %v\n", err)
        return
    }
    fmt.Printf("Usuário: %+v\n", novoUsuario)
}
```

## Casos de Uso

### 1. APIs REST
```go
// Serializar resposta da API
func GetUsuario(w http.ResponseWriter, r *http.Request) {
    usuario := Usuario{Nome: "Mike", Idade: 20, Email: "mike@example.com"}
    jsonBytes, _ := json.Marshal(usuario)
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)
}

// Deserializar requisição
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
    var usuario Usuario
    json.NewDecoder(r.Body).Decode(&usuario)
    // Processar usuário...
}
```

### 2. Configurações
```go
type Config struct {
    Port     int    `json:"port"`
    Database string `json:"database"`
    Debug    bool   `json:"debug"`
}

// Carregar configuração de arquivo JSON
func LoadConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    
    var config Config
    err = json.Unmarshal(data, &config)
    return &config, err
}
```

### 3. Armazenamento de Dados
```go
// Salvar struct em arquivo JSON
func SaveUsuario(usuario Usuario, filename string) error {
    jsonBytes, err := json.MarshalIndent(usuario, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, jsonBytes, 0644)
}
```

### 4. Comunicação entre Serviços
```go
// Enviar dados via HTTP
func SendData(url string, data interface{}) error {
    jsonBytes, err := json.Marshal(data)
    if err != nil {
        return err
    }
    
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    return nil
}
```

## Exemplos Avançados

### Struct com Campos Opcionais
```go
type Usuario struct {
    ID       int     `json:"id"`
    Nome     string  `json:"nome"`
    Email    *string `json:"email,omitempty"`  // Ponteiro permite nil
    Telefone *string `json:"telefone,omitempty"`
}

func main() {
    nome := "Mike"
    usuario := Usuario{
        ID:   1,
        Nome: nome,
        // Email e Telefone são nil, serão omitidos
    }
    
    jsonBytes, _ := json.Marshal(usuario)
    fmt.Println(string(jsonBytes))
    // Output: {"id":1,"nome":"Mike"}
}
```

### Custom Marshal/Unmarshal
```go
import "time"

type DataCustomizada struct {
    time.Time
}

func (d DataCustomizada) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf(`"%s"`, d.Time.Format("2006-01-02"))), nil
}

type Evento struct {
    Nome string         `json:"nome"`
    Data DataCustomizada `json:"data"`
}
```

### Trabalhando com JSON Raw
```go
type Mensagem struct {
    Tipo string          `json:"tipo"`
    Dados json.RawMessage `json:"dados"`  // Permite dados JSON arbitrários
}

func main() {
    jsonString := `{
        "tipo": "usuario",
        "dados": {"nome":"Mike","idade":20}
    }`
    
    var msg Mensagem
    json.Unmarshal([]byte(jsonString), &msg)
    
    // Processar dados baseado no tipo
    if msg.Tipo == "usuario" {
        var usuario Usuario
        json.Unmarshal(msg.Dados, &usuario)
        fmt.Println(usuario)
    }
}
```

## OBSERVAÇÕES IMPORTANTES

1. **Campos Exportados**: Apenas campos com letra maiúscula (exportados) são serializados
2. **Ponteiros**: Use ponteiros para campos opcionais que podem ser `nil`
3. **Tratamento de Erros**: Sempre trate erros ao usar Marshal/Unmarshal
4. **Performance**: `json.Marshal` é mais rápido que `json.MarshalIndent`, mas menos legível
5. **Tags**: Use tags para controlar a serialização (omitempty, -, string, etc.)
6. **Tipos Suportados**: Go suporta automaticamente: strings, números, booleans, arrays, maps, structs, ponteiros
7. **Valores Zero**: Campos com valores zero são incluídos por padrão, use `omitempty` para omitir

## Exemplo Completo

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Endereco struct {
    Rua    string `json:"rua"`
    Numero int    `json:"numero"`
    Cidade string `json:"cidade"`
}

type Usuario struct {
    ID       int      `json:"id"`
    Nome     string   `json:"nome"`
    Email    string   `json:"email,omitempty"`
    Senha    string   `json:"-"`
    Ativo    bool     `json:"ativo"`
    Endereco Endereco `json:"endereco"`
    Hobbies  []string `json:"hobbies"`
}

func main() {
    // Criar usuário
    usuario := Usuario{
        ID:    1,
        Nome:  "Mike",
        Email: "mike@example.com",
        Senha: "123456",
        Ativo: true,
        Endereco: Endereco{
            Rua:    "Rua das Flores",
            Numero: 123,
            Cidade: "São Paulo",
        },
        Hobbies: []string{"programação", "leitura"},
    }
    
    // Marshal - Struct para JSON
    jsonBytes, err := json.MarshalIndent(usuario, "", "  ")
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }
    fmt.Println("JSON Formatado:")
    fmt.Println(string(jsonBytes))
    
    // Unmarshal - JSON para Struct
    jsonString := `{
        "id": 2,
        "nome": "João",
        "email": "joao@example.com",
        "ativo": true,
        "endereco": {
            "rua": "Av. Paulista",
            "numero": 456,
            "cidade": "São Paulo"
        },
        "hobbies": ["esporte", "música"]
    }`
    
    var novoUsuario Usuario
    err = json.Unmarshal([]byte(jsonString), &novoUsuario)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }
    
    fmt.Println("\nUsuário Deserializado:")
    fmt.Printf("ID: %d\n", novoUsuario.ID)
    fmt.Printf("Nome: %s\n", novoUsuario.Nome)
    fmt.Printf("Email: %s\n", novoUsuario.Email)
    fmt.Printf("Ativo: %v\n", novoUsuario.Ativo)
    fmt.Printf("Endereço: %s, %d - %s\n", 
        novoUsuario.Endereco.Rua, 
        novoUsuario.Endereco.Numero, 
        novoUsuario.Endereco.Cidade)
    fmt.Printf("Hobbies: %v\n", novoUsuario.Hobbies)
}
```

## Resumo

- **Marshal**: Converte Go struct → JSON (bytes)
- **Unmarshal**: Converte JSON (bytes) → Go struct
- **Tags JSON**: Controlam serialização (`json:"nome"`, `omitempty`, `-`)
- **Tratamento de Erros**: Sempre verifique erros
- **MarshalIndent**: Produz JSON formatado (mais legível)
- **Suporta**: Structs, arrays, slices, maps, tipos primitivos

