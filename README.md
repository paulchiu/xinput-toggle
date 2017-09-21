# xinput-toggle

`xinput-toggle` is a simple command for toggling the enabled/disabled state
of a given device id. For list of devices, run `xinput`

# Usage

`xinput-toggle [device-id]`

# Example

This simple application was written to allow easy toggling of the touchpad on my
ThinkPad.

To get the device id, type `xinput`, you should see output such as:

```
$ xinput
⎡ Virtual core pointer                          id=2    [master pointer  (3)]
⎜   ↳ Virtual core XTEST pointer                id=4    [slave  pointer  (2)]
⎜   ↳ TPPS/2 IBM TrackPoint                     id=12   [slave  pointer  (2)]
⎜   ↳ MX Master                                 id=16   [slave  pointer  (2)]
⎜   ↳ SynPS/2 Synaptics TouchPad                id=11   [slave  pointer  (2)]
⎣ Virtual core keyboard                         id=3    [master keyboard (2)]
    ↳ Virtual core XTEST keyboard               id=5    [slave  keyboard (3)]
    ↳ Power Button                              id=6    [slave  keyboard (3)]
```

Given the above output, to toggle the touchpad, execute `xinput-toggle 11`, you
should see output:

```
$ xinput-toggle 11
Device id 11 disabled
```

To toggle the device again, run the same command and you should see:

```
$ xinput-toggle 11
Device id 11 enabled
```

The command can then be bound to a shortcut key or media key for quick access.

# Installation

Install the binary. See [releases][rels]. Please note that `xinput` needs to be
installed for the command to work.

[rels]: https://github.com/paulchiu/xinput-toggle/releases
