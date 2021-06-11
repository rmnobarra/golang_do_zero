package main

import (
	"bytes"
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
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))
}
