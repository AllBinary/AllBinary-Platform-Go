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

type JavaHelpUtil struct {
    LogUtil *LogUtil
    Point *Point
    Dimension *Dimension
    ContextSensitiveHelpActionEvent *ActionEvent
    CommonStrings *CommonStrings
}

var JavaHelpUtil_Instance *JavaHelpUtil = NewJavaHelpUtil()

func (this *JavaHelpUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Point = NewPoint(0, 0)
    this.Dimension = NewDimension(640, 480)
    this.ContextSensitiveHelpActionEvent = NewActionEvent(NewJLabel(), ActionEvent_ACTION_FIRST, nil)
    this.CommonStrings = CommonStrings_GetInstance()
}


func JavaHelpUtil_GetInstance() *JavaHelpUtil {
        return JavaHelpUtil_Instance
}


func NewJavaHelpUtil() *JavaHelpUtil {
    this := &JavaHelpUtil{}
    this.initFields()
    return this
}


func (this *JavaHelpUtil) GetHelpSet(filePath string) *HelpSet {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *HelpSet) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, JavaHelpUtil_GetInstance(), "set", e)
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var helpSetURI *URI = NewFile(filePath).ToURI(); _ = helpSetURI
            var helpSetURL *URL = helpSetURI.ToURL(); _ = helpSetURL
            tryControl1, tryResult1 = 1, this.GetHelpSet(helpSetURL); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// Java overloaded method: Go needs a different name
func (this *JavaHelpUtil) GetHelpSet(url *URL) *HelpSet {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *HelpSet) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, JavaHelpUtil_GetInstance(), "set", e)
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, NewHelpSet(nil, url); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *JavaHelpUtil) Show(helpSet *HelpSet) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, JavaHelpUtil_GetInstance(), "show", e)
                }
            }()
            var helpBroker *HelpBroker = helpSet.CreateHelpBroker(); _ = helpBroker
            helpBroker.SetLocation(this.Point)
            helpBroker.SetSize(this.Dimension)
            NewDisplayHelpFromSource(helpBroker).ActionPerformed(this.ContextSensitiveHelpActionEvent)
            this.LogUtil.PutF("CSH Action", JavaHelpUtil_GetInstance(), "show")
            return
        }(); tryControl1 == 1 {
            return
        }
}


func JavaHelpUtil_Main(args []string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
                }
            }()
            var javaHelpUtil *JavaHelpUtil = JavaHelpUtil_GetInstance(); _ = javaHelpUtil
            javaHelpUtil.Show(javaHelpUtil.GetHelpSet("G:/mnt/bc/mydev/working/automation/InputAutomationJavaApplication/AllBinaryInputAutomationHelp/AllBinaryInputAutomation.hs"))
            return
        }(); tryControl1 == 1 {
            return
        }
}

