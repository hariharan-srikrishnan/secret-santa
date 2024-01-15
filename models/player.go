package models

type Player struct {
	name string
	email string
}

func (p *Player) Name() string {
	if p != nil {
		return p.name
	}
	return ""
}

func (p *Player) Email() string {
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