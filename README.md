# Goat

GUI toolkit for the Go programming language based on Tcl/Tk.

## Overview

Goat is a fully cross-platform GUI toolkit based on
[Tcl/Tk](http://www.tcl.tk/). Goat allows you to build GUI applications easily
and with the knowledge of a consistent, native look and feel on every platform.

### Why Tcl/Tk?

Goat development was initiated based on the performance and uptake of the
[Tkinter](https://wiki.python.org/moin/TkInter) toolkit distributed as a
standard part of the [Python](https://www.python.org/) programming language.
Tkinter allows developers easy access to GUI programming with very little
learning. Being the _de facto_ GUI toolkit of Python has introduced more
developers to GUI application programming and increased the popularity of the
language as a whole. Goat is an attempt to provide Go with the same resource.

### Supported platforms

* Linux
* MacOS

#### Windows

It would be quite trivial to support windows but I don't have access to a
Windows machine for development. Pull requests are welcome in this respect.

## Dependencies

### Install Tcl and Tk header files.

#### Linux

```bash
sudo apt install tcl-dev tk-dev
```

#### MacOS

_They may already be installed by default but will need verification._

#### Windows

These would be provided by DLLs.
See https://github.com/nomad-software/tkd#windows-1
