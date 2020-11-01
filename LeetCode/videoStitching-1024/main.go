package main

import (
	"fmt"
	"math"
)

/*
	你将会获得一系列视频片段，这些片段来自于一项持续时长为T秒的体育赛事。这些片段可能有所重叠，也可能长度不一。
	视频片段clips[i]都用区间进行表示：开始于clips[i][0]并于clips[i][1]结束。
	我们甚至可以对这些片段自由地再剪辑，例如片段[0, 7]可以剪切成[0, 1] +[1, 3] + [3, 7]三部分。
	我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, T]）
	返回所需片段的最小数目，如果无法完成该任务，则返回-1 。
*/

func main() {
	clips := [][]int{{0, 2},{4,6},{8,10},{1,9},{1,5},{5,9}}
	T:=10
	fmt.Println(videoStitching(clips,T))
}

func videoStitching(clips [][]int, T int) int {
	const inf = math.MaxInt64 - 1
	dp := make([]int, T+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0
	for i := 1; i <= T; i++ {
		for _, clip := range clips {
			l, r := clip[0], clip[1]
			// 若能剪出子区间 [l,i]，则可以从 dp[l] 转移到 dp[i]
			if l < i && i <= r && dp[l] + 1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}
	if dp[T] == inf {
		return -1
	}
	return dp[T]
}

