package activities

import (
	"math/rand"
	"time"

	"github.com/rippinrobr/lunch-n-learn/internal/coladas"
)

// Clean tracks the coladas who will clean
// and cleaned previously
type Clean struct {
	Cleaner string
}

// GetCleaner returns the name of the person who
// is the next person to clean up after coladas
func (c *Clean) GetCleaner() string {
	return c.Cleaner
}

func (c *Clean) setCleaner(name string) {
	c.Cleaner = name
}

// PickNextCleaner determines who will clean the cups and other colada paraphernalia
func (c *Clean) PickNextCleaner(cleaners []*coladas.Drinker, barista, previousCleaner string) string {
	// This filters out the previous barista, current barista and prevoius cleaner since he/she can't make
	// them twice in a row
	potentialCleaners := filterCleaners(cleaners, func(name string) bool {
		return name != previousCleaner && name != barista
	})

	numCandidates := len(potentialCleaners)
	rand.Seed(time.Now().UTC().UnixNano())
	idx := rand.Intn(numCandidates)

	cleaner := potentialCleaners[idx]
	newCleanerName := cleaner.GetName()
	c.setCleaner(newCleanerName)

	return c.GetCleaner()
}

// filterCleaners takes a function that will be executed against each member of the slice
// if the comparison being done returns true then the cleaner is added to to the potential
// brewers slice. This is here to prevent a person from cleaning up back to back
func filterCleaners(cleaners []*coladas.Drinker, f func(string) bool) []*coladas.Drinker {
	potentialCleaners := make([]*coladas.Drinker, 0)
	for _, c := range cleaners {
		if f(c.GetName()) {
			potentialCleaners = append(potentialCleaners, c)
		}
	}
	return potentialCleaners
}
