# litra

A CLI tool for controlling the [Logitech Litra Glow](https://www.logitech.com/en-us/products/lighting/litra-glow.946-000001.html) USB key light.

## Linux USB Permissions

On Linux, accessing USB HID devices requires either running as root or setting up a udev rule. Create the following file to grant access to all users:

**`/etc/udev/rules.d/99-litra.rules`**
```
SUBSYSTEM=="hidraw", ATTRS{idVendor}=="046d", ATTRS{idProduct}=="c900", MODE="0666"
```

Then reload the rules and re-plug the device:

```bash
sudo udevadm control --reload-rules && sudo udevadm trigger
```

## Installation

**Install system-wide** (to `/usr/local/bin/litra`):

```bash
make install
```

**Or build the binary locally:**

```bash
make build
```

## Usage

```
litra [command]
```

### Turn on / off

```bash
litra on
litra off
```

### Brightness

Set, increase, or decrease brightness. Value is a percentage from `0` to `100`.

```bash
# Set brightness to 80%
litra brightness 80

# Increase brightness by 10%
litra brightness 10 --increase

# Decrease brightness by 10%
litra brightness 10 --decrease
```

### Color Temperature

Set, increase, or decrease color temperature. Value is in Kelvin between `2700` (warm) and `6500` (cool).

```bash
# Set temperature to 4000 K
litra temperature 4000

# Increase temperature by 200 K
litra temperature 200 --increase

# Decrease temperature by 200 K
litra temperature 200 --decrease
```

## Building

```bash
# Build binary to ./litra
make build

# Run without building
make run

# Remove built binary
make clean
```
