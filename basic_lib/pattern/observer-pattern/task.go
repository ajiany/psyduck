package main

import "fmt"

// Task 表示一个任务
type Task struct {
	ID          string
	Description string
	Status      string
}

// TaskCenter 任务中心，作为主题(Subject)
type TaskCenter struct {
	processors  []TaskProcessor
	currentTask *Task
}

// NewTaskCenter 创建新的任务中心
func NewTaskCenter() *TaskCenter {
	return &TaskCenter{
		processors: make([]TaskProcessor, 0),
	}
}

// RegisterProcessor 注册任务处理器
func (tc *TaskCenter) RegisterProcessor(processor TaskProcessor) {
	tc.processors = append(tc.processors, processor)
	fmt.Printf("处理器 [%s] 已注册\n", processor.GetName())
}

// RemoveProcessor 移除任务处理器
func (tc *TaskCenter) RemoveProcessor(processor TaskProcessor) {
	for i, p := range tc.processors {
		if p.GetName() == processor.GetName() {
			tc.processors = append(tc.processors[:i], tc.processors[i+1:]...)
			fmt.Printf("处理器 [%s] 已移除\n", processor.GetName())
			break
		}
	}
}

// PublishTask 发布新任务
func (tc *TaskCenter) PublishTask(task *Task) {
	tc.currentTask = task
	fmt.Printf("\n发布新任务: [%s] - %s\n", task.ID, task.Description)
	tc.notifyProcessors()
}

// notifyProcessors 通知所有处理器
func (tc *TaskCenter) notifyProcessors() {
	for _, processor := range tc.processors {
		processor.ProcessTask(tc.currentTask)
	}
}
