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

type BasicHudFactory struct {
    DIRECTION_EXCEPTION string
    VERTICAL int
    HORIZONTAL int
    BOTTOMLEFT int
    BOTTOMRIGHT int
    TOPLEFT int
    TOPRIGHT int
    BOTTOMCENTER int
    TOPCENTER int
    ABSOLUTE int
}

var BasicHudFactory_Instance *BasicHudFactory = NewBasicHudFactory()

func (this *BasicHudFactory) initFields() {
    this.DIRECTION_EXCEPTION = "Only Horizontal Direction is Allowed Currently"
    this.VERTICAL = 0
    this.HORIZONTAL = 1
    this.BOTTOMLEFT = 0
    this.BOTTOMRIGHT = 1
    this.TOPLEFT = 2
    this.TOPRIGHT = 3
    this.BOTTOMCENTER = 4
    this.TOPCENTER = 5
    this.ABSOLUTE = 6
}


func BasicHudFactory_GetInstance() *BasicHudFactory {
        return BasicHudFactory_Instance
}


func NewBasicHudFactory() *BasicHudFactory {
    this := &BasicHudFactory{}
    this.initFields()
    return this
}

