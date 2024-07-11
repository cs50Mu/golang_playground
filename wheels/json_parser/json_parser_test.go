package jsonparser

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestJsonParser(t *testing.T) {
	// jsonRaw := `{"hello": "world", "x": 12, "y": true, "z": null, "arr": [1, "2", true]}`
	// jsonRaw := `{"hello": "world", "x": 12, "y": true, "z": null, "arr": [1, 2, 3]}`
	jsonRaw := `{
  "hello": "world",
  "device": {
    "name": "xyz",
    "type": "001",
    "isOnline": true
  },
  "x": 12,
  "y": true,
  "z": null,
  "arr": [
    1,
    2,
    3,
    {
      "this": "is",
      "cool": true
    }
  ]
}`
	tokens, err := lex(jsonRaw)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tokens: %+v", spew.Sdump(tokens))

	idx, jsonVal, err := parse(tokens, 0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("idx: %v\njsonVal: %+v\n", idx, spew.Sdump(jsonVal))

	fmt.Printf("%v\n", prettyPrint(jsonVal, 1))
}
