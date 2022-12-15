package student

type Student struct { // Описываем структуру
	name string

	age int

	grade int
}

func New(n string, a, g int) *Student {
	s := Student{n, a, g}
	return &s
}

func (s Student) Name() string { // Фукция по возвращению имени
	return s.name
}

func (s Student) Age() int { // Фукция по возвращению возраста
	return s.age
}

func (s Student) Grade() int { // Фукция по возвращению курс
	return s.grade
}
