package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	emailFlag := flag.String("email", "<AUTHOR_EMAIL>", "Your email here to be put in the code of conduct document.")
	versionFlag := flag.String("version", "", "Version of the contributor covenant. Only English 1.4 and 2.0 is supported. Default is 2.0")

	flag.Parse()

	if emailFlag == nil {
		panic(errors.New("you have not input your email yet"))
	}

	if *versionFlag == "" {
		*versionFlag = "2.0"
	}

	var covenantTemplate string
	var mdFile []byte

	switch *versionFlag {
	case "1.4":
		mdFile, err := ioutil.ReadFile("./template-1.4.md")
		if err != nil {
			panic(err.Error())
		}

		covenantTemplate = string(mdFile)
		covenantTemplate = strings.Replace(covenantTemplate, "[INSERT EMAIL ADDRESS]", *emailFlag, 1)

	case "2.0":
		mdFile, err := ioutil.ReadFile("./template-2.0.md")
		if err != nil {
			panic(err.Error())
		}

		covenantTemplate = string(mdFile)
		covenantTemplate = strings.Replace(covenantTemplate, "[INSERT CONTACT METHOD]", *emailFlag, 1)
	}

	mdFile = []byte(covenantTemplate)
	ioutil.WriteFile("./CODE_OF_CONDUCT.md", mdFile, os.FileMode(0600))

	fmt.Println("Thank you for making a better open source community by adhering to the Contributor Covenant in your project! 😊")
}
