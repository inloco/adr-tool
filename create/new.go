package create

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"

	"github.com/Machiel/slugify"
)

const (
	newADRFormat       = "%s/%s-%s.md"
	templateFileFormat = "%s/template.md"

	titleRegex  = "^#[^#].*\n"
	statusRegex = "\\*[ ]Status.*\\n"
	dateRegex   = "\\*[ ]Date.*\\n"

	titleFormat  = "# %s\n"
	statusFormat = "* Status: proposed\n"
	dateFormat   = "* Date: 2006-01-02\n"
)

func extractNumber(s string) (int, error) {
	re := regexp.MustCompile("^[0-9]+")
	number := re.FindString(s)

	return strconv.Atoi(number)
}

func getNextADRNumber(team string) (string, error) {
	files, err := ioutil.ReadDir(team)
	if err != nil {
		return "", err
	}

	highest := 0
	for _, f := range files {
		if i, err := extractNumber(f.Name()); err == nil && i > highest {
			highest = i
		}
	}

	number := highest + 1
	paddedNumber := fmt.Sprintf("%04d", number)

	return paddedNumber, nil
}

func getTemplate(team string) (string, error) {
	path := fmt.Sprintf(templateFileFormat, team)

	templateBytes, err := ioutil.ReadFile(path)
	return string(templateBytes), err
}

func replaceInTemplate(template string, title string) string {
	t := regexp.MustCompile(titleRegex)
	mdTitle := fmt.Sprintf(titleFormat, title)
	template = t.ReplaceAllString(template, mdTitle)

	s := regexp.MustCompile(statusRegex)
	mdStatus := statusFormat
	template = s.ReplaceAllString(template, mdStatus)

	d := regexp.MustCompile(dateRegex)
	mdDate := time.Now().Format(dateFormat)
	template = d.ReplaceAllString(template, mdDate)

	return template
}

func Create(team string, title string) {
	slug := slugify.Slugify(title)
	number, err := getNextADRNumber(team)
	if err != nil {
		fmt.Println(err)
		return
	}

	path := fmt.Sprintf(newADRFormat, team, number, slug)

	template, err := getTemplate(team)
	if err != nil {
		fmt.Printf("error getting template: %s\n", err)
		return
	}

	adr := fmt.Sprintf(replaceInTemplate(template, title))

	err = ioutil.WriteFile(path, []byte(adr), 0644)
	if err != nil {
		fmt.Printf("error writing new template: %s", err)
		return
	}
}
