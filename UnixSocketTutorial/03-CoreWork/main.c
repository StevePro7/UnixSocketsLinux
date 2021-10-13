#include <sys/types.h>
#include <sys/socket.h>
#include <stdio.h>

int main(int argc, char **argv)
{
    fd_set fds;
    struct timeval tv;

    // do socket initialization
    tv.tv_sec = 1;
    tv.tv_usec = 500000;

    // tv now represents 1.5 seconds
    FD_ZERO(&fds);

    // adds sock to the file descriptor set.
    //FD_SET(sock, &fds);

    // wait 1.5 seconds for any data to be read from any single socket
    //select(socket + 1, &fds, NULL, NULL, &tv);

    //if (FD_ISSET(sock, &fds))
    //{
    //    recvfrom(s, buffer, buffer_len, 0, &sa, &sa_len);
    //}

    printf(" the end!!\n");
    return 0;
}