// 中介者模式（Mediator Pattern）是用来降低多个对象和类之间的通信复杂性。这种模式提供了一个中介类，
// 该类通常处理不同类之间的通信，并支持松耦合，使代码易于维护。中介者模式属于行为型模式。
// 用一个中介对象来封装一系列的对象交互，中介者使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互。
package mediator

import (
	"fmt"
)

type Country interface {
	SendMessage(message string)
	GetMessage(message string)
}
type USA struct {
	UnitedNations
}

func (usa USA) SendMessage(message string) {
	usa.UnitedNations.ForwardMessage(message, usa)
}

func (usa USA) GetMessage(message string) {
	fmt.Printf("美国收到对方消息: %s\n", message)
}

type Iraq struct {
	UnitedNations
}

func (iraq Iraq) SendMessage(message string) {
	iraq.UnitedNations.ForwardMessage(message, iraq)
}

func (iraq Iraq) GetMessage(message string) {
	fmt.Printf("伊拉克收到对方消息: %s\n", message)
}

type UnitedNations interface {
	ForwardMessage(message string, country Country)
}

type UnitedNationsSecurityCouncil struct {
	USA
	Iraq
}

func (unsc *UnitedNationsSecurityCouncil) ForwardMessage(message string, country Country) {
	switch country.(type) {
	case USA:
		unsc.Iraq.GetMessage(message)
	case Iraq:
		unsc.USA.GetMessage(message)
	default:
		fmt.Println("The country is not a member of UNSC")
	}

}
