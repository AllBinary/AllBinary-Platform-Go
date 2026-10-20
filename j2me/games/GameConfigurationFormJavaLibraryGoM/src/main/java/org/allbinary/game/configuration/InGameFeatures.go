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

type InGameFeatures struct {
    Init
}


func NewInGameFeatures() *InGameFeatures {
    this := &InGameFeatures{}
    this.Init = *NewInit()
    return this
}


func (this *InGameFeatures) Init() {
        var LABEL string = "Screen Buttons"; _ = LABEL
        var orientationData *OrientationData = OrientationData_GetInstance(); _ = orientationData
        var exclusiveOrientationSensorVector IBasicArrayList = any(GameFeatureChoiceGroups_GetExclusiveInstance().Get().Get(any(orientationData.ORIENTATION_SENSOR_INPUT).(any))).(IBasicArrayList); _ = exclusiveOrientationSensorVector
        var inGameFeatureChoiceGroups *InGameFeatureChoiceGroups = InGameFeatureChoiceGroups_GetExclusiveInstance(); _ = inGameFeatureChoiceGroups
        if exclusiveOrientationSensorVector != nil {
            var inGameExclusiveOrientationSensorVectorCanBeNull any = inGameFeatureChoiceGroups.Get().Get(any(orientationData.ORIENTATION_SENSOR_INPUT).(any)); _ = inGameExclusiveOrientationSensorVectorCanBeNull
            if inGameExclusiveOrientationSensorVectorCanBeNull == nil {
                this.AddToInGameMenu()
            } else {
                var inGameExclusiveOrientationSensorVector IBasicArrayList = any(inGameExclusiveOrientationSensorVectorCanBeNull).(IBasicArrayList); _ = inGameExclusiveOrientationSensorVector
                if inGameExclusiveOrientationSensorVector.Size() == 0 {
                    this.AddToInGameMenu()
                }
            }
        }
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(TouchFeatureFactory_GetInstance().SHOW_SCREEN_BUTTONS) || features.IsFeature(TouchFeatureFactory_GetInstance().AUTO_HIDE_SHOW_SCREEN_BUTTONS) || features.IsFeature(TouchFeatureFactory_GetInstance().HIDE_SCREEN_BUTTONS) {
            var exclusiveScreenButtonsVector IBasicArrayList = NewBasicArrayListD(); _ = exclusiveScreenButtonsVector
            var touchFeatureFactory *TouchFeatureFactory = TouchFeatureFactory_GetInstance(); _ = touchFeatureFactory
            exclusiveScreenButtonsVector.Add(touchFeatureFactory.AUTO_HIDE_SHOW_SCREEN_BUTTONS)
            exclusiveScreenButtonsVector.Add(touchFeatureFactory.SHOW_SCREEN_BUTTONS)
            exclusiveScreenButtonsVector.Add(touchFeatureFactory.HIDE_SCREEN_BUTTONS)
            inGameFeatureChoiceGroups.Add(LABEL, exclusiveScreenButtonsVector)
        }
}


func (this *InGameFeatures) AddToInGameMenu() {
        var orientationData *OrientationData = OrientationData_GetInstance(); _ = orientationData
        var inGameExclusiveOrientationSensorVector IBasicArrayList = NewBasicArrayListD(); _ = inGameExclusiveOrientationSensorVector
        inGameExclusiveOrientationSensorVector.Add(SensorFeatureFactory_GetInstance().ORIENTATION_SENSORS)
        inGameExclusiveOrientationSensorVector.Add(SensorFeatureFactory_GetInstance().NO_ORIENTATION)
        InGameFeatureChoiceGroups_GetExclusiveInstance().Add(orientationData.ORIENTATION_SENSOR_INPUT, inGameExclusiveOrientationSensorVector)
}


func (this *InGameFeatures) IsAny() bool {
        var features *Features = Features_GetInstance(); _ = features
        var touchFeatureFactory *TouchFeatureFactory = TouchFeatureFactory_GetInstance(); _ = touchFeatureFactory
        if features.IsFeature(touchFeatureFactory.SHOW_SCREEN_BUTTONS) || features.IsFeature(touchFeatureFactory.AUTO_HIDE_SHOW_SCREEN_BUTTONS) || features.IsFeature(touchFeatureFactory.HIDE_SCREEN_BUTTONS) {
            return true
        }
        var exclusiveOrientationSensorVector IBasicArrayList = any(GameFeatureChoiceGroups_GetExclusiveInstance().Get().Get(any(OrientationData_GetInstance().ORIENTATION_SENSOR_INPUT).(any))).(IBasicArrayList); _ = exclusiveOrientationSensorVector
        if exclusiveOrientationSensorVector != nil {
            return true
        }
        return false
}

