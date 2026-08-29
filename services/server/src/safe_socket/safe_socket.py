import socket

# TODO: Complete with a short-read/short-write tolerant implementation


def recv_all(socket: socket.socket, size):
    read = 0
    message_parts = []

    while read < size:
        received_message = socket.recv(size - read)
        if len(received_message) == 0:
            raise Exception("The other end is closed") # Segun la docu si se recibe un '', la conexion se cerro -> Chequear
        
        read += len(received_message)
        message_parts.append(received_message)

    return b"".join(message_parts)


def send_all(socket: socket.socket, bytes):
    total_sent = 0

    while total_sent < len(bytes):
        sent = socket.send(bytes[total_sent:])
        if sent == 0:
            raise Exception("This should not occur") # Según la docu, si se envian 0, la conexion está rota
        total_sent += sent

    return total_sent
