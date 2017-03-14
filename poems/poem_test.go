package poem

import (
	"fmt"
	"os"
	"testing"
)

func cmp(s []Stanza, o []Stanza) bool {
	if len(s) != len(o) {
		fmt.Fprintf(os.Stderr, "Lengths don't match: %d stanzas v. %d stanzas.\n",
			len(s), len(o))
		return false
	}

	for i := range s {
		if len(s[i]) != len(o[i]) {
			fmt.Fprintf(os.Stderr, "Stanza %d lengths don't match: %d lines v. %d lines.\n",
				i, len(s[i]), len(o[i]))
			return false
		}

		for j := range s[i] {
			if s[i][j] != o[i][j] {
				fmt.Fprintf(os.Stderr, "Line %d in stanza %d doesn't match:\n\t%s\n\t%s\n",
					j, i, s[i][j], o[i][j])
				return false
			}
		}
	}
	return true
}

func TestNewStanza(t *testing.T) {
	var text = `Do not go gentle into that good night,
Old age should burn and rave at close of day;
Rage, rage against the dying of the light.

Though wise men at their end know dark is right,
Because their words had forked no lightning they
Do not go gentle into that good night.

Though wise men at their end know dark is right,
Because their words had forked no lightning they
Do not go gentle into that good night.

Good men, the last wave by, crying how bright
Their frail deeds might have danced in a green bay,
Rage, rage against the dying of the light.
`

	expected := []Stanza{
		[]string{
			"Do not go gentle into that good night,",
			"Old age should burn and rave at close of day;",
			"Rage, rage against the dying of the light.",
		},
		[]string{
			"Though wise men at their end know dark is right,",
			"Because their words had forked no lightning they",
			"Do not go gentle into that good night.",
		},
		[]string{
			"Though wise men at their end know dark is right,",
			"Because their words had forked no lightning they",
			"Do not go gentle into that good night.",
		},
		[]string{
			"Good men, the last wave by, crying how bright",
			"Their frail deeds might have danced in a green bay,",
			"Rage, rage against the dying of the light.",
		},
	}

	stanzas := readStanzas(text)
	if !cmp(expected, stanzas) {
		fmt.Fprintf(os.Stderr, "%v\n", stanzas)
		t.Fail()
	}

}
