package progress

import (
	"fmt"
	"os"
)

func goToLine(y int) {
	// go to line
	line := fmt.Sprintf("\033[%d;%dH", y, 0)
	write(line)

	// erase line
	write("\033[2K")
}

func clearScreen() {
	write("\033[2J")
	write("\033[0;0H")
}

func write(s string) {
	os.Stdout.Write([]byte(s))
}
