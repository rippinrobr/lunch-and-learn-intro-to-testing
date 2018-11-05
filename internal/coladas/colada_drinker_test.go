package coladas

import (
	"testing"
)

func TestCreateColadaDrinker(t *testing.T) {
	name := "tester mctestface"
	imgPath := "/imgs/tester.png"
	canMake := true
	drinker, err := CreateColadaDrinker(1, name, canMake, imgPath)

	if err != nil {
		t.Errorf("CreateColadaDrinker err was not nil! err: '%s'", err.Error())
	}

	if drinker.Name != name {
		t.Errorf("drinker.name '%s' does not equal name: '%s'", drinker.Name, name)
	}

	if drinker.CanMake != canMake {
		t.Errorf("drinker.canMake '%t' does not equal canMake: '%t'", drinker.CanMake, canMake)
	}

	if drinker.HeadshotPath != imgPath {
		t.Errorf("drinker.headshotPath '%s' does not equal imgPath: '%s'", drinker.HeadshotPath, imgPath)
	}
}

func TestShouldThrowErrorWithEmptyName(t *testing.T) {
	_, err := CreateColadaDrinker(1, "", false, "/a/path")

	if err == nil {
		t.Errorf("CreateColadaDrinker should throw an error when empty name is provided")
	}
}

func TestGetNameShouldEqualDrinkersName(t *testing.T) {
	name := "tester mctestface"
	imgPath := "/imgs/tester.png"
	canMake := true

	drinker, err := CreateColadaDrinker(1, name, canMake, imgPath)
	if err != nil {
		t.Errorf("An error occurred while creating a drinker")
	}

	if drinker.GetName() != name {
		t.Errorf("drinker.GetName(): '%s' doesn't equal name: '%s'", drinker.GetName(), name)
	}

}

func TestMakesColadasShouldEqualCanMake(t *testing.T) {
	name := "tester mctestface"
	imgPath := "/imgs/tester.png"
	canMake := true

	drinker, err := CreateColadaDrinker(1, name, canMake, imgPath)
	if err != nil {
		t.Errorf("An error occurred while creating a drinker")
	}

	if drinker.MakesColadas() != canMake {
		t.Errorf("drinker.MakesColadas(): '%t' doesn't equal name: '%t'", drinker.MakesColadas(), canMake)
	}

}

func TestGetHeadshotPathShouldEqualDrinkersHeadshotPath(t *testing.T) {
	name := "tester mctestface"
	imgPath := "/imgs/tester.png"
	canMake := true

	drinker, err := CreateColadaDrinker(1, name, canMake, imgPath)
	if err != nil {
		t.Errorf("An error occurred while creating a drinker")
	}

	if drinker.GetHeadshotPath() != imgPath {
		t.Errorf("drinker.GetHeadshotPath(): '%s' doesn't equal imgPath: '%s'", drinker.GetHeadshotPath(), imgPath)
	}

}
