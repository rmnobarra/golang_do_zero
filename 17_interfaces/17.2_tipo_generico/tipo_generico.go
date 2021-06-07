package main

import "fmt"

func generica(interfac interface{}) {
	fmt.Println(interfac)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(1234453))

	mapa := map[interface{}]interface{}{
		1:            "string",
		float64(100): true,
		"string":     "string",
		true:         float64(12),
	}

	fmt.Println(mapa)

}
