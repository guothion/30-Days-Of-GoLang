package model

// 首字母小写不能导出
type student struct {
	Name  string
	score float64
}

// 通过工厂模式来解决别的地方调用
func NewStudent(Name string, Score float64) *student {
	return &student{
		Name:  Name,
		score: Score,
	}
}

func (s *student) GetScore() float64 {
	return s.score
}
