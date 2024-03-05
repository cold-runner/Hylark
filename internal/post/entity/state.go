package entity

type state int32

const (
	Audit state = 1 << iota

	Pass

	Denounce

	Hot
)
