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

type AnalogLocationInputFactory struct {
    List IBasicArrayList
}

var AnalogLocationInputFactory_Instance *AnalogLocationInputFactory = NewAnalogLocationInputFactory()

func (this *AnalogLocationInputFactory) initFields() {
    this.List = NewBasicArrayListD()
}


func AnalogLocationInputFactory_GetInstance() *AnalogLocationInputFactory {
        return AnalogLocationInputFactory_Instance
}


func NewAnalogLocationInputFactory() *AnalogLocationInputFactory {
    this := &AnalogLocationInputFactory{}
    this.initFields()
    return this
}


func (this *AnalogLocationInputFactory) AddPlayer(playerInputId int) {
        var analogLocationInput *AnalogLocationInput = AnalogLocationInput_NULL_ANALOG_LOCATION_INPUT; _ = analogLocationInput
        if playerInputId < this.List.Size() {
            analogLocationInput = any(this.List.Get(playerInputId)).(*AnalogLocationInput)
        }
        if analogLocationInput == AnalogLocationInput_NULL_ANALOG_LOCATION_INPUT {
            analogLocationInput = NewAnalogLocationInput(playerInputId)
            this.List.Add(analogLocationInput)
        }
}


func (this *AnalogLocationInputFactory) GetInstanceForPlayer(playerInputId int) *AnalogLocationInput {
        return any(this.List.Get(playerInputId)).(*AnalogLocationInput)
}


func (this *AnalogLocationInputFactory) GetList() IBasicArrayList {
        return this.List
}

