package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p Product) TotalValue() float64 {
	return p.Price * float64(p.Stock)
}

func (p Product) IsInStock() bool {
	return p.Stock > 0
}

func (p Product) Info() string {
	return fmt.Sprintf("商品: %v, 单价: ¥%v, 库存: %v件", p.Name, p.Price, p.Stock)
}

func (p *Product) Restock(amount int) {
	p.Stock += amount
}

func (p *Product) Sell(amount int) (success bool, message string) {
	if p.Stock <= 0 || p.Stock-amount <= 0 {
		return false, "库存不足"
	}

	p.Stock -= amount
	return true, "售卖成功"
}
