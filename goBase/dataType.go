package goBase

import "fmt"

/*
 * 学习go语言的数据类型
 */
func StudyDataType() {
	var (
		a int       //整形
		b float32   //浮点
		c complex64 //复数
		d byte      //字符
		e string    //字符串
		f bool      //布尔
	)
	fmt.Printf("%T %v \n ", a, a) //通用打印格式
	fmt.Println(b, c, d, e, f)

	//运算符：　＋，　＋＝，　－，　－＝，　＊，　＊＝，　／，　／＝，　％，
	//　％＝，　＝＝，　！＝，　＋＋，　－－，　＞，　＜，　＞＝，　＜＝，
	//＞＞，　＜＜，　＞＞＝，　＜＜＝，　＆＆，　｜｜，　！

	//&a　指向ａ这个值
	//+的地址
	//*b　指向ｂ地址所对应的值

	//值类型：int、float、bool、string、array这些类型都属于值类型，
	//值类型的变量直接指向存在内存中的值，值类型的变量的值存储在栈中。
	//引用类型：特指slice、map、channel这三种预定义类型。

	//arrry、slice、map的声明格式
	//array: var name [len]type //定长数组，type表示类型
	//	   var name [...]type //不定长数组
	//	   var name [len][len]type //二维数组
	//	   name := [len]type
	//slice: var name []type
	//	   var name []type = make([]type,len)
	//	   name := make([]type,len)
	//	   name := []type{
	//	   	...
	//       }
	//map:   var name map[keyT]valueT  //T == type
	//	   name := make(map[keyT]vlaueT)
	//	   var name = map[keyT]vlaueT{
	//	   	   "..." : "..."
	//	   }
}
