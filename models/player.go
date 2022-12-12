package models

type Player struct {
	name string
	email string
}

func (p *Player) GetName() string {
	if p != nil {
		return p.name
	}
	return ""
}

func (p *Player) GetEmail() string {
	if p != nil {
		return p.email
	}
	return ""
}

func CreatePlayer(name, email string) *Player {
    return &Player{
		name: name,
		email: email,
	}
}