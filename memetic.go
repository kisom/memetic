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

func checkPoem(p poem.Poem) {
	fmt.Printf("Selected: %s by %s (%d)\n", p.Title, p.Author, p.Year)

	for i := config.start; i < config.end; i++ {
		fmt.Println("Stanza", i)
		for j := 0; j < len(p.Stanzas[i]); j++ {
			if config.guide {
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

				if config.noResetStanza {
					j--
				} else {
					j = -1
				}
				continue
			}
		}
	}

	fmt.Println("COMPLETE")
}

var config struct {
	guide         bool
	noResetStanza bool
	start, end    int
}

func checkArgs(p poem.Poem) {
	if config.start < 1 {
		fmt.Fprintf(os.Stderr, "Can't start from a stanza less than 1.")
		os.Exit(1)
	}

	if config.end < 0 {
		fmt.Fprintf(os.Stderr, "Can't end from a stanza less than 0.")
		os.Exit(1)
	}

	if config.start > len(p.Stanzas) {
		config.start = len(p.Stanzas)
	}

	if config.start > len(p.Stanzas) {
		config.end = len(p.Stanzas)
	}
}

func main() {
	flag.BoolVar(&config.guide, "g", false, "guided mode")
	flag.BoolVar(&config.noResetStanza, "n", false, "don't reset from the beginning of the stanza")
	flag.IntVar(&config.start, "s", 1, "starting stanza (normal people style, not array style)")
	flag.IntVar(&config.end, "e", 0, "ending stanza (normal people style, not array style; 0 means end-of-poem)")
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

	if config.end == 0 {
		config.end = 1
	}

	checkArgs(p)

	config.start-- // Go from natural to array style.
	config.end--

	checkPoem(p)
}
