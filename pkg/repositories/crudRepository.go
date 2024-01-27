package repositories

type BaseRepository[T any] struct {

}

func (r *BaseRepository[T]) Create(model T) error {
	return nil
}

// type CrudRepository[T any] interface {
// 	Create(model T) error
// 	GetAll() ([]T, error)
// 	GetById(id string) (T, error)
// 	Update(id string, model T) error
// 	Delete(id string) error
// }

