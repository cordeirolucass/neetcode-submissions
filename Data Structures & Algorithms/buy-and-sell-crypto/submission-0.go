func maxProfit(prices []int) int {
    maxProfit := 0
    minPrice := 101
    for i:= 0; i < len(prices); i++ {
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else {
            lucroHoje := prices[i] - minPrice
            if lucroHoje > maxProfit {
                maxProfit = lucroHoje
            }
        }
    }
    return maxProfit
}
