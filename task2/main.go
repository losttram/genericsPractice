package main

import (
	"errors"
	"fmt"
)

var (
	errElementExists       = errors.New("element already exists")
	errElementDoesNotExist = errors.New("element does not exist")
)

type Map[K comparable, V any] struct {
	data map[K]V
}

func (m *Map[K, V]) Add(key K, value V) error {
	if _, ok := m.data[key]; !ok {
		m.data[key] = value
	} else {
		return fmt.Errorf("%w: key == %v; value == %v", errElementExists, key, m.data[key])
	}
	return nil
}

func (m *Map[K, _]) Delete(key K) error {
	if _, ok := m.data[key]; ok {
		delete(m.data, key)
	} else {
		return fmt.Errorf("%w: key == %v", errElementDoesNotExist, key)
	}
	return nil
}

func (m *Map[K, V]) Update(key K, value V) error {
	if _, ok := m.data[key]; ok {
		m.data[key] = value
	} else {
		return fmt.Errorf("%w: key == %v", errElementDoesNotExist, key)
	}
	return nil
}

func (m *Map[K, V]) Get(key K) (V, error) {
	if value, ok := m.data[key]; ok {
		return value, nil
	} else {
		return value, fmt.Errorf("%w: key == %v", errElementDoesNotExist, key)
	}
}

// Task: create Map[K, V] with methods Add(K, V), Delete(K), Update(K, V), Get(K)
func main() {
	m := Map[int, string]{
		data: make(map[int]string),
	}
	m.Add(1, "ono")
	m.Update(1, "one")
	value, _ := m.Get(1)
	fmt.Println(value)
	m.Delete(1)
}
