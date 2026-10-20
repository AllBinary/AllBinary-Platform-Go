/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type RepaintBehavior struct {
    LogUtil *LogUtil
    NAME string
}

// The Java class RepaintBehavior has subclasses so its variables are this interface
type IRepaintBehavior interface {
    RepaintBehavior_() *RepaintBehavior
    Repaint(canvas ICanvas)
    OnChangeRepaint(canvas ICanvas)
}

func (this *RepaintBehavior) RepaintBehavior_() *RepaintBehavior {
    return this
}

var RepaintBehavior_Instance IRepaintBehavior = NewRepaintBehavior()

func (this *RepaintBehavior) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.NAME = "RepaintBehavior"
}

func NewRepaintBehavior() *RepaintBehavior {
    this := &RepaintBehavior{}
    this.initFields()
    return this
}


func RepaintBehavior_GetInstance() IRepaintBehavior {
        return RepaintBehavior_Instance
}


func (this *RepaintBehavior) Repaint(canvas ICanvas) {
}


func (this *RepaintBehavior) OnChangeRepaint(canvas ICanvas) {
        var features *Features = Features_GetInstance(); _ = features
        var openGLFeatureFactory *OpenGLFeatureFactory = OpenGLFeatureFactory_GetInstance(); _ = openGLFeatureFactory
        if features.IsFeature(openGLFeatureFactory.OPENGL) {
            DisplayInfoSingleton_GetInstance().Process()
        } else {
            var thread IThread = NewThread(/* Java anonymous class body: implement it as a named type
                

            func (this *RepaintBehavior) Run() {
                    if tryControl1 := func() (tryControl1 int) {
                        defer func() {
                            if recovered := recover(); recovered != nil {
                                e := Java_ToError(recovered)
                                _ = e
                        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                        logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.RUN, e)
                            }
                        }()
                        canvas.Repaint()
                        DisplayInfoSingleton_GetInstance().Process()
                        return
                    }(); tryControl1 == 1 {
                        return
                    }
            }
*/ NewARunnable(), this.NAME); _ = thread
            thread.Start()
        }
}

