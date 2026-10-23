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

type PositionStrings struct {
    X string
    Y string
    X_LABEL string
    Y_LABEL string
    Z_LABEL string
    DX_LABEL string
    DY_LABEL string
    DZ_LABEL string
}

var PositionStrings_Instance *PositionStrings = NewPositionStrings()

func (this *PositionStrings) initFields() {
    this.X = "x"
    this.Y = "y"
    this.X_LABEL = "x: "
    this.Y_LABEL = "y: "
    this.Z_LABEL = "z: "
    this.DX_LABEL = "dx: "
    this.DY_LABEL = "dy: "
    this.DZ_LABEL = "dz: "
}

func NewPositionStrings() *PositionStrings {
    this := &PositionStrings{}
    this.initFields()
    return this
}


func PositionStrings_GetInstance() *PositionStrings {
        return PositionStrings_Instance
}

