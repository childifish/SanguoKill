package info_center

import "fmt"

type Player struct {
	Name         string
	Hp           int
	ID           int
	Hero         Hero
	HandCard     []Card
	AttackNum    int
	CheckingList []Card
	Equipments   []Card
	Situation    string
}


func (p *Player)Do()  {
	NowPlayer.PrintPlayer()
	//回合开始
	p.Situation = "回合开始阶段"
	p.AttackNum = 1
	skill := p.CheckHeroSkill()
	if skill{
		p.Hero.Skill.HeroDo()
	}

	//判定阶段
	p.Situation = "判定阶段"
	p.PrintSituation()
	p.CheckChecking()

	//抽牌阶段--判定技能
	p.Situation = "抽牌阶段"
	p.Draw(2)

	//出牌阶段
	p.Situation = "出牌阶段"
	p.PlayCards()

	//弃牌阶段
	p.Situation = "弃牌阶段"
	p.Discarding()

	//回合结束
	p.Situation = "回合开始阶段"

	p.Situation = "sleep"


}

func (p *Player)PlayCards()  {
	for {
		p.PrintSituation()
		//获取要用的牌
		cardID := p.ChooseCard()
		switch cardID {
		case 1024:
			break
		case 2048:
			p.ChooseSkill()

		default:

		}

		p.UseCard(cardID)

		if p.Ask("是否结束出牌？0--不结束 1--结束"){
			break
		}
	}
}

func (p *Player)CheckChecking()  {
	if len(p.CheckingList)>0{
		fmt.Println("存在判定牌")
		for i := 0; i < len(p.CheckingList); i++ {
			fmt.Printf("进行判定:%s---\n",p.CheckingList[i])
		}
	}
}


func (p *Player)ChooseSkill()  {
	fmt.Printf("发动了%s的技能\n",p.Hero.HeroName)
	choose := p.Hero.Skill.Choose(*Players)
	if len(choose) == 0{
		fmt.Println("无可选择目标")
		return
	}

	for i := 0; i < len(choose); i++ {
		fmt.Print("可选择的有,序号：",i)
		PrintCard(choose[i].(Card))
	}
	var j int
	_, err := fmt.Scanln(&j)
	if err != nil{
		return
	}
	p.Hero.Skill.AskAndEffect(&choose[j])

}

func (p *Player)Discarding() {
	for {
		if p.Ok2Pass(){
			break
		}
		fmt.Printf("手牌数量超过生命值，还须弃置%d张牌\n",len(p.HandCard)-p.Hp)
		p.Discard()

	}
}

func ChooseCard(c []Card)int  {
	PrintCards(c)
	var i int
	_, err := fmt.Scan(&i)
	if err != nil{
		return ChooseCard(c)
	}
	if len(c)<i{
		return ChooseCard(c)
	}
	return i
}


//弃牌
func (p *Player)Discard()  {
	i := ChooseCard(p.HandCard)
	p.HandCard = append(p.HandCard[:i-1], p.HandCard[i:]...)
}

//判断是否需要弃牌
func (p *Player)Ok2Pass()bool  {
	if p.Hp >= len(p.HandCard){
		return true
	}
	return false
}

//抽牌
func (p *Player)Draw(i int)[]Card {
	for _, i3 := range Deck.GetCard(i) {
		i3.User = p.Name
		p.HandCard = append(p.HandCard,i3)
	}
	return p.HandCard


}

//判断是否可以被选择
func (p *Player)ChooseAble()bool   {
 	return true
}

//进行响应
func (p *Player)Response(targeter Targeter)bool   {
	fmt.Println("需要",targeter.Need()[0].SelfNameIs(),"进行响应，输入0放弃响应")
	able := p.FindResponse(targeter)
	if len(able) == 0{
		fmt.Println("无可以响应的牌")
		return false
	}
	i := ChooseCard(able)
	if i == 0{
		fmt.Println("取消响应")
		return false
	}
	ok := able[i-1].Use()
	if ok {
		fmt.Println("响应完成")
		if len(p.HandCard) == 0{
			p.HandCard = nil
		}
		p.HandCard = append(p.HandCard[:i], p.HandCard[i+1:]...)
		return true
	}
	return false
}

func (p *Player)NameIs()string  {
	return  p.Name
}

//寻找可以响应的牌
func (p *Player)FindResponse(targeter Targeter)(re []Card)  {
	//在手牌区找有无可以响应的牌
	println(p.HandCard)
	for _, card := range p.HandCard {
		for _, resp := range card.Effect.AbleResponse() {
			if targeter.Self().NameIs() == resp.NameIs(){
				fmt.Println("finding ...",card.Name,"可以响应:",resp.NameIs())
				re = append(re,card)
			}
		}
	}
	//在装备区找有无可以响应的牌
	for _, card := range p.Equipments {
		for _, resp := range card.Effect.AbleResponse() {
			if targeter.Self().NameIs() == resp.NameIs(){
				re = append(re,card)
			}
		}
	}
	return re
}

//询问是/否
func (p *Player)Ask(s string)bool  {
	var i int
	fmt.Println(s)
	_, err := fmt.Scan(&i)
	if err != nil{
		return p.Ask(s)
	}
	switch i {
	case 0:
		return false
	case 1:
		return true
	default:
		return p.Ask(s)
	}
}

//打印手牌
func (p *Player)PrintHandCard()  {
	PrintCards(p.HandCard)
}

//使用手牌
func (p *Player)UseCard(i int)  {
	//使用牌
	if i >= len(p.HandCard) {
		return
	}

	ok := p.HandCard[i].Use()
	if !ok{
		fmt.Println("使用失败")
		return
	}

	//删除选择的手牌
	p.HandCard = append(p.HandCard[:i], p.HandCard[i+1:]...)
}

//选择手牌
func (p *Player)ChooseCard()int {
	PrintCards(p.HandCard)
	var id int
	fmt.Println("输入使用牌的序号,输入1024以终止出牌阶段,输入2048以进入技能选择,当前攻击次数剩余：",p.AttackNum)
	_, err := fmt.Scan(&id)
	if err!= nil{
		return p.ChooseCard()
	}
	if id==1024||id==2048{
		return id
	}
	if len(p.HandCard)<id||id<1{
		fmt.Println("序号错误")
		return p.ChooseCard()
	}	
	return id-1
}

func (p *Player)PrintSituation()  {
	fmt.Printf("-------现在是%s的%s阶段--------\n",p.Name,p.Situation)
	fmt.Printf("攻击次数为%d\n",p.AttackNum)
}

func (p *Player)PrintPlayer()  {
	fmt.Printf("玩家姓名：%s 选择英雄：%s,id:%d 剩余HP: %d\n",p.Name,p.Hero.HeroName,p.ID,p.Hp)
}

//加血
func (p *Player)Heal(i int)bool  {
	if p.Hp+i > p.Hero.HeroHp {
		return false
	}
	p.Hp += i
	return true
}

//受伤
func (p *Player)Hurt(i int)  {
	Result := p.Hp - i
	p.Hp = Result
	if p.Hp <= 0 {
		//找有没有桃
		Players.Killed(p.ID)
	}
}

func (p *Player)CheckHeroSkill()bool  {
	for i := 0; i < len(p.Hero.SkillTimer); i++ {
		if p.Situation == p.Hero.SkillTimer[i]{
			return true
		}
	}
	return false
}