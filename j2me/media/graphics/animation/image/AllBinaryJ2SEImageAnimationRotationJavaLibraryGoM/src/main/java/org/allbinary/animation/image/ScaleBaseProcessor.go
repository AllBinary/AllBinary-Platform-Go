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

type ScaleBaseProcessor struct {
}

// The Java class ScaleBaseProcessor has subclasses so its variables are this interface
type IScaleBaseProcessor interface {
    ScaleBaseProcessor_() *ScaleBaseProcessor
    Update(originalImage IImage, originalImageArray []IImage, imageToShowArray []IImage, index int, scaleX float32, scaleY float32, maxScaleX float32, maxScaleY float32)
}

func (this *ScaleBaseProcessor) ScaleBaseProcessor_() *ScaleBaseProcessor {
    return this
}

var ScaleBaseProcessor_Instance IScaleBaseProcessor = NewScaleBaseProcessor()

func NewScaleBaseProcessor() *ScaleBaseProcessor {
    this := &ScaleBaseProcessor{}
    return this
}


func ScaleBaseProcessor_GetInstance() IScaleBaseProcessor {
        return ScaleBaseProcessor_Instance
}


func (this *ScaleBaseProcessor) Update(originalImage IImage, originalImageArray []IImage, imageToShowArray []IImage, index int, scaleX float32, scaleY float32, maxScaleX float32, maxScaleY float32) {
}

