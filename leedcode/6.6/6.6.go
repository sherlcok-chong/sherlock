package __6

import (
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
	lens := len(strs)
	dp := make([][][]int,lens+1)
	for i := range dp{
		dp[i] = make([][]int,m+1)
		for j:=range dp[i] {
			dp[i][j]=make([]int ,n+1)
		}
	}
	for i,val:= range strs{
		zero := strings.Count(val,"0")
		one := len(val)-zero
		for j:=0;j<=m;j++{
			for k:=0;k<=n;k++ {
				dp[i+1][j][k]=dp[i][j][k]
				if j>=zero&&k>=n {
					dp[i+1][j][k]=max(dp[i][j][k],dp[i][j-zero][k-one]+1)
				}
			}
		}
	}
	return dp[lens][m][n]
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}