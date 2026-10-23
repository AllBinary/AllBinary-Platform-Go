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

type AngleInfo struct {
    AngleIncrementInfo *AngleIncrementInfo
    Angle int16
    FrameUtil *FrameUtil
}

var AngleInfo_ANGLE string = "Angle: "

func (this *AngleInfo) initFields() {
    
    this.Angle = 0
    this.FrameUtil = FrameUtil_GetInstance()
}


func AngleInfo_GetInstance(angleIncrement int16) *AngleInfo {
        return NewAngleInfo(AngleIncrementInfoFactory_GetInstance().GetAt(angleIncrement))
}


func NewAngleInfo(angleIncrementInfo *AngleIncrementInfo) *AngleInfo {
    this := &AngleInfo{}
    this.initFields()
        this.AngleIncrementInfo = angleIncrementInfo
    return this
}


func (this *AngleInfo) AdjustAngle(frame int) {
        var newAngle int = int(this.AngleIncrementInfo.GetAngleIncrement()) * frame - 90; _ = newAngle
        this.SetAngle(int16(this.FrameUtil.AdjustAngleToFrameAngle(newAngle)))
}


func (this *AngleInfo) GetAngle() int16 {
        return this.Angle
}


func (this *AngleInfo) SetAngle(angle int16) {
        this.Angle = angle
}


func (this *AngleInfo) GetAngleIncrementInfo() *AngleIncrementInfo {
        return this.AngleIncrementInfo
}


func (this *AngleInfo) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(AngleInfo_ANGLE)
        stringBuffer.Appendshort(this.Angle)
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(StringUtil_GetInstance().ToString(this.GetAngleIncrementInfo()))
        return fmt.Sprint(stringBuffer)
}

func (this *AngleInfo) String() string {
    return this.ToString()
}

