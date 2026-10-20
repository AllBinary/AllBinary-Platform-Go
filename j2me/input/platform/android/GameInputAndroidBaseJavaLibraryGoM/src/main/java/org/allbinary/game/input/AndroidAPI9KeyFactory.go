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

type AndroidAPI9KeyFactory struct {
    AndroidKeyFactory
}

// The Java class AndroidAPI9KeyFactory has subclasses so its variables are this interface
type IAndroidAPI9KeyFactory interface {
    IAndroidKeyFactory
    AndroidAPI9KeyFactory_() *AndroidAPI9KeyFactory
}

func (this *AndroidAPI9KeyFactory) AndroidAPI9KeyFactory_() *AndroidAPI9KeyFactory {
    return this
}

var AndroidAPI9KeyFactory_Instance any = NullUtil_GetInstance().NULL_OBJECT


func AndroidAPI9KeyFactory_GetInstance() IAndroidAPI9KeyFactory {
        if AndroidAPI9KeyFactory_Instance == NullUtil_GetInstance().NULL_OBJECT {
            AndroidAPI9KeyFactory_Instance = NewAndroidAPI9KeyFactory()
        }
        return any(AndroidAPI9KeyFactory_Instance).(IAndroidAPI9KeyFactory)
}


func NewAndroidAPI9KeyFactory() *AndroidAPI9KeyFactory {
    this := &AndroidAPI9KeyFactory{}
    this.AndroidKeyFactory = *NewAndroidKeyFactory()
        this.KEYCODE_BUTTON_L1 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_L1, "BUTTON L1")
        this.KEYCODE_BUTTON_R1 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_R1, "BUTTON R1")
        this.KEYCODE_BUTTON_THUMBR = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_THUMBR, "BUTTON THUMBR")
        this.KEYCODE_BUTTON_THUMBL = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_THUMBL, "BUTTON THUMBL")
        this.KEYCODE_BUTTON_START = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_START, "BUTTON START")
        this.KEYCODE_BUTTON_MODE = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_MODE, "BUTTON MODE")
        this.KEYCODE_BUTTON_B = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_B, "BUTTON B")
        this.KEYCODE_BUTTON_A = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_A, "BUTTON A")
        this.KEYCODE_BUTTON_X = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_X, "BUTTON X")
        this.KEYCODE_BUTTON_Y = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_Y, "BUTTON Y")
        this.KEYCODE_BUTTON_L2 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_L2, "BUTTON L2")
        this.KEYCODE_BUTTON_R2 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_R2, "BUTTON R2")
        this.KEYCODE_BUTTON_SELECT = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_SELECT, "BUTTON SELECT")
        this.KEYCODE_BUTTON_C = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_C, "BUTTON C")
        this.KEYCODE_BUTTON_Z = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_Z, "BUTTON Z")
    return this
}

