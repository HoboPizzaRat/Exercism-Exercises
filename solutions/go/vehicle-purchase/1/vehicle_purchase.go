package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	} else {
		return option2 + " is clearly the better choice."
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
// For a rough estimate, assume if the vehicle is less than 3 years old, it costs 80%
// of the original price it had when it was brand new. If it is at least 10 years old,
// it costs 50%. If the vehicle is at least 3 years old but less than 10 years, it costs
// 70% of the original price.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resellPrice float64 = 0
	if age < 3 {
		resellPrice = originalPrice * 0.8
	} else if age < 10 {
		resellPrice = originalPrice * 0.7
	} else {
		resellPrice = originalPrice * 0.5
	}
	return resellPrice
}
