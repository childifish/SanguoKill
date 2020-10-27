package info_center

import (
	"fmt"
)

const (
	//基础牌
	//判定牌
	//锦囊牌
)

var CadrIndex = []CardEffect{Kill{},Evade{},Peach{},PeachesUnion{},Waitertight{}}

//只是牌本身的信息，没有其他有关技能，效果的信息
type Card struct {
	Poker
	User   string
	Name   string
	Effect CardEffect
}

func PrintCards(c []Card)  {
	for i, i2 := range c {
		fmt.Printf("序号%d:花色%s,卡牌名称:%v\n",i+1,i2.PrintPoker(),i2.Name)
	}
}
func PrintCard(c Card)  {
	fmt.Printf("花色%s,卡牌名称:%v\n",c.PrintPoker(),c.Name)
}

func (c *Card)Use()bool{

	var nowTarget Targeter
	is,targeter := c.Effect.SelfIsTargeter()
	//使用卡牌非当前回合玩家--》可能在响应
	//这里得改改。。。
	if c.User != NowPlayer.Name &&(c.Name == "闪"||c.Name == "无懈可击"){
		return true
	}
	if !is{
		fmt.Println(c.Name,"不能这样使用")
		return false
	}

	nowTarget = targeter

	ok := nowTarget.Check()

	if !ok{
		return false
	}

	nowTarget.Do()

	Targets := nowTarget.Choose(*Players)
	if len(Targets)<1{
		fmt.Println("无可选择对象")
		return false
	}

	realTarget :=  ChooseTarget(Targets)

	nowTarget.AskAndEffect(&realTarget)

	return  true

}

func (c Card)ChooseAble()bool  {
	return true
}

func (c Card)Response(Targeter)bool  {
	return true
}

func ChooseTarget(targets []Target) Target {
	var i int
	fmt.Println("输入你想指定的目标")
	for _, i3 := range targets {
		player := i3.(*Player)
		player.PrintPlayer()
	}
	scanln, err := fmt.Scanln(&i)
	if err != nil{
		fmt.Println("选取错误")
		return nil
	}
	return targets[scanln-1]
}