package problems

// CoinChange returns the number of ways to make change for a given amount
func CoinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    dp[0] = 1
    for _, coin := range coins {
        for i := coin; i <= amount; i++ {
            dp[i] += dp[i-coin]
        }
    }
    return dp[amount]
}

// CoinChangeRecursive returns the number of ways to make change for a given amount
func CoinChangeRecursive(coins []int, amount int) int {
    memo := make([]int, amount+1)
    return coinChangeRecursive(coins, amount, memo)
}

func coinChangeRecursive(coins []int, amount int, memo []int) int {
    if amount == 0 {
        return 1
    }
    if amount < 0 {
        return 0
    }
    if memo[amount] != 0 {
        return memo[amount]
    }
    var res int
    for _, coin := range coins {
        res += coinChangeRecursive(coins, amount-coin, memo)
    }
    memo[amount] = res
    return res
}
