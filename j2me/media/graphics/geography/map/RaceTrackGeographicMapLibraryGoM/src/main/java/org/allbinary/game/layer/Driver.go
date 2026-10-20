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

type Driver struct {
    TurnSpeed int
    TrackPathSelectionSkill int
    Aggression *Aggression
}

var Driver_NULL_DRIVER *Driver = NewDriver(SmallIntegerSingletonFactory_GetInstance().GetAt(0), SmallIntegerSingletonFactory_GetInstance().GetAt(0), NewAggression(SmallIntegerSingletonFactory_GetInstance().GetAt(0), SmallIntegerSingletonFactory_GetInstance().GetAt(0), SmallIntegerSingletonFactory_GetInstance().GetAt(0)))

func (this *Driver) initFields() {
    this.TurnSpeed = SmallIntegerSingletonFactory_GetInstance().GetAt(0)
    this.TrackPathSelectionSkill = SmallIntegerSingletonFactory_GetInstance().GetAt(0)
    
}


func NewDriver(turnSpeed int, trackPathSelectionSkill int, aggression *Aggression) *Driver {
    this := &Driver{}
    this.initFields()
        this.SetTurnSpeed(turnSpeed)
        this.SetTrackPathSelectionSkill(trackPathSelectionSkill)
        this.Aggression = aggression
    return this
}


func (this *Driver) GetTurnSpeed() int {
        return this.TurnSpeed
}


func (this *Driver) SetTurnSpeed(turnSpeed int) {
        this.TurnSpeed = turnSpeed
}


func (this *Driver) GetTrackPathSelectionSkill() int {
        return this.TrackPathSelectionSkill
}


func (this *Driver) SetTrackPathSelectionSkill(trackPathSelectionSkill int) {
        this.TrackPathSelectionSkill = trackPathSelectionSkill
}


func (this *Driver) GetAggression() *Aggression {
        return this.Aggression
}


func (this *Driver) SetAggression(aggression *Aggression) {
        this.Aggression = aggression
}

