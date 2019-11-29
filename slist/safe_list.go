package slist

import (
	"container/list"
	"sync"
)

type safeList struct {
	list   *list.List
	locker *sync.RWMutex
}

func newSafeList() *safeList {
	return &safeList{
		list:   list.New(),
		locker: new(sync.RWMutex),
	}
}

func (s *safeList) Len() int {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.list.Len()
}

func (s *safeList) PushBack(e interface{}) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.list.PushBack(e)
}

func (s *safeList) RemoveFront() interface{} {
	s.locker.Lock()
	defer s.locker.Unlock()

	if s.list.Len() <= 0 {
		return nil
	}
	return s.list.Remove(s.list.Front())
}
