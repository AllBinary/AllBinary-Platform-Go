/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type NullKeyInputViewComposite struct {
}

// Implements KeyInputViewCompositeInterface

var NullKeyInputViewComposite_NULL_KEY_INPUT_VIEW_COMPOSITE *NullKeyInputViewComposite = NewNullKeyInputViewComposite()

func NewNullKeyInputViewComposite() *NullKeyInputViewComposite {
    this := &NullKeyInputViewComposite{}
    return this
}


func (this *NullKeyInputViewComposite) SetKeyInputView(keyInputView *View) {
}


func (this *NullKeyInputViewComposite) GetRootView() *View {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullKeyInputViewComposite) GetView() *View {
        panic(errors.New(fmt.Sprint()))
}


func (this *NullKeyInputViewComposite) RunOnUiThread(action Runnable) {
}

