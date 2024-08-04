package banner

import "fmt"

func PrintBanner() {

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

⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⠴⠖⠛⠛⠛⠛⠛⠛⠟⠛⠛⠛⠛⠛⠛⠻⠿⢿⣿⠽⠽⠿⢷⣒⠦⢄⣀⣀⣀⣀⣀⣀⡀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⠞⠉⠀⠀⠀⠀⠩⡑⠈⢓⠀⠃⠒⡐⢆⡐⣃⣒⠀⢠⠏⡟⠀⡀⠀⢀⢈⢹⣛⢮⡻⣏⡝⢿⡿⢿⡇⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⠋⠀⠀⢀⣄⠀⠀⢂⠀⠨⠀⢄⢤⢄⢠⠀⢀⠀⠊⠠⣠⠏⣸⡗⡄⡄⠄⠀⢈⠋⢧⣀⡹⣎⡷⣿⣀⣁⣃⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠂⠄⠀⠀⠀⠀⠀⣈⣀⣀⣒⣤⣴⠾⢯⠦⠴⠤⠤⠴⠥⢤⣥⣬⡯⠡⢓⠄⠤⠄⠓⣁⣈⣁⣉⣘⣞⣓⣟⣧⣔⣅⡈⠹⠖⠚⠛⠳⣍⣀⣀⡤⠤⠝⠛⢷⡀⠀
⠐⢀⠀⠀⠀⠀⠀⢠⣀⣭⠤⠶⠛⠋⠉⠙⠋⠁⠀⠀⠀⠀⠀⡉⠉⠃⠌⠉⡉⠷⣬⠀⣀⡨⢤⠴⣞⠚⠛⢉⣉⠀⠀⠈⡇⠙⠋⠀⣀⣀⠤⠴⠒⠋⠋⡇⠁⠁⠀⣀⢀⡈⢧⠀
⠂⠀⠀⠀⣠⣶⣾⡿⠋⠇⠀⠀⠁⠀⠀⠀⠀⠀⠠⠂⠀⠀⠂⠘⠈⣤⣀⣬⠾⠓⢟⣉⠤⣿⣿⣢⠵⠛⠉⠙⡈⡳⣍⣡⡷⠖⠉⠉⣥⢂⠀⠀⠐⠀⢀⠗⠀⠀⣤⡿⢿⣿⣾⣿
⠀⠂⣴⢿⡽⠋⢹⠒⢶⠓⠒⠒⡻⠮⠦⡧⠭⠐⠨⡭⡭⠬⠔⡛⠋⠙⣨⡬⣖⠏⣻⣥⡾⡛⠏⣑⣵⣷⣯⣗⣦⡀⠻⡀⢰⣞⡀⣀⠂⡀⡀⣀⡐⡀⣸⠄⠀⢰⣿⣿⡿⣿⣿⠸
⢠⡾⣩⡿⣡⣴⣶⣗⣒⣂⣬⣴⣭⣯⣏⣩⡡⣴⣔⠤⠄⠴⠀⣤⡾⣻⡷⡟⣿⣷⠟⣁⣄⣠⣴⣿⣿⣿⡋⠻⣷⣵⡠⢹⣸⠄⠀⠀⠁⠂⠐⠀⠂⢂⡇⠀⠀⢸⣿⣿⣧⣿⡽⠀
⢸⣷⣿⣼⡟⠙⣻⠿⠿⡁⠀⢱⢾⠿⠿⠿⠿⠿⢛⡖⠂⣄⡾⠷⠶⢞⠛⠛⣿⠟⠉⠁⠀⣺⣿⣧⣾⣞⣯⠱⣾⣿⣇⠀⢿⠀⠀⠀⠂⠀⠀⠀⠀⣼⣀⡀⠀⣿⣿⣿⣯⣿⡇⢰
⣸⣶⣿⣿⠉⢉⠓⠚⢛⠺⠟⠗⡳⠿⠿⡤⠾⠴⣾⣷⠈⠀⡇⠀⠙⠀⠆⣸⡟⠀⠂⠅⢐⣿⡇⣦⢹⣷⢿⣀⢚⣻⡇⠀⠺⣄⠐⣀⣢⢤⠴⣞⣫⣯⠤⠗⠒⡿⣿⣻⣿⡏⠀⠘
⣸⣻⣿⣿⣿⣿⠿⣿⣷⣧⣦⣤⣬⣤⣤⣤⠵⢄⣸⣿⠈⠉⢻⣥⣅⣤⣴⣾⡇⠀⠀⠄⢸⡟⣷⣿⣿⣿⠺⠛⢭⣿⡇⠀⢀⣏⠭⠽⠚⣛⣉⡥⠶⠒⠚⠛⠉⠁⠸⠿⠿⠀⠀⠀
⢻⠠⢽⣿⡻⠿⣾⡿⣿⡿⣿⣿⣿⣿⣿⣿⡿⡀⣸⣿⠀⠀⠸⠿⡿⠿⠿⠋⡇⢀⡰⣠⣾⡇⡿⣋⣀⣏⣻⡦⣼⣾⠓⠋⣁⣴⠴⠛⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠸⠿⣿⣿⣾⣮⣯⣭⣿⣛⣻⣻⣿⠿⠿⠥⠼⢿⠛⣿⣧⣶⣶⣿⣥⣭⣉⡹⠽⠟⠋⠁⣸⡇⢷⢿⡏⢸⡆⢁⣼⠛⠚⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠉⠉⠙⠛⠓⠛⠿⠿⠿⠿⢿⣿⣾⣿⣿⣟⣿⣷⣖⣀⣀⣐⣐⣐⣖⣊⣩⣉⣍⣳⠘⢦⡣⠬⣵⣾⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠁⠉⠏⠩⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠛⠓⠒⠛⠒⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀

𝙰 𝙲𝚘𝚖𝚖𝚊𝚗𝚍 𝚊𝚗𝚍 𝙲𝚘𝚗𝚝𝚛𝚘𝚕 𝙵𝚛𝚊𝚖𝚎𝚠𝚘𝚛𝚔 𝚝𝚘 𝚌𝚘𝚗𝚝𝚛𝚘𝚕 𝚑𝚊𝚌𝚔𝚎𝚍 𝚌𝚊𝚛𝚜 𝚛𝚎𝚖𝚘𝚝𝚎𝚕𝚢.

								𝙱𝚞𝚒𝚕𝚝 𝚋𝚢: 𝚅𝚊𝚕𝚎𝚛𝚒𝚊𝚗 𝙶𝚛𝚒𝚏𝚏𝚒𝚝𝚑𝚜
								𝙶𝚒𝚝𝚑𝚞𝚋: 𝚑𝚝𝚝𝚙𝚜://𝚐𝚒𝚝𝚑𝚞𝚋.𝚌𝚘𝚖/𝚟𝚎𝚗𝚐𝚎𝚊𝚗𝚌𝚎
`+ reset)
}