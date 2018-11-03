package activities

import (
	"testing"

	"github.com/rippinrobr/lunch-n-learn/internal/coladas"
)

func TestPickNextCleanerShouldNotChooseThePreviousCleanerOrBarista(t *testing.T) {
	prevCleaner := "c"
	barista := "a"
	a, _ := coladas.CreateColadaDrinker(1, "a", true, "/img/a.png")
	b, _ := coladas.CreateColadaDrinker(2, "b", true, "/img/b.png")
	c, _ := coladas.CreateColadaDrinker(3, "c", true, "/img/c.png")
	d, _ := coladas.CreateColadaDrinker(4, "d", true, "/img/a.png")
	e, _ := coladas.CreateColadaDrinker(5, "e", true, "/img/b.png")
	f, _ := coladas.CreateColadaDrinker(6, "f", true, "/img/c.png")
	cleaners := []*coladas.Drinker{a, b, c, d, e, f}

	clean := Clean{}
	cleaner := clean.PickNextCleaner(cleaners, barista, prevCleaner)
	if cleaner == prevCleaner {
		t.Errorf("The newly selected cleaner '%s' was the previous cleaner '%s'", cleaner, prevCleaner)
	}

	if cleaner == barista {
		t.Errorf("The newly selected cleaner '%s' was the person who just made the colada '%s'", cleaner, barista)
	}
}
