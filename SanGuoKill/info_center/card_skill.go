package info_center

type CardEffect interface {
	Need()[]Responser
	AbleResponse()[]Targeter
	SelfIsTargeter()(bool, Targeter)
}

type CardType interface{
}


//判定牌
type DetermineCard interface {
	//Begin(card Card, self playerchan.Player) // 每回合开始的时候做的事情
	Operation
}

//锦囊牌
type KitsCard interface {
	Operation
}

