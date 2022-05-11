package ebike_test

import (
	"fmt"
	"testing"

	"github.com/ahrberg/ebike"
)

func TestEbike(t *testing.T) {

	rider := ebike.Rider{
		MaxGear:   6,
		CrankHigh: 100,
		CrankLow:  10,
		MaxVel:    velConv(25),
		MaxTorque: 20,
	}

	frontWhl := ebike.Wheel{
		Inert:   0.147,
		Radius:  0.355, // 28'
		RollRes: 0.0055,
	}

	rearWhl := ebike.Wheel{
		Inert:   0.147,
		Radius:  0.355, // 28'
		RollRes: 0.0055,
	}

	drive := ebike.Drive{
		GearRatio: []float64{
			0.632, 0.741, 0.843, 0.989, 1.145, 1.335, 1.545,
		},
		WheelRatio: []float64{0.084, 0.084, 0.084, 0.084, 0.084, 0.084, 0.084},
		Efficiency: 0.9,
	}

	bike := ebike.Bike{
		FrontWhl: frontWhl,
		RearWhl:  rearWhl,
		Drive:    drive,
		Rider:    rider,
		TotMass:  100,
	}

	for i := 0; i < 60; i++ {
		o := bike.Run(0)

		fmt.Printf("%g\n", o.Vel*3.6)

	}
}

func velConv(vel float64) float64 {
	return vel * 3.6
}
