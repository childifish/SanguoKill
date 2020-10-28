package info_center

import (
	"fmt"
)

var Magics = []Targeter{PeachesUnion{},Waitertight{},CardAppearInVoid{}}

//todo: 锦囊牌
type Magic struct {
	Card
}

func (ma *Magic)AskAll(t Targeter)bool  {
	for i := 0; i < len(Players.Players); i++ {
		if Players.Players[i].Name == ma.User{
			continue
		}
		fmt.Printf("现在是%s进行响应：",Players.Players[i].Name)
		response := Players.Players[i].Response(t)
		if response {
			return false
		}
	}
	return true
}

type PeachesUnion struct {
	Magic
}

func (p PeachesUnion) Choose(chain PlayerChain)(re []Target) {
	re = append(re, NowPlayer)
	return
}
func (p PeachesUnion) AskAndEffect(target *Target) {
	ok := p.AskAll(p)
	if !ok{
		return
	}
	for i := 0; i < len(Players.Players); i++ {
		fmt.Println(Players.Players[i].Name,"现在", Players.Players[i].Hp)
		Players.Players[i].Heal(1)
		fmt.Println("治愈了", Players.Players[i].Hp)
	}
}
func (p PeachesUnion) Self() Targeter {
	return p
}
func (p PeachesUnion) NameIs() string {
	return "桃园结义"
}
func (p PeachesUnion) Use()bool  {
	return true
}
func (p PeachesUnion) Need()(re []Responser) {
	re = append(re, &Waitertight{})
	return
}
func (p PeachesUnion) AbleResponse()[]Targeter {
	return nil
}
func (p PeachesUnion) SelfIsTargeter()(bool, Targeter){
	return true,&p
}
func (p PeachesUnion) Check()bool  {
	return true
}
func (p PeachesUnion) Do() {

}


//无懈可击
type Waitertight struct {
	Magic
}

func (w Waitertight) SelfIsTargeter()(bool, Targeter){
	return true,&w
}
func (w Waitertight) Check() bool {
	return true
}
func (w Waitertight) Do() {

}
func (w Waitertight) Choose(chain PlayerChain) []Target {
	return nil
}
func (w Waitertight) AskAndEffect(target *Target) {
	ok := w.AskAll(w)
	if !ok{
		return
	}
}
func (w Waitertight) Self() Targeter {
	return w
}
func (w Waitertight) Need()(re []Responser) {
	re = append(re,&Waitertight{})
	return
}
func (w Waitertight) NameIs() string {
	return "无懈可击"
}
func (w Waitertight) AbleResponse()(ta []Targeter) {
	ta = append(ta,Magics...)
	return
}
func (w Waitertight) SelfNameIs() string {
	return "无懈可击"
}


//无中生有
type CardAppearInVoid struct {
	Magic
}

func (caiv CardAppearInVoid) Check() bool {
	return true
}
func (caiv CardAppearInVoid) Do() {
}
func (caiv CardAppearInVoid) Choose(chain PlayerChain)(re []Target) {
	re = append(re, NowPlayer)
	return
}
func (caiv CardAppearInVoid) AskAndEffect(target *Target) {
	ok := caiv.AskAll(caiv)
	if !ok{
		return
	}
	NowPlayer.Draw(2)
}
func (caiv CardAppearInVoid) Self() Targeter {
	return caiv
}
func (caiv CardAppearInVoid) Need()(re []Responser) {
	re = append(re, &Waitertight{})
	return
}
func (caiv CardAppearInVoid) NameIs() string {
	return "无中生有"
}
func (caiv CardAppearInVoid) SelfIsTargeter()(bool, Targeter){
	return true,&caiv
}
func (caiv CardAppearInVoid) AbleResponse()[]Targeter {
	return nil
}
