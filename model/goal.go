package model


type Goal struct {
	ID uint
	Name string
	Description string
	Mark bool
}

// Get done from all actions
func (g Goal) done() {

}
