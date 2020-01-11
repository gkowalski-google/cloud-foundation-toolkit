package launchpad

// diagram represents a GCP draw diagram
type diagram struct {
	paths []*diagramPath
	cards []*diagramCard
}

// diagramPath represents a path within a GCP diagram
type diagramPath struct {
	from      *card
	to        *card
	connector string
}

// diagramCard represents a card within a GCP diagram
type diagramCard struct {
	name string
}

// addCard adds a card into a diagram
func (d *diagram) addCard(displayName string) (error {
	card := diagramCard{
		name: name
	}
	d.cards = d.cards.append(card)
}

// draw generates diagram(s) for a given org
func (ao *assembledOrg) draw() error {
	orgDiagram := diagram{}

	ao.org.draw(orgDiagram)
}
