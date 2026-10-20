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

type UnitTechnologySingletonFactory struct {
    Damage int
    Health int
    Reload int
    Repair int
    Speed int
}

var UnitTechnologySingletonFactory_Instance *UnitTechnologySingletonFactory = NewUnitTechnologySingletonFactory()

func (this *UnitTechnologySingletonFactory) initFields() {
    
    
    
    
    
}


func UnitTechnologySingletonFactory_GetInstance() *UnitTechnologySingletonFactory {
        return UnitTechnologySingletonFactory_Instance
}


func NewUnitTechnologySingletonFactory() *UnitTechnologySingletonFactory {
    this := &UnitTechnologySingletonFactory{}
    this.initFields()
    return this
}


func (this *UnitTechnologySingletonFactory) Init() {
        this.SetDamage(0)
        this.SetHealth(0)
        this.SetReload(0)
        this.SetRepair(0)
        this.SetSpeed(0)
}


func (this *UnitTechnologySingletonFactory) GetDamage() int {
        return this.Damage
}


func (this *UnitTechnologySingletonFactory) SetDamage(damage int) {
        this.Damage = damage
}


func (this *UnitTechnologySingletonFactory) GetHealth() int {
        return this.Health
}


func (this *UnitTechnologySingletonFactory) SetHealth(health int) {
        this.Health = health
}


func (this *UnitTechnologySingletonFactory) GetReload() int {
        return this.Reload
}


func (this *UnitTechnologySingletonFactory) SetReload(reload int) {
        this.Reload = reload
}


func (this *UnitTechnologySingletonFactory) GetRepair() int {
        return this.Repair
}


func (this *UnitTechnologySingletonFactory) SetRepair(repair int) {
        this.Repair = repair
}


func (this *UnitTechnologySingletonFactory) GetSpeed() int {
        return this.Speed
}


func (this *UnitTechnologySingletonFactory) SetSpeed(speed int) {
        this.Speed = speed
}

