package model

type TaxiService struct {
	CustomerId string `json:"customer_id"`
	VehicleId  string `json:"vehicle_id"`
	DriverId   string `json:"driver_id"`
}
