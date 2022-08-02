package hiramekishare

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"

	"github.com/google/uuid"
)

type Idea struct {
	ID      uint32   `id`
	Members []string `members`
}

func (idea *Idea) Entry(m string) {
	idea.Members = append(idea.Members, m)
}

func (idea *Idea) Save() error {
	if idea.ID == 0 {
		u, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		idea.ID = u.ID()
	}

	s, err := json.Marshal(idea)
	if err != nil {
		return err
	}

	filename := strconv.Itoa(int(idea.ID)) + ".txt"
	return os.WriteFile(filename, s, 0600)
}

func (a *Idea) Delete() error {
	id := strconv.Itoa(int(a.ID))
	err := os.Remove(id + ".txt")
	return err
}

func (a *Idea) Equals(b *Idea) bool {
	equals := (a.ID == b.ID) && (reflect.DeepEqual(a.Members, b.Members))
	return equals
}

func LoadIdea(id int) (*Idea, error) {
	filename := strconv.Itoa(id) + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var idea Idea
	err = json.Unmarshal(body, &idea)
	return &idea, err
}
