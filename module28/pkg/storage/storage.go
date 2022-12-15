package storage

import (
	"module28/pkg/student"
)

type Storage map[string]*student.Student

func New() Storage {
	var st Storage
	st = make(map[string]*student.Student)
	return st
}

func (st Storage) Get(name string) *student.Student { // Метод get
	return st[name]
}
func (st Storage) Put(s *student.Student) { // Метод put
	st[s.Name()] = s
}
