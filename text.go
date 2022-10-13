package main

import (
	"bytes"
	"embed"
	"github.com/jbuchbinder/gg"
)

//go:embed assets
var fonts embed.FS

func TextOnImg(opt Options) (*bytes.Buffer, error) {
	var f map[string]Color
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
	var font string
	if opt.Italic {
		font = "assets/Victor-Mono-SemiBold-Italic-NF.ttf"
	} else {
		font = "assets/Victor-Mono-SemiBold-NF.ttf"
	}
	if err := dc.LoadFontFaceFS(fonts, font, 96); err != nil {
		panic(err)
	}
	w, h := dc.MeasureMultilineString(opt.Text, 1.4)

	// heights for the new image, + padding around
	nw := int(w) + 200
	nh := int(h) + 200

	dc = gg.NewContext(nw, nh)
	if err := dc.LoadFontFaceFS(fonts, font, 96); err != nil {
		panic(err)
	}
	hex := f["base"].Hex
	dc.SetHexColor(hex)
	dc.Clear()
	dc.SetHexColor(f[opt.Color].Hex)
	dc.DrawStringWrapped(opt.Text, float64(nw/2), float64(nh/2), 0.5, 0.5, 65536.0, 1.4, gg.AlignCenter)

	buffer := new(bytes.Buffer)
	err := dc.EncodePNG(buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
