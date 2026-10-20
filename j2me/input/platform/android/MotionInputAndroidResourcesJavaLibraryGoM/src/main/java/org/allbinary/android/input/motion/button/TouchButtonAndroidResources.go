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

type TouchButtonAndroidResources struct {
    Init
}


func NewTouchButtonAndroidResources() *TouchButtonAndroidResources {
    this := &TouchButtonAndroidResources{}
    this.Init = *NewInit()
    return this
}


func (this *TouchButtonAndroidResources) Init() {
        var resourceUtil *ResourceUtil = ResourceUtil_GetInstance(); _ = resourceUtil
        var androidResources *AndroidResources = AndroidResources_GetInstance(); _ = androidResources
        var RESOURCE_STRING_ARRAY []ITouchButtonResource = []ITouchButtonResource{TouchButtonBlankResource_GetInstance(), TouchButtonGenericActionResource_GetInstance(), TouchButtonStartResource_GetInstance(), TouchButtonUpResource_GetInstance(), TouchButtonDownResource_GetInstance(), TouchButtonTurnLeftResource_GetInstance(), TouchButtonTurnRightResource_GetInstance(), TouchButtonStrafeLeftResource_GetInstance(), TouchButtonStrafeRightResource_GetInstance()}; _ = RESOURCE_STRING_ARRAY
        var RESOURCE_ARRAY []int = make([]int, len(RESOURCE_STRING_ARRAY)); _ = RESOURCE_ARRAY
        if CommonButtons_GetInstance().STANDARD_BUTTON_SIZE == 128 {
            RESOURCE_ARRAY[0] = androidResources.Raw.Touch_button_128_by_128
            RESOURCE_ARRAY[1] = androidResources.Raw.Touch_button_generic_action_128_by_128
            RESOURCE_ARRAY[2] = androidResources.Raw.Touch_button_start_128_by_128
            RESOURCE_ARRAY[3] = androidResources.Raw.Touch_button_up_arrow_128_by_128
            RESOURCE_ARRAY[4] = androidResources.Raw.Touch_button_down_arrow_128_by_128
            RESOURCE_ARRAY[5] = androidResources.Raw.Touch_button_turn_left_arrow_128_by_128
            RESOURCE_ARRAY[6] = androidResources.Raw.Touch_button_turn_right_arrow_128_by_128
            RESOURCE_ARRAY[7] = androidResources.Raw.Touch_button_strafe_left_arrow_128_by_128
            RESOURCE_ARRAY[8] = androidResources.Raw.Touch_button_strafe_right_arrow_128_by_128
            if OperatingSystemFactory_GetInstance().GetOperatingSystemInstance().IsOverScan() {
                resourceUtil.AddResource(RESOURCE_STRING_ARRAY[2].HINT, fmt.Sprint(androidResources.Raw.Ouya_o))
            }
        } else {
            RESOURCE_ARRAY[0] = androidResources.Raw.Touch_button_64_by_64
            RESOURCE_ARRAY[1] = androidResources.Raw.Touch_button_generic_action_64_by_64
            RESOURCE_ARRAY[2] = androidResources.Raw.Touch_button_start_64_by_64
            RESOURCE_ARRAY[3] = androidResources.Raw.Touch_button_up_arrow_64_by_64
            RESOURCE_ARRAY[4] = androidResources.Raw.Touch_button_down_arrow_64_by_64
            RESOURCE_ARRAY[5] = androidResources.Raw.Touch_button_turn_left_arrow_64_by_64
            RESOURCE_ARRAY[6] = androidResources.Raw.Touch_button_turn_right_arrow_64_by_64
            RESOURCE_ARRAY[7] = androidResources.Raw.Touch_button_strafe_left_arrow_64_by_64
            RESOURCE_ARRAY[8] = androidResources.Raw.Touch_button_strafe_right_arrow_64_by_64
        }
        for index := int(0); index < len(RESOURCE_ARRAY); index++ {
            resourceUtil.AddResource(RESOURCE_STRING_ARRAY[index].RESOURCE, fmt.Sprint(RESOURCE_ARRAY[index]))
        }
}

