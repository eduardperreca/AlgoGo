package main

func main(){

	var dp = make([]int, 10)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < 10; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	println(dp[9])
	
}