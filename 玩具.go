package main

import "fmt"

//小R过年收到了很多压岁钱，所以他想拿去买玩具。小R总共收到了 m 元压岁钱。商店里有10^9种玩具，种类编号为1~10^9，第i种玩具的价格为i元。但是小R已经有了其中n种玩具了，他不喜欢重复，所以每种玩具他最多只想拥有一个，已经有的就不会再买了，没有的也只会买最多一个。现在小R想知道他最多能再买多少种玩具。
//输入描述
//
//第一行两个整数 n,m，表示已有玩具的种类数和压岁钱数量。
//
//接下来一行n个整数a1,a2,…an，表示小R拥有的所有玩具的编号。
//
//对于 100% 的数据，2<=n<=100000,1<=ai,m<=10^9，保证ai,互不相同。
//输出描述
//
//输出一个整数表示小R最多能再买多少种玩具。
//样例输入
//
//5 16
//
//2 5 9 10 100

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	nums := make(map[int]bool)
	var v int
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		nums[v] = true
	}
	fmt.Println(backtrack(m, 0, []int{}))
}

func backtrack(m int, start int, path []int) {
	if Sum(path) {
		return
	}
	for i := start; i < m; i++ {

	}
}

//5 16
//2 5 9 10 100
