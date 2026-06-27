package models

type PricingStrategy string

const (
	Hourly PricingStrategy = "hourly"
	Daily  PricingStrategy = "daily"
)
