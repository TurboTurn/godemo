package main

import (
	"errors"
	"fmt"
	//"github.com/pkg/errors"
)

func main() {
	fmt.Println("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちは世界")
	var vname1, vname2, vname3 = "name1", "name2", "name3"
	fmt.Println(vname1, vname2, vname3)

	const pi float32 = 3.1415926
	fmt.Println(pi)

	var enable, disabled bool = true, false
	fmt.Println(enable, disabled)

	/**
	！Go还支持复数。它的默认类型是complex128（64位实数+64位虚数）。如果需要小一些的，也
	有complex64(32位实数+32位虚数)。
	*/
	var c complex64 = 5 + 5i
	fmt.Println(c)

	var s string = "ahello" //字符串是不可变的
	var ch = []byte(s)      // 将字符串 s 转换为 []byte 类型
	ch[0] = 'h'
	s = string(ch) // 再转换回 string 类型
	fmt.Println(s)

	s = "hello"
	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Println(s)

	//如果要声明一个多行的字符串怎么办？可以通过`来声明：
	m := `hello
	world`
	fmt.Println(m)
	//括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。

	err := errors.New("emit macho dwarf")
	fmt.Println(err)

	f := float32(3.14)
	fmt.Println(f)

	var i byte = 'd'
	fmt.Println(i)

	var(s1  = 5
		s2 float32 )
	fmt.Println(s1,s2)

	const(
		x = iota
		y = iota
		z = iota
	)
	fmt.Println(x,y,z)
}
