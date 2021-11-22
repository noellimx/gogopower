package person

import "fmt"

type Ranger struct {
	ShoutValue string
}

func (r Ranger) Shout() string {
	return r.ShoutValue
}

func (r *Ranger) ChangeShoutValue(newVal string) {
	r.ShoutValue = newVal
}

func (r Ranger) ChangeShoutValue2(newVal string) {
	r.ShoutValue = newVal
}

func shout(r *Ranger) {

}

type Shouter interface {
	Shout() string
}

func triggerShout(shouter Shouter) {
	fmt.Print(shouter.Shout())
}
