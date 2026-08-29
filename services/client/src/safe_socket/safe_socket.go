package safe_socket

import (
	"io"
)

// TODO: Complete with a short-read/short-write tolerant implementation
const MAX_ATTEMPTS = 1

func SendAll(socket io.Writer, bytes []byte) error {
	written := 0
	attempts := MAX_ATTEMPTS

	for len(bytes)-written > 0 {
		n, err := socket.Write(bytes[written:])
		written += n

		if err == io.ErrShortWrite && MAX_ATTEMPTS > 0 { // Es considerado valido que no se hayan escrito bytes?
			// Buscar otros posibles errores que puedan necesitar volver a intentar
			// o en qué casos sería válido reintentar
			attempts -= 1
			continue
		}

		if err != nil {
			return err
		}
	}
	return nil
}

func RecvAll(socket io.Reader, size int) ([]byte, error) {
	buff := make([]byte, size)
	read := 0

	for size-read > 0 {
		n, err := socket.Read(buff[read:])
		read += n

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}

		if err != nil {
			return nil, err
		}
	}
	return buff[:read], nil
}
