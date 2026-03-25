package main

import "fmt"

/*
1. ride matching strategy
2. concrete implemetations of the algorithm
3. strategy is decided during runtime
*/

type RideMatching interface {
	Match(riderLocation string)
}

type NearestDriver struct{}

func NewNearestDriver() *NearestDriver {
	return &NearestDriver{}
}

func (nd *NearestDriver) Match(riderLocation string) {
	fmt.Printf("NearestDriver: matching nearest driver with rider from %s \n", riderLocation)
}

type AirPortQueue struct{}

func NewAirPortQueue() *AirPortQueue {
	return &AirPortQueue{}
}

func (aq *AirPortQueue) Match(riderLocation string) {
	fmt.Printf("AirPortQueue: matching nearest driver with rider from terminal:%s \n", riderLocation)
}

type SurgePriority struct{}

func NewSurgePriority() *SurgePriority {
	return &SurgePriority{}
}

func (sp *SurgePriority) Match(riderLocation string) {
	fmt.Printf("SurgePriority: matching nearest driver with rider from %s \n", riderLocation)
}

type RideMatchingService struct {
	strategy RideMatching
}

func NewRideMatchingService(strategy RideMatching) *RideMatchingService {
	return &RideMatchingService{
		strategy: strategy,
	}
}

func (rs *RideMatchingService) SetStrategy(strategy RideMatching) {
	rs.strategy = strategy
}

func (rs *RideMatchingService) MatchDriver(location string) {
	if rs.strategy == nil {
		fmt.Printf("strategy not set \n")
		return
	}
	rs.strategy.Match(location)
}

func main() {
	nearestDriverStrategy := NewRideMatchingService(NewNearestDriver())
	nearestDriverStrategy.MatchDriver("Anna Nagar")

	matchingStrategy := NewRideMatchingService(NewAirPortQueue())
	matchingStrategy.MatchDriver("Terminal 1")

	matchingStrategy.SetStrategy(NewSurgePriority())
	matchingStrategy.MatchDriver("Anna Nagar")
}
