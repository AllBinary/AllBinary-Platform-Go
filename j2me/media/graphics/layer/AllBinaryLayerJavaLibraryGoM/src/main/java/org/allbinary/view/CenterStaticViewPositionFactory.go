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

type CenterStaticViewPositionFactory struct {
}

// The Java class CenterStaticViewPositionFactory has subclasses so its variables are this interface
type ICenterStaticViewPositionFactory interface {
    CenterStaticViewPositionFactory_() *CenterStaticViewPositionFactory
    GetInstance(z int) IViewPositionBase
}

func (this *CenterStaticViewPositionFactory) CenterStaticViewPositionFactory_() *CenterStaticViewPositionFactory {
    return this
}



func NewCenterStaticViewPositionFactory() *CenterStaticViewPositionFactory {
    this := &CenterStaticViewPositionFactory{}
    return this
}


func (this *CenterStaticViewPositionFactory) GetInstance(z int) IViewPositionBase {
        var displayInfoSingleton *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfoSingleton
        return NewStaticViewPosition((displayInfoSingleton.GetLastWidth() >> 1), (displayInfoSingleton.GetLastHeight() >> 1), z)
}

