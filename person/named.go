package person

type Named interface {
	GetName() string
	SetName(string) Named
}
