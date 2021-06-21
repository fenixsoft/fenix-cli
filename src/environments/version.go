package environments

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"io"
)

// DO NOT modify this, will inject form Git tags by CI/CD
var Version = "Development Build"

func Logo(args []string, out io.Writer) {
	logo := "\n" +
		"[dark_gray] ________  ________  ____  _____  _____  ____  ____        ______  _____     _____  \n" +
		"[dark_gray]|_   __  ||_   __  ||_   \\|_   _||_   _||_  _||_  _|      .' ___  ||_   _|   |_   _| \n" +
		"[light_gray]  | |_ \\_|  | |_ \\_|  |   \\ | |    | |    \\ \\  / /______ / .'   \\_|  | |       | |   \n" +
		"[light_gray]  |  _|     |  _| _   | |\\ \\| |    | |     > `' <|______|| |         | |   _   | |   \n" +
		"[reset] _| |_     _| |__/ | _| |_\\   |_  _| |_  _/ /'`\\ \\_      \\ `.___.'\\ _| |__/ | _| |_  \n" +
		"[white]|_____|   |________||_____|\\____||_____||____||____|      `.____ .'|________||_____|\n[reset]"

	logo += fmt.Sprintf("%84v\n", Version) +
		"                                              [blue][underline]https://github.com/fenixsoft/fenix-cli[reset]\n" +
		"                                              Press key [yellow]<F1>[reset] to get help information\n" +
		"                                                                                     \n"
	_, _ = colorstring.Println(logo)
}
