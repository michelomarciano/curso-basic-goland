# RODAR PROJETO

Para rodar o projeto você pode executar os seguintes comandos:

| Descrição | Comando |
|-----------|---------|
| Rodar projeto diretamente | `go run <arquivo>.go` |
| Rodar projeto compilado | `./modulo` |

## Casos de Uso
- **Desenvolvimento rápido**: Usar `go run` para testar código rapidamente durante desenvolvimento
- **Produção**: Compilar com `go build` para criar binário otimizado para produção
- **Distribuição**: Compilar para diferentes plataformas (Windows, Linux, macOS)
- **CI/CD**: Compilar em pipelines de integração contínua
- **Debugging**: Compilar com flags de debug para análise

**Exemplo prático:**
```bash
# Desenvolvimento - execução rápida
go run main.go

# Compilar para produção
go build -o app

# Compilar para outra plataforma
GOOS=linux GOARCH=amd64 go build

# Executar binário compilado
./app
```

