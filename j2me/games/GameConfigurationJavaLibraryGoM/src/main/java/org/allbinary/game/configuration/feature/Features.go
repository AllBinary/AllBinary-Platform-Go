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

type Features struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    List IBasicArrayList
    DefaultList IBasicArrayList
}

var Features_SINGLETON any = NullUtil_GetInstance().NULL_OBJECT

func (this *Features) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func Features_GetInstance() *Features {
        if Features_SINGLETON == NullUtil_GetInstance().NULL_OBJECT {
            Features_SINGLETON = NewFeatures()
        }
        return any(Features_SINGLETON).(*Features)
}


func NewFeatures() *Features {
    this := &Features{}
    this.initFields()
        this.List = NewBasicArrayListD()
        this.DefaultList = NewBasicArrayListD()
        this.Init()
    return this
}


// throws Exception
func (this *Features) AddDefault(gameFeature IFeature) {
        this.Add(gameFeature)
        if !this.DefaultList.Contains(gameFeature) {
            this.DefaultList.Add(gameFeature)
        }
}


// throws Exception
func (this *Features) Add(gameFeature IFeature) {
        if !this.List.Contains(gameFeature) {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Append(fmt.Sprint(gameFeature))), this, this.CommonStrings.ADD)
            this.List.Add(gameFeature)
            GameFeatureEventHandler_GetInstance().FireEvent(NewGameFeatureEvent(gameFeature, fmt.Sprint(gameFeature)))
        }
}


// throws Exception
func (this *Features) RemoveDefault(gameFeature IFeature) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().START_LABEL).Append(fmt.Sprint(gameFeature))), this, "removeDefault")
        this.Remove(gameFeature)
        this.DefaultList.Remove(gameFeature)
}


// throws Exception
func (this *Features) Remove(gameFeature IFeature) {
        if this.List.Contains(gameFeature) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.REMOVE)
            this.List.Remove(gameFeature)
            GameFeatureEventHandler_GetInstance().FireEvent(NewGameFeatureEvent(gameFeature, fmt.Sprint(gameFeature)))
        }
}


func (this *Features) Init() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
                }
            }()
            var gameFeatureFactory *GameFeatureFactory = GameFeatureFactory_GetInstance(); _ = gameFeatureFactory
            var inputFeatureFactory *InputFeatureFactory = InputFeatureFactory_GetInstance(); _ = inputFeatureFactory
            var sensorFeatureFactory *SensorFeatureFactory = SensorFeatureFactory_GetInstance(); _ = sensorFeatureFactory
            var touchFeatureFactory *TouchFeatureFactory = TouchFeatureFactory_GetInstance(); _ = touchFeatureFactory
            this.AddDefault(gameFeatureFactory.ARTIFICIAL_INTELLEGENCE_PROCESSOR)
            this.AddDefault(gameFeatureFactory.COLLIDABLE_INTERFACE_LAYER_PROCESSOR)
            this.AddDefault(gameFeatureFactory.GAME_INPUT_LAYER_PROCESSOR)
            this.AddDefault(gameFeatureFactory.TICKABLE_LAYER_PROCESSOR)
            this.AddDefault(inputFeatureFactory.MULTI_KEY_PRESS)
            this.AddDefault(inputFeatureFactory.REMOVE_DUPLICATE_KEY_PRESSES)
            this.AddDefault(gameFeatureFactory.SCREEN_SHAKE)
            this.AddDefault(gameFeatureFactory.POST_IMAGE_LOADING_MODIFICATION)
            var operatingSystemInterface IGenericOperatingSystem = OperatingSystemFactory_GetInstance().GetOperatingSystemInstance(); _ = operatingSystemInterface
            if operatingSystemInterface.IsOverScan() {
                this.AddDefault(sensorFeatureFactory.NO_ORIENTATION)
                this.AddDefault(touchFeatureFactory.HIDE_SCREEN_BUTTONS)
            } else {
                this.AddDefault(sensorFeatureFactory.ORIENTATION_SENSORS)
                this.AddDefault(touchFeatureFactory.AUTO_HIDE_SHOW_SCREEN_BUTTONS)
            }
            this.AddDefault(touchFeatureFactory.TOUCH_ENABLED)
            this.AddDefault(MainFeatureFactory_GetInstance().FULL_SCREEN)
            this.AddDefault(MainFeatureFactory_GetInstance().LOAD_ONDEMAND)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *Features) IsDefault(gameFeature IFeature) bool {
        return this.DefaultList.Contains(gameFeature)
}


func (this *Features) IsFeature(gameFeature IFeature) bool {
        return this.List.Contains(gameFeature)
}


// throws Exception
func (this *Features) Toggle(gameFeature IFeature) {
        if this.IsFeature(gameFeature) {
            this.Remove(gameFeature)
        } else {
            this.Add(gameFeature)
        }
}

