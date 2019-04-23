package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/udhos/conbox/applets/cat"
	"github.com/udhos/conbox/applets/echo"
)

const conboxVersion = "0.0"

func main() {

	appletTable := loadApplets()

	// 1. try basename
	appletName := filepath.Base(os.Args[0])
	if applet, found := appletTable[appletName]; found {
		run(applet, os.Args[1:])
		return
	}

	if appletName != "conbox" {
		showVersion()
		fmt.Printf("conbox: basename: applet '%s' not found\n", appletName)
		return
	}

	// 2. try arg 1
	if len(os.Args) > 1 {
		arg := os.Args[1]
		switch arg {
		case "-h":
			usage(appletTable)
			return
		case "-l":
			listApplets(appletTable, "\n")
			return
		}
		appletName = arg
		if applet, found := appletTable[appletName]; found {
			run(applet, os.Args[2:])
			return
		}
		showVersion()
		fmt.Printf("conbox: arg 1: applet '%s' not found\n", appletName)
	}

	usage(appletTable)
	os.Exit(1)
}

func usage(tab map[string]appletFunc) {
	fmt.Println("usage: conbox APPLET [ARG]... : run APPLET")
	fmt.Println("       conbox -h              : show command-line help")
	fmt.Println("       conbox -l              : list applets")
	fmt.Println()
	registeredApplets(tab)
}

func registeredApplets(tab map[string]appletFunc) {
	fmt.Println("conbox: registered applets:")
	listApplets(tab, " ")
	fmt.Println()
}

func listApplets(tab map[string]appletFunc, sep string) {
	for n := range tab {
		fmt.Printf("%s%s", n, sep)
	}
}

func showVersion() {
	fmt.Printf("conbox: version %s runtime %s GOMAXPROC=%d OS=%s ARCH=%s\n", conboxVersion, runtime.Version(), runtime.GOMAXPROCS(0), runtime.GOOS, runtime.GOARCH)
}

func run(applet appletFunc, args []string) {
	exitCode := applet(args)
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

type appletFunc func(args []string) int

func loadApplets() map[string]appletFunc {
	tab := map[string]appletFunc{}
	tab["cat"] = cat.Run
	tab["echo"] = echo.Run
	return tab
}
