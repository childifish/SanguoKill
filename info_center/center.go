package info_center

import (
	"fmt"
)

var Deck CardDeck
var Players *PlayerChain

var NowPlayer *Player

var turn int


func Start()  {



	Deck = InitCardDeck(40)
	Deck.PrintDeck()

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
		fmt.Println(Players.Players[0],"获得了胜利")
		return true
	}
	return false
}