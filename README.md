# GeradorCPF
Gerador de CPF válido em Go, com filtro por estado e exportação para arquivo.

> ⚠️ Este projeto tem finalidade educacional

## Funcionalidades

- Geração de CPF válido (dígitos verificadores corretos)
- Opção de gerar CPF por estado (SP, RJ, PE, etc.)
- Geração em lote com exportação para arquivo `.txt`
- Exibição de progresso a cada 10 CPFs gerados
- Binário standalone (sem dependências externas)


## Compilação

Requer Go 1.20+.

```bash
go build -o geradorcpf
