package ebike

type Rider struct {
	MaxGear   int
	CrankHigh float64
	CrankLow  float64
	MaxVel    float64
	MaxTorque float64
}
type RiderInp struct {
	Gear  int
	Vel   float64 // m/s
	Slope float64 // %
	Crank float64 // rpm
}
type RiderOut struct {
	Torque float64 // Nm should also use Rider Power, T = P / n
	Brake  bool
	Gear   int
}

func (r *Rider) Run(inp RiderInp) RiderOut {

	return RiderOut{
		Gear:   r.selGear(inp.Gear, inp.Crank),
		Brake:  false,
		Torque: r.calcTorque(inp.Vel, inp.Slope, inp.Crank),
	}
}

func (r *Rider) selGear(gear int, crank float64) int {
	if gear == 0 {
		return 1
	}

	if crank > r.CrankHigh && gear < r.MaxGear {
		return gear + 1
	}

	if crank < r.CrankLow && gear > 1 {
		return gear - 1
	}

	return gear

}

func (r *Rider) calcTorque(vel, slope, crank float64) float64 {

	crankOvrSpd := r.CrankHigh + r.CrankHigh*0.1
	slopeFreeRroll := -1.0

	if crank > crankOvrSpd {
		return 0
	}

	if vel > r.MaxVel {
		return 0
	}

	if slope < slopeFreeRroll {
		return 0
	}

	return r.MaxTorque
}
