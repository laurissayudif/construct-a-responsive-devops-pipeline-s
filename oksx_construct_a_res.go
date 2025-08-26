package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// DevOpsPipelineSimulator represents a DevOps pipeline simulator
type DevOpsPipelineSimulator struct {
	PipelineName string            `json:"pipeline_name"`
	Stages       []Stage           `json:"stages"`
	Artifacts    map[string]string `json:"artifacts"`
}

// Stage represents a stage in the pipeline
type Stage struct {
	StageName  string   `json:"stage_name"`
	Tasks      []Task   `json:"tasks"`
	DependsOn  []string `json:"depends_on"`
	Artifact   string   `json:"artifact"`
	Status     string   `json:"status"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	ErrorMsg   string   `json:"error_msg"`
}

// Task represents a task in a stage
type Task struct {
	TaskName    string `json:"task_name"`
	TaskType   string `json:"task_type"`
	Status     string `json:"status"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	ErrorMsg   string   `json:"error_msg"`
}

// NewDevOpsPipelineSimulator returns a new DevOps pipeline simulator
func NewDevOpsPipelineSimulator(pipelineName string) *DevOpsPipelineSimulator {
	return &DevOpsPipelineSimulator{
		PipelineName: pipelineName,
		Stages:       []Stage{},
		Artifacts:    map[string]string{},
	}
}

// AddStage adds a stage to the pipeline
func (d *DevOpsPipelineSimulator) AddStage(stageName string, dependsOn []string) *Stage {
	stage := Stage{
		StageName:  stageName,
		Tasks:      []Task{},
		DependsOn:  dependsOn,
		Artifact:   "",
		Status:     "pending",
		StartTime:  time.Time{},
		EndTime:    time.Time{},
		ErrorMsg:   "",
	}
	d.Stages = append(d.Stages, stage)
	return &stage
}

// AddTask adds a task to a stage
func (s *Stage) AddTask(taskName string, taskType string) *Task {
	task := Task{
		TaskName:    taskName,
		TaskType:   taskType,
		Status:     "pending",
		StartTime:  time.Time{},
		EndTime:    time.Time{},
		ErrorMsg:   "",
	}
	s.Tasks = append(s.Tasks, task)
	return &task
}

// Run runs the pipeline simulator
func (d *DevOpsPipelineSimulator) Run() {
	for _, stage := range d.Stages {
		fmt.Printf("Running stage: %s\n", stage.StageName)
		for _, task := range stage.Tasks {
			task.Status = "running"
			task.StartTime = time.Now()
			time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
			task.Status = "success"
			task.EndTime = time.Now()
			fmt.Printf("  Task: %s - %s\n", task.TaskName, task.Status)
		}
		stage.Status = "success"
		stage.EndTime = time.Now()
	}
}

func main() {
	simulator := NewDevOpsPipelineSimulator("my-pipeline")
	stage1 := simulator.AddStage("build", []string{})
	stage1.AddTask("compile", "build")
	stage1.AddTask("test", "test")
	stage2 := simulator.AddStage("deploy", []string{"build"})
	stage2.AddTask("deploy-to-prod", "deploy")
	simulator.Run()

	jsonData, _ := json.MarshalIndent(simulator, "", "  ")
	fmt.Println(string(jsonData))
}