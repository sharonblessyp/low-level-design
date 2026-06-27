package models

type VehicleType string

// enums for vehicle types
const (
	Car  VehicleType = "car"
	Bike VehicleType = "bike"
)

// check for valid vehicle types
func (vt VehicleType) IsValid() bool {
	switch vt {
	case Car, Bike:
		return true
	default:
		return false
	}
}

type Vehicle struct {
	ID   string
	Type VehicleType
}
