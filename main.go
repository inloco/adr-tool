package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/inloco/adr-tool/new"
	newTeam "github.com/inloco/adr-tool/newteam"
)

var (
	cli = kingpin.New("adr", "Assists creating new ADRs at In Loco")

	newTeamCmd  = cli.Command("new-team", "Create a new directory with a copy of the template")
	newTeamCode = newTeamCmd.Arg("code", "Code that identifies your team. Max 6 chars.").Required().String()

	newCmd      = cli.Command("new", "Create a new ADR")
	newCmdTeam  = newCmd.Flag("team", "Code that identifies your team.").String()
	newCmdTitle = newCmd.Arg("title", "Title of your new ADR.").Required().String()
)

func main() {
	switch kingpin.MustParse(cli.Parse(os.Args[1:])) {
	case newTeamCmd.FullCommand():
		newTeam.NewTeam(*newTeamCode)
	case newCmd.FullCommand():
		new.New(*newCmdTeam, *newCmdTitle)
	}
}
