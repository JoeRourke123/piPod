import socket


def run_client():
    # create a socket object
    client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

    server_ip = "127.0.0.1"  # replace with the server's IP address
    server_port = 9090  # replace with the server's port number
    # establish connection with server
    client.bind((server_ip, server_port))
    client.setblocking(0)

    while True:
        try:
            # input message and send it to the server
            #msg = input("Enter message: ")
            #client.send(msg.encode("utf-8")[:1024])

            # receive message from the server
            response = client.recv(128)
            print(response[0])
            print(response[1])
            print(response[2])
            print("______")
        except:
            pass

    # close client socket (connection to the server)
    client.close()
    print("Connection to server closed")


if __name__ == "__main__":
    run_client()
