package v1

import (
	"fmt"
	"os/exec"
	"strings"
)

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797

func generalOut(in string) {
	fmt.Print(in)
}

func Out(OutputToScreen string) {

	for k, v := range ColourMap {
		OutputToScreen = strings.ReplaceAll(OutputToScreen, "["+k+"]", v)
	}
	generalOut(OutputToScreen)
}

func MoveTo(x int, y int) {
	generalOut(fmt.Sprintf("\033[%d;%dH", y, x))
}

func MoveRelative(x int, y int) {
	generalOut(fmt.Sprintf("\033[%d;%dH", y, x))
}

func MoveDown(line int) {
	generalOut(fmt.Sprintf("\033[%dB", line))
}

func MoveDownAndBeginning(line int) {
	generalOut("\n")
	// generalOut(fmt.Sprintf("\033[%dE", line))
}

func MoveY(y int) {
	generalOut(fmt.Sprintf("\033[%dA", y))
}

func MoveX(x int) {
	generalOut(fmt.Sprintf("\033[%dC", x))
}

/*
func GetCurrent() (int, int) {
    generalOut("\033[6n")

    var buf [16]byte
    var x, y int

    fd := int(os.Stdin.Fd())
    var readfds syscall.FdSet
    readfds.Bits[fd/64] |= 1 << (uint(fd) % 64)

    timeout := syscall.NsecToTimeval(100 * time.Millisecond.Nanoseconds())
    err := syscall.Select(fd+1, &readfds, nil, nil, &timeout)
    if err != nil {
        return 0, 0
    }

    n, _ := os.Stdin.Read(buf[:])
    fmt.Sscanf(string(buf[:n]), "\033[%d;%dR", &y, &x)

    return x, y
}

*/

func Clear() {

	// disable input buffering
	_ = exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	_ = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	generalOut("\033[2J")
}

func ShowCursor() {
	generalOut("\033[?25h")
}

func HideCursor() {
	generalOut("\033[?25l")

}

func Reset() {
	Out("[F-NORMAL][F-RESET]")
}
