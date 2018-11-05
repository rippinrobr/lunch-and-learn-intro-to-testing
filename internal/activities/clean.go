package activities

import (
	"math/rand"
	"time"

	"github.com/rippinrobr/lunch-n-learn/internal/coladas"
	"github.com/rippinrobr/lunch-n-learn/internal/history"
)

// Clean tracks the coladas who will clean
// and cleaned previously
type Clean struct {
	//Cleaner *coladas.Drinker
}

// PickNextCleaner determines who will clean the cups and other colada paraphernalia
func (c *Clean) PickNextCleaner(cleaners []*coladas.Drinker, barista *coladas.Drinker, lastDraw *history.DrawingResult) *coladas.Drinker {
	// This filters out the previous barista, current barista and prevoius cleaner since he/she can't make
	// them twice in a row
	potentialCleaners := filterCleaners(cleaners, func(drinker *coladas.Drinker) bool {
		return drinker.UID != lastDraw.CleanerID && drinker.UID != barista.UID
	})

	numCandidates := len(potentialCleaners)
	rand.Seed(time.Now().UTC().UnixNano())
	idx := rand.Intn(numCandidates)

	return potentialCleaners[idx]
}

// filterCleaners takes a function that will be executed against each member of the slice
// if the comparison being done returns true then the cleaner is added to to the potential
// brewers slice. This is here to prevent a person from cleaning up back to back
func filterCleaners(cleaners []*coladas.Drinker, f func(*coladas.Drinker) bool) []*coladas.Drinker {
	potentialCleaners := make([]*coladas.Drinker, 0)
	for _, c := range cleaners {
		if f(c) {
			potentialCleaners = append(potentialCleaners, c)
		}
	}
	return potentialCleaners
}
