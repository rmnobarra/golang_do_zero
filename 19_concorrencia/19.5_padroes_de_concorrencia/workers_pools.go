package main

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

}

func worker()

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
