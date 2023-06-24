package main

import "fmt"

func hi(name string) string {
	// fmt.Println("hi!" + name)
	msg := "hi!" + name
	return msg
}

func main() {
	// 型
	a := 10
	b := 12.3
	c :="hoge"
	var d bool
	fmt.Printf("a: %d, b:%f, c:%s, d:%t\n",a,b,c,d)

	// const
	const (
		sun = iota //1
		mon //2
		tue //3
	)
	fmt.Println(sun, mon, tue)

	// 演算
	var x int
	x = 10 % 3
	x += 3
	fmt.Println(x)

	var s string
	s = "hello" + "world"
	fmt.Println(s)

	ae := true
	be := false
	fmt.Println(ae && be) //and
	fmt.Println(ae || be) //or
	fmt.Println(!ae) //not

	// ポインタ
	// アドレスを指し示す変数
	// 演算はできない
	ap := 5
	var pap *int
	pap = &ap //&ap = apのアドレス
	// papの領域にあるデータの値 = *pap
	fmt.Println(pap)
	fmt.Println(*pap)

	// 関数
	// hi ("taguchi")
	fmt.Println(hi("taguchi"))
}


