# Desktop GUI for LG Ultrafine Display Brightness adjustment

`lg-ultrafine-brightness-switcher` is a GUI Application build from [wails](https://wails.app/) to interfaces with the LG UltraFine 4K monitor in order to adjust brightness from Linux.(Because this monitor was designed specifically for Apple computers and thus has no physical buttons. Without the built-in features of macOS, adjusting brightness on this monitor wasn't possible.)

Note: I have tested the UltraFine 4K 23.7" model 24MD4KL. The old 4k 22 inch model or old&new 27 inch 5k models are not tested.
You may need to manually change the Device ID in `go-lguf` library.

## Usage

> Wails does not support building release packages for linux distros yet so you have to manually build this project and run it throuth terminal.

```
// install all dependencies wails need and wails itself.

// cd this project

// build the binary
$wails build
// run application
$sudo build/ultrafine-brightness-slider

```

## If this doesn't work, try change Device ID

in terminal run
```
$lsusb | grep LG
```

As my case the results contain `Bus 003 Device 006: ID 043e:9a63 LG Electronics USA, Inc. LG UltraFine Display Controls`

in `lguf` library change do the following modification:
```
// lguf.go

// before
const (
	MaxBrightness uint16 = 54000
	MinBrightness uint16 = 540

	VendorID  gousb.ID = 0x043e
	ProductID gousb.ID = 0x9a40
)

// after
const (
	MaxBrightness uint16 = 54000
	MinBrightness uint16 = 540

	VendorID  gousb.ID = 0x043e
	ProductID gousb.ID = 0x9a63
)
```


## License & Credits

This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).

This project would not exist if it wasn't for [felicianotech](https://github.com/felicianotech)'s [go-lguf repository](https://github.com/felicianotech/go-lguf/).
It's a Go library that interfaces with the LG UltraFine 4K monitor in order to adjust brightness from Linux. 

Also check out for [Jean-François Beauchamp](https://github.com/velum)'s [repository](https://github.com/velum/lguf-brightness).
It's a C++ based, light-weight version of `go-lguf`.