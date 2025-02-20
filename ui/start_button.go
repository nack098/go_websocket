package ui

type startButton struct {
	name string
}

func (b *startButton) render() {
}

func (b *startButton) action() {
}

func (b *startButton) getName() *string {
	return &b.name
}
