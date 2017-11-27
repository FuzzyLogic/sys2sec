sys2sec - Convert system call traces to a unique syscall list for seccomp
======
**Synopsis** strace output -> sys2sec -> seccomp input

## Table of Contents
* Description
* Building and Usage
* TODOs

## Description
sys2sec is extremely simple. Just pipe your strace output into it, define a trigger and obtain your list of required system calls.

## Building, Testing and Installing
sys2sec only requires a go version >= 1.6.3 to be built.
This is based on the one used when developing this project.

#### Building and Usage
Building couldn't be simpler.

```bash
$ go build -o sys2sec main.go
```

Using it is just as simple. The input is expected on STDIN. The only argument to sys2sec is a trigger. The trigger can be any substring in your STDIN stream.
The following example shows how to get all syscalls when ping is used to send two packets to localhost.

```bash
$ sudo strace ping localhost -c 2 2>&1 >/dev/null | ./sys2sec "execve"
Triggered on: execve("/bin/ping", ["ping", "localhost", "-c", "2"], [/* 26 vars */]) = 0

brk
access
mmap
open
fstat
close
read
mprotect
arch_prctl
munmap
capget
capset
prctl
getuid
setuid
geteuid
socket
getpid
stat
connect
getsockname
setsockopt
getsockopt
ioctl
rt_sigaction
rt_sigprocmask
sendmsg
recvmsg
write
sched_yield
setitimer
exit_group
```

The trigger is set to the first occurrence of the execve() syscall.

#### TODOs
Make sys2sec more sophisticated and add some tests.
