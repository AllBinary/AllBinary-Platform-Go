/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type CharacterFrictionProperties struct {
    FrictionProperties
    WATER_FRICTION_NOMINATOR int
    AIR_FRICTION_NOMINATOR int
    COLLISION_FRICTION_NOMINATOR int
}


func (this *CharacterFrictionProperties) initFields() {
    
    
    
}


func NewCharacterFrictionProperties(airFriction int, waterFriction int, collisionFriction int) *CharacterFrictionProperties {
    this := &CharacterFrictionProperties{}
    this.FrictionProperties = *NewFrictionProperties()
    this.initFields()
        this.AIR_FRICTION_NOMINATOR = FrictionData_GetFrictionDenominator() - airFriction
        this.COLLISION_FRICTION_NOMINATOR = FrictionData_GetFrictionDenominator() - collisionFriction
        this.WATER_FRICTION_NOMINATOR = FrictionData_GetFrictionDenominator() - waterFriction
    return this
}


func (this *CharacterFrictionProperties) GetWATER_FRICTION_NOMINATOR() int {
        return this.WATER_FRICTION_NOMINATOR
}


func (this *CharacterFrictionProperties) GetAIR_FRICTION_NOMINATOR() int {
        return this.AIR_FRICTION_NOMINATOR
}


func (this *CharacterFrictionProperties) GetCOLLISION_FRICTION_NOMINATOR() int {
        return this.COLLISION_FRICTION_NOMINATOR
}

