package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	_ "embed"

	"github.com/urfave/cli"
)

//go:embed embeds/*.plist
var embed_plist []byte

//go:embed embeds/*.sh
var embed_shell []byte

var (
	VersionString = "MISSING build version [git hash]"

	plist_name = "com.utils.toggleairport"

	path_shell = "/Library/Scripts/toggleAirport.sh"
	path_plist = "/Library/LaunchAgents/" + plist_name + ".plist"
)

func init() {
	if runtime.GOOS != "darwin" {
		log.Printf("# Err: Plz Use with [MacOS]")
		os.Exit(1)
	}
	if currentUser, err := user.Current(); err != nil || currentUser.Username != "root" {
		log.Printf("# Err: Plz Use with [sudo], %+v", err)
		os.Exit(1)
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "switcher"
	app.Usage = "mac-auto-xlan-switch"
	app.Version = fmt.Sprintf("Git:[%s] (%s)", strings.ToUpper(VersionString), runtime.Version())
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start Auto xLan Switcher",
			Action: func(c *cli.Context) (err error) {
				if info, err := os.Stat(path_shell); err != nil || info.IsDir() {
					os.Remove(path_shell)
					if err = os.WriteFile(path_shell, embed_shell, 0o755); err != nil {
						return err
					}
				}
				log.Println("Installed: " + path_shell)

				if info, err := os.Stat(path_plist); err != nil || info.IsDir() {
					os.Remove(path_plist)
					if err = os.WriteFile(path_plist, embed_plist, 0o644); err != nil {
						return err
					}
					os.Chown(path_plist, os.Getuid(), os.Getgid())
					os.Chown(path_plist, 0, 0) // root wheels
				}
				log.Println("Installed: " + path_plist)

				var out []byte
				command := exec.Command("launchctl", "bootstrap", "system", path_plist)
				err = command.Run()

				command = exec.Command("launchctl", "list", plist_name)
				out, err = command.CombinedOutput()
				log.Println("Status:\n", string(out))
				return err
			},
		},
		{
			Name:  "stop",
			Usage: "Stop Auto xLan Switcher",
			Action: func(c *cli.Context) (err error) {
				if _, err = os.Stat(path_plist); err == nil {
					var out []byte
					command := exec.Command("launchctl", "list", plist_name)
					out, err = command.CombinedOutput()
					log.Println("Status:\n", string(out))

					command = exec.Command("launchctl", "bootout", "system", path_plist)
					err = command.Run()
				} else {
					err = nil
				}

				os.Remove(path_shell)
				log.Println("Removed: " + path_shell)

				os.Remove(path_plist)
				log.Println("Removed: " + path_plist)

				return err
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
