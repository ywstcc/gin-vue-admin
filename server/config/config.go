package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql  Mssql           `mapstructure:"mssql" json:"mssql" yaml:"mssql"`
	Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Oracle Oracle          `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	Sqlite Sqlite          `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`

	DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
