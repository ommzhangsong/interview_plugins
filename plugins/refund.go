package plugins

import "awesomeProject/plugin"

func init() {
	plugin.RegisterGlobal(&RefundPlugin{})
}

type RefundPlugin struct{}

func (r *RefundPlugin) Name() string {
	return "refund"
}

func (r *RefundPlugin) Version() string {
	return "1.0.0"
}

func (r *RefundPlugin) Run(data map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"msg": "退款处理完成",
	}, nil
}
