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

type BasicEmailQueue struct {
    BasicQueue
    LogUtil *LogUtil
}

// Implements EmailQueueInterface

var BasicEmailQueue_Instance *BasicEmailQueue = NewBasicEmailQueue()
var BasicEmailQueue_BasicEmailQueue *BasicEmailQueue

func (this *BasicEmailQueue) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewBasicEmailQueue() *BasicEmailQueue {
    this := &BasicEmailQueue{}
    this.BasicQueue = *NewBasicQueue()
    this.initFields()
    return this
}


// synchronized
func BasicEmailQueue_GetInstance() *BasicEmailQueue {
        if BasicEmailQueue_BasicEmailQueue == nil {
            BasicEmailQueue_BasicEmailQueue = NewBasicEmailQueue()
        }
        return BasicEmailQueue_BasicEmailQueue
}


// throws Exception
// synchronized
func (this *BasicEmailQueue) Offer(emailInterface EmailInterface) bool {
        this.Offer(any(emailInterface).(any))
        this.Post(emailInterface)
        return true
}


// synchronized
func (this *BasicEmailQueue) Remove(emailInterface EmailInterface) {
        this.Remove(any(emailInterface).(any))
}


// Java overloaded method: Go needs a different name
// synchronized
func (this *BasicEmailQueue) Remove() EmailInterface {
        return nil
}


// throws Exception
// synchronized
func (this *BasicEmailQueue) Post(emailInterface EmailInterface) {
        if this.Send(emailInterface) {
            this.Remove(emailInterface)
            this.ProcessAllUnsent()
        }
}


// throws Exception
// synchronized
func (this *BasicEmailQueue) ProcessAllUnsent() {
        var size int = this.QueueVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var emailInterface EmailInterface = any(this.QueueVector.Get(index)).(EmailInterface); _ = emailInterface
            if this.Send(emailInterface) {
                this.Remove(emailInterface)
                this.ProcessAllUnsent()
            }
        }
}


// synchronized
func (this *BasicEmailQueue) Send(emailInterface EmailInterface) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.Put("Failed Email Send Debug: " + fmt.Sprint(emailInterface.GetDebugInfo()), this, "send", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.PutF("Sending: " + fmt.Sprint(emailInterface.Log()), this, "send")
            }
            Transport_Send(emailInterface.GetMimeMessage())
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.PutF("Email Send Debug: " + fmt.Sprint(emailInterface.GetDebugInfo()), this, "send")
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

