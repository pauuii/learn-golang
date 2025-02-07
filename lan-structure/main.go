// 声明包
package main

//引入包
import (
	p1 "basis/lan-structure/package1"
	_ "basis/lan-structure/package2"
	"fmt"
)

// 初始化函数，golang 每个包的引用会优先调用该函数
func init() {
	fmt.Println("调用 lan-structure init")
}

// 函数（main 函数为整个程序的入口）
/*
多行注释
*/
func main() {
	var i = 10
	// 语句或表达式
	fmt.Println("调用 lan-structure main")
	fmt.Println(fmt.Sprintf("打印参数i ： %d", i))
	// 包名可与包引用目录不一致(包名和文件夹命名可以不同)
	// golang 公有成员与私有成员通过成员的标识的首字母来区分
	// 首字母大写表示公有成员，首字母小写表示私有成员，只有两种情况
	p1.F1()
}
