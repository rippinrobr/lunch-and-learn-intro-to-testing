package activities

import (
	"testing"

	"github.com/rippinrobr/lunch-n-learn/internal/coladas"
)

func TestPickNextBaristaShouldNotChooseThePreviousBarista(t *testing.T) {
	prevBarista := "c"
	a, _ := coladas.CreateColadaDrinker(1, "a", true, "/img/a.png")
	b, _ := coladas.CreateColadaDrinker(2, "b", true, "/img/b.png")
	c, _ := coladas.CreateColadaDrinker(3, "c", true, "/img/c.png")
	coladas := []*coladas.Drinker{a, b, c}

	brew := Brew{}
	barista := brew.PickNextBarista(coladas, prevBarista)
	if barista == prevBarista {
		t.Errorf("The newly selected barista '%s' was the previous barista '%s'", barista, prevBarista)
	}
}

func TestPickNextBaristaDoesNotChooseSomeoneWhoDoesntMakeColadas(t *testing.T) {
	prevBarista := "a"
	a, _ := coladas.CreateColadaDrinker(1, "a", true, "/img/a.png")
	b, _ := coladas.CreateColadaDrinker(2, "b", false, "/img/b.png")
	c, _ := coladas.CreateColadaDrinker(3, "c", true, "/img/c.png")
	d, _ := coladas.CreateColadaDrinker(4, "d", true, "/img/d.png")

	coladas := []*coladas.Drinker{a, b, c, d}

	brew := Brew{}
	barista := brew.PickNextBarista(coladas, prevBarista)
	if barista == b.GetName() {
		t.Errorf("The newly selected barista '%s' should not have been selected, MakesColadas is false\n%+v\n", barista, b)
	}

}

func TestGetBaristaReturnsTheBaristasName(t *testing.T) {
	barista := "testing"
	b := Brew{
		Barista: barista,
	}

	if barista != b.GetBarista() {
		t.Errorf("GetBarista() should have returned '%s' but returned '%s' instead", barista, b.GetBarista())
	}
}
