# go-fbdev

Package **fbdev** provides tools for working with the Frame Buffer Device (fbdev) (that is common on Linux based on operating systems), for the Go programming language.

The _Frame Buffer Device_ is an _easy_ way to do pixel oriented graphics programming.

It is easy because you can treat the _Frame Buffer Device_ as a file (that you can read from, and write to), or as memory (that you can also read from, and write to).


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fbdev

[![GoDoc](https://godoc.org/github.com/reiver/go-fbdev?status.svg)](https://godoc.org/github.com/reiver/go-fbdev)


## Introduction

Documentation on _The Frame Buffer Device_ introduces it as:

> The frame buffer device provides an abstraction for the graphics hardware. It
> represents the frame buffer of some video hardware and allows application
> software to access the graphics hardware through a well-defined interface, so
> the software doesn't need to know anything about the low-level (hardware
> register) stuff.
> 
> The device is accessed through special device nodes, usually located in the
> /dev directory, i.e. /dev/fb*.

Examples of `/dev/fb*` device nodes are:

* `/dev/fb0`
* `/dev/fb1`
* `/dev/fb2`
* etc.


## Treating a Frame Buffer Device as a File

A _Frame Buffer Device_ can be accessed as a file, such as the file `/dev/fb0`, `/dev/fb1`, `/dev/fb2`, `/dev/fb3`, etc.

For instance, in Go, you might do something such as the following (if you wanted to open the _Frame Buffer Device_ at `/dev/fb0`):
```go
fbdev, err := os.Open("/dev/fb0")
```
... or:
```go
fbdev, err := os.OpenFile("/dev/fb0", O_WRONLY)
```
... or even:
```go
fbdev, err := os.OpenFile("/dev/fb0", O_RDWR)
```

(If you wanted to open a different _Frame Buffer Device_, you would just replace the `name` in the `os.Open()` or `os.OpenFile()` call, from `/dev/fb0` to `/dev/fb1`, `/dev/fb2`, `/dev/fb3`, etc.)

Once you have done something such as that, you can then read from one of these types of files, to determine what color specific pixel values are.

For example, in Go, you might do something such as:
```go
var buffer [307200]byte

var p []byte = buffer[:]

// ...

n, err := fbdev.Read(p)
```
Or:
```go
var offset int64
var whence int

// ...

ret, err := fbdev.Seek(offset, whence)

// ...

n, err := fbdev.Read(p)
```
Or alternatively:
```go
var offset int64

// ...

n, err := fbdev.ReadAt(p, offset)

```

For example, you might use `Read` or `ReadAt` in that way if you wanted to create the ability for a program to take a _screenshot_.

**(Although, the astute reader might have noticed that we don't know the _width_ or _height_ of the resulting image, and don't know how many bytes make up a single pixel, and don't know what color format each pixel is using.)**

(But we can figure those things out using `syscall.Syscall(syscall.SYS_IOCTL)`.
More on that later.)

To write to (i.e., to _draw_ to) the _Frame Buffer Device_, we can use `Write` and `WriteAt`, as in the following examples:
```go
n, err := fbdev.Write(b)
```
Or:
```go
var offset int64
var whence int

// ...

ret, err := fbdev.Seek(offset, whence)

// ...

n, err := fbdev.Write(b)
```
Or alternatively:
```go
var offset int64

// ...

n, err := fbdev.WriteAt(b, offset)

```

**(But again, the astute reader might have noticed that we don't know the _width_ or _height_ of the resulting image, and don't know how many bytes make up a single pixel, and don't know what color format each pixel is using.)**

(And again we can figure those things out using `syscall.Syscall(syscall.SYS_IOCTL)`.
More on that later.)


## Memory

The _frame buffer_ can also be accessed as memory, 


## Troubleshooting #1: Testing The Framebuffer Device

One way to see if the Frame Buffer Device is working on your computer is to try to run the following command:

```
cat /dev/urandom >/dev/fb0
```

If you see this error message:
```
-bash: /dev/fb0: Permission denied
```

... then you need to run that command as `root`.


But anyways...

**If this worked, then you should see some colored pixels.**
**At least in the top-left corner, if not filling the screen.**

But, if that didn't work, keep on reading....


## Troubleshooting #2: Virtual Consoles

Note that, if you are in X11, and this test (i.e., running `cat /dev/urandom >/dev/fb0` from a terminal in X11) didn't work for you,
then you may need to switch to a different _virtual console_ to get this to work.

If you are in X11, then you are probably in the _virtual console_ that can be reached by pressing `[CTRL]+[ALT]+[F7]`.

So, to try the test again, first switch to a different _virtual console_.

Ex: press `[CTRL]+[ALT]+[F6]`.

(After pressing that, you should have been switched to a different _virtual console_, and should see a terminal screen, that takes up the whole screen.
It will probably be a a lot of black, with some white text in the top-left corner.)

At this point, you will probably see a prompt that is asking you to login. So... login.

(Once you are logged in....)

Next try running the same command from before. I.e.,:
```
cat /dev/urandom >/dev/fb0
```

And again, if you see this error message:
```
-bash: /dev/fb0: Permission denied
```

... then you need to need to run the program as `root`.


And just like before, **if this worked, then you should see some colored pixels.**
**At least in the top-left corner, if not filling the screen.**

(And then, once you are done, logout of that _virtual console_, and switch back to your original _virtual console_.
Probably by pressing `[CTRL]+[ALT]+[F7]`.)


## Troubleshooting #3: X11

Note that, if you tried that command (i.e,. `cat /dev/urandom >/dev/fb0`) from X11, but it didn't work for you, there is a way to make it work from X11.

First you install the _Frame Buffer Device_ driver for X11, with something like this:
```
sudo apt-get install xserver-xorg-video-fbdev
```

And then you create the configuration file:
```
/usr/share/X11/xorg.conf.d/99-fbdev.conf
```

... with contents such as:
```
Section "Device"
  Identifier "myfb"
  Driver "fbdev"
  Option "fbdev" "/dev/fb1"
EndSection
```

And then restart X11.

The easiest way (to restart X11) if you are running a typical Linux based operating system
(such as [elementary OS](https://elementary.io/), [Fedora](https://getfedora.org/), [openSUSE](https://www.opensuse.org/), [Mint](https://www.linuxmint.com/), [Ubuntu](https://www.ubuntu.com/], etc, or even [Linux From Scratch](http://www.linuxfromscratch.org/)),
(on a laptop, desptop, or tower computer)
that automatically starts X11 for you, is to just to first (completely) _shut down_ you computer (i.e., don't just _suspend_ it, it needs to be completely _shut down_), and then _turn it back on_.

(Alternatively, if you are running a system that doesn't automatically start into X11 for you, you can start X11 with the command: `startx`.)

(Also, if you are running a system that doesn't automatically start into X11 for you, and you don't want to create that configuration file, then you can make X11 use the _Frame Buffer Device_ by starting X11 with the command: `FRAMEBUFFER=/dev/fb1 startx`.)


## See Also

* [The Frame Buffer Device](https://www.kernel.org/doc/Documentation/fb/framebuffer.txt)
* [The Frame Buffer Device API](https://www.kernel.org/doc/Documentation/fb/api.txt)
* [The Frame Buffer Device Documentation](https://www.kernel.org/doc/Documentation/fb/)
* [Framebuffer use](https://github.com/notro/fbtft/wiki/Framebuffer-use)
