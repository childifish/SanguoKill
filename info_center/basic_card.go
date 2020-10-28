package info_center

import (
	"fmt"
)

//杀（行为）
type Attack interface {
}

//闪避（行为）
//可能可以给八卦阵之类的用到
type Dodge interface {
}

//杀
type Kill struct {
	Card
	Attack
}

//只能取距离为一的牌
func (t Kill)Choose(chain PlayerChain)(re []Target)  {
	for i := 0; i < len(chain.Players); i++ {
		if chain.Players[i].Name!= NowPlayer.Name{
			re = append(re,&chain.Players[i])
		}
	}
	return re
}
//检查是否可行
func (t Kill)Check()bool  {
	if NowPlayer.AttackNum >=1{
		return true
	}
	fmt.Println("没有攻击次数")
	return false
}
//具体操作
func (t Kill)Do()  {

}

func (t Kill)NameIs()string  {
	return "杀"
}

func (t Kill)AskAndEffect(target *Target)  {
	NowPlayer.AttackNum--
	var value = *target
	Player,ok := value.(*Player)
	able := Player.ChooseAble()
	if !able {
		fmt.Println("不合法的目标")
		return
	}
	response := Player.Response(t)
	//没有反应
	if !response {
		if ok{
			fmt.Println("成功造成伤害")
			fmt.Println(Player.Hp)
			Player.Hurt(1)
			fmt.Println(Player.Hp)
			return
		}
		fmt.Println("有问题")
	}
	fmt.Println("阻止成功")
}

//返回本身
func (t Kill)Self() Targeter {
	return &t
}

//返回本身
func (t Kill)SelfIsTargeter()(bool, Targeter)  {
	return true,&t
}

//返回可以响应该牌的对象
func (t Kill)Need()(re []Responser)  {
	re = append(re, Evade{})
	return re
}

//返回该牌能够响应的牌
func (t Kill)AbleResponse()(ta []Targeter)  {
	return ta
}

//闪
type Evade struct {
	Card
	Dodge
}


func (e Evade)Use()bool  {
	return true
}

//可以响应的牌
func (e Evade)AbleResponse()(ta []Targeter)  {
	ta = append(ta, Kill{})
	return
}

//能被响应
func (e Evade)Need()(re []Responser)  {
	return
}

//是主动并选择对象的牌吗
func (e Evade)SelfIsTargeter()(bool, Targeter)  {
	return false,nil
}

//返回自己的名字
func (e Evade)SelfNameIs()string  {
	return "闪"
}

type Peach struct {
	Card
}

func (p Peach) Choose(chain PlayerChain)(re []Target) {
	if NowPlayer.Hp+1 > NowPlayer.Hero.HeroHp{
		return re
	}
	re = append(re, NowPlayer)
	return
}
func (p Peach) AskAndEffect(target *Target) {
	fmt.Println(NowPlayer.Name,"现在", NowPlayer.Hp)
	NowPlayer.Heal(1)
	fmt.Println("治愈了", NowPlayer.Hp)

}
func (p Peach) Self() Targeter {
	return p
}
func (p Peach) NameIs() string {
	return "Peach"
}
func (p Peach) Use()bool  {
	return true
}
func (p Peach) Need()[]Responser {
	return nil
}
func (p Peach) AbleResponse()[]Targeter {
	return nil
}
func (p Peach) SelfIsTargeter()(bool, Targeter){
	return true,&p
}
func (p Peach) Check()bool  {
	return true
}
func (p Peach) Do() {

}
