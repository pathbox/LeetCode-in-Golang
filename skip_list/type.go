package skiplist

import (
	"math/rand"
)

// Return true if lhs greater than rhs
type GreaterThanFunc func(lhs, rhs interface{}) bool

// Return true if lhs less than rhs
type LessThanFunc GreaterThanFunc

type defaultRandSource struct{}

type Comparable interface{
	Descending() bool
	Compare(lhs, rhs interface{}) bool
}

type elementNode struct {
	next []*Element  // Element 数组  一个链表结构
}

type Element struct {
	elementNode
	key, Value interface{}
	score float64
}

type SkipList struct {
	elementNode
	level int
	length int
	keyFunc Comparable
	randSource rand.Source
	reversed bool

	prevNodesCache []*elementNode // a cache for Set/Remove
}

type Scorable interface {
	Score() float64
}

func (r defaultRandSource) Int63() int64 {
	return rand.Int63
}

func (r defaultRandSource) Seed(seed int64) {
	rand.Seed(seed)
}

func (f GreaterThanFunc) Descending() bool {
	return false
}

func (f GreaterThanFunc) Compare(lhs, rhs interface{}) bool {
	return f(lhs, rhs)
}

func (f LessThanFunc) Descending() bool {
	return true
}

func (f LessThanFunc) Compare(lhs, rhs interface{}) bool {
	return f(lhs, rhs)
}

