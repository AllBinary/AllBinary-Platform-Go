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

type ImageTool struct {
    LogUtil *LogUtil
}


func (this *ImageTool) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewImageTool() *ImageTool {
    this := &ImageTool{}
    this.initFields()
    return this
}


func ImageTool_Main(args []string) {
        java_Awt.EventQueue.InvokeLater(/* Java anonymous class body: implement it as a named type
            

        func (this *ImageTool) Run() {
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
                    logUtil.PutF(commonStrings.START, this, commonStrings.RUN)
                    NewImageToolJFrame().SetVisible(true)
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewARunnable())
}

