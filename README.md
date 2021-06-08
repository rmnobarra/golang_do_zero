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

#### numeros inteiros

```go
int8, int16, int32 e int64
```

* o int sem especificar o tipo, utiliza a arquitetura do computador como base. Ex: Se o pc for 64 bits, int64

* uint (unsigned int) um int que desconsidera o sinal.

* alias int32 = rune

* alias uint8 = byte

#### numeros reais

float32, float64

#### string

var str string "texto"

* go não tem char (1 caractere)

#### valor zero 

* valor atribuido a uma variavel quando não é inicializada

* Em go todo tipo de dado tem o valor zero, o valor inicial

#### booleano
true / false

#### erro

### funções

uma função é basicamente uma série de instruções que o programa irá seguir.
Ela pode ter parâmetros e retornos

### operadores

* aritméticos
* atribuição
* relacionais
* lógicos
* unários
* ternários

### struct

mais próximo de uma classe

uma coleção de campos, cada campo tem um nome e um tipo

### ponteiros

* como se fosse uma variavel mas ao invés de guardar um valor, um endereço de memória de algo

* uma referência de memória

* desferenciação


### arrays e slices

array é uma lista de valores

se não especificar o tamanho do array, ele se torna um slice


### urfave cli

https://github.com/urfave/cli


### goroutines

* paralelismo ocorre quando duas ou mais tarefas são executadas exatamente ao mesmo tempo

*  concorrencia não necessariamente são executadas ao mesmo tempo


## testes automatizados

uma função que testará outra função verificando se a execução ocorre como esperado