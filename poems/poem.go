package poem

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// A Stanza is a poetic stanza.
type Stanza []string

// A Poem contains metadata about a poem.
type Poem struct {
	Title  string
	Author string
	Year   int

	Stanzas []Stanza
}

// The Library contains the poems I know.
var Library = map[string]Poem{}

func readStanzas(text string) []Stanza {
	buf := bytes.NewBufferString(text)
	scanner := bufio.NewScanner(buf)

	var stanzas []Stanza
	var stanza Stanza

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			if len(stanza) != 0 {
				stanzas = append(stanzas, stanza)
				stanza = Stanza{}
			}
		} else {
			stanza = append(stanza, line)
		}
	}

	if len(stanza) != 0 {
		stanzas = append(stanzas, stanza)
	}

	return stanzas
}

func Titles() []string {
	var titles = make([]string, 0, len(Library))
	for _, poem := range Library {
		line := fmt.Sprintf("%s (%s, %d)", poem.Title, poem.Author, poem.Year)
		titles = append(titles, line)
	}

	sort.Strings(titles)
	return titles
}

func register(poem Poem) {
	var mapTitle = strings.ToLower(poem.Title)
	Library[mapTitle] = poem
}
