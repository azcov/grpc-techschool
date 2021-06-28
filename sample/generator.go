package sample

import (
	"github.com/azcov/grpc-techschool/pb"
	"github.com/golang/protobuf/ptypes"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	cores := randomInt(2, 8)
	minGhz := randomFloat64(2.0, 3.5)
	return &pb.CPU{
		Brand:         brand,
		Name:          randomCPUName(brand),
		NumberCores:   uint32(cores),
		NumberThreads: uint32(randomInt(cores, cores*2)),
		MinGhz:        minGhz,
		MaxGhz:        randomFloat64(minGhz, minGhz*2),
	}
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	minGhz := randomFloat64(1.0, 1.5)
	return &pb.GPU{
		Brand:  brand,
		Name:   randomGPUName(brand),
		MinGhz: minGhz,
		MaxGhz: randomFloat64(minGhz, minGhz*2),
		Memory: &pb.Memory{
			Value: uint64(randomInt(2, 8)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
}

func NewRAM() *pb.Memory {
	return &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}
}

func NewSSD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
}

func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
}

func NewScreen() *pb.Screen {

	return &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	return &pb.Laptop{
		Id:       randomUUID().String(),
		Brand:    brand,
		Name:     randomLaptopname(brand),
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(7, 15),
		ReleaseYear: uint32(randomInt(2018, 2021)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
}
