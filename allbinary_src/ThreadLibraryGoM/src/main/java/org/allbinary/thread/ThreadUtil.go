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

type ThreadUtil struct {
    LogUtil *LogUtil
    JOIN string
}

var ThreadUtil_Instance *ThreadUtil = NewThreadUtil()

func (this *ThreadUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.JOIN = "join"
}

func NewThreadUtil() *ThreadUtil {
    this := &ThreadUtil{}
    this.initFields()
    return this
}


func ThreadUtil_GetInstance() *ThreadUtil {
        return ThreadUtil_Instance
}


func (this *ThreadUtil) IsRunning(thread IThread) bool {
        if thread != nil && thread.IsAlive() {
            return true
        } else {
            return false
        }
}


// throws Exception
func (this *ThreadUtil) Join(thread IThread) {
        if ThreadUtil_GetInstance().IsRunning(thread) {
            this.LogUtil.PutF("Waiting for Thread To Join/End", this, this.JOIN)
            thread.Join()
        }
}
