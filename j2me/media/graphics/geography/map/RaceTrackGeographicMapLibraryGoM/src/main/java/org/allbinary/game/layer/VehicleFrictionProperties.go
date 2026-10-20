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

type VehicleFrictionProperties struct {
    FrictionProperties
    TIRE_FRICTION_NOMINATOR int
    BRAKE_FRICTION_NOMINATOR int
    AIR_FRICTION_NOMINATOR int
    CRASH_FRICTION_NOMINATOR int
}


func (this *VehicleFrictionProperties) initFields() {
    
    
    
    
}


func NewVehicleFrictionProperties(tireFriction int, brakeFriction int, airFriction int, crashFriction int) *VehicleFrictionProperties {
    this := &VehicleFrictionProperties{}
    this.FrictionProperties = *NewFrictionProperties()
    this.initFields()
        this.TIRE_FRICTION_NOMINATOR = FrictionData_GetFrictionDenominator() - tireFriction
        this.BRAKE_FRICTION_NOMINATOR = FrictionData_GetFrictionDenominator() - brakeFriction
        this.AIR_FRICTION_NOMINATOR = FrictionData_GetFrictionDenominator() - airFriction
        this.CRASH_FRICTION_NOMINATOR = FrictionData_GetFrictionDenominator() - crashFriction
    return this
}


func (this *VehicleFrictionProperties) GetTireFrictionNominator() int {
        return this.TIRE_FRICTION_NOMINATOR
}


func (this *VehicleFrictionProperties) GetBrakeFrictionNominator() int {
        return this.BRAKE_FRICTION_NOMINATOR
}


func (this *VehicleFrictionProperties) GetAirFrictionNominator() int {
        return this.AIR_FRICTION_NOMINATOR
}


func (this *VehicleFrictionProperties) GetCrashFrictionNominator() int {
        return this.CRASH_FRICTION_NOMINATOR
}

