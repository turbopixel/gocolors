// Copyright 2022 Nico Hemkes (aká turbopixel). All rights served.
// This code is licensed under the MIT license

package main

import (
	"fmt"
	"image/color"
	"testing"
)

func TestRgb2Hex(t *testing.T) {
	want := "#000000"
	raw := RGB{0, 0, 0}
	converted := Rgb2Hex(raw)

	if converted != want {
		t.Errorf("Expected result: %s. Actual result: %q", want, converted)
	} else {
		fmt.Println(fmt.Sprintf("converted RGB(%d,%d,%d) to HEX %s", raw.Red, raw.Green, raw.Blue, converted))
	}
}

func TestHex2RGB(t *testing.T) {
	var test = color.RGBA{255, 255, 255, 0}
	var rgb color.RGBA

	// white
	var hex = "#fff"
	rgb, _ = Hex2RGB(hex)

	if rgb.R != test.R {
		t.Errorf("Expected result: RGB(%d,%d,%d). Actual result: RGB(%d,%d,%d)", test.R, test.G, test.B, rgb.R, rgb.G, rgb.B)
	} else {
		fmt.Println(fmt.Sprintf("converted %s to RGB(%d,%d,%d)", hex, rgb.R, rgb.G, rgb.B))
	}

	// red
	var hex2 = "#f00"
	test = color.RGBA{255, 0, 0, 0}
	rgb, _ = Hex2RGB(hex2)

	if rgb.R != test.R {
		t.Errorf("Expected result: RGB(%d,%d,%d). Actual result: RGB(%d,%d,%d)", test.R, test.G, test.B, rgb.R, rgb.G, rgb.B)
	} else {
		fmt.Println(fmt.Sprintf("converted %s to RGB(%d,%d,%d)", hex2, rgb.R, rgb.G, rgb.B))
	}
}

func BenchmarkRgb2Hex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rgb2Hex(RGB{0, 100, 255})
		_, _ = Hex2RGB("#f00f00")
	}
}
