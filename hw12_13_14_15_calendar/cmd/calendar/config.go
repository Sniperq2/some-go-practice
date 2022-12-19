package main

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger LoggerConf
}

type LoggerConf struct {
	Level    string `yaml: level`
	Filename string `yaml: filename`
}

func NewConfig() Config {
	return Config{}
}
