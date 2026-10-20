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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type LurchAI struct {
    BasicAI
    CurrentRelativeAngle int
    CurrentSpeed int
}

// The Java class LurchAI has subclasses so its variables are this interface
type ILurchAI interface {
    IBasicAI
    LurchAI_() *LurchAI
    ProcessAI(allBinaryLayerManager IAllBinaryLayerManager)
    Reverse()
    Accelerate()
    Drop()
}

func (this *LurchAI) LurchAI_() *LurchAI {
    return this
}


func (this *LurchAI) initFields() {
    this.CurrentRelativeAngle = 0
    
}


func NewLurchAI(ownerLayerInterface IAllBinaryLayer, gameInput IGameInput) *LurchAI {
    this := &LurchAI{}
        this.BasicAI = *NewBasicAI(ownerLayerInterface, gameInput)
        this.initFields()
    return this
}


// throws Exception
func (this *LurchAI) ProcessAI(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.CurrentSpeed == -1 {
            this.CurrentSpeed = 5
        } else {
            this.CurrentSpeed = -1
        }
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        var x int = ownerLayerInterface.GetXP(); _ = x
        if ownerLayerInterface.GetXP() - this.CurrentSpeed <= 0 {
            this.Reverse()
            this.Drop()
        }
        if ownerLayerInterface.GetX2() + this.CurrentSpeed > DisplayInfoSingleton_GetInstance().GetLastWidth() {
            this.Reverse()
            this.Accelerate()
            this.Drop()
        }
        if this.CurrentRelativeAngle == 0 {
            x += this.CurrentSpeed
        } else if this.CurrentRelativeAngle == 180 {
            x -= this.CurrentSpeed
        }
        ownerLayerInterface.SetPosition(x, ownerLayerInterface.GetYP(), ownerLayerInterface.GetZP())
        if this.CurrentRelativeAngle == 0 {
            this.BasicAI.ProcessKeyAI(Canvas_KEY_NUM0)
        } else if this.CurrentRelativeAngle == 180 {
            this.BasicAI.ProcessKeyAI(Canvas_KEY_POUND)
        }
}


func (this *LurchAI) Reverse() {
        if this.CurrentRelativeAngle == 180 {
            this.CurrentRelativeAngle = 0
        } else if this.CurrentRelativeAngle == 0 {
            this.CurrentRelativeAngle = 180
        }
}


func (this *LurchAI) Accelerate() {
        if this.CurrentSpeed < 20 {
            this.CurrentSpeed++
        }
}


func (this *LurchAI) Drop() {
        var ownerLayerInterface IAllBinaryLayer = this.GetOwnerLayerInterface(); _ = ownerLayerInterface
        var y int = ownerLayerInterface.GetYP(); _ = y
        if ownerLayerInterface.GetY2() + ownerLayerInterface.GetHeight() > DisplayInfoSingleton_GetInstance().GetLastHeight() {
            y = 0
        } else {
            y += ownerLayerInterface.GetHeight() + 1
        }
        ownerLayerInterface.SetPosition(ownerLayerInterface.GetXP(), y, ownerLayerInterface.GetZP())
}

