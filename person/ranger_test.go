package person

import "testing"

func TestRangerShout(t *testing.T) {
	ranger := Ranger{}
	ranger.Shout()

	rangerPtr := &ranger
	rangerPtr.Shout()

	triggerShout(ranger)
	triggerShout(rangerPtr)

	shout(&ranger)

	ranger.ChangeShoutValue("shout!")
	t.Logf("first %s", ranger.Shout())

	ranger.ChangeShoutValue2("s22hout!")
	t.Logf("first %s", ranger.Shout())
}
