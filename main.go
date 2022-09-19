package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Fprintf(os.Stderr, "usage: go-new name\n")
		os.Exit(1)
	}
	name := os.Args[1]
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	directory := path + "/go-" + name
	_, err = os.Stat(directory)
	if os.IsNotExist(err) {
		os.Mkdir(directory, 0755)
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	main := `package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}
`
	err = ioutil.WriteFile(directory+"/main.go", []byte(main), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	mod := fmt.Sprintf(`module %s

go 1.19
`, name)
	err = ioutil.WriteFile(directory+"/go.mod", []byte(mod), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	readme := fmt.Sprintf(`# %s
`, name)
	err = ioutil.WriteFile(directory+"/README.md", []byte(readme), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
