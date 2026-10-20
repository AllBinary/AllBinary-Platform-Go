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

type AndroidNumberKeyToJ2MECanvasNumberKey struct {
}

var AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap []int = make([]int, 17)


func AndroidNumberKeyToJ2MECanvasNumberKey_Init() {
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_0] = Canvas_KEY_NUM0
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_1] = Canvas_KEY_NUM1
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_2] = Canvas_KEY_NUM2
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_3] = Canvas_KEY_NUM3
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_4] = Canvas_KEY_NUM4
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_5] = Canvas_KEY_NUM5
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_6] = Canvas_KEY_NUM6
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_7] = Canvas_KEY_NUM7
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_8] = Canvas_KEY_NUM8
        AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[KeyEvent_KEYCODE_9] = Canvas_KEY_NUM9
}


func NewAndroidNumberKeyToJ2MECanvasNumberKey() *AndroidNumberKeyToJ2MECanvasNumberKey {
    this := &AndroidNumberKeyToJ2MECanvasNumberKey{}
    return this
}


func AndroidNumberKeyToJ2MECanvasNumberKey_GetKey(key int) int {
        if key < len(AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap) {
            var value int = AndroidNumberKeyToJ2MECanvasNumberKey_AndroidtoJavaMicroEditionKeyMap[key]; _ = value
            if value != 0 {
                return value
            }
        }
        return key
}

