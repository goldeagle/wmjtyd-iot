package core

import (
	"context"
	"sync"
)

// Lifecycle 管理应用的生命周期
type Lifecycle struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewLifecycle 创建新的Lifecycle实例
func NewLifecycle() *Lifecycle {
	ctx, cancel := context.WithCancel(context.Background())
	return &Lifecycle{
		ctx:    ctx,
		cancel: cancel,
	}
}

// Start 启动生命周期管理
func (lc *Lifecycle) Start(initFunc func() error) error {
	if err := initFunc(); err != nil {
		return err
	}
	return nil
}

// Stop 停止生命周期管理
func (lc *Lifecycle) Stop(cleanupFunc func()) {
	lc.cancel()
	lc.wg.Wait()
	cleanupFunc()
}

// Context 返回生命周期上下文
func (lc *Lifecycle) Context() context.Context {
	return lc.ctx
}

// AddTask 添加后台任务
func (lc *Lifecycle) AddTask(task func(ctx context.Context)) {
	lc.wg.Add(1)
	go func() {
		defer lc.wg.Done()
		task(lc.ctx)
	}()
}
