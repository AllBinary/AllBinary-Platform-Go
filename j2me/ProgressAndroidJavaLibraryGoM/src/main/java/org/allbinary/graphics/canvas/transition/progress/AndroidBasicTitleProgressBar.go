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

type AndroidBasicTitleProgressBar struct {
    ProgressCanvas
    ShowTitleProgressBarRunnable IARunnable
    DismissTitleProgressBarRunnable IARunnable
    TitleProgressDialogPortionSetProgressRunnable IARunnable
    TitleProgressDialogSetProgressRunnable IARunnable
    MidletActivity *Activity
    Portion int
    IMAGE []IImage
    Image IImage
    Animation IAnimation
}

// Implements DisplayChangeEventListener

var AndroidBasicTitleProgressBar_RESOURCE string = "ProgressImage"
var AndroidBasicTitleProgressBar_Background int

func (this *AndroidBasicTitleProgressBar) initFields() {
    this.ShowTitleProgressBarRunnable = NullRunnable_GetInstance()
    this.DismissTitleProgressBarRunnable = NullRunnable_GetInstance()
    this.TitleProgressDialogPortionSetProgressRunnable = NullRunnable_GetInstance()
    this.TitleProgressDialogSetProgressRunnable = NullRunnable_GetInstance()
    this.MidletActivity = AndroidUtil_NULL_ACTIVITY
    this.Portion = 0
    this.IMAGE = make([]IImage, 4)
    this.Image = NullImage_NULL_IMAGE
    this.Animation = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
}


func NewAndroidBasicTitleProgressBar(title string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *AndroidBasicTitleProgressBar {
    this := &AndroidBasicTitleProgressBar{}
        this.ProgressCanvas = *NewProgressCanvas(title, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
        this.Init()
        DisplayChangeEventHandler_GetInstance().AddListenerInterface(this)
        var size int = len(this.IMAGE); _ = size
        for index := int(0); index < size; index++ {
            this.IMAGE[index] = NullImage_NULL_IMAGE
        }
    return this
}


func (this *AndroidBasicTitleProgressBar) Init() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
                }
            }()
            if AndroidBasicTitleProgressBar_Background != 0 {
                ResourceUtil_GetInstance().AddResource(AndroidBasicTitleProgressBar_RESOURCE, int(AndroidBasicTitleProgressBar_Background))
                GameFeatureImageCacheFactory_Init()
                this.Image = ImageCacheFactory_GetInstance().GetWithKey(AndroidBasicTitleProgressBar_RESOURCE)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) InitActivity(activity *Activity) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.INIT, e)
                }
            }()
            if this.MidletActivity != activity {
                this.MidletActivity = activity
                this.ShowTitleProgressBarRunnable = NewShowTitleProgressBarRunnable(this.MidletActivity, this)
                this.DismissTitleProgressBarRunnable = NewDismissTitleProgressBarRunnable(this.MidletActivity, this)
                this.TitleProgressDialogSetProgressRunnable = NewTitleProgressBarSetProgressRunnable(this.MidletActivity, this)
                this.TitleProgressDialogPortionSetProgressRunnable = NewTitleProgressBarPortionSetProgressRunnable(this.MidletActivity, this)
                this.LoadProgressImages()
            }
            this.UpdateCurrent()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) UpdateCurrent() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.UPDATE, e)
                }
            }()
            if !this.IsBackground() {
                this.Animation = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
            } else if AndroidBasicTitleProgressBar_Background != 0 {
                var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
                var currentImage IImage; _ = currentImage
                if displayInfo.IsPortrait() {
                    currentImage = this.GetImage(0)
                } else {
                    currentImage = this.GetImage(2)
                }
                if currentImage == nil {
                    this.Animation = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
                } else {
                    this.Animation = NewImageAnimation(currentImage, AnimationBehavior_GetInstance())
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *AndroidBasicTitleProgressBar) OnDisplayChangeEvent(displayChangeEvent *DisplayChangeEvent) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF(this.CommonStrings.EXCEPTION_LABEL + ExceptionUtil_GetInstance().GetStackTrace(e), this, this.CanvasStrings.ON_DISPLAY_CHANGE_EVENT)
            this.Animation = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
                }
            }()
            this.LoadProgressImages()
            this.UpdateCurrent()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) LoadProgressImages() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.PutF("IllegalArgumentException " + fmt.Sprint(ExceptionUtil_GetInstance().GetStackTrace(e)), this, "loadProgressImages")
            this.Animation = NullAnimationFactory_GetFactoryInstance().GetInstance(0)
                }
            }()
            var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
            var lastWidth int = displayInfo.GetLastWidth(); _ = lastWidth
            var lastHeight int = displayInfo.GetLastHeight(); _ = lastHeight
            if displayInfo.IsPortraitWH(lastWidth, lastHeight) {
                this.SetImages(0, lastWidth, lastHeight)
            } else {
                this.SetImages(2, lastWidth, lastHeight)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) IsInitialized() bool {
        if this.MidletActivity != AndroidUtil_NULL_ACTIVITY {
            return true
        } else {
            return false
        }
}


func (this *AndroidBasicTitleProgressBar) Start() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.START_METHOD_NAME, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.START_METHOD_NAME)
            this.ProgressCanvas.Start()
            this.MidletActivity.RunOnUiThread(this.ShowTitleProgressBarRunnable)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) End() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.END_METHOD_NAME, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.END_METHOD_NAME)
            this.MidletActivity.RunOnUiThread(this.DismissTitleProgressBarRunnable)
            this.ProgressCanvas.End()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) AddEarlyPortion(value int, text string, index int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.ADD_PORTION, e)
                }
            }()
            this.Portion = value
            this.ProgressCanvas.AddEarlyPortion(value, text, index)
            if this.MidletActivity != AndroidUtil_NULL_ACTIVITY {
                this.MidletActivity.RunOnUiThread(this.TitleProgressDialogPortionSetProgressRunnable)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) AddPortion(value int, text string, index int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.ADD_PORTION, e)
                }
            }()
            this.Portion = value
            this.ProgressCanvas.AddPortion(value, text, index)
            this.MidletActivity.RunOnUiThread(this.TitleProgressDialogPortionSetProgressRunnable)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) AddNormalPortion(value int, text string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.ADD_PORTION, e)
                }
            }()
            this.Portion = value
            this.ProgressCanvas.AddNormalPortion(value, text)
            this.MidletActivity.RunOnUiThread(this.TitleProgressDialogPortionSetProgressRunnable)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) SetValue(value int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setValue", e)
                }
            }()
            this.ProgressCanvas.SetValue(value)
            this.MidletActivity.RunOnUiThread(this.TitleProgressDialogSetProgressRunnable)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidBasicTitleProgressBar) WaitUntilDisplayed() {
}


// throws Exception
func (this *AndroidBasicTitleProgressBar) SetImages(index int, lastWidth int, lastHeight int) {
        var image IImage = this.Image; _ = image
        if image != NullImage_NULL_IMAGE {
            if Features_GetInstance().IsFeature(MainFeatureFactory_GetInstance().FULL_SCREEN) {
                if this.IMAGE[index] == NullImage_NULL_IMAGE {
                    this.IMAGE[index] = ImageScaleUtil_GetInstance().CreateImage2(ImageCacheFactory_GetInstance(), image, float32(lastWidth), float32(image.GetWidth()), float32(lastHeight) - 20.0, float32(image.GetHeight()), false)
                }
            } else {
                var nextIndex int = index + 1; _ = nextIndex
                if this.IMAGE[nextIndex] == NullImage_NULL_IMAGE {
                    this.IMAGE[nextIndex] = ImageScaleUtil_GetInstance().CreateImage2(ImageCacheFactory_GetInstance(), image, float32(lastWidth), float32(image.GetWidth()), float32(lastHeight) - 28.0, float32(image.GetHeight()), false)
                }
            }
        }
}


// throws Exception
func (this *AndroidBasicTitleProgressBar) InitOpenGL(graphics IGraphics) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.INIT)
        this.Image = GameFeatureImageCacheFactory_GetInstance().GetWithKey(AndroidBasicTitleProgressBar_RESOURCE)
        var preResourceImageUtil *PreResourceImageUtil = PreResourceImageUtil_GetInstance(); _ = preResourceImageUtil
        var size int = len(this.IMAGE); _ = size
        for index := int(0); index < size; index++ {
            if this.IMAGE[index] != NullImage_NULL_IMAGE {
                this.IMAGE[index] = preResourceImageUtil.Encapsulate(this.IMAGE[index])
            }
        }
        this.UpdateCurrent()
}


// throws Exception
func (this *AndroidBasicTitleProgressBar) Update(graphics IGraphics) {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.UPDATE)
        this.InitOpenGL(graphics)
        this.Image = GameFeatureImageCacheFactory_GetInstance().GetWithKey(AndroidBasicTitleProgressBar_RESOURCE)
        var preResourceImageUtil *PreResourceImageUtil = PreResourceImageUtil_GetInstance(); _ = preResourceImageUtil
        var size int = len(this.IMAGE); _ = size
        for index := int(0); index < size; index++ {
            if this.IMAGE[index] != NullImage_NULL_IMAGE {
                preResourceImageUtil.Update(graphics, this.IMAGE[index])
            }
        }
}


// throws Exception
func (this *AndroidBasicTitleProgressBar) GetImage(index int) IImage {
        var image IImage = NullImage_NULL_IMAGE; _ = image
        if Features_GetInstance().IsFeature(MainFeatureFactory_GetInstance().FULL_SCREEN) {
            image = this.IMAGE[index]
        } else {
            image = this.IMAGE[index + 1]
        }
        return image
}


func (this *AndroidBasicTitleProgressBar) Paint2(graphics IGraphics) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CanvasStrings.PAINT, e)
                }
            }()
            this.Animation.PaintXY(graphics, 0, 20)
            this.ProgressCanvas.Paint2(graphics)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func AndroidBasicTitleProgressBar_SetBackgroundResource(background int) {
        AndroidBasicTitleProgressBar_Background = background
}


func (this *AndroidBasicTitleProgressBar) SetBackground(background bool) {
        this.ProgressCanvas.SetBackground(background)
        this.UpdateCurrent()
}


func (this *AndroidBasicTitleProgressBar) SetPortion(portion int) {
        this.Portion = portion
}


func (this *AndroidBasicTitleProgressBar) GetPortion() int {
        return this.Portion
}

