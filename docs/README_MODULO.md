# MÓDULO

Um módulo Go é um conjunto de pacotes relacionados que são versionados juntos. O arquivo `go.mod` na raiz do projeto define o módulo e gerencia as dependências, similar ao `package.json` do Node.js.

Quando você executa `go build` na raiz do projeto, ele compila todos os pacotes e cria um arquivo binário executável (com o nome do módulo ou especificado). Com isso, você pode executar todos os pacotes com `./modulo` (ou o nome do binário criado).

**OBSERVAÇÃO**: O binário não se atualiza sozinho. Ele precisa sempre ser recompilado com `go build` após alterações no código.

| Descrição | Comando |
|-----------|---------|
| Compilar o projeto e criar binário | `go build` |
| Criar um novo módulo | `go mod init <nome>` |
| Executar projeto compilado | `./modulo` |
| Remover dependências não utilizadas | `go mod tidy` | 

## Casos de Uso
- **Gerenciamento de dependências**: Controlar versões de bibliotecas externas
- **Projetos grandes**: Organizar código em múltiplos pacotes dentro de um módulo
- **Versionamento**: Gerenciar diferentes versões do seu projeto
- **Builds reproduzíveis**: Garantir que todos usem as mesmas versões de dependências
- **CI/CD**: Facilitar builds automatizados com dependências consistentes

**Exemplo prático:**
```bash
# Criar um novo projeto
go mod init meu-projeto

# Adicionar dependência externa
go get github.com/gin-gonic/gin

# Atualizar dependências
go get -u ./...

# Limpar dependências não usadas
go mod tidy
```

