package _map

import (
	"errors"
)

type LinkedHashMap[K comparable, V any] struct {
	items []LinkedHashMapItem[K, V]
}

type LinkedHashMapItem[K comparable, V any] struct {
	Key   K
	Value V
}

func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		items: []LinkedHashMapItem[K, V]{},
	}
}

func (m *LinkedHashMap[K, V]) Get(i K) V {
	for _, item := range m.items {
		if item.Key == i {
			return item.Value
		}
	}
	panic(errors.New("the key does not exist"))
}

func (m *LinkedHashMap[K, V]) Has(i K) bool {
	if m.Length() == 0 {
		return false
	}
	for _, item := range m.items {
		if item.Key == i {
			return true
		}
	}
	return false
}

func (m *LinkedHashMap[K, V]) Put(k K, v V) {
	m.items = append(m.items, LinkedHashMapItem[K, V]{
		Key:   k,
		Value: v,
	})
}

func (m *LinkedHashMap[K, V]) List() []LinkedHashMapItem[K, V] {
	if m.Length() == 0 {
		return []LinkedHashMapItem[K, V]{}
	}
	return m.items
}

func (m *LinkedHashMap[K, V]) Length() int {
	if m == nil {
		return 0
	}
	return len(m.items)
}

func (m *LinkedHashMap[K, V]) Keys() []K {
	if m.Length() == 0 {
		return []K{}
	}
	var keys = make([]K, m.Length())
	for _, item := range m.items {
		keys = append(keys, item.Key)
	}
	return keys
}

func (m *LinkedHashMap[K, V]) Values() []V {
	if m.Length() == 0 {
		return []V{}
	}
	var values = make([]V, m.Length())
	for _, item := range m.items {
		values = append(values, item.Value)
	}
	return values
}

func (m *LinkedHashMap[K, V]) ToHashMap() map[K]V {
	if m.Length() == 0 {
		return nil
	}
	var hashMap = make(map[K]V, m.Length())
	for _, item := range m.items {
		hashMap[item.Key] = item.Value
	}
	return hashMap
}
