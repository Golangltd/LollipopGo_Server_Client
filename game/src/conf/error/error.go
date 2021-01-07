package error

type Code int

const (
	CorrectCode  Code = 0     //操作正确码
	NoEnoughCoin Code = 11002 //金币不足
)
