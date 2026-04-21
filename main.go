package main

import (
	"awesomeProject/manager"
	"awesomeProject/plugin"
	"fmt"

	_ "awesomeProject/plugins" // 匿名导入，自动加载所有插件
)

func main() {
	// 1. 创建插件管理器
	pm := manager.NewPluginManager()

	// 2. 自动注册所有插件
	for _, p := range plugin.GetAll() {
		pm.Register(p)
	}

	// 3. 查看所有插件信息
	fmt.Println("=== 所有插件列表 ===")
	for _, info := range pm.List() {
		fmt.Printf("插件：%s | 版本：%s | 状态：%s\n",
			info["name"], info["version"], info["status"])
	}

	// 4. 执行插件
	fmt.Println("\n=== 执行 order_pay 插件 ===")
	data := map[string]interface{}{"order_id": "ORDER_1001"}
	result, err := pm.Run("order_pay", data)
	if err != nil {
		fmt.Println("执行失败：", err)
	} else {
		fmt.Println("执行成功：", result)
	}

	// 5. 禁用插件
	fmt.Println("\n=== 禁用 refund 插件 ===")
	pm.Disable("refund")

	// 6. 查看禁用后的状态
	name, version, status, _ := pm.GetPluginInfo("refund")
	fmt.Printf("插件：%s | 版本：%s | 状态：%s\n", name, version, status)

	// 7. 尝试执行已禁用的插件（会报错，验证隔离）
	fmt.Println("\n=== 尝试执行禁用的 refund ===")
	_, err = pm.Run("refund", nil)
	fmt.Println("执行结果：", err)
}
