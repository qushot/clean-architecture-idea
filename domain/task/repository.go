package task

type Repository interface {
	Save(entity *Entity) error
	Delete(id ID) error
}
