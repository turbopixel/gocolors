# gocolors
Tiny color converter written in [golang](https://go.dev/).

With this small library color values can be converted from RGB to HEX or vice versa.

## Requirements

* go 1.18

## Example use

```golang
package main

import (
  "fmt"
)

func example() {
	
  // CONVERT COLOR FROM RGB IN HEX
  rgbcolor := RGB{0, 0, 0}
  converted := Rgb2Hex(rgbcolor) // returns hex color

  fmt.Println(fmt.Sprintf("converted RGB(%d,%d,%d) to HEX %s", rgbcolor.Red, rgbcolor.Green, rgbcolor.Blue, converted))
  // output: converted RGB(0,0,0) to HEX #000000


  // CONVERT COLOR FROM HEX IN RGB
  hex := "#feefee"
  rgb, _ = Hex2RGB(hex) // returns RGB

  fmt.Println(fmt.Sprintf("converted %s to RGB(%d,%d,%d)", hex, rgb.R, rgb.G, rgb.B))
  // output: converted #feefee to RGB(254,239,238)  
}

```

## License

Open [LICENSE.md](https://github.com/turbopixel/gocolors/blob/master/LICENSE.md) for more informations about the license.
