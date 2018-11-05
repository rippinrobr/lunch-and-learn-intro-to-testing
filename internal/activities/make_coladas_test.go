package activities

import (
	"testing"

	"github.com/rippinrobr/lunch-n-learn/internal/drawing"

	"github.com/rippinrobr/lunch-n-learn/internal/coladas"
)

func TestPickNextBaristaShouldNotChooseThePreviousBarista(t *testing.T) {
	a, _ := coladas.CreateColadaDrinker(1, "a", true, "/img/a.png")
	b, _ := coladas.CreateColadaDrinker(2, "b", true, "/img/b.png")
	c, _ := coladas.CreateColadaDrinker(3, "c", true, "/img/c.png")
	coladas := []*coladas.Drinker{a, b, c}

	brew := Brew{}
	lastDraw := drawing.Result{
		Barista:    c.Name,
		BaristaID:  c.UID,
		BaristaImg: c.HeadshotPath,
	}
	barista := brew.PickNextBarista(coladas, &lastDraw)
	if barista.UID == lastDraw.BaristaID {
		t.Errorf("The newly selected barista '%s' was the previous barista '%s'", barista.Name, lastDraw.Barista)
	}
}

func TestPickNextBaristaDoesNotChooseSomeoneWhoDoesntMakeColadas(t *testing.T) {
	a, _ := coladas.CreateColadaDrinker(1, "a", true, "/img/a.png")
	b, _ := coladas.CreateColadaDrinker(2, "b", false, "/img/b.png")
	c, _ := coladas.CreateColadaDrinker(3, "c", true, "/img/c.png")
	d, _ := coladas.CreateColadaDrinker(4, "d", true, "/img/d.png")

	lastDraw := drawing.Result{
		Barista:    a.Name,
		BaristaID:  a.UID,
		BaristaImg: a.HeadshotPath,
	}
	coladas := []*coladas.Drinker{a, b, c, d}

	brew := Brew{}
	barista := brew.PickNextBarista(coladas, &lastDraw)
	if barista.UID == b.UID {
		t.Errorf("The newly selected barista '%s' should not have been selected, MakesColadas is false\n%+v\n", barista.Name, b)
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
