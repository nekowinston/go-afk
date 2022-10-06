package main

import (
	_ "embed"
	"encoding/json"
)

func UnmarshalPalette(data []byte) (Palette, error) {
	var r Palette
	err := json.Unmarshal(data, &r)
	return r, err
}

type Palette struct {
	Latte     map[string]Colour `json:"latte"`
	Frappe    map[string]Colour `json:"frappe"`
	Macchiato map[string]Colour `json:"macchiato"`
	Mocha     map[string]Colour `json:"mocha"`
}

type Colour struct {
	Hex string `json:"hex"`
	RGB string `json:"rgb"`
	Hsl string `json:"hsl"`
}

var palette Palette

//go:embed palette.json
var paletteJSON []byte

func init() {
	palette, _ = UnmarshalPalette(paletteJSON)
}
