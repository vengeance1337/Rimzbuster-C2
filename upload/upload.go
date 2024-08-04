package upload

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

func UploadFile(conn net.Conn, localPath string, remoteDir string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Extract the file name from the local path
	fileName := filepath.Base(localPath)
	remotePath := filepath.Join(remoteDir, fileName)

	conn.Write([]byte(fmt.Sprintf("upload %s\n", remotePath)))

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		conn.Write(buf[:n])
	}

	conn.Write([]byte("EOF\n"))
	return nil
}