package domain

type Branch struct {
	name    string
	surname string
}

type BranchView struct {
	Name    string
	Surname string
}

func (b *Branch) setName(name string) {
	b.name = name
}

func (b *Branch) Name() string {
	return b.name
}

func (b *Branch) Surname() string {
	return "TEST"
}

func (b *Branch) AsView() BranchView {
	return BranchView{
		b.name,
		"TEST",
	}
}

func newBranch(name string) Branch {
	return Branch{
		name: name,
	}
}
