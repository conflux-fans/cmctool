package configs

type Config struct {
	Mail struct {
		Sender struct {
			FromAlias string `yaml:"fromAlias"`
			Host      string `yaml:"host"`
			Port      int    `yaml:"port"`
			Username  string `yaml:"username"`
			Password  string `yaml:"password"`
		} `yaml:"sender"`
		Receivers []string `yaml:"receivers"`
	} `yaml:"mail"`
}
