package info_center

import "fmt"

type Position int

func (p Position)CalculatePos(targetID int,chain PlayerChain)int  {
	return 1
}

//寻找周围的人，want是可以的距离
//todo: 还没用，应该在Check中使用
func (p Position)FindNearBy(player Player,chain PlayerChain,want int)[]Player {
	var pos int
	for i, i2 := range chain.Players {
		if i2.ID == player.ID {
			pos = i
		}
	}
	fmt.Println("现在处于",pos)
	return chain.Players
}
