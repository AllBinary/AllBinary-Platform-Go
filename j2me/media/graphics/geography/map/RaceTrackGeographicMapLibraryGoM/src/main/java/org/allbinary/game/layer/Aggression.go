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

type Aggression struct {
    Ram int
    Projectile int
    Drops int
}


func (this *Aggression) initFields() {
    this.Ram = SmallIntegerSingletonFactory_GetInstance().GetAt(0)
    this.Projectile = SmallIntegerSingletonFactory_GetInstance().GetAt(0)
    this.Drops = SmallIntegerSingletonFactory_GetInstance().GetAt(0)
}


func NewAggression(ram int, projectile int, drops int) *Aggression {
    this := &Aggression{}
    this.initFields()
        this.SetRam(ram)
        this.SetProjectile(projectile)
        this.SetDrops(drops)
    return this
}


func (this *Aggression) GetRam() int {
        return this.Ram
}


func (this *Aggression) SetRam(ram int) {
        this.Ram = ram
}


func (this *Aggression) GetProjectile() int {
        return this.Projectile
}


func (this *Aggression) SetProjectile(projectile int) {
        this.Projectile = projectile
}


func (this *Aggression) GetDrops() int {
        return this.Drops
}


func (this *Aggression) SetDrops(drops int) {
        this.Drops = drops
}

