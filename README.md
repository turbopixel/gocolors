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
	
  // CONVERT FROM RGB COLOR IN HEX
  raw := RGB{0, 0, 0}
  converted := Rgb2Hex(raw) // returns hex color

  fmt.Println(fmt.Sprintf("converted RGB(%d,%d,%d) to HEX %s", raw.Red, raw.Green, raw.Blue, converted))

  // CONVERT FROM HEX IN RGB COLOR 
  hex := "#feefee"
  rgb, _ = Hex2RGB(hex) // returns RGB

  fmt.Println(fmt.Sprintf("converted %s to RGB(%d,%d,%d)", hex, rgb.R, rgb.G, rgb.B))
  
}

```

## License

Open [LICENSE.md](https://github.com/turbopixel/gocolors/blob/master/LICENSE.md) for more informations about the license.
