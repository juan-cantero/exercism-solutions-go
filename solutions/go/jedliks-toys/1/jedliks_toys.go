package jedlik
import (
    "fmt"
    "math"
    )


func (c *Car) Drive() {
    if c.battery - c.batteryDrain < 0  {
        return
    }
    c.battery -= c.batteryDrain
    c.distance += c.speed
}

func (c *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %v meters", c.distance)
}

func (c *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %v%%",c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
    moves := int(math.Ceil(float64(trackDistance) / float64(c.speed)))
	batteryDrainForDistance := moves * c.batteryDrain
	return c.battery >= batteryDrainForDistance
}

