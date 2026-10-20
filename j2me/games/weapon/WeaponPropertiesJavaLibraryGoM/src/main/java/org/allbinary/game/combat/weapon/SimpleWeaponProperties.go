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

type SimpleWeaponProperties struct {
    Damage int
    Dissipation int16
    Range int
}

// The Java class SimpleWeaponProperties has subclasses so its variables are this interface
type ISimpleWeaponProperties interface {
    SimpleWeaponProperties_() *SimpleWeaponProperties
    Set(simpleWeaponProperties ISimpleWeaponProperties)
    SetDamage(damage int)
    GetDamage() int
    SetDissipation(dissipation int16)
    GetDissipation() int16
    GetRange() int
    SetRange(range_ int)
}

func (this *SimpleWeaponProperties) SimpleWeaponProperties_() *SimpleWeaponProperties {
    return this
}


func (this *SimpleWeaponProperties) initFields() {
    
    
    
}

func NewSimpleWeaponProperties() *SimpleWeaponProperties {
    this := &SimpleWeaponProperties{}
    this.initFields()
    return this
}


func (this *SimpleWeaponProperties) Set(simpleWeaponProperties ISimpleWeaponProperties) {
        this.SetDamage(simpleWeaponProperties.GetDamage())
        this.SetDissipation(simpleWeaponProperties.GetDissipation())
        this.SetRange(simpleWeaponProperties.GetRange())
}


func (this *SimpleWeaponProperties) SetDamage(damage int) {
        this.Damage = damage
}


func (this *SimpleWeaponProperties) GetDamage() int {
        return this.Damage
}


func (this *SimpleWeaponProperties) SetDissipation(dissipation int16) {
        this.Dissipation = dissipation
}


func (this *SimpleWeaponProperties) GetDissipation() int16 {
        return this.Dissipation
}


func (this *SimpleWeaponProperties) GetRange() int {
        return this.Range
}


func (this *SimpleWeaponProperties) SetRange(range_ int) {
        this.Range = range_
}

