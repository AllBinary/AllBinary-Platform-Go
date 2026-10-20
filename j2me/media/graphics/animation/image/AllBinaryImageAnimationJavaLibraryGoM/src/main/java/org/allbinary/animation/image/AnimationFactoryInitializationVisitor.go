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

type AnimationFactoryInitializationVisitor struct {
    Width int
    Height int
    Dx int
    Dy int
    OriginalDx int
    OriginalDy int
}


func (this *AnimationFactoryInitializationVisitor) initFields() {
    
    
    
    
    
    
}

func NewAnimationFactoryInitializationVisitor() *AnimationFactoryInitializationVisitor {
    this := &AnimationFactoryInitializationVisitor{}
    this.initFields()
    return this
}


func (this *AnimationFactoryInitializationVisitor) Visit() {
}


func (this *AnimationFactoryInitializationVisitor) ToString() string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        var commonLabels *CommonLabels = CommonLabels_GetInstance(); _ = commonLabels
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(PositionStrings_GetInstance().DX_LABEL).Appendfloat(float32(this.Dx))
        stringBuffer.Append(PositionStrings_GetInstance().DY_LABEL).Appendfloat(float32(this.Dy))
        stringBuffer.Append(commonLabels.WIDTH_LABEL)
        stringBuffer.Appendint(this.Width)
        stringBuffer.Append(commonSeps.SPACE)
        stringBuffer.Append(commonLabels.HEIGHT_LABEL)
        stringBuffer.Appendint(this.Height)
        return fmt.Sprint(stringBuffer)
}

func (this *AnimationFactoryInitializationVisitor) String() string {
    return this.ToString()
}

