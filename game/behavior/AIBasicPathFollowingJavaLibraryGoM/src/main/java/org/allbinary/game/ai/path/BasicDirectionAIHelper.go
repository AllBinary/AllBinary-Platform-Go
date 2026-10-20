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

type BasicDirectionAIHelper struct {
    Name string
    AngleInfo *AngleInfo
}


func (this *BasicDirectionAIHelper) initFields() {
    
    
}


func NewBasicDirectionAIHelper(name string, angleInfo *AngleInfo) *BasicDirectionAIHelper {
    this := &BasicDirectionAIHelper{}
    this.initFields()
        this.Name = name
        this.AngleInfo = angleInfo
    return this
}


func (this *BasicDirectionAIHelper) TurnAI(frame int) int {
        var keyDirection int = -1; _ = keyDirection
        var directionAngle int = this.AngleInfo.GetAngleIncrementInfo().GetFrameAngle(frame); _ = directionAngle
        var angle int16 = this.AngleInfo.GetAngle(); _ = angle
        var degrees int = max(directionAngle - int(angle), -(directionAngle - int(angle))); _ = degrees
        if degrees < 180 {
            if int(angle) > directionAngle {
                keyDirection = Canvas_LEFT
            } else if int(angle) < directionAngle {
                keyDirection = Canvas_RIGHT
            }
        } else {
            if int(angle) > directionAngle {
                keyDirection = Canvas_RIGHT
            } else if int(angle) < directionAngle {
                keyDirection = Canvas_LEFT
            }
        }
        return keyDirection
}


func (this *BasicDirectionAIHelper) GetAIKeyPressed(frame int) int {
        var angle int16 = int16(this.AngleInfo.GetAngleIncrementInfo().GetFrameAngle(frame.IntValue())); _ = angle
        if int(this.AngleInfo.GetAngle()) != int(angle) {
            return this.TurnAI(frame.IntValue())
        } else {
            return Canvas_UP
        }
}


func (this *BasicDirectionAIHelper) GetAIKeyPressedFromDirection(geographicMapDirectionData *Direction) int {
        var keyDirection int = -1; _ = keyDirection
        var angleIncrementInfo *AngleIncrementInfo = this.AngleInfo.GetAngleIncrementInfo(); _ = angleIncrementInfo
        var directionFactory *DirectionFactory = DirectionFactory_GetInstance(); _ = directionFactory
        if geographicMapDirectionData == directionFactory.DOWN {
            keyDirection = this.GetAIKeyPressed(angleIncrementInfo.DOWN_FRAME)
        } else if geographicMapDirectionData == directionFactory.UP {
            keyDirection = this.GetAIKeyPressed(angleIncrementInfo.UP_FRAME)
        } else if geographicMapDirectionData == directionFactory.LEFT {
            keyDirection = this.GetAIKeyPressed(angleIncrementInfo.LEFT_FRAME)
        } else if geographicMapDirectionData == directionFactory.RIGHT {
            keyDirection = this.GetAIKeyPressed(angleIncrementInfo.RIGHT_FRAME)
        }
        return keyDirection
}


func (this *BasicDirectionAIHelper) GetAngleInfoP() *AngleInfo {
        return this.AngleInfo
}


func (this *BasicDirectionAIHelper) GetName() string {
        return this.Name
}

