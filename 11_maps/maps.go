package main

import "fmt"

func main() {
	fmt.Println("Maps")
	fmt.Println("-------------------")

	usuario := map[string]string{
		"nome":      "Macaco",
		"sobrenome": "Louco",
	}

	fmt.Println(usuario["nome"], usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Florzinha",
			"ultimo":   "Poderosa",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 69",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "touro",
	}
	fmt.Println(usuario2)
}
