package smap

import "sync"

type SMap[K comparable, V any] struct {
	data map[K]V
	lock sync.RWMutex
}

func NewSMap[K comparable, V any]() *SMap[K, V] {
	return &SMap[K, V]{
		data: make(map[K]V),
	}
}

func (s *SMap[K, V]) Set(k K, v V) (exist bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, exist = s.data[k]
	if exist {
		return true
	}
	s.data[k] = v
	return false
}

func (s *SMap[K, V]) Get(k K) (V, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	v, ok := s.data[k]
	return v, ok
}

func (s *SMap[K, V]) Pop(k K) (V, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	v, ok := s.data[k]
	if ok {
		delete(s.data, k)
	}
	return v, ok
}

func (s *SMap[K, V]) Delete(k K) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.data, k)
}

func (s *SMap[K, V]) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.data)
}

func (s *SMap[K, V]) Range(f func(k K, v V) bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for k, v := range s.data {
		if !f(k, v) {
			break
		}
	}
}

func (s *SMap[K, V]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = make(map[K]V)
}

func (s *SMap[K, V]) Keys() []K {
	s.lock.RLock()
	defer s.lock.RUnlock()
	keys := make([]K, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s *SMap[K, V]) Values() []V {
	s.lock.RLock()
	defer s.lock.RUnlock()
	values := make([]V, 0, len(s.data))
	for _, v := range s.data {
		values = append(values, v)
	}
	return values
}

func (s *SMap[K, V]) Has(k K) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.data[k]
	return ok
}
