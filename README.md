# colorscape

colorscape is a color code conversion tool.

## Overview

```
NAME:
   colorscape - colorscape is a color code conversion tool.

USAGE:
   colorscape.exe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR:
   ysbrothersk

COMMANDS:
     code, c, hex, h  convert to rgb from hex color code
     rgb, r           convert to hex color code from rgb
     help, h          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Usage

Convert to hex color code from rgb

```
> colorscape code FFFFFF
255 255 255
```

Convert to rgb from hex color code

```
> colorscape rgb 255 255 255
FFFFFF
```

## How to Development

### Dependency

* Go
* dep

### Install library

```
> dep ensure
```