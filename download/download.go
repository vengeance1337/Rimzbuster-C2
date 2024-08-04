package download

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func DownloadFile(conn net.Conn, remotePath string, localPath string) error {
	// Determine the final file path
	var finalPath string
	if stat, err := os.Stat(localPath); err == nil && stat.IsDir() {
		// If localPath is a directory, append the file name from remotePath
		fileName := filepath.Base(remotePath)
		finalPath = filepath.Join(localPath, fileName)
	} else {
		// If localPath is a file path, use it as is
		finalPath = localPath
	}

	// Send the download request to the client
	conn.Write([]byte(fmt.Sprintf("download %s\n", remotePath)))

	file, err := os.Create(finalPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to receive file: %v", err)
		}

		content := string(buf[:n])
		if strings.Contains(content, "EOF\n") {
			content = strings.Replace(content, "EOF\n", "", 1)
			if len(content) > 0 {
				file.Write([]byte(content))
			}
			break
		}

		file.Write(buf[:n])
	}

	return nil
}