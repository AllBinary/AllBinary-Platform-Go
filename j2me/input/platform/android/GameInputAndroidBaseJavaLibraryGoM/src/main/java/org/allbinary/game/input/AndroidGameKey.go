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

type AndroidGameKey struct {
    Input
}



func NewAndroidGameKey(keyCode int, name string) *AndroidGameKey {
    this := &AndroidGameKey{}
        this.Input = *NewInput(keyCode, name)
        var inputFactory *InputFactory = InputFactory_GetInstance(); _ = inputFactory
        inputFactory.InputIntegerArray[this.GetId()] = this
    return this
}


func (this *AndroidGameKey) ToString() string {
        return "AndroidGameKey: " + this.Input.ToString()
}

func (this *AndroidGameKey) String() string {
    return this.ToString()
}


func AndroidGameKey_GetString(keyCode int) string {
        var input IInput = InputFactory_GetInstance().GetInstanceById(keyCode); _ = input
        if input != nil {
            return input.GetName()
        } else {
            return StringUtil_GetInstance().EMPTY_STRING
        }
}

