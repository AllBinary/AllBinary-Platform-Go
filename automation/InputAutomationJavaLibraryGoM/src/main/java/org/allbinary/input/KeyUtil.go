/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
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

type KeyUtil struct {
}



func NewKeyUtil() *KeyUtil {
    this := &KeyUtil{}
    return this
}


func KeyUtil_IsNormallyDisplayed(keyCode int) bool {
        switch keyCode {
            case KeyEvent_VK_SPACE, KeyEvent_VK_COMMA, KeyEvent_VK_MINUS, KeyEvent_VK_PERIOD, KeyEvent_VK_SLASH, KeyEvent_VK_0, KeyEvent_VK_1, KeyEvent_VK_2, KeyEvent_VK_3, KeyEvent_VK_4, KeyEvent_VK_5, KeyEvent_VK_6, KeyEvent_VK_7, KeyEvent_VK_8, KeyEvent_VK_9, KeyEvent_VK_SEMICOLON, KeyEvent_VK_EQUALS, KeyEvent_VK_A, KeyEvent_VK_B, KeyEvent_VK_C, KeyEvent_VK_D, KeyEvent_VK_E, KeyEvent_VK_F, KeyEvent_VK_G, KeyEvent_VK_H, KeyEvent_VK_I, KeyEvent_VK_J, KeyEvent_VK_K, KeyEvent_VK_L, KeyEvent_VK_M, KeyEvent_VK_N, KeyEvent_VK_O, KeyEvent_VK_P, KeyEvent_VK_Q, KeyEvent_VK_R, KeyEvent_VK_S, KeyEvent_VK_T, KeyEvent_VK_U, KeyEvent_VK_V, KeyEvent_VK_W, KeyEvent_VK_X, KeyEvent_VK_Y, KeyEvent_VK_Z, KeyEvent_VK_OPEN_BRACKET, KeyEvent_VK_BACK_SLASH, KeyEvent_VK_CLOSE_BRACKET, KeyEvent_VK_NUMPAD0, KeyEvent_VK_NUMPAD1, KeyEvent_VK_NUMPAD2, KeyEvent_VK_NUMPAD3, KeyEvent_VK_NUMPAD4, KeyEvent_VK_NUMPAD5, KeyEvent_VK_NUMPAD6, KeyEvent_VK_NUMPAD7, KeyEvent_VK_NUMPAD8, KeyEvent_VK_NUMPAD9, KeyEvent_VK_MULTIPLY, KeyEvent_VK_ADD, KeyEvent_VK_SEPARATER, KeyEvent_VK_SUBTRACT, KeyEvent_VK_DECIMAL, KeyEvent_VK_DIVIDE:
                return true
            default:
                return false
        }
}

