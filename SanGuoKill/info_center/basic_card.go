package info_center

import "fmt"

//杀（行为）
type Attack interface {
}

//闪避（行为）
type Dodge interface {
}

//杀
type Kill struct {
	Card
	Attack
}

//只能取距离为一的牌
func (t Kill)Choose(chain PlayerChain)(re []Target)  {
	for _, i2 := range chain.Players {
		if i2.Name!= NowPlayer.Name{
			re = append(re,&i2)
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

func (t Kill)AskAndEffect(target Target)  {
	player,ok := target.(*Player)
	able := player.ChooseAble()
	if !able {
		fmt.Println("不合法的目标")
		return
	}
	response := player.Response(t)
	//没有反应
	if !response {
		if ok{
			fmt.Println("成功造成伤害")
			fmt.Println(player.Hp)
			player.Hurt(1)
			fmt.Println(player.Hp)
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

//不能被直接使用
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

func (p Peach)Use()bool  {
	return true
}
