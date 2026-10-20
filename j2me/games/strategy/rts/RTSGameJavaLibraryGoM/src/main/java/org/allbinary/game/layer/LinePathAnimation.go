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

type LinePathAnimation struct {
}

// The Java class LinePathAnimation has subclasses so its variables are this interface
type ILinePathAnimation interface {
    LinePathAnimation_() *LinePathAnimation
    Paint(graphics IGraphics, point IGPoint, nextPoint IGPoint, tiledLayer ILayer)
}

func (this *LinePathAnimation) LinePathAnimation_() *LinePathAnimation {
    return this
}

var LinePathAnimation_Instance ILinePathAnimation = NewLinePathAnimation()

func NewLinePathAnimation() *LinePathAnimation {
    this := &LinePathAnimation{}
    return this
}


func LinePathAnimation_GetInstance() ILinePathAnimation {
        return LinePathAnimation_Instance
}


func (this *LinePathAnimation) Paint(graphics IGraphics, point IGPoint, nextPoint IGPoint, tiledLayer ILayer) {
        graphics.DrawLine(nextPoint.GetX(), nextPoint.GetY(), point.GetX(), point.GetY())
}

