package main

func main() {
	// 创建任务中心
	taskCenter := NewTaskCenter()

	// 创建不同的处理器
	logProcessor := &LogProcessor{name: "日志处理器"}
	notificationProcessor := &NotificationProcessor{name: "通知处理器"}
	dataProcessor := &DataProcessor{name: "数据处理器"}

	// 注册处理器
	taskCenter.RegisterProcessor(logProcessor)
	taskCenter.RegisterProcessor(notificationProcessor)
	taskCenter.RegisterProcessor(dataProcessor)

	// 创建并发布任务
	task1 := &Task{
		ID:          "TASK-001",
		Description: "处理用户订单",
		Status:      "进行中",
	}
	taskCenter.PublishTask(task1)

	// 移除一个处理器
	taskCenter.RemoveProcessor(notificationProcessor)

	// 发布另一个任务
	task2 := &Task{
		ID:          "TASK-002",
		Description: "生成月度报表",
		Status:      "待处理",
	}
	taskCenter.PublishTask(task2)
}
