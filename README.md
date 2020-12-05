# go-utils

Monorepo for a bunch of small command line utilities, written in [go](https://golang.org).



## Building

to build all tools, simply run the `build` script:

*in **CMD**:*

```cmd
build.cmd 1.0.0
```

*in **Powershell**:*

```powershell
.\build.cmd 1.0.0
```



You will then find the outputs in the `bin` directory.





## Current Tools:

### datetime

A simple utility to print the current [Unix Time](https://en.wikipedia.org/wiki/Unix_time).

**Usage:**

```cmd
datetime
```





### listener

A command line utility that will echo anything it receives on the provided port or named pipe.

**Usage:**

```cmd
# listen on a named pipe
listener \\.\pipe\someNamedPipe
```

```cmd
# listen on a port
listener 22
```



---

