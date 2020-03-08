package model

type Decor interface {
	Decorate(meta *Meta) *Meta
}
