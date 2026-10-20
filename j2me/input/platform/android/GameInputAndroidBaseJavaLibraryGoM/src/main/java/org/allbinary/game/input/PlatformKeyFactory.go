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

type PlatformKeyFactory struct {
    AndroidKeyFactory IAndroidKeyFactory
}

var PlatformKeyFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *PlatformKeyFactory) initFields() {
    this.AndroidKeyFactory = ActivityFractureControllerUtilFactory_GetInstance()
}

func NewPlatformKeyFactory() *PlatformKeyFactory {
    this := &PlatformKeyFactory{}
    this.initFields()
    return this
}


func PlatformKeyFactory_GetInstance() *PlatformKeyFactory {
        if PlatformKeyFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            PlatformKeyFactory_Instance = NewPlatformKeyFactory()
        }
        return any(PlatformKeyFactory_Instance).(*PlatformKeyFactory)
}


func (this *PlatformKeyFactory) GetString(keyCode int) string {
        return AndroidGameKey_GetString(keyCode)
}


func (this *PlatformKeyFactory) IsSubmission(input IInput) bool {
        return this.AndroidKeyFactory.IsSubmission(input)
}


func (this *PlatformKeyFactory) IsDelete(input IInput) bool {
        return this.AndroidKeyFactory.IsDelete(input)
}


func (this *PlatformKeyFactory) IsBackSpace(input IInput) bool {
        return this.AndroidKeyFactory.IsBackSpace(input)
}


func (this *PlatformKeyFactory) IsLeft(input IInput) bool {
        return this.AndroidKeyFactory.IsLeft(input)
}


func (this *PlatformKeyFactory) IsRight(input IInput) bool {
        return this.AndroidKeyFactory.IsRight(input)
}


func (this *PlatformKeyFactory) IsUp(input IInput) bool {
        return this.AndroidKeyFactory.IsUp(input)
}


func (this *PlatformKeyFactory) IsDown(input IInput) bool {
        return this.AndroidKeyFactory.IsDown(input)
}


func (this *PlatformKeyFactory) IsEnter(input IInput) bool {
        return this.AndroidKeyFactory.IsEnter(input)
}

