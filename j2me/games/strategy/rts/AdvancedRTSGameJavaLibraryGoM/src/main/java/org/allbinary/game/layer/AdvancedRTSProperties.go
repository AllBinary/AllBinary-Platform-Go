/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type AdvancedRTSProperties struct {
    ShakeListener IShakeAnimationListener
    Vibration IAllBinaryVibrationME
    Duration int
    WaypointBehaviorBase TickableInterface
}


func (this *AdvancedRTSProperties) initFields() {
    
    
    
    
}


func AdvancedRTSProperties_CreatePropertiesSimulated() *AdvancedRTSProperties {
        return NewAdvancedRTSProperties(NoShakeAnimationListener_NO_SHAKE_ANIMATION_LISTENER, AllBinaryNoVibration_NO_VIBRATION, 0, NewWaypointBehaviorBase())
}


func AdvancedRTSProperties_CreateProperties() *AdvancedRTSProperties {
        return NewAdvancedRTSProperties(ShakeAnimationListenerFactory_GetInstance(), AllBinaryVibration_GetInstance(), int(GameConfigurationCentral_GetInstance().VIBRATION.GetValue()) * 100, NewWaypointBehaviorBase())
}


func NewAdvancedRTSProperties(shakeListener IShakeAnimationListener, vibration IAllBinaryVibrationME, duration int, waypointBehaviorBase TickableInterface) *AdvancedRTSProperties {
    this := &AdvancedRTSProperties{}
    this.initFields()
        this.ShakeListener = shakeListener
        this.Vibration = vibration
        this.Duration = duration
        this.WaypointBehaviorBase = waypointBehaviorBase
    return this
}

