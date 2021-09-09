package sample

import (
	"math/rand"

	"github.com/grpc-go/examples/techSchool/pb"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1: 
		return pb.Keyboard_QWERTY
	case 2: 
		return pb.Keyboard_QWETZ
	default: 
		return pb.Keyboard_AZERTY
	}

}

func randomBool() bool {
	return rand.Intn(2) == 1 
	
}


