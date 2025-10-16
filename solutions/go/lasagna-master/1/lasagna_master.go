package lasagna


func PreparationTime(layers []string, minutes int ) int {
    if minutes == 0 {
        minutes = 2
    }
    return len(layers) * minutes
}


func Quantities(layers []string ) (int, float64) {
    noodles := 0
    sauce := 0.0
    for _,v := range layers {
        switch v {
            case "noodles":
            	noodles += 50
            case "sauce":
            	sauce += 0.2
        }
    }
    return noodles, sauce
	
    
}

func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList) -1] = friendsList[len(friendsList) -1]
}



func ScaleRecipe(quantities []float64, portions int) []float64 {
    scale := float64(portions) / 2.0
    scaledQuantities := make([]float64, len(quantities))
    for i,v := range quantities {
        scaledQuantities[i] = v * scale 
    }
    return scaledQuantities
}


