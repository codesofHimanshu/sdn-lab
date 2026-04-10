#include <iostream>
#include <cstring>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <unistd.h>

using namespace std;

int main() {
    int sock = 0;
    struct sockaddr_in serv_addr;
    char buffer[1024] = {0};

    // Create socket
    sock = socket(AF_INET, SOCK_STREAM, 0);

    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(5000);

    // Convert IP address
    inet_pton(AF_INET, "10.0.0.2", &serv_addr.sin_addr);

    // Connect to server
    connect(sock, (struct sockaddr *)&serv_addr, sizeof(serv_addr));

    // Send message
    send(sock, "Hi", 2, 0);

    // Receive response
    read(sock, buffer, 1024);
    cout << "Server says: " << buffer << endl;

    close(sock);

    return 0;
}