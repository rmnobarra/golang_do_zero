# golang do zero

## instalando

1. Obter o link do pacote .tar em https://golang.org/

2. Download do pacote e instalação no path (versão 1.16.4 no exemplo)

```bash
sudo rm -rf /usr/local/go && sudo wget -c https://golang.org/dl/go1.16.4.linux-amd64.tar.gz -O - | sudo tar -xzf -C /usr/local
```

3. Add o path do go nas variáveis de ambiente

sudo echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile

4. para validar
```bash
go version
```

```bash
go env
```

## executando

para executar o hello world

```go
go run ./hello_world/primeiro.go
```

## módulos

* conjunto de pacotes que compoe um projeto

### para iniciar modulo

```go
go mod init <NOME DO MODULO>
```

### para buildar
```go
go build
```

### para executar o arquivo binário executável, gerado pelo build

```bash
./modulo
```

### adicionando pacotes externos

```go
go get github.com/badoux/checkmail
```

### para remover pacotes não utilizados

```go
go mod tidy
```

## variaveis

Duas formas de declarar variaveis, explicitamente e implicitamente

explicitamente:

```go
var variavel1 string = "variavel 1"
```

implicitamente:

```go
variavel2 := "variavel 2"
```

### tipos de dados

numeros inteiros

```go
int8, int16, int32 e int64
```