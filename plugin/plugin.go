package plugin

// 插件三种状态信息
const (
	StatusEnabled  = "Enabled"
	StatusDisabled = "Disabled"
	StatusError    = "Error"
)

// Plugin 插件接口,api first定义plugin接口
type Plugin interface {
	Name() string // 返回插件名字（消息类型：order_pay / refund）
	Version() string
	Run(data map[string]interface{}) (map[string]interface{}, error) // 执行逻辑
}
type PluginWrapper struct {
	Instance Plugin
	Status   string // 启用/禁用/异常
}
