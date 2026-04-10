#include <iostream>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <unistd.h>

using namespace std;

int main() {

    int sock;
    struct sockaddr_in serv_addr;

    int a, b, result;
    char op;

    // Create socket
    sock = socket(AF_INET, SOCK_STREAM, 0);

    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(5000);

    inet_pton(AF_INET, "10.0.0.2", &serv_addr.sin_addr);

    // Connect to server
    connect(sock, (struct sockaddr*)&serv_addr, sizeof(serv_addr));

    // Input from user
    cout << "Enter first number: ";
    cin >> a;

    cout << "Enter operator (+ - * /): ";
    cin >> op;

    cout << "Enter second number: ";
    cin >> b;

    // Send values separately (FIXED)
    send(sock, &a, sizeof(a), 0);
    send(sock, &op, sizeof(op), 0);
    send(sock, &b, sizeof(b), 0);

    // Receive result
    read(sock, &result, sizeof(result));

    cout << "Result from server: " << result << endl;

    close(sock);

    return 0;
}