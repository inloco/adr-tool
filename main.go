package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	newTeam "github.com/inloco/adr-tool/newteam"
)

var (
	cli = kingpin.New("adr", "Assists creating new ADRs at In Loco")

	newTeamCmd  = cli.Command("new-team", "Create a new directory with a copy of the template")
	newTeamCode = newTeamCmd.Arg("code", "Code that identifies your team. Max 6 chars.").Required().String()
)

func main() {
	switch kingpin.MustParse(cli.Parse(os.Args[1:])) {
	case newTeamCmd.FullCommand():
		newTeam.NewTeam(*newTeamCode)
	}
}
