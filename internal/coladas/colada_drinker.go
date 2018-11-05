package coladas

import (
	"errors"
)

// Drinker models people who partake in the drinking of
// coladas.
type Drinker struct {
	UID          int64  `json:"uid"`
	Name         string `json:"name"`
	CanMake      bool   `json:"canMake"`
	HeadshotPath string `json:"headshotPath"`
}

// GetName returns the name of the drinker
func (c *Drinker) GetName() string {
	return c.Name
}

// MakesColadas shows if the drinker is also a maker
func (c *Drinker) MakesColadas() bool {
	return c.CanMake
}

// GetHeadshotPath returns the path to the person's headshot
func (c *Drinker) GetHeadshotPath() string {
	return c.HeadshotPath
}

// CreateColadaDrinker is a way to create new drinkers and hopefully new makers
// as well
func CreateColadaDrinker(uid int64, name string, canMake bool, headshot string) (*Drinker, error) {
	if name == "" {
		return nil, errors.New("the name parameter cannot be empty")
	}

	return &Drinker{
		UID:          uid,
		Name:         name,
		CanMake:      canMake,
		HeadshotPath: headshot,
	}, nil
}
