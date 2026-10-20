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

type DamageFloaters struct {
    Paintable
}

// The Java class DamageFloaters has subclasses so its variables are this interface
type IDamageFloaters interface {
    IPaintable
    DamageFloaters_() *DamageFloaters
    Add(damage int)
}

func (this *DamageFloaters) DamageFloaters_() *DamageFloaters {
    return this
}

var DamageFloaters_SINGLETON IDamageFloaters = NewDamageFloaters()

func NewDamageFloaters() *DamageFloaters {
    this := &DamageFloaters{}
    this.Paintable = *NewPaintable()
    return this
}


func DamageFloaters_GetInstance() IDamageFloaters {
        return DamageFloaters_SINGLETON
}


func (this *DamageFloaters) Add(damage int) {
}

