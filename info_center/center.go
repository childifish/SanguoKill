package info_center

import (
	"fmt"
)

var Deck CardDeck
var Players *PlayerChain

var NowPlayer *Player

var turn int


func Start()  {

	//初始化卡池
	Deck = InitCardDeck(80)
	Deck.PrintDeck()

	//人数
	Players = InitPlayer(4)
	Players.PrintPlayer()

	for i := 0; i < len(Players.Players); i++ {
		Players.Players[i].Draw(4)
	}

	NowPlayer = &Players.Players[0]

	for {

		NowPlayer.Do()

		NowPlayer = Players.ChooseNextPlayer(NowPlayer)

		turn++

		//if turn >10 {
		//	Players.Killed(Players.Len()-1)
		//}

		if WinCondition(){
			//结束游戏
			break
		}
	}


}

//胜利条件
//现在是活到最后的赢
func WinCondition()bool  {
	if Players.Len()==1{
		Players.Players[0].PrintPlayer()
		fmt.Println("获得了胜利")
		return true
	}
	return false
}