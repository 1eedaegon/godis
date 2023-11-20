package main

// ** SUDO CODES

// Server
// fd = socket() // 디스크립터 integer
// bind(fd, address) // 주소 바인딩
// listen(fd) // 1. make queue, 2. transite to passive socket, 3. Wait call accept()
// while True:
//     conn_fd = accept(fd) // wait something
//     do_something_with(conn_fd) //do something
//     close(conn_fd) // socket closing

// Client
// fd = socket() // 디스크립터 integer
// connect(fd, address) //
// do_something_with(fd)
// close(fd)

// int val = 1;
// setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &val, sizeof(val));

// struct sockaddr_in addr = {};
// addr.sin_family = AF_INET;
// addr.sin_port = ntohs(1234);
// addr.sin_addr.s_addr = ntohl(0);    // wildcard address 0.0.0.0
// int rv = bind(fd, (const sockaddr *)&addr, sizeof(addr));
// if (rv) {
// 	die("bind()");
// }

// // listen
// rv = listen(fd, SOMAXCONN);
// if (rv) {
// 	die("listen()");
// }
