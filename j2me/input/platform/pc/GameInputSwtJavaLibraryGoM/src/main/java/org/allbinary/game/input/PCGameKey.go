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

type PCGameKey struct {
    Input
}



func NewPCGameKey(keyCode int, name string) *PCGameKey {
    this := &PCGameKey{}
        this.Input = *NewInput(keyCode, name)
        var id int = this.GetId(); _ = id
        if id >= 0 {
            InputFactory_GetInstance().Add(id, this)
        }
    return this
}


func (this *PCGameKey) ToString() string {
        return "PCGameKey: " + this.Input.ToString()
}

func (this *PCGameKey) String() string {
    return this.ToString()
}


func PCGameKey_GetString(keyCode int) string {
        var input IInput = InputFactory_GetInstance().GetInstanceById(keyCode); _ = input
        if input != nil {
            return input.GetName()
        } else {
            return StringUtil_GetInstance().EMPTY_STRING
        }
}

