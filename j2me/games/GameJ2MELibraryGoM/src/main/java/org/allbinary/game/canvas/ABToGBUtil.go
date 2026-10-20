/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type ABToGBUtil struct {
    AbCanvas ICanvas
    AllBinaryGameLayerManager IAllBinaryGameLayerManager
}

var ABToGBUtil_Instance *ABToGBUtil = NewABToGBUtil()

func (this *ABToGBUtil) initFields() {
    this.AbCanvas = NullCanvas_NULL_CANVAS
    this.AllBinaryGameLayerManager = AllBinaryGameLayerManager_GetNullInstance()
}

func NewABToGBUtil() *ABToGBUtil {
    this := &ABToGBUtil{}
    this.initFields()
    return this
}


func ABToGBUtil_GetInstance() *ABToGBUtil {
        return ABToGBUtil_Instance
}

