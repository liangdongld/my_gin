package autoload

type AmapConfig struct {
	Key string `ini:"key" yaml:"key"`
}

var Amap = AmapConfig{}
