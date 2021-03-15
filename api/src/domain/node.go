package domain

type INode interface {
	GetID() int
	GetCreatedAtUnix() int
}
