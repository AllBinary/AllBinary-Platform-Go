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

type ImageJ2SEUtil struct {
    TRANSPARENT_COLOR *Color
}

var ImageJ2SEUtil_Instance *ImageJ2SEUtil = NewImageJ2SEUtil()

func (this *ImageJ2SEUtil) initFields() {
    this.TRANSPARENT_COLOR = NewColor(0, 0, 0, 0)
}

func NewImageJ2SEUtil() *ImageJ2SEUtil {
    this := &ImageJ2SEUtil{}
    this.initFields()
    return this
}


func ImageJ2SEUtil_GetInstance() *ImageJ2SEUtil {
        return ImageJ2SEUtil_Instance
}

