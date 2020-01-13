package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/inloco/adr-tool/create"
	createTeam "github.com/inloco/adr-tool/createteam"
)

var (
	cli = kingpin.New("adr", "Assists creating new ADRs at In Loco")

	createTeamCmd  = cli.Command("create-team", "Create a new directory with a copy of the template")
	createTeamCode = createTeamCmd.Arg("code", "Code that identifies your team. Max 6 chars.").Required().String()

	createCmd      = cli.Command("create", "Create a new ADR")
	createCmdTeam  = createCmd.Flag("team", "Code that identifies your team.").String()
	createCmdTitle = createCmd.Arg("title", "Title of your new ADR.").Required().String()
)

func main() {
	switch kingpin.MustParse(cli.Parse(os.Args[1:])) {
	case createTeamCmd.FullCommand():
		createTeam.CreateTeam(*createTeamCode)
	case createCmd.FullCommand():
		create.Create(*createCmdTeam, *createCmdTitle)
	}
}
