package hiramekishare

import "testing"

func TestLoadIdea(t *testing.T) {
	idea, err := LoadIdea(123)
	if err != nil {
		t.Fatal(err)
	}

	if idea == nil {
		t.Fatal("idea is nil.")
	}
}

func TestIdeaEntry(t *testing.T) {
	idea := Idea{}
	m := "Grime"
	idea.Entry(m)
	if idea.Members[0] != m {
		t.Fail()
	}
}

func TestIdeaSave(t *testing.T) {
	idea := Idea{
		Members: []string{"Mercy", "Sasha", "Ann"},
	}

	err := idea.Save()
	if err != nil {
		t.Fatal(err)
	}

	savedIdea, err := LoadIdea(int(idea.ID))
	if err != nil {
		t.Fatal(err)
	}

	if !idea.Equals(savedIdea) {
		t.Fail()
	}

	err = idea.Delete()
	if err != nil {
		t.Fatal(err)
	}
}

func TestIdeaEquals(t *testing.T) {
	a := &Idea{
		ID:      123456789,
		Members: []string{"Sprig", "Polly", "Hop Pop"},
	}
	b := &Idea{
		ID:      123456789,
		Members: []string{"Sprig", "Polly", "Hop Pop"},
	}

	if !a.Equals(b) {
		t.Fail()
	}
}

func TestIdeaDelete(t *testing.T) {
	i := &Idea{}
	err = i.Save()
	if err != nil {
		t.Fatal(err)
	}

	err = i.Delete()
	if err != nil {
		t.Fatal(err)
	}
}
