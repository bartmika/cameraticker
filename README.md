# CameraTicker
[![GoDoc](https://godoc.org/github.com/gomarkdown/markdown?status.svg)](https://pkg.go.dev/github.com/bartmika/cameraticker)
[![Go Report Card](https://goreportcard.com/badge/github.com/bartmika/cameraticker)](https://goreportcard.com/report/github.com/bartmika/cameraticker)
[![License](https://img.shields.io/github/license/bartmika/cameraticker)](https://github.com/bartmika/cameraticker/blob/master/LICENSE)
![Go version](https://img.shields.io/github/go-mod/go-version/bartmika/cameraticker)

CameraTicker is a Go application (with a wrapper over [`libcamera-still`](https://www.raspberrypi.com/documentation/accessories/camera.html)) to periodically capture latest frames from the Raspberry Pi camera.

Under the hood, it executes:

    $ libcamera-still --width xxx --height yyy --format zzz --filename abc

With the supported format values:
* png
* bmp
* rgb
* yuv420

## Usage

```text
Snap periodically scheduled photos from the Raspberry Pi camera

Usage:
  cameraticker [flags]
  cameraticker [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  snap        Snap a single photo with the camera
  start       Start the ticker to snap periodic photos.
  version     Print the version number

Flags:
  -h, --help   help for cameraticker

Use "cameraticker [command] --help" for more information about a command.
```

## Supported Hardware

* Raspberry Pi Camera Module

## Contributing

Found a bug? Want a feature to improve the package? Please create an [issue](https://github.com/bartmika/cameraticker/issues).

## License
Made with ❤️ by [Bartlomiej Mika](https://bartlomiejmika.com).   
The project is licensed under the [ISC License](LICENSE).

Resource used:

* [cgxeiji/picam](https://github.com/cgxeiji/picam) is another similar project which provides a Go developer interface over the Raspberry Pi camera to allow your app to access the `image.Image` or `[]uint8` data immediately.
* [alexellis/phototimer](https://github.com/alexellis/phototimer) is a `python` application written to accomplish the same functionality as `CameraTicker`. The code repository has a lot of useful information to learn from!
