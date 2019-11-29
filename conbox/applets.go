package main

import (
	"github.com/udhos/conbox/applets/cat"
	"github.com/udhos/conbox/applets/chmod"
	"github.com/udhos/conbox/applets/chown"
	"github.com/udhos/conbox/applets/cp"
	"github.com/udhos/conbox/applets/dd"
	"github.com/udhos/conbox/applets/echo"
	"github.com/udhos/conbox/applets/head"
	"github.com/udhos/conbox/applets/host"
	"github.com/udhos/conbox/applets/ls"
	"github.com/udhos/conbox/applets/mkdir"
	"github.com/udhos/conbox/applets/mv"
	"github.com/udhos/conbox/applets/printenv"
	"github.com/udhos/conbox/applets/ps"
	"github.com/udhos/conbox/applets/pwd"
	"github.com/udhos/conbox/applets/rm"
	"github.com/udhos/conbox/applets/rmdir"
	"github.com/udhos/conbox/applets/sh"
	"github.com/udhos/conbox/applets/shell"
	"github.com/udhos/conbox/applets/sleep"
	"github.com/udhos/conbox/applets/touch"
	"github.com/udhos/conbox/applets/which"
	"github.com/udhos/conbox/common"
)

func loadApplets() map[string]common.AppletFunc {
	tab := map[string]common.AppletFunc{
		"cat":      cat.Run,
		"chmod":    chmod.Run,
		"chown":    chown.Run,
		"cp":       cp.Run,
		"dd":       dd.Run,
		"echo":     echo.Run,
		"head":     head.Run,
		"host":     host.Run,
		"ls":       ls.Run,
		"mkdir":    mkdir.Run,
		"mv":       mv.Run,
		"printenv": printenv.Run,
		"pwd":      pwd.Run,
		"ps":       ps.Run,
		"rm":       rm.Run,
		"rmdir":    rmdir.Run,
		"sh":       sh.Run,
		"shell":    shell.Run,
		"sleep":    sleep.Run,
		"touch":    touch.Run,
		"which":    which.Run,
	}
	return tab
}
