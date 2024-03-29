color
=====
[![pkg.go.dev documentation](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/adrg/color)
[![License: MIT](https://img.shields.io/github/license/adrg/splash)](http://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/adrg/color)](https://goreportcard.com/report/github.com/adrg/color)

Color is a small CLI application which provides the ability style text
from the standard input.

## Installation

To build use:
```sh
git clone https://github.com/adrg/color.git
cd color
make
```

To install system wide use:
```sh
sudo make install
```

To uninstall use:
```sh
sudo make uninstall
```

### Usage
```sh
input | color [-h] [-v] STYLE...
```

### Input format
Parameters can be specified in the input text with '{index}' where index
is a positive number less than the number of styles passed in to color.
Parameter indices can be repeated. To reset the current style use '{r}'.
A reset is automatically applied before each parameter occurence in the
input text. The last '{r}' in the can be omitted as a reset is done at
the end of the input.

## Style format
```
foreground:background+attributes
```

**Colors**
```
black red green yellow blue magenta cyan white
```

**Attributes**
```
b - Bold
d - Dim
i - Italic
u - Underline
B - Blink
f - Fast blink
r - Reverse
h - Hidden
c - Crossed out
```

## Examples
```sh
echo "{0}green fg" | color green
echo "{0}blue bg{r}, default bg" | color :blue
echo "{0}yellow fg{r}, default fg, {0}yellow fg again" | color yellow
echo "{0}green fg, red bg, bold" | color green:red+b
echo "{0}underline{r}, {1}bold{r} and {2}reverse" | color +u +b +r
echo "{0}blue fg, red bg, bold, reverse" | color blue:red+br
````

<img alt="screenshot" src="https://raw.githubusercontent.com/adrg/adrg.github.io/master/assets/projects/color/screenshot.png" width="100%" />

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.  
See [CONTRIBUTING.MD](CONTRIBUTING.md).

## License

Copyright (c) 2014 Adrian-George Bostan.

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT).
See [LICENSE](LICENSE) for more details.
