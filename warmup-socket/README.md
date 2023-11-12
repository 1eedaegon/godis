# How to use TCP on golang after build redis packet

## what is socket with tcp on unix

The socket() syscall returns an fd. Here is a rough explanation of "fd" if you are unfamiliar with Unix systems: An fd is an integer that refers to something in the Linux kernel, like a TCP connection, a disk file, a listening port, or some other resources, etc.

The bind() and listen() syscall: the bind() associates an address to a socket fd, and the listen() enables us to accept connections to that address.

The accept() takes a listening fd, when a client makes a connection to the listening address, the accept() returns an fd that represents the connection socket. Here is the pseudo-code that explains the typical workflow of a server

```c
fd = socket()
bind(fd, address)
listen(fd)
while True:
    conn_fd = accept(fd)
    do_something_with(conn_fd)
    close(conn_fd)
```

system call?

The read() syscall receives data from a TCP connection. The write() syscall sends data. The close() syscall destroys the resource referred by the fd and recycles the fd number.

We have introduced the syscalls needed for server-side network programming. For the client side, the connect() syscall takes a socket fd and address and makes a TCP connection to that address. Here is the pseudo-code for the client

## server/client

## packet parser
