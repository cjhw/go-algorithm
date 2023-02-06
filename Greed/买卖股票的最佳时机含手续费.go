package Greed

func maxProfit1(prices []int, fee int) int {
	minBuy := prices[0]
	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minBuy {
			minBuy = prices[i]
		}

		if prices[i] > minBuy && prices[i] < fee+minBuy {
			continue
		}

		if prices[i] > fee+minBuy {
			res += prices[i] - fee - minBuy
			//更新最小值（如果还在收获利润的区间里，表示并不是真正的卖出，而计算利润每次都要减去手续费，所以要让minBuy = prices[i] - fee;，这样在明天收获利润的时候，才不会多减一次手续费！）
			minBuy = prices[i] - fee
		}
	}

	return res
}
