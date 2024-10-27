package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/exp/rand"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

// Task 任务接口
type Task interface {
	Execute(ctx context.Context) error
}

// ShellTask 实现了 Task 接口的具体任务
type ShellTask struct {
	taskTimeout time.Duration
	scriptPath  string
}

func (t *ShellTask) Execute(ctx context.Context) error {
	// 检查脚本文件是否存在
	if _, err := os.Stat(t.scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("脚本文件不存在: %s", t.scriptPath)
	}
	env := os.Environ()
	env = append(env, "PACKAGE_ROOT=/tmp/acheck_tool")

	if err := os.Chmod(t.scriptPath, 0755); err != nil {
		return fmt.Errorf("无法给脚本赋予权限: %w", err)
	}

	cmd := exec.CommandContext(ctx, "bash", t.scriptPath)
	cmd.Env = env
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("启动脚本时出错: %w", err)
	}

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("脚本执行失败，退出状态码: %d，异常内容: %s", exitError.ExitCode(), err)
		}
		return fmt.Errorf("执行脚本出错: %w", err)
	}
	return nil
}

// NewShellTask 构造函数
func NewShellTask(scriptPath string) (*ShellTask, error) {
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("脚本文件不存在: %s", scriptPath)
	}
	return &ShellTask{scriptPath: scriptPath}, nil
}

// TaskManager 负责管理任务的执行
type TaskManager struct {
	mu           sync.Mutex
	taskExecuted bool
	taskTimeout  time.Duration
}

func (tm *TaskManager) StartTask(task Task) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.taskExecuted {
		return fmt.Errorf("任务已执行，跳过")
	}

	tm.taskExecuted = true
	go func() {
		defer func() {
			tm.mu.Lock()
			tm.taskExecuted = false
			tm.mu.Unlock()
		}()
		ctx, cancel := context.WithTimeout(context.Background(), tm.taskTimeout)

		defer cancel() // 确保资源释放

		// 捕获任务执行中的异常
		defer func() {
			if r := recover(); r != nil {
				log.Println("任务执行异常:", r)
			}
		}()

		if err := task.Execute(ctx); err != nil {
			if errors.Is(context.DeadlineExceeded, ctx.Err()) {
				log.Println("任务超时取消", rand.Intn(100))
			} else {
				log.Println("执行错误：", err)
			}
		} else {
			log.Println("后台任务完成!")
		}
	}()
	return nil
}

var (
	taskManager = &TaskManager{
		taskTimeout: 10 * time.Second,
	}
)

// startTaskHandler 处理 HTTP 请求以启动后台任务
func startTaskHandler(w http.ResponseWriter, r *http.Request) {
	task, err := NewShellTask("/home/all_projects/coroutines_test/check.sh")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := taskManager.StartTask(task); err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

	log.Println("后台任务已启动.")
	fmt.Fprintln(w, "后台任务已启动.")
}

func main() {
	// 初始化日志文件
	logFile, err := os.OpenFile("/home/all_projects/coroutines_test/operations.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("无法打开日志文件:", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	http.HandleFunc("/start-task", startTaskHandler)
	log.Println("服务器正在运行，访问 http://localhost:8080/start-task 来启动任务...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("服务器启动失败:", err)
	}
}
