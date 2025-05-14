package repository

import (
	"cmd/internal/models"
	"errors"
	"sync"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository struct {
	tasks map[string]models.Task
	mutex sync.Mutex
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]models.Task),
	}
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var tasks []models.Task
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) Create(task models.Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.tasks[task.ID] = task
	return nil
}

func (r *TaskRepository) GetByID(id string) (models.Task, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		return models.Task{}, ErrTaskNotFound
	}
	return task, nil
}

func (r *TaskRepository) Update(task models.Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return ErrTaskNotFound
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *TaskRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}
	delete(r.tasks, id)
	return nil
}
