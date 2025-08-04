package main

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

// 接口的继承
type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Stu struct {
}

func (stu Stu) test01() {

}
func (stu Stu) test02() {

}
func (stu Stu) test03() {

}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()
	a.test02()
	a.test03()
}
