package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/adrg/splash"
)

func isatty(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, syscall.TCGETS,
		uintptr(unsafe.Pointer(&termios)), 0, 0, 0)

	return err == 0
}

func color(s string, styles []splash.Style) string {
	for i, style := range styles {
		param := fmt.Sprintf("{%d}", i)
		value := fmt.Sprintf("%s%s", "{r}", style)
		s = strings.Replace(s, param, value, -1)
	}

	return strings.Replace(s, "{r}", splash.Reset.String(), -1)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Print(usage)
		return
	}

	styles := []splash.Style{}
	isTty := isatty(os.Stdout.Fd())

	for _, arg := range os.Args[1:] {
		switch strings.TrimSpace(arg) {
		case "-h", "--help":
			fmt.Print(help)
			return

		case "-v", "--version":
			fmt.Print(version)
			return

		default:
			style := splash.NewStyle()
			if isTty {
				style = splash.ParseStyle(arg)
			}

			styles = append(styles, style)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(color(scanner.Text(), styles))
	}

	fmt.Print(splash.Reset)
}
