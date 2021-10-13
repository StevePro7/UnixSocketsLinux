#include <arpa/inet.h>
#include <memory.h>
#include <stdio.h>

int main(int argc, char **argv)
{
    struct sockaddr_in dest;
    uint32_t value;

    memset(&dest, '\0', sizeof(dest));
    dest.sin_addr.s_addr = inet_addr("68.178.157.132");

    value = dest.sin_addr.s_addr;
    printf("value = %d\n", value);

    return 0;
}