package info_center

import "fmt"

type GuanYuSkill struct {
	Hero
}

//选一张牌，将他变为杀
//不对头
func (g GuanYuSkill)AskAndEffect(target *Target)  {
	var value = *target
	card := value.(Card)
	card.Effect = Kill{}
	card.Use()
}

func (g GuanYuSkill)HeroDo()  {

}

func (g GuanYuSkill)MightBeCard()(cf []CardEffect)  {
	cf = append(cf,Kill{})
	return
}

func (g GuanYuSkill)Choose(p PlayerChain)(re []Target) {
	for i := 0; i < len(NowPlayer.HandCard); i++ {
		if NowPlayer.HandCard[i].Flower==1||NowPlayer.HandCard[i].Flower==3{
			re = append(re,NowPlayer.HandCard[i])
		}
	}
	return
}

type ZhenjiSkill struct {
	Hero
}

func (z ZhenjiSkill)HeroDo()  {
	fmt.Println("可以发动洛神")
	for  {
		top := Deck.CheckTop()
		if top.Poker.Flower == 0||top.Poker.Flower == 2{
			fmt.Println("发动洛神")
			NowPlayer.HandCard = append(NowPlayer.HandCard,top)
			continue
		}else {
			fmt.Println("发动失败")
			return
		}
	}

}

func (z ZhenjiSkill)AskAndEffect(target *Target) {
	fmt.Println("不能直接使用")
}

func (z ZhenjiSkill)Choose(p PlayerChain)(re []Target)  {
	for i := 0; i < len(NowPlayer.HandCard); i++ {
		if NowPlayer.HandCard[i].Flower==0||NowPlayer.HandCard[i].Flower==2{
			re = append(re,NowPlayer.HandCard[i])
		}
	}
	return
}

func (z ZhenjiSkill)MightBeCard()(cf []CardEffect)  {
	cf = append(cf,Evade{})
	return
}

//这里加选择英雄
func ChooseHero(string2 string)Hero  {
	switch string2 {
	case "关羽":
		return Hero{
			HeroName: "关羽",
			User:     "",
			HeroHp:   4,
			Skill:    GuanYuSkill{},
			SkillTimer: nil,
		}
	case "甄姬":
		return Hero{
			HeroName: "甄姬",
			User:     "",
			HeroHp:   3,
			Skill:    ZhenjiSkill{},
			SkillTimer: []string{"回合开始阶段"},
		}
	default:
		return Hero{
			HeroName: "关羽",
			User:     "",
			HeroHp:   4,
			Skill:    GuanYuSkill{},
		}
	}
}