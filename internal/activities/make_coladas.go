package activities

import (
	"math/rand"
	"time"

	"github.com/rippinrobr/lunch-n-learn/internal/coladas"
	"github.com/rippinrobr/lunch-n-learn/internal/history"
)

// Brew models the act of creating the coladas
type Brew struct {
	Barista string
}

// GetBarista returns the person who is making
// the next colada
func (b *Brew) GetBarista() string {
	return b.Barista
}

// PickNextBarista randomly selects the next prerson to brew
// the espresso.  The previous barista should not be picked
func (b *Brew) PickNextBarista(makers []*coladas.Drinker, lastDraw *history.LogEntry) *coladas.Drinker {
	// This filters out the previous barista since he/she can't make
	// them twice in a row
	brewers := filterBaristas(makers, func(drinker *coladas.Drinker) bool {
		return drinker.UID != lastDraw.BaristaID
	})

	numBrewers := len(brewers)
	rand.Seed(time.Now().UTC().UnixNano())
	idx := rand.Intn(numBrewers)

	return brewers[idx]
}

// filterBaristas takes a function that will be executed against each member of the slice
// if the comparison being done returns true then the brewer is added to to the potential
// brewers slice. THis is so that the person that just made a colada doesn't make it again
func filterBaristas(brewers []*coladas.Drinker, f func(*coladas.Drinker) bool) []*coladas.Drinker {
	potentialBrewers := make([]*coladas.Drinker, 0)
	for _, b := range brewers {
		if b.MakesColadas() && f(b) {
			potentialBrewers = append(potentialBrewers, b)
		}
	}
	return potentialBrewers
}
