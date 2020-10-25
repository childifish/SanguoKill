package info_center

import (
	"fmt"
)

const (
	//基础牌
	//判定牌
	//锦囊牌
)

//只是牌本身的信息，没有其他有关技能，效果的信息
type Card struct {
	Poker
	User   string
	Name   string
	Effect CardEffect
}

func PrintCards(c []Card)  {
	for i, i2 := range c {
		fmt.Printf("序号%d:花色%s,卡牌信息:%v\n",i+1,i2.PrintPoker(),i2)
	}
}

func (c *Card)Use()bool{

	var nowTarget Targeter
	is,targeter := c.Effect.SelfIsTargeter()
	//使用卡牌非当前回合玩家--》可能在响应
	if c.User != NowPlayer.Name &&c.Name == "闪"{
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

	realTarget :=  ChooseTarget(Targets)

	nowTarget.AskAndEffect(realTarget)

	//todo: 不能成功掉血
	//fmt.Println(Players)

	return  true

}

func ChooseTarget(targets []Target) Target {
	var i int
	fmt.Println("输入你想指定的目标")
	for _, i3 := range targets {
		fmt.Println(i3)
	}
	scanln, err := fmt.Scanln(&i)
	if err != nil{
		fmt.Println("选取错误")
		return nil
	}
	return targets[scanln-1]
}