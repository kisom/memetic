package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	poem "git.metacircular.net/kyle/memetic/poems"
)

func readLine(prompt string) (line string, err error) {
	fmt.Printf(prompt)
	rd := bufio.NewReader(os.Stdin)
	line, err = rd.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.TrimSpace(line)
	return
}

func checkPoem(guide bool, p poem.Poem) {
	fmt.Printf("Selected: %s by %s (%d)\n", p.Title, p.Author, p.Year)

	for i := 0; i < len(p.Stanzas); i++ {
		fmt.Println("Stanza", i)
		for j := 0; j < len(p.Stanzas[i]); j++ {
			if guide {
				fmt.Printf("%04d: %s\n", j+1, p.Stanzas[i][j])
			}

			line, err := readLine(fmt.Sprintf("%04d> ", j+1))
			if err != nil {
				if err == io.EOF {
					os.Exit(0)
				}
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(1)
			}

			if line != p.Stanzas[i][j] {
				fmt.Println(p.Stanzas[i][j])
				j--
				continue
			}
		}
	}

	fmt.Println("COMPLETE")
}

func main() {
	var guide bool
	flag.BoolVar(&guide, "g", false, "guided mode")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("Library:\n\n")
		for _, title := range poem.Titles() {
			fmt.Println("+", title)
		}
		fmt.Printf("\n")
		os.Exit(0)
	}

	selection := strings.ToLower(flag.Arg(0))
	p, ok := poem.Library[selection]
	if !ok {
		fmt.Fprintf(os.Stderr, "Title \"%s\" wasn't found in the library.\n", selection)
		os.Exit(1)
	}
	checkPoem(guide, p)
}
