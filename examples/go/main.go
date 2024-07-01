package main

import (
	"encoding/json"
	"fmt"
	"log"

	gen "github.com/superpowerdotcom/medplum-gen/build/go"
)

func main() {
	od := &gen.ObservationDefinition{
		Id: StrPtr("123"),
	}

	data, err := json.Marshal(od)
	if err != nil {
		log.Fatalf("unable to marshal: %v", err)
	}

	fmt.Println("Marshalled data: ", string(data))
}

func StrPtr(s string) *string {
	return &s
}
