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

type ScaleProperties struct {
    ScaleX float32
    ScaleY float32
    ScaleWidth int
    ScaleHeight int
    ShouldScale bool
}

var ScaleProperties_Instance *ScaleProperties = NewScaleProperties()

func (this *ScaleProperties) initFields() {
    this.ScaleX = 1.0
    this.ScaleY = 1.0
    
    
    
}

func NewScaleProperties() *ScaleProperties {
    this := &ScaleProperties{}
    this.initFields()
    return this
}


func (this *ScaleProperties) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().COLON_SEP).Appendboolean(this.ShouldScale).Append("scaleX: ").Appendfloat(this.ScaleX).Append("scaleY: ").Appendfloat(this.ScaleY).Append("scaleWidth: ").Appendint(this.ScaleWidth).Append("scaleHeight: ").Appendint(this.ScaleHeight))
}

func (this *ScaleProperties) String() string {
    return this.ToString()
}

