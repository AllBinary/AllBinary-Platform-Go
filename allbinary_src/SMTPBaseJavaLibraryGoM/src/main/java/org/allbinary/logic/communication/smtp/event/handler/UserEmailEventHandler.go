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

type UserEmailEventHandler struct {
    EmailVector *BasicArrayList
    EmailInfo *EmailInfo
    UserEmailEventNameData *UserEmailEventNameData
}


func (this *UserEmailEventHandler) initFields() {
    
    
    
}


func NewUserEmailEventHandler() *UserEmailEventHandler {
    this := &UserEmailEventHandler{}
    this.initFields()
        this.EmailVector = NewBasicArrayListD()
    return this
}


// throws Exception
// synchronized
func (this *UserEmailEventHandler) ReceiveEmailInfo(userEmailEventNameData *UserEmailEventNameData, emailInfo *EmailInfo) {
        this.UserEmailEventNameData = userEmailEventNameData
        this.EmailInfo = emailInfo
        this.FireEmailEvent()
}


// synchronized
func (this *UserEmailEventHandler) AddListener(vector *BasicArrayList) {
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var userEmailEventListenerInterface UserEmailEventListenerInterface = any(vector.Get(index)).(UserEmailEventListenerInterface); _ = userEmailEventListenerInterface
            this.AddListener(userEmailEventListenerInterface)
        }
}


// Java overloaded method: Go needs a different name
// synchronized
func (this *UserEmailEventHandler) AddListener(emailEventListenerInterface UserEmailEventListenerInterface) {
        this.EmailVector.Add(emailEventListenerInterface)
}


// synchronized
func (this *UserEmailEventHandler) RemoveListener(emailEventListenerInterface UserEmailEventListenerInterface) {
        this.EmailVector.Remove(emailEventListenerInterface)
}


// throws Exception
// synchronized
func (this *UserEmailEventHandler) FireEmailEvent() {
        var emailEvent *EmailEvent = NewEmailEvent(this, this.UserEmailEventNameData, this.EmailInfo, 0); _ = emailEvent
        var size int = this.EmailVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var emailEventListenerInterface UserEmailEventListenerInterface = any(this.EmailVector.Get(index)).(UserEmailEventListenerInterface); _ = emailEventListenerInterface
            emailEventListenerInterface.OnEmailSendRequest(emailEvent)
        }
}

