# Problema de Josephus

Solução para o [Problema de Josephus](https://en.wikipedia.org/wiki/Josephus_problem) para a disciplina de Estrutura de Dados.

# Requisitos

- Deve-se usar uma estrutura de lista ligada circular;

- Receber os dados do problemas via stdin e apresentar via stdout;

## Exemplo

- Entrada:

```zsh
3
10
1
10
2
10
3
```

- Saída esperada:

```zsh
Usando n=10, m=1, resultado=5
Usando n=10, m=2, resultado=10
Usando n=10, m=3, resultado=6
```

## Solução

- Instale [Go](https://golang.org/);

- Faço build para gerar um executável:

```zsh
go build -o bin/josephus main.go
```

> [Compile](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies) para arquitetura de sua preferência.

- Execute o binário gerado na pasta `/bin`:

```zsh
# no linux:
./bin/josephus
```

- Passe **um** argumento por linha, sendo o primeiro n=1 o número de casos de testes, n+1 o número de pessoas e n+2 o passo.
