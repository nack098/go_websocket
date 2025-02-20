package ui

type menuItem interface {
	render()
	action()
	getName() *string
}
