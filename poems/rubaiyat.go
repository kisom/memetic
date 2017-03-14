package poem

var rubaiyat = Poem{
	Title:  "The Rubaiyat of Omar Khayyam",
	Author: "Edward Fitzgerald",
	Year:   1889,
	Stanzas: readStanzas(`WAKE! For the Sun, who scatter'd into flight
 The Stars before him from the Field of Night,
   Drives Night along with them from Heav'n, and strikes
 The Sultan's Turret with a Shaft of Light.
`),
}

func init() {
	register(rubaiyat)
}
