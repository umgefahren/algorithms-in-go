package algorithms_in_go

import "testing"

func TestDh(t *testing.T) {
	sp := NewSharedParams()
	aS := NewSecret(sp)
	bS := NewSecret(sp)
	yA := aS.computeY()
	yB := bS.computeY()
	kAB := aS.computeShared(yB)
	kBA := bS.computeShared(yA)
	if kBA.Cmp(kAB) != 0 {
		t.Errorf("The shared secret %v and %v are not equal.", kAB, kBA)
	}
}
