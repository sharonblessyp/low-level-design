package strategy

import (
	"math"
	"parking_lot/models"
)

type FeeStartegy interface {
	CalculateFee(ticket models.Ticket) models.Receipt
}

// pick the factory to calculate fees
func ChooseFeeStrategy(feeType models.PricingStrategy) FeeStartegy {
	switch feeType {
	case models.Hourly:
		return &HourlyStartegy{RatePerHour: 10.00}
	case models.Daily:
		return &DailyStartegy{RatePerDay: 100.00}
	}
	return nil
}

type HourlyStartegy struct {
	RatePerHour float64
}

func (h *HourlyStartegy) CalculateFee(ticket models.Ticket) models.Receipt {
	duration := ticket.ExitTime.Sub(ticket.EntryTime)
	return models.Receipt{
		Charges:   h.RatePerHour * math.Ceil(duration.Hours()),
		EntryTime: ticket.EntryTime,
		ExitTime:  ticket.ExitTime,
		Duration:  duration,
	}
}

type DailyStartegy struct {
	RatePerDay float64
}

func (d *DailyStartegy) CalculateFee(ticket models.Ticket) models.Receipt {
	return models.Receipt{
		Charges:   d.RatePerDay,
		EntryTime: ticket.EntryTime,
		ExitTime:  ticket.ExitTime,
		Duration:  ticket.ExitTime.Sub(ticket.EntryTime),
	}
}
