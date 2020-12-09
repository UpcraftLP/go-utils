# go-utils

Monorepo for a bunch of small command line utilities, written in [go](https://golang.org).



## Building

to build all tools, simply run the `build` script:

in **CMD**:

```cmd
build.cmd 1.0.0
```

in **Powershell**:

```powershell
.\build.cmd 1.0.0
```



on **Linux**:

###### *Note: This currently still only builds the windows targets, due to npipe not being available for linux.*

```sh
chmod +x build
./build 1.0.0
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





### time

A utility to record the execution time of another process or command.

**Usage:**

```cmd
# parameters are treated as arguments for the application,
# so this launches 'cmd /c exit 1'
#
# Output: Process finished after 16.9926ms with exit code 1
time cmd /c exit 1
```

