package info_center

//基础操作
type Operation interface {
	Choose(PlayerChain)[]Target //选择目标
	AskAndEffect(Target)        //询问对象响应--若有响应，
	//AskAndEffect(Target)
}

//作为目标
type Target interface {
	ChooseAble()bool 	//是否能够被选取
	Response(Targeter)bool
}

//在阶段切换时 进行一个“行为”，比如洛神（检查是否处于出牌阶段，）
//技能和卡牌要声明自己的有效阶段
//比如出了无懈可击的锦囊牌，只能在出牌阶段使用
type Action interface {
	Check()bool    //预先检查--（阶段是否正确，）
	Do()   	//实际操作//没有目标的不需要Do
}

//能指定目标的操作(主动牌/技能)
type Targeter interface {
	Action
	Operation
	Self() Targeter
	Need()[]Responser
	NameIs()string
}

//能够响应的牌
type Responser interface {
	Use()bool
	AbleResponse()[]Targeter
	SelfNameIs()string
}

type Result interface {

}