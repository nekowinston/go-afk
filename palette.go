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
	Latte     map[string]Color `json:"latte"`
	Frappe    map[string]Color `json:"frappe"`
	Macchiato map[string]Color `json:"macchiato"`
	Mocha     map[string]Color `json:"mocha"`
}

type Color struct {
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
