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

type ScaleProcessor struct {
    ScaleBaseProcessor
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ImageScaleUtil *ImageScaleUtil
}

var ScaleProcessor_InstanceC *ScaleProcessor = NewScaleProcessor()

func (this *ScaleProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ImageScaleUtil = ImageScaleUtil_GetInstance()
}

func NewScaleProcessor() *ScaleProcessor {
    this := &ScaleProcessor{}
    this.ScaleBaseProcessor = *NewScaleBaseProcessor()
    this.initFields()
    return this
}


func ScaleProcessor_GetInstance() IScaleBaseProcessor {
        return ScaleProcessor_InstanceC
}


func (this *ScaleProcessor) Update(originalImage IImage, originalImageArray []IImage, ximageToShowArray []IImage, unused int, scaleX float32, scaleY float32, maxScaleX float32, maxScaleY float32) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.UPDATE, e)
                }
            }()
            this.ImageScaleUtil.Scale2(originalImage, originalImageArray, ximageToShowArray, unused, scaleX, scaleY, maxScaleX, maxScaleY)
            return
        }(); tryControl1 == 1 {
            return
        }
}

