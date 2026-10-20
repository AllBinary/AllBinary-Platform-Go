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

type AnalogInputBuilder struct {
}

var AnalogInputBuilder_Instance *AnalogInputBuilder = NewAnalogInputBuilder()

func NewAnalogInputBuilder() *AnalogInputBuilder {
    this := &AnalogInputBuilder{}
    return this
}


func AnalogInputBuilder_GetInstance() *AnalogInputBuilder {
        return AnalogInputBuilder_Instance
}


func (this *AnalogInputBuilder) Disable(inputProcessorArray []IGameInputProcessor, collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer) {
        inputProcessorArray[Canvas_UP] = NewSpecialUpGameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_DOWN] = NewSpecialDownGameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_RIGHT] = NewSpecialRightGameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_LEFT] = NewSpecialLeftGameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_KEY_NUM0] = NewSpecial1GameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_KEY_NUM5] = NewSpecial3GameInputProcessor(collidableDestroyableDamageableLayer)
}


func (this *AnalogInputBuilder) Enable(inputProcessorArray []IGameInputProcessor, collidableDestroyableDamageableLayer ICollidableDestroyableDamageableLayer) {
        inputProcessorArray[Canvas_UP] = NewSpecialAnalogUpGameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_DOWN] = NewSpecialAnalogDownGameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_RIGHT] = NewSpecialAnalogRightGameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_LEFT] = NewSpecialAnalogLeftGameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_KEY_NUM0] = NewSpecialAnalogLeftTriggerGameInputProcessor(collidableDestroyableDamageableLayer)
        inputProcessorArray[Canvas_KEY_NUM5] = NewSpecialAnalogRightTriggerGameInputProcessor(collidableDestroyableDamageableLayer)
}

