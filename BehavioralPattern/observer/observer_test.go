package observer

import "testing"

func Test_Observer(t *testing.T) {
	station := WeatherStation{}
	u1 := User{Name: "zhangsan"}
	u2 := User{Name: "lisi"}

	station.RegisterObserver(&u1)
	station.RegisterObserver(&u2)

	station.SendMsg(8)
	station.SendMsg(0)
}
