package info_center

type Hero struct {
	HeroName string
	User   string
	HeroHp int
	Skill Skill
	SkillTimer []string
}

//锁定技--触发条件达到就必须触发
//比较特殊
type Passive interface {
	Situation()
}

//限定技--全局只能触发一次
//在游戏开始时声明该变量
type OnlySkill interface {
	IsAble()bool
	Operation
}

//觉醒技=限定技+锁定技
type AwakeningSkill interface {
	Passive
	OnlySkill
}

//基础技能
type BasicSkill interface {
	Operation
}

type Skill interface {
	Operation
	MightBeCard()[]CardEffect
	HeroDo()
}
