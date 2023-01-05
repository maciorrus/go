package person

type Person struct {
	fname string
	lname string
	login string
	pass  string
	email string
}

func (p *Person) GetName() string {
	return p.fname
}

func (p *Person) SetName(name string) Named {
	p.fname = name
	return p
}
