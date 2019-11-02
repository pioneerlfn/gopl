package gopl

type Counter interface {
	Counter() (int, error)
}
