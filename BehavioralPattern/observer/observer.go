package observer

import "fmt"

type Observer interface {
	RevMsg(wd int)
}

type User struct {
	Name string
}

func (u User) RevMsg(wd int) {
	fmt.Printf("%s 现在温度 %d\n", u.Name, wd)
}

type Subject interface {
	SendMsg(wd int)
	Notify()
	RegisterObserver(Observer)
}

type WeatherStation struct {
	observerList []Observer
	wd           int
}

func (w *WeatherStation) SendMsg(wd int) {
	w.wd = wd
	w.Notify()
}

func (w *WeatherStation) Notify() {
	for _, observer := range w.observerList {
		observer.RevMsg(w.wd)
	}
}

func (w *WeatherStation) RegisterObserver(observer Observer) {
	w.observerList = append(w.observerList, observer)
}
