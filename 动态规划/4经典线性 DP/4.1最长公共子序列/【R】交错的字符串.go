package main

import "fmt"

// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
// 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空
// 子字符串
// ：
// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
//
// 注意：a + b 意味着字符串 a 和 b 连接。
func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}

// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 输出：true

// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// 输出：false

// 输入：s1 = "", s2 = "", s3 = ""
// 输出：true
func isInterleave(s1 string, s2 string, s3 string) bool {

	/*
	   考虑将求长度为m的整个s1和长度为n的整个s2能否交错构成整个s3这一全局问题，拆分为s1的前i个字符组成的子字符串，和s2的前j个字符组成的子字符串，能否交错构成s3的前i+j个字符组成的子字符串这一子问题；当i==m，j==n时，该子问题的解就是全局问题的解，因此使用动态规划
	   状态：
	       构建一个大小为m+1 * n+1的状态数组dp
	       dp[i][j]为bool类型，表示s1的前i个字符组成的子字符串，和s2的前j个字符组成的子字符串，能否交错构成s3的前i+j个字符组成的子字符串
	   起始状态：
	       当i==0，j==0：此时s1和s2都是空串，能够组成s3空串，因此dp[0][0]=true
	       当i==0，j!=0：此时s1为空串，只需要比较s2和s3；如果s2的前j-1个字符与s3的前j-1个字符匹配，且s2的第j个字符和s3的第j个字符相等，则s2的前j个字符与s3的前j个字符匹配；即：dp[0][j] = dp[0][j-1] && s2[j-1]==s3[j-1]
	       当i!=0，j==0：此时s2为空串，只需要比较s1和s3；如果s1的前i-1个字符与s3的前i-1个字符匹配，且s1的第i个字符和s3的第i个字符相等，则s1的前i个字符与s3的前i个字符匹配；即：dp[i][0] = dp[i-1][0] && s1[i-1]==s3[i-1]
	   状态转移：
	       情况1：如果s1的前i-1个字符与s2的前j个字符能够匹配s3的前i+j-1个字符，且s1的第i个字符与s3的第i+j个字符相等，则s1的前i个字符与s2的前j个字符能够匹配s3的前i+j个字符；即dp[i][j] = dp[i-1][j] && s1[i-1]==s3[i+j-1]
	       情况2：如果s1的前i个字符与s2的前j-1个字符能够匹配s3的前i+j-1个字符，且s2的第j个字符与s3的第i+j个字符相等，则s1的前i个字符与s2的前j个字符能够匹配s3的前i+j个字符；即dp[i][j] = dp[i][j-1] && s2[j-1]==s3[i+j-1]
	       上述两种情况只需要满足其中一种即可，因此dp[i][j] = (dp[i-1][j] && s1[i-1]==s3[i+j-1]) || (dp[i][j-1] && s2[j-1]==s3[i+j-1])
	   复杂度分析：
	       时间复杂度：遍历m*n的矩阵进行状态运算，因此时间复杂度为O(mn)
	       空间复杂度：构建一个大小为m*n的状态数组，因此空间复杂度为O(mn)
	*/
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true
	for i := 1; i < len(dp); i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < len(dp[0]); j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			// 下面两个条件满足一个即可
			// 如果  s1[i-1] == s3[i+j-1] 则查看上一个 s1[i-1] 字符是否为 ture
			if s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i-1][j]
			}
			// 如果  s2[j-1] == s3[i+j-1] 则查看上一个 s2[i-1][j-1] 字符是否为 ture
			if s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j-1] || dp[i][j]
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
