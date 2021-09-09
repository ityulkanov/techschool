package sample

import "github.com/ityulkanov/techschool/pb"

func NewKeyboard *pb.Keyboard {
	keyboard := &pb.Keyboard {
		Layout: randomKeyboardLayout(), 
		Backlit: randomBool();
	}

	return keyboard
}

func NewCPU() *pb.CPU {
	cpu := &pb.CPU{
		Brand: randomCPUBrand()
	}
	return cpu
}