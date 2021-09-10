package sample

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ityulkanov/techschool/pb"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCUPBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberThreads: uint32(numberThreads),
		NumberCores:   uint32(numberCores),
		MinGhz:        float32(minGhz),
		MaxGhz:        float32(maxGhz),
	}
	return cpu
}

func newGPU() *pb.GPU {
	minGhz := randomFloat64(2, 8)
	maxGhz := randomFloat64(minGhz, 24)
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

func newRam() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}

func newSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

func newHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}

func newScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitoch:  randomBool(),
	}
	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:          uuid.New().String(),
		Brand:       brand,
		Name:        name,
		Cpu:         NewCPU(),
		Ram:         newRam(),
		Gpus:        []*pb.GPU{newGPU()},
		Storages:    []*pb.Storage{newSSD(), newHDD()},
		Screen:      newScreen(),
		Keyboard:    NewKeyboard(),
		Weight:      &pb.Laptop_WeightKg{WeightKg: randomFloat64(1.0, 3.0)},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: float64(randomInt(2000, 2021)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return laptop
}
