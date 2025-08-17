package task

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(task *Task) error {
	return r.db.Create(task).Error
}

func (r *Repository) Update(task *Task) error {
	result := r.db.Model(task).
		Where("id = ?", task.ID).
		Updates(map[string]interface{}{
			"title": task.Title,
		})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) Delete(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}

func (r *Repository) GetbyID(id uint) (*Task, error) {
	var task Task

	err := r.db.First(&task, id).Error

	return &task, err
}

func (r *Repository) GetTaskbyUserID(user_id uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", user_id).Find(&tasks).Error
	return tasks, err
}

func (r *Repository) List() ([]Task, error) {
	var tasks []Task

	err := r.db.Find(&tasks).Error

	return tasks, err
}
