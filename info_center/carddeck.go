package info_center

import (
	"fmt"
	"math/rand"
)

type CardDeck struct {
	MainDeck []Card
}

//初始化卡组
func InitCardDeck(i int)(deck CardDeck)  {
	deck.MainDeck = deck.NewDeck(i)
	return
}

//获取牌库里的所有牌
func (c *CardDeck)GetCardDeck()[]Card {
	return c.MainDeck
}

//获取n张牌
func (c *CardDeck)GetCard(num int)(re []Card)  {
	for i:=0;i<num;i++{
		c.CheckShuffle()
		re = append(re,c.MainDeck[len(c.MainDeck)-1])
		c.MainDeck = c.MainDeck[:len(c.MainDeck)-1]
	}
	return re
}

//查看牌顶.//判定
func (c *CardDeck)CheckTop() Card {
	re := c.MainDeck[len(c.MainDeck)-1]
	c.MainDeck = c.MainDeck[:len(c.MainDeck)-1]
	fmt.Println("判定结果")
	PrintCard(re)
	return re
}

//检查是否需要洗牌
func (c *CardDeck)CheckShuffle()  {
	if len(c.MainDeck)<=1{
		//这样可以保证洗牌时最上方卡为剩余的那张
		last := c.MainDeck[0]
		c.MainDeck = c.NewDeck(40)
		c.MainDeck = append(c.MainDeck,last)
	}
}

//返回一个新的牌组
func (c *CardDeck)NewDeck(j int)(deck []Card)  {
	for i := 0; i < j; i++ {
		deck = append(deck, RandCard())
	}
	return deck
}

func (c *CardDeck)PrintDeck()  {
	PrintCards(c.MainDeck)
}

func RandCard() Card {
	intn := rand.Intn(6)
	switch intn{
	case 0:
		return Card{
			Poker:  RandPoker(),
			User:   "",
			Name:   "杀",
			Effect: Kill{},
		}
	case 1:
		return Card{
			Poker:  RandPoker(),
			User:   "",
			Name:   "闪",
			Effect: Evade{},
		}
	case 2:
		return Card{
			Poker:  RandPoker().Goto("red"),
			User:   "",
			Name:   "桃",
			Effect: Peach{},
		}
	case 3:
		return Card{
			Poker:  RandPoker().Goto("red"),
			User:   "",
			Name:   "桃园结义",
			Effect: PeachesUnion{},
		}
	case 4:
		return Card{
			Poker:  RandPoker(),
			User:   "",
			Name:   "无懈可击",
			Effect: Waitertight{},
		}
	case 5:
		return Card{
			Poker:  RandPoker(),
			User:   "",
			Name:   "无中生有",
			Effect: CardAppearInVoid{},
		}
	default:
		return Card{}
	}
}

