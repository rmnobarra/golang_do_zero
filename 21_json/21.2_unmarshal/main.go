package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:none`
	Raca  string `json:raca`
	Idade uint   `json:idade`
}

func main() {
	cachorroEmJSON := `{"Nome":"Rex","Raca":"Dalmata","Idade":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)
}
