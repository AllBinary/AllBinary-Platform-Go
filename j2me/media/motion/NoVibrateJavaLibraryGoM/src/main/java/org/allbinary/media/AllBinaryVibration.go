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

type AllBinaryVibration struct {
    AllBinaryVibrationME
}

var AllBinaryVibration_VIBRATION IAllBinaryVibrationME = AllBinaryNoVibration_NO_VIBRATION


func NewAllBinaryVibration() *AllBinaryVibration {
    this := &AllBinaryVibration{}
    this.AllBinaryVibrationME = *NewAllBinaryVibrationME()
    return this
}


func AllBinaryVibration_GetInstance() IAllBinaryVibrationME {
        return AllBinaryVibration_VIBRATION
}


func AllBinaryVibration_Init() {
        if int(GameConfigurationCentral_GetInstance().VIBRATION.GetValue()) == 0 {
            AllBinaryVibration_VIBRATION = AllBinaryNoVibration_NO_VIBRATION
        } else {
            AllBinaryVibration_VIBRATION = NewAllBinaryVibration()
        }
}


func (this *AllBinaryVibration) Vibrate(duration int, type_ int, volume int) {
}

