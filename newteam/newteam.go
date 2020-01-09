package newteam

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	templateFilePath  = "_base/template.md"
	newTemplateFormat = "%s/template.md"
)

func NewTeam(teamCode string) {
	if len(teamCode) > 6 {
		fmt.Println("Code is too long! Keep it to a max of 6 characters")
		return
	}

	teamCode = strings.ToUpper(teamCode)

	templateBytes, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		fmt.Printf("error reading template file (are you in the adr repo?): %s", err)
		return
	}

	err = os.Mkdir(teamCode, 0755)
	if err != nil {
		fmt.Printf("error making a dir: %s", err)
		return
	}

	newTemplatePath := fmt.Sprintf(newTemplateFormat, teamCode)

	err = ioutil.WriteFile(newTemplatePath, templateBytes, 0644)
	if err != nil {
		fmt.Printf("error writing new template: %s", err)
		return
	}
}
