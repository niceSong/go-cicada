package exception

type Throwable struct {
	Title  string `json:"title,omitempty" yaml:"title" bson:"title" mapstructure:"Title"`
	Status int32  `json:"status,omitempty" yaml:"status" bson:"status" mapstructure:"Status"`
	Detail string `json:"detail,omitempty" yaml:"detail" bson:"detail" mapstructure:"Detail"`
	Cause  string `json:"cause,omitempty" yaml:"cause" bson:"cause" mapstructure:"Cause"`
	Code   int    `json:"code,omitempty" yaml:"code" bson:"code" mapstructure:"Code"`
}
