// Copyright 2022 Nico Hemkes (aká turbopixel). All rights served.
// This code is licensed under the MIT license

package main

import (
	"fmt"
	"image/color"
)

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

// Hex2RGB Converts hex color into rgb format.
// original code found on https://stackoverflow.com/a/54200713
func Hex2RGB(hex string) (c color.RGBA, err error) {
	c.A = 0xff

	switch len(hex) {
	case 7:
		_, err = fmt.Sscanf(hex, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(hex, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}

	return c, err
}

// Rgb2Hex converts rgb color into the hex format.
func Rgb2Hex(rgb RGB) string {
	return fmt.Sprintf("#%02x%02x%02x", rgb.Red, rgb.Green, rgb.Blue)
}
