package plugin

var globalPlugins []Plugin

func RegisterGlobal(p Plugin) {
	globalPlugins = append(globalPlugins, p)
}

func GetAll() []Plugin {
	return globalPlugins
}
