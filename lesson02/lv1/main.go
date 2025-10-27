package main

import "fmt"

func main() {
	p := &Product{
		Name:  "Go圣经",
		Price: 89.5,
		Stock: 10,
	}

	success, _ := p.Sell(5)
	if success {
		fmt.Printf(" 售卖%v本: 成功, 剩余库存: %v\n", 5, p.Stock)
	} else {
		fmt.Println("失败, 库存不足")
	}

	p.Restock(20)
	fmt.Printf("进货20本, 当前库存: %v\n", p.Stock)

	success, _ = p.Sell(30)
	if success {
		fmt.Printf(" 售卖%v本: 成功, 剩余库存: %v\n", 30, p.Stock)
	} else {
		fmt.Println("失败, 库存不足")
	}

	fmt.Printf("%v\n", p.Info())
	fmt.Printf("库存总价值: ¥%v", p.TotalValue())
}
