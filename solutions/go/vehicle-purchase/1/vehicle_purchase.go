package purchase


func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}


func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	}
	return option2 + " is clearly the better choice."
}


func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.80
	} else if age >=3 && age < 10 {
		return originalPrice * 0.70
	} else {
		return originalPrice * 0.50
	}
}
