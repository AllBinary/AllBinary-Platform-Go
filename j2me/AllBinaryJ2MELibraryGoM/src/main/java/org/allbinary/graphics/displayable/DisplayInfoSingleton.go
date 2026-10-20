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

type DisplayInfoSingleton struct {
    LogUtil *LogUtil
    CommonLabels *CommonLabels
    CommonSeps *CommonSeps
    CommonStrings *CommonStrings
    ORIENTATION string
    ADJUSTING_FOR_SCALING_IN_PORTRAIT string
    ADJUSTING_FOR_SCALING_IN_LANDSCAPE string
    Last []int
    LastHalf []int
    Full []int
    Top int
    Left int
    XOffset int
    YOffset int
    ScaleLargestTo float32
    WIDTH int
    HEIGHT int
    CUSTOM_WIDTH int
    CUSTOM_HEIGHT int
    ScalableListener IBaseScalable
    DisplayRatio float32
    Ratio float32
    SET_LAST_SIZE_METHOD_NAME string
    FULL_WIDTH string
    FULL_HEIGHT string
    LAST_WIDTH string
    LAST_HEIGHT string
    PORTRAIT string
    LANDSCAPE string
    DisplayChangeEvent *DisplayChangeEvent
    FIRE_METHOD_NAME string
    REASON string
    UPDATE_FROM_ORIENTATION_CHANGE string
    List IBasicArrayList
    DISPLAY_INFO string
    FULL string
    LAST string
    LAST_HALF string
}

var DisplayInfoSingleton_SINGLETON *DisplayInfoSingleton = NewDisplayInfoSingleton()

func (this *DisplayInfoSingleton) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonLabels = CommonLabels_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ORIENTATION = "ORIENTATION"
    this.ADJUSTING_FOR_SCALING_IN_PORTRAIT = "Adjusting for Scaling in portrait display ratio: "
    this.ADJUSTING_FOR_SCALING_IN_LANDSCAPE = "Adjusting for Scaling in landscape display ratio: "
    this.Last = make([]int, 4)
    this.LastHalf = make([]int, 4)
    this.Full = make([]int, 4)
    
    
    
    
    
    this.WIDTH = 0
    this.HEIGHT = 1
    this.CUSTOM_WIDTH = 2
    this.CUSTOM_HEIGHT = 3
    this.ScalableListener = NewBaseScalable()
    
    this.Ratio = 1.0
    this.SET_LAST_SIZE_METHOD_NAME = "setLastSize"
    this.FULL_WIDTH = " FullWidth: "
    this.FULL_HEIGHT = " FullHeight: "
    this.LAST_WIDTH = "LastWidth: "
    this.LAST_HEIGHT = " LastHeight: "
    this.PORTRAIT = "Portrait"
    this.LANDSCAPE = "Landscape"
    this.DisplayChangeEvent = NewDisplayChangeEvent(this)
    this.FIRE_METHOD_NAME = "fire"
    this.REASON = "reason: "
    this.UPDATE_FROM_ORIENTATION_CHANGE = "Updating from Orientation Change "
    this.List = NewBasicArrayListD()
    this.DISPLAY_INFO = "Display Info: "
    this.FULL = "full"
    this.LAST = "last"
    this.LAST_HALF = "lastHalf"
}


func DisplayInfoSingleton_GetInstance() *DisplayInfoSingleton {
        return DisplayInfoSingleton_SINGLETON
}


func NewDisplayInfoSingleton() *DisplayInfoSingleton {
    this := &DisplayInfoSingleton{}
    this.initFields()
        var size int = len(this.Last); _ = size
        for index := int(0); index < size; index++ {
            this.Last[index] = 0
            this.LastHalf[index] = 0
            this.Full[index] = 0
        }
        if AndroidUtil_IsAndroid() {
            this.ScaleLargestTo = 640.0
        } else {
            this.ScaleLargestTo = 1080.0
        }
    return this
}


func (this *DisplayInfoSingleton) GetLastHalf() []int {
        return this.LastHalf
}


func (this *DisplayInfoSingleton) GetLast() []int {
        return this.Last
}


func (this *DisplayInfoSingleton) GetFull() []int {
        return this.Full
}


func (this *DisplayInfoSingleton) GetScaleLargestTo() float32 {
        return this.ScaleLargestTo
}


func (this *DisplayInfoSingleton) SetScaleLargestTo(scaleLargestTo int) {
        this.ScaleLargestTo = float32(scaleLargestTo)
}


func (this *DisplayInfoSingleton) SetLastSize(aLastWidth int, aLastHeight int, reason string) {
        if this.Full[this.WIDTH] != aLastWidth || this.Full[this.HEIGHT] != aLastHeight {
            this.SetLastSizeForce(aLastWidth, aLastHeight, reason)
        }
}


func (this *DisplayInfoSingleton) SetLastSizeForce(aLastWidth int, aLastHeight int, reason string) {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(CommonLabels_GetInstance().START_LABEL).Append(reason)), this, this.SET_LAST_SIZE_METHOD_NAME)
        var aFullWidth int = aLastWidth; _ = aFullWidth
        var aFullHeight int = aLastHeight; _ = aFullHeight
        stringMaker.Delete(0, stringMaker.Length())
        this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.FULL_WIDTH).Appendint(aFullWidth).Append(this.FULL_HEIGHT).Appendint(aFullHeight).Append(this.ToString())), this, this.SET_LAST_SIZE_METHOD_NAME)
        var operatingSystemInterface IGenericOperatingSystem = OperatingSystemFactory_GetInstance().GetOperatingSystemInstance(); _ = operatingSystemInterface
        if operatingSystemInterface.IsOverScan() {
            aLastWidth = aLastWidth * operatingSystemInterface.GetOverScanXPercent() / 100
            aLastHeight = aLastHeight * operatingSystemInterface.GetOverScanYPercent() / 100
        }
        if operatingSystemInterface.IsScalable() {
            if this.IsPortraitWH(aLastWidth, aLastHeight) {
                if float32(aLastHeight) > this.ScaleLargestTo {
                    this.DisplayRatio = this.ScaleLargestTo / float32(aLastHeight)
                    this.Ratio = float32(aLastHeight) / this.ScaleLargestTo
                    stringMaker.Delete(0, stringMaker.Length())
                    this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.ADJUSTING_FOR_SCALING_IN_PORTRAIT).Appendfloat(this.DisplayRatio)), this, this.SET_LAST_SIZE_METHOD_NAME)
                    aLastWidth = int((float32(aLastWidth) * this.DisplayRatio))
                    aLastHeight = int((float32(aLastHeight) * this.DisplayRatio))
                    this.ScalableListener.Scale(this.Ratio)
                } else {
                    this.Ratio = 1.0
                    this.ScalableListener.Scale(this.Ratio)
                }
            } else {
                if float32(aLastWidth) > this.ScaleLargestTo {
                    this.DisplayRatio = this.ScaleLargestTo / float32(aLastWidth)
                    this.Ratio = float32(aLastWidth) / this.ScaleLargestTo
                    stringMaker.Delete(0, stringMaker.Length())
                    this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.ADJUSTING_FOR_SCALING_IN_LANDSCAPE).Appendfloat(this.DisplayRatio)), this, this.SET_LAST_SIZE_METHOD_NAME)
                    aLastWidth = int((float32(aLastWidth) * this.DisplayRatio))
                    aLastHeight = int((float32(aLastHeight) * this.DisplayRatio))
                    this.ScalableListener.Scale(this.Ratio)
                } else {
                    this.Ratio = 1.0
                    this.ScalableListener.Scale(this.Ratio)
                }
            }
        }
        stringMaker.Delete(0, stringMaker.Length())
        this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.LAST_WIDTH).Appendint(aLastWidth).Append(this.LAST_HEIGHT).Appendint(aLastHeight)), this, this.SET_LAST_SIZE_METHOD_NAME)
        this.XOffset = aFullWidth - aLastWidth
        this.YOffset = aFullHeight - aLastHeight
        this.Left = this.ScalableListener.GetLeft(this.XOffset)
        this.Top = this.ScalableListener.GetTop(this.YOffset)
        this.Full[this.WIDTH] = aFullWidth
        this.Full[this.HEIGHT] = aFullHeight
        this.Last[this.WIDTH] = aLastWidth
        this.LastHalf[this.WIDTH] = (this.Last[this.WIDTH] >> 1)
        this.Last[this.HEIGHT] = aLastHeight
        this.LastHalf[this.HEIGHT] = (this.Last[this.HEIGHT] >> 1)
        SWTJOGLProcessor_GetInstance().SetCustom(aLastWidth, aLastHeight, this.Ratio)
        this.Add(this.SET_LAST_SIZE_METHOD_NAME)
}


func (this *DisplayInfoSingleton) IsPortraitWH(lastWidth int, lastHeight int) bool {
        if lastHeight > lastWidth {
            return true
        } else {
            return false
        }
}


func (this *DisplayInfoSingleton) IsPortrait() bool {
        return this.IsPortraitWH(this.Last[this.WIDTH], this.Last[this.HEIGHT])
}


func (this *DisplayInfoSingleton) Add(reason string) {
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.REASON).Append(reason)), this, this.FIRE_METHOD_NAME)
        stringMaker.Delete(0, stringMaker.Length())
        this.LogUtil.PutF(this.ToStringAppend(stringMaker), this, this.FIRE_METHOD_NAME)
        this.List.Add(reason)
}


func (this *DisplayInfoSingleton) Process() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, this.FIRE_METHOD_NAME, e)
                }
            }()
            if this.List.Size() > 0 {
                this.ProcessForced()
            }
            this.List.Clear()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DisplayInfoSingleton) ProcessForced() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE(this.CommonStrings.EXCEPTION, this, this.FIRE_METHOD_NAME, e)
                }
            }()
            var swtJOGLProcessor *SWTJOGLProcessor = SWTJOGLProcessor_GetInstance(); _ = swtJOGLProcessor
            swtJOGLProcessor.Clear()
            DisplayChangeEventHandler_GetInstance().FireEvent(this.DisplayChangeEvent)
            LastDisplayChangeEventHandler_GetInstance().FireEvent(this.DisplayChangeEvent)
            swtJOGLProcessor.OnSurfaceChanged()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *DisplayInfoSingleton) Update(displayable IDisplayable, reason string) {
        var aLastWidth int = displayable.GetWidth(); _ = aLastWidth
        var aLastHeight int = displayable.GetHeight(); _ = aLastHeight
        var aFullWidth int = aLastWidth; _ = aFullWidth
        var aFullHeight int = aLastHeight; _ = aFullHeight
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(CommonLabels_GetInstance().START_LABEL).Append(reason).Append(this.LAST_WIDTH).Appendint(aLastWidth).Append(this.LAST_HEIGHT).Appendint(aLastHeight).Append(this.CommonSeps.SPACE).Append(this.ToString())), this, this.CommonStrings.UPDATE)
        if aLastWidth > 0 && aLastHeight > 0 {
            var features *Features = Features_GetInstance(); _ = features
            var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
            if !features.IsDefault(openGLFeatureFactory.OPENGL) || this.Last[this.WIDTH] != aLastWidth || this.Last[this.HEIGHT] != aLastHeight {
                stringMaker.Delete(0, stringMaker.Length())
                this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.UPDATE_FROM_ORIENTATION_CHANGE)), this, this.CommonStrings.UPDATE)
                var operatingSystemInterface IGenericOperatingSystem = OperatingSystemFactory_GetInstance().GetOperatingSystemInstance(); _ = operatingSystemInterface
                if operatingSystemInterface.IsOverScan() {
                    aLastWidth = aLastWidth * operatingSystemInterface.GetOverScanXPercent() / 100
                    aLastHeight = aLastHeight * operatingSystemInterface.GetOverScanYPercent() / 100
                }
                if operatingSystemInterface.IsScalable() {
                    if this.IsPortraitWH(aLastWidth, aLastHeight) {
                        if float32(aLastHeight) > this.ScaleLargestTo {
                            this.DisplayRatio = this.ScaleLargestTo / float32(aLastHeight)
                            this.Ratio = float32(aLastHeight) / this.ScaleLargestTo
                            stringMaker.Delete(0, stringMaker.Length())
                            this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.ADJUSTING_FOR_SCALING_IN_PORTRAIT).Appendfloat(this.DisplayRatio)), this, this.CommonStrings.UPDATE)
                            aLastWidth = int((float32(aLastWidth) * this.DisplayRatio))
                            aLastHeight = int((float32(aLastHeight) * this.DisplayRatio))
                            this.ScalableListener.Scale(this.Ratio)
                        } else {
                            this.Ratio = 1.0
                            this.ScalableListener.Scale(this.Ratio)
                        }
                    } else {
                        if float32(aLastWidth) > this.ScaleLargestTo {
                            this.DisplayRatio = this.ScaleLargestTo / float32(aLastWidth)
                            this.Ratio = float32(aLastWidth) / this.ScaleLargestTo
                            stringMaker.Delete(0, stringMaker.Length())
                            this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.ADJUSTING_FOR_SCALING_IN_LANDSCAPE).Appendfloat(this.DisplayRatio)), this, this.CommonStrings.UPDATE)
                            aLastWidth = int((float32(aLastWidth) * this.DisplayRatio))
                            aLastHeight = int((float32(aLastHeight) * this.DisplayRatio))
                            this.ScalableListener.Scale(this.Ratio)
                        } else {
                            this.Ratio = 1.0
                            this.ScalableListener.Scale(this.Ratio)
                        }
                    }
                }
                stringMaker.Delete(0, stringMaker.Length())
                this.LogUtil.PutF(fmt.Sprint(stringMaker.Append(this.UPDATE_FROM_ORIENTATION_CHANGE).Append(this.LAST_WIDTH).Appendint(aLastWidth).Append(this.LAST_HEIGHT).Appendint(aLastHeight)), this, this.CommonStrings.UPDATE)
                this.XOffset = aFullWidth - aLastWidth
                this.YOffset = aFullHeight - aLastHeight
                this.Left = this.ScalableListener.GetLeft(this.XOffset)
                this.Top = this.ScalableListener.GetTop(this.YOffset)
                this.Full[this.WIDTH] = aFullWidth
                this.Full[this.HEIGHT] = aFullHeight
                this.Last[this.WIDTH] = aLastWidth
                this.LastHalf[this.WIDTH] = (this.Last[this.WIDTH] >> 1)
                this.Last[this.HEIGHT] = aLastHeight
                this.LastHalf[this.HEIGHT] = (this.Last[this.HEIGHT] >> 1)
                SWTJOGLProcessor_GetInstance().SetCustom(aLastWidth, aLastHeight, this.Ratio)
                this.Add(this.CommonStrings.UPDATE)
            }
        }
}


func (this *DisplayInfoSingleton) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        return this.ToStringAppend(stringBuffer)
}

func (this *DisplayInfoSingleton) String() string {
    return this.ToString()
}


func (this *DisplayInfoSingleton) ToStringAppend(stringBuffer *StringMaker) string {
        stringBuffer.Append(this.DISPLAY_INFO)
        stringBuffer.Append(this.FULL)
        stringBuffer.Append(this.CommonLabels.WIDTH_LABEL)
        stringBuffer.Appendint(this.Full[this.WIDTH])
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.FULL)
        stringBuffer.Append(this.CommonLabels.HEIGHT_LABEL)
        stringBuffer.Appendint(this.Full[this.HEIGHT])
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.LAST)
        stringBuffer.Append(this.CommonLabels.WIDTH_LABEL)
        stringBuffer.Appendint(this.Last[this.WIDTH])
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.LAST)
        stringBuffer.Append(this.CommonLabels.HEIGHT_LABEL)
        stringBuffer.Appendint(this.Last[this.HEIGHT])
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.LAST_HALF)
        stringBuffer.Append(this.CommonLabels.WIDTH_LABEL)
        stringBuffer.Appendint(this.LastHalf[this.WIDTH])
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.LAST_HALF)
        stringBuffer.Append(this.CommonLabels.HEIGHT_LABEL)
        stringBuffer.Appendint(this.LastHalf[this.HEIGHT])
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(func() string { if this.IsPortrait() { return this.PORTRAIT }; return this.LANDSCAPE }())
        return fmt.Sprint(stringBuffer)
}


func (this *DisplayInfoSingleton) GetTop() int {
        return this.Top
}


func (this *DisplayInfoSingleton) GetLeft() int {
        return this.Left
}


func (this *DisplayInfoSingleton) GetLastWidth() int {
        return this.Last[this.WIDTH]
}


func (this *DisplayInfoSingleton) GetLastHeight() int {
        return this.Last[this.HEIGHT]
}


func (this *DisplayInfoSingleton) GetLastHalfWidth() int {
        return this.LastHalf[this.WIDTH]
}


func (this *DisplayInfoSingleton) GetLastHalfHeight() int {
        return this.LastHalf[this.HEIGHT]
}


func (this *DisplayInfoSingleton) GetCustomLastWidth() int {
        return this.Last[this.CUSTOM_WIDTH]
}


func (this *DisplayInfoSingleton) GetCustomLastHeight() int {
        return this.Last[this.CUSTOM_HEIGHT]
}


func (this *DisplayInfoSingleton) GetRawLastWidth() int {
        return this.Last[this.WIDTH]
}


func (this *DisplayInfoSingleton) GetRawLastHeight() int {
        return this.Last[this.HEIGHT]
}


func (this *DisplayInfoSingleton) SetCustom(width int, height int) {
        this.Last[this.CUSTOM_WIDTH] = width
        this.LastHalf[this.CUSTOM_WIDTH] = (this.Last[this.CUSTOM_WIDTH] >> 1)
        this.Last[this.CUSTOM_HEIGHT] = height
        this.LastHalf[this.CUSTOM_HEIGHT] = (this.Last[this.CUSTOM_HEIGHT] >> 1)
}


func (this *DisplayInfoSingleton) SetScalableListener(scalableListener IBaseScalable) {
        this.ScalableListener = scalableListener
}


func (this *DisplayInfoSingleton) GetScalableListener() ScalableListener {
        return this.ScalableListener
}


func (this *DisplayInfoSingleton) GetRatio() float32 {
        return this.Ratio
}


func (this *DisplayInfoSingleton) GetDisplayRatio() float32 {
        return this.DisplayRatio
}


func (this *DisplayInfoSingleton) SetOffset(left int, top int) {
        this.Left = left
}


func (this *DisplayInfoSingleton) GetxOffset() int {
        return this.XOffset
}


func (this *DisplayInfoSingleton) GetyOffset() int {
        return this.YOffset
}

