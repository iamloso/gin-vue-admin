package config

type Excel struct {
	Dir string `mapstructure:"dir" json:"dir" yaml:"dir"`
	UserPic string `mapstructure:"userPic" json:"userPic" yaml:"userPic"`
}
