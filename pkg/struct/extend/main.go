package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) ShowInfo() {
	fmt.Printf("Name:%s, Age:%d, Score:%d\n", s.Name, s.Age, s.Score)
}

func (s *Student) SetScore(score int) {
	s.Score = score
}

type Pupil struct {
	Student
}

func (p *Pupil) Testing() {
	fmt.Println("小学生正在考试...")
}

type Graduate struct {
	Student
}

func (g *Graduate) Testing() {
	fmt.Println("大学生正在测试...")
}

func main() {
	// 对于这种继承的我们如何使用
	var pupil = &Pupil{
		Student: Student{
			Score: 80,
		},
	}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 18

	pupil.ShowInfo()
	pupil.SetScore(100)
	pupil.Testing()
	pupil.ShowInfo()

	var graduate = &Graduate{
		Student: Student{
			Name:  "",
			Age:   0,
			Score: 50,
		},
	}
	graduate.ShowInfo()
	graduate.SetScore(90)
	graduate.Testing()
	graduate.ShowInfo()
}
