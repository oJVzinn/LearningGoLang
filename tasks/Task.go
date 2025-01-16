package tasks

var tasks []Task

type Task struct {
	Name        string
	InitialDate string
	Completed   bool
}

func AddNewTask(task Task) {
	tasks = append(tasks, task)
}

func RemoveTask(task Task) {
	for i := 0; i < len(tasks); i++ {
		taskFind := tasks[i]
		if taskFind.Name == task.Name {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
}

func UpdateTask(task Task) {
	for i, taskNow := range tasks {
		if taskNow.Name == task.Name {
			tasks[i].Completed = true
			break
		}
	}
}

func FindByName(taskName string) (*Task, bool) {
	for i, task := range tasks {
		if task.Name == taskName {
			return &tasks[i], true
		}
	}

	return nil, false
}
