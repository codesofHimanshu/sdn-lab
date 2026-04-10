#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>

using namespace std;

int main() {

    int server_fd, new_socket;
    struct sockaddr_in address;
    int addrlen = sizeof(address);

    int a, b, result;
    char op;

    // Create socket
    server_fd = socket(AF_INET, SOCK_STREAM, 0);

    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(5000);

    bind(server_fd, (struct sockaddr*)&address, sizeof(address));
    listen(server_fd, 3);

    cout << "Server waiting..." << endl;

    new_socket = accept(server_fd, (struct sockaddr*)&address, (socklen_t*)&addrlen);

    // Receive values separately (FIXED)
    read(new_socket, &a, sizeof(a));
    read(new_socket, &op, sizeof(op));
    read(new_socket, &b, sizeof(b));

    cout << "Received: " << a << " " << op << " " << b << endl;

    // Perform calculation
    switch(op) {
        case '+': result = a + b; break;
        case '-': result = a - b; break;
        case '*': result = a * b; break;
        case '/':
            if(b == 0) {
                cout << "Division by zero!" << endl;
                result = 0;
            } else {
                result = a / b;
            }
            break;
        default:
            cout << "Invalid operator!" << endl;
            result = 0;
    }

    // Send result
    send(new_socket, &result, sizeof(result), 0);

    close(new_socket);
    close(server_fd);

    return 0;
}