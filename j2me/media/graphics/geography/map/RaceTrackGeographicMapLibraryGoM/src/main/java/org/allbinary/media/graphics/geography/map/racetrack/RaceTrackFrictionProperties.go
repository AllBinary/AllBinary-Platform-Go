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

type RaceTrackFrictionProperties struct {
    ROAD_FRICTION_NOMINATOR int
    GROUND_FRICTION_NOMINATOR int
}


func (this *RaceTrackFrictionProperties) initFields() {
    
    
}


func NewRaceTrackFrictionProperties(roadFriction int, groundFriction int) *RaceTrackFrictionProperties {
    this := &RaceTrackFrictionProperties{}
    this.initFields()
        this.ROAD_FRICTION_NOMINATOR = FrictionData_GetFrictionDenominator() - roadFriction
        this.GROUND_FRICTION_NOMINATOR = FrictionData_GetFrictionDenominator() - groundFriction
    return this
}


func (this *RaceTrackFrictionProperties) GetRoadFrictionNominator() int {
        return this.ROAD_FRICTION_NOMINATOR
}


func (this *RaceTrackFrictionProperties) GetGroundFrictionNominator() int {
        return this.GROUND_FRICTION_NOMINATOR
}

