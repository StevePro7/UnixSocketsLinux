Unix Socket Tutorial
13/10/2021

https://www.tutorialspoint.com/unix_sockets/index.htm


https://www.tutorialspoint.com/unix_sockets/what_is_socket.htm

Sockets allow communication between two different processes on the same machine
or different machines

File descriptor
An integer associated with an open file or network connection, text file, terminal etc


https://www.tutorialspoint.com/unix_sockets/network_addresses.htm

Subnetting
branch off a network
control network traffic

basic idea of subnetting
partition network address into 2x parts:
1. network address
2. host    address


DNS
Hostname resolution
/etc/hosts


https://www.tutorialspoint.com/unix_sockets/client_server_model.htm

Middleware
perform security checks and load balancing


Concurrent server
simplest way to write concurrent server under Unix is to fork a child process to handle each client separately


CLIENT
socket()
connect()
read()
write()


SERVER
socket()
bind()
listen()
accept()
read()
write()



https://www.tutorialspoint.com/unix_sockets/socket_structures.htm

sockaddr_in 


in_addr 
netid / hostid

hostent

servent


Always pass socket structs by reference [pass pointer]
Always pass length by reference so can be updated
Always sets structure to NULL using memset otherwise may have garbage values


Ports
https://www.tutorialspoint.com/unix_sockets/ports_and_services.htm

/etc/services


port service functions
getservbyname
getservbyport


Network Byte Order
https://www.tutorialspoint.com/unix_sockets/network_byte_orders.htm

Endianness


examples of byte ordering functions
e.g.
host to short		htos
host to long		htol
network to host short	ntos
network to host long	ntol


Functions
https://www.tutorialspoint.com/unix_sockets/ip_address_functions.htm
inet_aton
inet_addr
inet_ntoa
