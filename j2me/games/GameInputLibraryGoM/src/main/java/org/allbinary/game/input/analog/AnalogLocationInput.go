// Generated Code Do Not Modify

package allbinary

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var _ = errors.New
var _ = fmt.Sprint
var _ = math.Abs
var _ = rand.Float64
var _ = os.Exit
var _ = reflect.TypeOf
var _ = runtime.GC
var _ = strconv.Itoa
var _ = strings.Contains
var _ = time.Now
var _ = unicode.IsDigit

type AnalogLocationInput struct {
    PlayerId int
    CustomGPoint *CustomGPoint
    Rz int
    RightTrigger int
    LeftTrigger int
}

var AnalogLocationInput_NULL_ANALOG_LOCATION_INPUT *AnalogLocationInput = NewAnalogLocationInput(-1)

func (this *AnalogLocationInput) initFields() {
    
    
    
    
    
}


func NewAnalogLocationInput(playerId int) *AnalogLocationInput {
    this := &AnalogLocationInput{}
    this.initFields()
        this.PlayerId = playerId
        this.CustomGPoint = NewCustomGPoint(0, 0, 0)
    return this
}


func (this *AnalogLocationInput) GetPlayerId() int {
        return this.PlayerId
}


func (this *AnalogLocationInput) GetCustomGPoint() *CustomGPoint {
        return this.CustomGPoint
}


func (this *AnalogLocationInput) GetRz() int {
        return this.Rz
}


func (this *AnalogLocationInput) SetRz(rz int) {
        this.Rz = rz
}


func (this *AnalogLocationInput) GetRightTrigger() int {
        return this.RightTrigger
}


func (this *AnalogLocationInput) SetRightTrigger(rightTrigger int) {
        this.RightTrigger = rightTrigger
}


func (this *AnalogLocationInput) GetLeftTrigger() int {
        return this.LeftTrigger
}


func (this *AnalogLocationInput) SetLeftTrigger(leftTrigger int) {
        this.LeftTrigger = leftTrigger
}

