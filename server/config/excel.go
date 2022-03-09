package config

type Excel struct {
	Dir string `mapstructure:"dir" json:"dir" yaml:"dir"`
	UserPic string `mapstructure:"userPic" json:"userPic" yaml:"userPic"`
	UserPay string `mapstructure:"userPay" json:"userPic" yaml:"userPay"`
	UserCertify string `mapstructure:"userCertify" json:"userPic" yaml:"userCertify"`
}
