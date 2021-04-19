package task

import "fmt"

type Entity struct {
	ID          ID
	Name        Name
	Description Description
}

func NewEntity(name Name, description Description) (*Entity, error) {
	id := ID("id")

	entity := &Entity{
		ID:          id,
		Name:        name,
		Description: description,
	}

	if err := entity.validate(); err != nil {
		return nil, err
	}

	return entity, nil
}

func (r *Entity) validate() error {
	if err := r.ID.validate(); err != nil {
		return err
	}

	if err := r.Name.validate(); err != nil {
		return err
	}

	if err := r.Description.validate(); err != nil {
		return err
	}

	return nil
}

func (r *Entity) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Description: %s", r.ID, r.Name, r.Description)
}
