package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string] int {
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen":6,
        "dozen":12,
        "small_gross":120,
        "gross":144,
        "great_gross":1728,
        
    }
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int {}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value,ok := units[unit]
    if !ok {
        return false
    }
  	bill[item] += value
    return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemVal, exists := bill[item]
    if !exists {
        return false
    }
    unitVal, exists := units[unit]
    if !exists {
        return false
    }
	difference := itemVal - unitVal
    if difference < 0 {
        return false
    }else if difference == 0 {
        delete(bill,item)
    } else {
        bill[item] -= unitVal
    }
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val,ok := bill[item]
    if !ok {
        return 0, false
    }
    return val , true
}
