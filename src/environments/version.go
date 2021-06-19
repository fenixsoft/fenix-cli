package environments

import (
	"fmt"
	"io"
)

// DO NOT modify this, will inject form Git tags by CI/CD
var Version = "Development Build"

func Logo(args []string, out io.Writer) {
	_, _ = out.Write([]byte("\n ________  ________  ____  _____  _____  ____  ____         ______  _____     _____  \n" +
		"|_   __  ||_   __  ||_   \\|_   _||_   _||_  _||_  _|      .' ___  ||_   _|   |_   _| \n" +
		"  | |_ \\_|  | |_ \\_|  |   \\ | |    | |    \\ \\  / /______ / .'   \\_|  | |       | |   \n" +
		"  |  _|     |  _| _   | |\\ \\| |    | |     > `' <|______|| |         | |   _   | |   \n" +
		" _| |_     _| |__/ | _| |_\\   |_  _| |_  _/ /'`\\ \\_      \\ `.___.'\\ _| |__/ | _| |_  \n" +
		"|_____|   |________||_____|\\____||_____||____||____|      `.____ .'|________||_____| \n" +
		fmt.Sprintf("%84v\n", Version) +
		"                                              https://github.com/fenixsoft/fenix-cli\n" +
		"                                              Press key <F1> to get help information\n" +
		"                                                                                     \n"))
}
