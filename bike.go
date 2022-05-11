package ebike

import "math"

type Wheel struct {
	Inert   float64 // kg*m²
	RollRes float64 // N
	Radius  float64 // m
}

type Drive struct {
	GearRatio  []float64
	Efficiency float64
	WheelRatio []float64
}

type Bike struct {
	FrontWhl Wheel
	RearWhl  Wheel
	Drive    Drive
	Rider    Rider
	TotMass  float64
	gear     int
	vel      float64
	dist     float64
}

type BikeOut struct {
	Vel  float64
	Dist float64
}

func (b *Bike) Run(slope float64) BikeOut {

	if b.gear < 1 {
		b.gear = 1
	}

	rider := b.Rider.Run(RiderInp{
		Gear:  b.gear,
		Vel:   b.vel,
		Crank: b.velToCrank(),
		Slope: slope,
	})

	// Drive forces
	fd := b.crankToForce(rider.Torque)

	// Repelling forces
	fr := b.FrontWhl.rollRes(b.vel, b.TotMass) + b.RearWhl.rollRes(b.vel, b.TotMass) + parallelCompInSlope(b.TotMass*9.81, slope)

	// Resulting force
	f := fd - fr

	// Motion calculation
	a := b.TotMass / f // f = ma
	b.vel += a         // v = v0 + at
	b.dist += b.vel    // s = s0 + vt

	return BikeOut{
		Vel:  b.vel,
		Dist: b.dist,
	}
}

// crankToForce returns the resulting wheel force caused by crank torque
func (b Bike) crankToForce(torque float64) float64 {
	return torque / b.Drive.totRatio(b.gear) * b.Drive.Efficiency
}

// velToCrank returns the resulting crank speed cased by the bike velocity
func (b Bike) velToCrank() float64 {
	wheelSpeed := b.vel / b.FrontWhl.Radius * 60
	return wheelSpeed / b.Drive.totRatio(b.gear)
}

func (d Drive) totRatio(gear int) float64 {
	return d.GearRatio[gear-1] * d.WheelRatio[gear-1]
}

func slopeInRad(slopePerc float64) float64 {
	return math.Atan(slopePerc / 100)
}

// parallelCompInSlope returns the parallel force component in a slope
func parallelCompInSlope(force float64, slope float64) float64 {
	return math.Sin(slope) * force

}

func (w Wheel) rollRes(vel float64, totMass float64) float64 {
	return w.RollRes * totMass * 9.81 * vel
}
