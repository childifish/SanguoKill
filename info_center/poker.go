package info_center

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)


type Poker struct {
	Num int
	Flower int
}

func RandPoker()(p Poker)  {
	p.Num = rand.Intn(12)
	p.Flower =rand.Intn(3)
	return p
}

//变成
func (p Poker)Goto(para string)Poker  {
	switch para {
	case "red":
		if time.Now().Unix()%2==0 {
			p.Flower = 1
		}else {
			p.Flower = 3
		}
	case "black":
		if time.Now().Unix()%2==0 {
			p.Flower = 0
		}else {
			p.Flower = 2
		}
	}
	return p
}

//拼点
func (p Poker)Comparison(poker Poker)bool  {
	if p.Num > poker.Num{
		return true
	}
	if p.Num == poker.Num&&p.Flower < poker.Flower{
		return true
	}
	return false
}

func (p Poker)PrintPoker()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 12:
		buffer += "2"
	case 11:
		buffer += "A"
	case 10:
		buffer += "K"
	case 9:
		buffer += "Q"
	case 8:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}