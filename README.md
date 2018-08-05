# go-fb

Package **fb** provides tools for working with the Frame Buffer Device (fbdev) (that is common on Linux based on operating systems), for the Go programming language.

The _Frame Buffer Device_ is an _easy_ way to do pixel oriented graphics programming.

It is easy because you can treat the _Frame Buffer Device_ as a file (that you can read from, and write to), or as memory (that you can also read from, and write to).


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-fb

[![GoDoc](https://godoc.org/github.com/reiver/go-fb?status.svg)](https://godoc.org/github.com/reiver/go-fb)


## Description

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


## Testing

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


## Virtual Consoles

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


## See Also

* [The Frame Buffer Device](https://www.kernel.org/doc/Documentation/fb/framebuffer.txt)
* [The Frame Buffer Device API](https://www.kernel.org/doc/Documentation/fb/api.txt)
* [The Frame Buffer Device Documentation](https://www.kernel.org/doc/Documentation/fb/)
* [Framebuffer use](https://github.com/notro/fbtft/wiki/Framebuffer-use)
