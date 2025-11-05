package memento

type Node interface {
	SetState(state string)
	GetState() string
}

type TextNode struct {
	state string
}

func (t *TextNode) SetState(state string) {
	t.state = state
}

func (t *TextNode) GetState() string {
	return t.state
}

func (t *TextNode) Save() Memento {
	return &TextMemento{
		state: t.state,
	}
}

type Memento interface {
	GetState() string
}
type TextMemento struct {
	state string
}

func (t TextMemento) GetState() string {
	return t.state
}

type Manage struct {
	states []Memento
}

func (m *Manage) Save(t Memento) {
	m.states = append(m.states, t)
}

func (m *Manage) Back(index int) Memento {
	return m.states[index]
}
