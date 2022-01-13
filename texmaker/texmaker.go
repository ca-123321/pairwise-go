package texmaker

import (
	"fmt"
	"log"
	"os"

	"github.com/rwestlund/gotex"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Example of how to use interpolation in the contents of the pdf
func MakeContents() string {
	mycard := "card"

	var document = `
        \documentclass[12pt]{article}
        \begin{document}
        %s
        \end{document}
        `
	return fmt.Sprintf(document, mycard)
}

func MakeTeX() {
	contents := MakeContents()

	var pdf, err = gotex.Render(contents, gotex.Options{
		Command:   "/usr/bin/pdflatex",
		Runs:      1,
		Texinputs: "/my/asset/dir:/my/other/asset/dir"})

	if err != nil {
		log.Println("render failed ", err)
	} else {
		d1 := []byte(pdf)
		// TODO: create ./etc/, ask user if it's fine
		err := os.WriteFile("etc/test.pdf", d1, 0644)
		check(err)
	}
}
