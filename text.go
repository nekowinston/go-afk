package main

import (
	"github.com/fogleman/gg"
	"image"
)

func TextOnImg(opt Options) (image.Image, error) {
	var f map[string]Colour
	switch opt.Flavor {
	case "latte":
		f = palette.Latte
	case "frappe":
		f = palette.Frappe
	case "macchiato":
		f = palette.Macchiato
	case "mocha":
		f = palette.Mocha
	default:
		f = palette.Frappe
	}

	// measure the text on a tmp context
	dc := gg.NewContext(4096, 4096)
	if err := dc.LoadFontFace("assets/Victor-Mono SemiBold-Nerd Font.ttf", 96); err != nil {
		panic(err)
	}
	w, h := dc.MeasureString(opt.Text)

	// heights for the new image, + padding around
	nw := int(w) + 100
	nh := int(h) + 100

	dc = gg.NewContext(nw, nh)
	if err := dc.LoadFontFace("assets/Victor-Mono SemiBold-Nerd Font.ttf", 96); err != nil {
		panic(err)
	}
	hex := f["base"].Hex
	dc.SetHexColor(hex)
	dc.Clear()
	dc.SetHexColor(f[opt.Color].Hex)
	dc.DrawStringAnchored(opt.Text, float64(nw/2), float64(nh/2), 0.5, 0.5)

	return dc.Image(), nil
}
