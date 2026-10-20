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

type AngleIncrementInfo struct {
    AngleIncrement int16
    DOWN_FRAME int
    UP_FRAME int
    LEFT_FRAME int
    RIGHT_FRAME int
    FrameUtil *FrameUtil
}

var AngleIncrementInfo_INCREMENT string = "Inc: "

func (this *AngleIncrementInfo) initFields() {
    
    
    
    
    
    this.FrameUtil = FrameUtil_GetInstance()
}


func NewAngleIncrementInfo(angleIncrement int16) *AngleIncrementInfo {
    this := &AngleIncrementInfo{}
    this.initFields()
        this.AngleIncrement = angleIncrement
        var angleFactory *AngleFactory = AngleFactory_GetInstance(); _ = angleFactory
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        this.DOWN_FRAME = smallIntegerSingletonFactory.GetAt(int(angleFactory.DOWN.GetValue()) / int(this.AngleIncrement))
        this.UP_FRAME = smallIntegerSingletonFactory.GetAt(int(angleFactory.UP.GetValue()))
        this.LEFT_FRAME = smallIntegerSingletonFactory.GetAt(int(angleFactory.LEFT.GetValue()) / int(this.AngleIncrement))
        this.RIGHT_FRAME = smallIntegerSingletonFactory.GetAt(int(angleFactory.RIGHT.GetValue()) / int(this.AngleIncrement))
    return this
}


func (this *AngleIncrementInfo) GetFrameAngle(frame int) int {
        return this.FrameUtil.GetFrameAngle(frame, int(this.AngleIncrement))
}


func (this *AngleIncrementInfo) GetAngleIncrement() int16 {
        return this.AngleIncrement
}


func (this *AngleIncrementInfo) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(AngleIncrementInfo_INCREMENT).Appendshort(this.AngleIncrement))
}

func (this *AngleIncrementInfo) String() string {
    return this.ToString()
}


func (this *AngleIncrementInfo) GetClosestGeneralDirection(angle int16) int {
        var closestDirection int = SmallIntegerSingletonFactory_GetInstance().GetAt(360); _ = closestDirection
        if int(this.UP_FRAME) - int(angle) < int(closestDirection) - int(angle) {
            closestDirection = this.UP_FRAME
        }
        if int(this.DOWN_FRAME) - int(angle) < int(closestDirection) - int(angle) {
            closestDirection = this.DOWN_FRAME
        }
        if int(this.LEFT_FRAME) - int(angle) < int(closestDirection) - int(angle) {
            closestDirection = this.LEFT_FRAME
        }
        if int(this.RIGHT_FRAME) - int(angle) < int(closestDirection) - int(angle) {
            closestDirection = this.RIGHT_FRAME
        }
        return int(closestDirection)
}
