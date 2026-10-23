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

type AngleFactory struct {
    LogUtil *LogUtil
    TOTAL_ANGLE int16
    QUARTER_TOTAL_ANGLE int16
    AngleArray []IAngle
    NEGATIVE_ONE int16
    NOT_ANGLE *NamedAngle
    DOWN *NamedAngle
    UP *NamedAngle
    LEFT *NamedAngle
    RIGHT *NamedAngle
    FrameUtil *FrameUtil
}

var AngleFactory_Instance *AngleFactory = NewAngleFactory()

func (this *AngleFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.TOTAL_ANGLE = 360
    this.QUARTER_TOTAL_ANGLE = 90
    this.AngleArray = make([]IAngle, int(this.TOTAL_ANGLE))
    this.NEGATIVE_ONE = int16(-1)
    this.NOT_ANGLE = NewNamedAngle(this.NEGATIVE_ONE, CommonStrings_GetInstance().EMPTY)
    
    
    
    
    this.FrameUtil = FrameUtil_GetInstance()
}


func AngleFactory_GetInstance() *AngleFactory {
        return AngleFactory_Instance
}


func NewAngleFactory() *AngleFactory {
    this := &AngleFactory{}
    this.initFields()
        var commonPhoneStrings *CommonPhoneStrings = CommonPhoneStrings_GetInstance(); _ = commonPhoneStrings
        var total int16 = int16(len(this.AngleArray)); _ = total
        this.UP = NewNamedAngle(int16(0), commonPhoneStrings.UP)
        this.AngleArray[0] = this.UP
        for index := int(1); index < 90; index++ {
            this.AngleArray[index] = NewAngle(int16(index))
        }
        this.RIGHT = NewNamedAngle(int16(90), commonPhoneStrings.RIGHT)
        this.AngleArray[90] = this.RIGHT
        for index := int(91); index < 180; index++ {
            this.AngleArray[index] = NewAngle(int16(index))
        }
        this.DOWN = NewNamedAngle(int16(180), commonPhoneStrings.DOWN)
        this.AngleArray[180] = this.DOWN
        for index := int(181); index < 270; index++ {
            this.AngleArray[index] = NewAngle(int16(index))
        }
        this.LEFT = NewNamedAngle(int16(270), commonPhoneStrings.LEFT)
        this.AngleArray[270] = this.LEFT
        for index := int(271); index < int(total); index++ {
            this.AngleArray[index] = NewAngle(int16(index))
        }
    return this
}


func (this *AngleFactory) GetAt(index int) IAngle {
        var adjustedIndex int = this.FrameUtil.AdjustAngleToFrameAngle(index); _ = adjustedIndex
        return this.AngleArray[adjustedIndex]
}


func (this *AngleFactory) GetClosestDirection(angle int) IAngle {
        if (angle >= 315 && angle < 360) || (angle >= 0 && angle < 45) {
            return this.UP
        } else if angle >= 45 && angle < 135 {
            return this.RIGHT
        } else if angle >= 135 && angle < 225 {
            return this.DOWN
        } else if angle >= 225 && angle < 315 {
            return this.LEFT
        }
        panic(errors.New(fmt.Sprint()))
}


func (this *AngleFactory) GetGeneralDirection() IAngle {
        return this.NOT_ANGLE
}


func AngleFactory_Main(args []string) {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var angleFactory *AngleFactory = AngleFactory_GetInstance(); _ = angleFactory
        for index := int(0); index < 360; index++ {
            stringMaker.Appendint(index).Append(CommonSeps_GetInstance().FORWARD_SLASH).Appendshort(angleFactory.GetClosestDirection(index).GetValue()).Append(CommonSeps_GetInstance().NEW_LINE)
        }
        LogUtil_GetInstance().PutF(fmt.Sprint(stringMaker), "main", "main")
}

