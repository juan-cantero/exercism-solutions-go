package thefarm

import (
    "fmt"
    "errors"
    )

type InvalidCowsError struct {
    cows int
    message string
}
func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%v cows are invalid: %s",e.cows, e.message)
}


func DivideFood(calc FodderCalculator, cows int) (float64,error) {
    totalFodder, err := calc.FodderAmount(cows)
    if err != nil {
        return 0.0,err
    }
    factor, err := calc.FatteningFactor()
    if err != nil {
        return 0.0,err
    }
    return totalFodder / float64(cows) * factor, nil
    
}

func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64,error) {
    if cows <= 0 {
        return 0.0, errors.New("invalid number of cows")
    }
    return DivideFood(calc, cows)
    
}

func ValidateNumberOfCows(cows int) error {
    if cows < 0 {
        return &InvalidCowsError{cows:cows, message:"there are no negative cows"}
    } else if cows == 0 {
        return &InvalidCowsError{cows:cows, message:"no cows don't need food"}
    }
    return nil
    
}

