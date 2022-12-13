package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

// *implemented
func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	var tasks []entity.Task
	query, err := r.db.Model(&entity.Task{}).Select("*").Where("user_id = ?", id).Rows()
	defer query.Close()
	for query.Next() {
		err = r.db.ScanRows(query, &tasks)
	}
	if err != nil {
		return []entity.Task{}, err
	}
	return tasks, nil
}

// *implemented
func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	res := r.db.Create(&task)
	if res.Error != nil {
		return 0, res.Error
	}
	return task.ID, nil
}

// *implemented
func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var task entity.Task
	query := r.db.Model(&entity.Task{}).First(&task, id)
	if query.Error != nil {
		return entity.Task{}, query.Error
	}
	return task, nil
}

// *implemented
func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	var tasks []entity.Task
	query := r.db.Model(&entity.Task{}).First(&tasks, "category_id = ?", catId)
	if query.Error != nil {
		return []entity.Task{}, nil
	}
	return tasks, nil
}

// *implemented
func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	// fmt.Println(task)
	res := r.db.Model(&entity.Task{}).Where("id = ?", &task.ID).Updates(&task)
	return res.Error
}

// *implemented
func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	res := r.db.Delete(&entity.Task{}, id)
	return res.Error
}
