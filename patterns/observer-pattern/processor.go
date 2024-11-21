package main

import "fmt"

// TaskProcessor 任务处理器接口（观察者）
type TaskProcessor interface {
	ProcessTask(*Task)
	GetName() string
}

// LogProcessor 日志处理器
type LogProcessor struct {
	name string
}

func (lp *LogProcessor) ProcessTask(task *Task) {
	fmt.Printf("日志处理器: 记录任务 [%s] 的状态变更为 %s\n", task.ID, task.Status)
}

func (lp *LogProcessor) GetName() string {
	return lp.name
}

// NotificationProcessor 通知处理器
type NotificationProcessor struct {
	name string
}

func (np *NotificationProcessor) ProcessTask(task *Task) {
	fmt.Printf("通知处理器: 发送任务 [%s] 的通知消息\n", task.ID)
}

func (np *NotificationProcessor) GetName() string {
	return np.name
}

// DataProcessor 数据处理器
type DataProcessor struct {
	name string
}

func (dp *DataProcessor) ProcessTask(task *Task) {
	fmt.Printf("数据处理器: 保存任务 [%s] 到数据库\n", task.ID)
}

func (dp *DataProcessor) GetName() string {
	return dp.name
}
