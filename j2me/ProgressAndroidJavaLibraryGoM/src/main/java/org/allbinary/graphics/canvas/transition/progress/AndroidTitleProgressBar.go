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

type AndroidTitleProgressBar struct {
    ProgressCanvas
    ShowTitleProgressBarRunnable *ShowTitleProgressBarRunnable
    DismissTitleProgressBarRunnable *DismissTitleProgressBarRunnable
    ProgressDialogPortionSetProgressRunnable *TitleProgressBarPortionSetProgressRunnable
    ProgressDialogSetProgressRunnable *TitleProgressBarSetProgressRunnable
    ProgressActivity ProgressActivityInterface
    Portion int
}


func (this *AndroidTitleProgressBar) initFields() {
    this.ShowTitleProgressBarRunnable = NewShowTitleProgressBarRunnable(this)
    this.DismissTitleProgressBarRunnable = NewDismissTitleProgressBarRunnable(this)
    this.ProgressDialogPortionSetProgressRunnable = NewTitleProgressBarPortionSetProgressRunnable(this)
    this.ProgressDialogSetProgressRunnable = NewTitleProgressBarSetProgressRunnable(this)
    this.ProgressActivity = NullProgressActivity_NULL_PROGRESS_ACTIVITY
    this.Portion = 0
}


func NewAndroidTitleProgressBar(title string, backgroundBasicColor *BasicColor, foregroundBasicColor *BasicColor) *AndroidTitleProgressBar {
    this := &AndroidTitleProgressBar{}
        this.ProgressCanvas = *NewProgressCanvas(title, backgroundBasicColor, foregroundBasicColor)
        this.initFields()
    return this
}


func (this *AndroidTitleProgressBar) Init(activity ProgressActivityInterface) {
        this.ProgressActivity = activity
}


func (this *AndroidTitleProgressBar) IsInitialized() bool {
        if this.ProgressActivity != NullProgressActivity_NULL_PROGRESS_ACTIVITY {
            return true
        } else {
            return false
        }
}


func (this *AndroidTitleProgressBar) Start() {
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
            this.ProgressActivity.RunOnUiThread(this.ShowTitleProgressBarRunnable)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidTitleProgressBar) End() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.END_METHOD_NAME, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.END_METHOD_NAME)
            this.ProgressActivity.RunOnUiThread(this.DismissTitleProgressBarRunnable)
            this.ProgressCanvas.End()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidTitleProgressBar) AddPortion(value int, text string, index int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.ADD_PORTION, e)
                }
            }()
            this.ProgressCanvas.AddPortion(value, text, index)
            this.Portion = value
            this.ProgressActivity.RunOnUiThread(this.ProgressDialogPortionSetProgressRunnable)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidTitleProgressBar) AddNormalPortion(value int, text string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.ADD_PORTION, e)
                }
            }()
            this.ProgressCanvas.AddNormalPortion(value, text)
            this.Portion = value
            this.ProgressActivity.RunOnUiThread(this.ProgressDialogPortionSetProgressRunnable)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidTitleProgressBar) SetValue(value int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setValue", e)
                }
            }()
            this.ProgressCanvas.SetValue(value)
            this.ProgressActivity.RunOnUiThread(this.ProgressDialogSetProgressRunnable)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *AndroidTitleProgressBar) WaitUntilDisplayed() {
}


func (this *AndroidTitleProgressBar) Paint(graphics IGraphics) {
}

type TitleProgressBarSetProgressRunnable struct {
    AndroidTitleProgressBar *AndroidTitleProgressBar
}

// Implements Runnable


func (this *TitleProgressBarSetProgressRunnable) initFields() {
        
}
    

    func NewAndroidTitleProgressBar(androidTitleProgressBar *AndroidTitleProgressBar) *AndroidTitleProgressBar {
        this := &AndroidTitleProgressBar{}
        this.initFields()
            this.AndroidTitleProgressBar = androidTitleProgressBar
        return this
    }
    

    func (this *AndroidTitleProgressBar) Run() {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                    }
                }()
                var value int = int(this.AndroidTitleProgressBar.GetValue()); _ = value
                this.AndroidTitleProgressBar.ProgressActivity.OnTitleProgressBarSetProgress(value)
                return
            }(); tryControl1 == 1 {
                return
            }
    }

type TitleProgressBarPortionSetProgressRunnable struct {
    AndroidTitleProgressBar *AndroidTitleProgressBar
}

// Implements Runnable


func (this *TitleProgressBarPortionSetProgressRunnable) initFields() {
        
}
    

    func NewAndroidTitleProgressBar(androidTitleProgressBar *AndroidTitleProgressBar) *AndroidTitleProgressBar {
        this := &AndroidTitleProgressBar{}
        this.initFields()
            this.AndroidTitleProgressBar = androidTitleProgressBar
        return this
    }
    

    func (this *AndroidTitleProgressBar) Run() {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                    }
                }()
                var value int = int((this.AndroidTitleProgressBar.GetValue() + this.AndroidTitleProgressBar.GetMaxValue() / this.AndroidTitleProgressBar.Portion)); _ = value
                this.AndroidTitleProgressBar.ProgressActivity.OnTitleProgressBarSetProgress(value)
                return
            }(); tryControl1 == 1 {
                return
            }
    }

type ShowTitleProgressBarRunnable struct {
    AndroidTitleProgressBar *AndroidTitleProgressBar
}

// Implements Runnable


func (this *ShowTitleProgressBarRunnable) initFields() {
        
}
    

    func NewAndroidTitleProgressBar(androidTitleProgressBar *AndroidTitleProgressBar) *AndroidTitleProgressBar {
        this := &AndroidTitleProgressBar{}
        this.initFields()
            this.AndroidTitleProgressBar = androidTitleProgressBar
        return this
    }
    

    func (this *AndroidTitleProgressBar) Run() {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                    }
                }()
                var maxValue int = int(this.AndroidTitleProgressBar.GetMaxValue()); _ = maxValue
                this.AndroidTitleProgressBar.ProgressActivity.OnShowTitleProgressBar(maxValue, false)
                return
            }(); tryControl1 == 1 {
                return
            }
    }

type DismissTitleProgressBarRunnable struct {
    AndroidTitleProgressBar *AndroidTitleProgressBar
}

// Implements Runnable


func (this *DismissTitleProgressBarRunnable) initFields() {
        
}
    

    func NewAndroidTitleProgressBar(androidTitleProgressBar *AndroidTitleProgressBar) *AndroidTitleProgressBar {
        this := &AndroidTitleProgressBar{}
        this.initFields()
            this.AndroidTitleProgressBar = androidTitleProgressBar
        return this
    }
    

    func (this *AndroidTitleProgressBar) Run() {
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            if tryControl1 := func() (tryControl1 int) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e := Java_ToError(recovered)
                        _ = e
                logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                    }
                }()
                this.AndroidTitleProgressBar.ProgressActivity.OnDismissTitleProgressBar()
                return
            }(); tryControl1 == 1 {
                return
            }
    }

