package config

type Lark struct {
	Id     string `default:"${ID}" json:"id,omitempty" validate:"required"`
	Secret string `default:"${SECRET}" json:"secret,omitempty" validate:"required"`
}
