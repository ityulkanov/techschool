package sample

import (
	"math/rand"

	"github.com/ityulkanov/techschool/pb"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return gen.Keyboard_AZERTY
	case 2:
		return pb.Keyboard_QWETZ
	default:
		return pb.Keyboard_AZERTY
	}

}

func randomBool() bool {
	return rand.Intn(2) == 1

}
