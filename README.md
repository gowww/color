# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) color [![GoDoc](https://godoc.org/github.com/gowww/color?status.svg)](https://godoc.org/github.com/gowww/color) [![Build](https://travis-ci.org/gowww/color.svg?branch=master)](https://travis-ci.org/gowww/color) [![Coverage](https://coveralls.io/repos/github/gowww/color/badge.svg?branch=master)](https://coveralls.io/github/gowww/color?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/color)](https://goreportcard.com/report/github.com/gowww/color) ![Status Stable](https://img.shields.io/badge/status-stable-brightgreen.svg)

Package [color](https://godoc.org/github.com/gowww/color) provides color manipulation and format conversion.

## Installing

1.  Get package:

    ```Shell
    go get -u github.com/gowww/color
    ```

2.  Import it in your code:

    ```Go
    import "github.com/gowww/color"
    ```

## Usage

### From hexadecimal

To get a [Color](https://godoc.org/github.com/gowww/color#Color) instance from an hexadecimal representation, use [NewHex](https://godoc.org/github.com/gowww/color#NewHex):

```Go
color.NewHex("#004bff")
```

### From RGB

To get a [Color](https://godoc.org/github.com/gowww/color#Color) instance from red, green and blue values, use [NewRGB](https://godoc.org/github.com/gowww/color#NewRGB):

```Go
color.NewRGB(0, 75, 255)
```

To get a [Color](https://godoc.org/github.com/gowww/color#Color) instance from red, green, blue and alpha values, use [NewRGBA](https://godoc.org/github.com/gowww/color#NewRGBA):

```Go
color.NewRGBA(0, 75, 255, .5)
```

### To hexadecimal

To get the hexadecimal representation of a [Color](https://godoc.org/github.com/gowww/color#Color) without alpha, use [Color.Hex](https://godoc.org/github.com/gowww/color#Color.Hex):

```Go
color.NewRGBA(0, 75, 255, .5).Hex() // "#004bff"
```

To get the hexadecimal representation of a [Color](https://godoc.org/github.com/gowww/color#Color) with alpha, use [Color.HexAlpha](https://godoc.org/github.com/gowww/color#Color.HexAlpha) or [Color.AlphaHex](https://godoc.org/github.com/gowww/color#Color.AlphaHex):

```Go
color.NewRGBA(0, 75, 255, .5).HexAlpha() // "#004bff7f"
color.NewRGBA(0, 75, 255, .5).HexAlpha() // "#7f004bff"
```
