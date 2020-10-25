package info_center

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

func (g GuanYuSkill)Choose(p PlayerChain)(re []Target) {
	for i := 0; i < len(NowPlayer.HandCard); i++ {
		if NowPlayer.HandCard[i].Flower==1||NowPlayer.HandCard[i].Flower==3{
			re = append(re,NowPlayer.HandCard[i])
		}
	}
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