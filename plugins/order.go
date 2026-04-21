package plugins

import "awesomeProject/plugin"

func init() {
	plugin.RegisterGlobal(&OrderPayPlugin{})
}

type OrderPayPlugin struct{}

func (o *OrderPayPlugin) Name() string {
	return "order_pay"
}

func (o *OrderPayPlugin) Version() string {
	return "1.0.0"
}

func (o *OrderPayPlugin) Run(data map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"msg": "订单支付处理完成",
	}, nil
}
