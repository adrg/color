package main

const usage = `Usage: color [-h] [-v] STYLE...
Try 'color --help' for more information.
`

const help = `Usage: color [-h] [-v] STYLE...
Stye text from the standard input using each STYLE specified.

Input format:
  Parameters can be specified in the input text with '{index}' where index
  is a positive number less than the number of styles passed in to color.
  Parameter indices can be repeated. To reset the current style use '{r}'.
  A reset is automatically applied before each parameter occurrence in the
  input text. The last '{r}' in the can be omitted as a reset is done at
  the end of the input.

Style format: foreground:background+attributes
  Foreground and background can have the following values: black, red,
  green, yellow, blue, magenta, cyan, white. The supported attribute values
  are: b - bold, d - dim, i - italic, u - underline, b - blink, r - reverse,
  f - fast blink, h - hidden,  c - crossed out.

Examples:
  echo "{0}green fg" | color green
  echo "{0}blue bg{r} default bg" | color :blue
  echo "{0}yellow fg{r} default fg {0}yellow fg again" | color yellow
  echo "{0}green fg red bg bold" | color green:red+b
  echo "{0}underline{r}, {1}bold{r} and {2}reverse" | color +u +b +r
  echo "{0}blue fg red bg bold reverse" | color blue:red+br

Optional arguments:
  -h, --help	 show this help message and exit
  -v, --version  show version information and exit

For reporting bugs, getting help or just browsing the source code of the
application visit: <https://github.com/adrg/color>
`

const version = `color 1.0
Copyright (C) 2014 Adrian-George Bostan <adrg@epistack.com>
License MIT: <http://opensource.org/licenses/MIT>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Adrian-George Bostan.
`
