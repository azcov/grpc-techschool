package sample

import (
	"math/rand"

	"github.com/azcov/grpc-techschool/pb"
	"github.com/google/uuid"
)

func randomStringFromSet(args ...string) string {
	n := len(args)
	if n == 0 {
		return ""
	}
	return args[rand.Intn(n)]
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "Apple")
}

func randomCPUName(brand string) string {
	switch brand {
	case "Intel":
		return randomStringFromSet(
			"Core i9-11xxxHK",
			"Core i7-11xxxH",
			"Core i5-11xxxF",
			"Core i3-11xxxF",
		)
	case "AMD":
		return randomStringFromSet(
			"Core i9-11xxxHK",
			"Core i7-11xxxH",
			"Core i5-11xxxF",
			"Core i3-11xxxF",
		)
	default:
		return randomStringFromSet(
			"Clone",
			"OEM",
			"Unknown",
			"Illegal",
		)
	}
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("Nvidia", "AMD", "Apple")
}

func randomGPUName(brand string) string {
	switch brand {
	case "Nvidia":
		return randomStringFromSet(
			"RTX 3080Ti",
			"RTX 2080Ti",
			"GTX 1660Ti",
			"GTX 1080Ti",
		)
	case "AMD":
		return randomStringFromSet(
			"Radeon RX 6900 XT",
			"Radeon RX 5900 XT",
			"Radeon RX 580",
			"Radeon RX 480",
		)
	default:
		return randomStringFromSet(
			"Clone",
			"OEM",
			"Unknown",
			"Illegal",
		)
	}
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	return &pb.Screen_Resolution{
		Width:  uint32(height * 16 / 9),
		Height: uint32(height),
	}
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomUUID() uuid.UUID {
	return uuid.New()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Asus", "Acer", "Lenovo")
}

func randomLaptopname(brand string) string {
	switch brand {
	case "Asus":
		return randomStringFromSet(
			"Republic Of Gamer",
			"Zenbook",
			"Vivobook")
	case "Acer":
		return randomStringFromSet(
			"Predator",
			"Swift",
			"Aspire")
	case "Lenovo":
		return randomStringFromSet(
			"Legion",
			"Yoga",
			"Ideapad")
	default:
		return randomStringFromSet(
			"Clone",
			"OEM",
			"Unknown",
			"Illegal",
		)
	}
}
