package manager

import (
	"awesomeProject/plugin"
	"fmt"
	"sync"
)

// 插件管理器
type PluginManager struct {
	plugins map[string]*plugin.PluginWrapper
	lock    sync.RWMutex
}

// NewPluginManager 创建管理器
func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make(map[string]*plugin.PluginWrapper),
	}
}

// 1. 注册插件（加载）
func (m *PluginManager) Register(p plugin.Plugin) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.plugins[p.Name()] = &plugin.PluginWrapper{
		Instance: p,
		Status:   plugin.StatusEnabled,
	}
}

// 2. 卸载插件
func (m *PluginManager) Unregister(name string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.plugins, name)
}

// 3. 启用插件
func (m *PluginManager) Enable(name string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	wrapper, ok := m.plugins[name]
	if !ok {
		return fmt.Errorf("插件不存在")
	}
	wrapper.Status = plugin.StatusEnabled
	return nil
}

// 4. 禁用插件
func (m *PluginManager) Disable(name string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	wrapper, ok := m.plugins[name]
	if !ok {
		return fmt.Errorf("插件不存在")
	}
	wrapper.Status = plugin.StatusDisabled
	return nil
}

// 5. 执行插件（统一用 Run）
func (m *PluginManager) Run(name string, data map[string]interface{}) (map[string]interface{}, error) {
	m.lock.RLock()
	wrapper, ok := m.plugins[name]
	m.lock.RUnlock()

	if !ok {
		return nil, fmt.Errorf("插件[%s]不存在", name)
	}

	// 插件被禁用，不执行
	if wrapper.Status != plugin.StatusEnabled {
		return nil, fmt.Errorf("插件[%s]未启用", name)
	}

	// 异常隔离：panic 不崩溃主程序
	defer func() {
		if err := recover(); err != nil {
			m.lock.Lock()
			wrapper.Status = plugin.StatusError
			m.lock.Unlock()
			fmt.Printf("插件崩溃已隔离: %s | %v\n", name, err)
		}
	}()

	// ✅ 调用插件的 Run 方法
	return wrapper.Instance.Run(data)
}

// 获取插件信息（名称、版本、状态）
func (m *PluginManager) GetPluginInfo(name string) (string, string, string, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	wrapper, ok := m.plugins[name]
	if !ok {
		return "", "", "", fmt.Errorf("插件不存在")
	}

	return wrapper.Instance.Name(),
		wrapper.Instance.Version(),
		wrapper.Status,
		nil
}

// 获取所有插件列表
func (m *PluginManager) List() []map[string]string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	var list []map[string]string
	for _, w := range m.plugins {
		list = append(list, map[string]string{
			"name":    w.Instance.Name(),
			"version": w.Instance.Version(),
			"status":  w.Status,
		})
	}
	return list
}
