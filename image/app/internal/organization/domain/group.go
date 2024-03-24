package domain

type Group struct {
	name string
}

func (group *Group) setName(name string) {
	group.name = name
}

func newGroup(name string) Group {
	return Group{
		name: name,
	}
}
