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

type AndroidAPI12KeyFactory struct {
    AndroidAPI9KeyFactory
}

var AndroidAPI12KeyFactory_Instance12 any = NullUtil_GetInstance().NULL_OBJECT


func AndroidAPI12KeyFactory_GetAPI12Instance() *AndroidAPI12KeyFactory {
        if AndroidAPI12KeyFactory_Instance12 == NullUtil_GetInstance().NULL_OBJECT {
            AndroidAPI12KeyFactory_Instance12 = NewAndroidAPI12KeyFactory()
        }
        return any(AndroidAPI12KeyFactory_Instance12).(*AndroidAPI12KeyFactory)
}


func NewAndroidAPI12KeyFactory() *AndroidAPI12KeyFactory {
    this := &AndroidAPI12KeyFactory{}
    this.AndroidAPI9KeyFactory = *NewAndroidAPI9KeyFactory()
        this.KEYCODE_BUTTON_1 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_1, "BUTTON 1")
        this.KEYCODE_BUTTON_10 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_10, "BUTTON 10")
        this.KEYCODE_BUTTON_11 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_11, "BUTTON 11")
        this.KEYCODE_BUTTON_12 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_12, "BUTTON 12")
        this.KEYCODE_BUTTON_13 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_13, "BUTTON 13")
        this.KEYCODE_BUTTON_14 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_14, "BUTTON 14")
        this.KEYCODE_BUTTON_15 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_15, "BUTTON 15")
        this.KEYCODE_BUTTON_16 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_16, "BUTTON 16")
        this.KEYCODE_BUTTON_2 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_2, "BUTTON 2")
        this.KEYCODE_BUTTON_3 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_3, "BUTTON 3")
        this.KEYCODE_BUTTON_4 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_4, "BUTTON 4")
        this.KEYCODE_BUTTON_5 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_5, "BUTTON 5")
        this.KEYCODE_BUTTON_6 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_6, "BUTTON 6")
        this.KEYCODE_BUTTON_7 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_7, "BUTTON 7")
        this.KEYCODE_BUTTON_8 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_8, "BUTTON 8")
        this.KEYCODE_BUTTON_9 = NewAndroidGameKey(KeyEvent_KEYCODE_BUTTON_9, "BUTTON 9")
    return this
}

