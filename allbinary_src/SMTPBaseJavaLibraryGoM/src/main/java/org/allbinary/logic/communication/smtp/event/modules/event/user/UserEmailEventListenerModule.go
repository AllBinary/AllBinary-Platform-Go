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

type UserEmailEventListenerModule struct {
    LogUtil *LogUtil
    UserInterface UserInterface
}

// Implements UserEmailEventListenerInterface


func (this *UserEmailEventListenerModule) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewUserEmailEventListenerModule(userInterface UserInterface) *UserEmailEventListenerModule {
    this := &UserEmailEventListenerModule{}
    this.initFields()
        this.UserInterface = userInterface
    return this
}


// throws Exception
func (this *UserEmailEventListenerModule) OnEmailSendFailure(emailEvent *EmailEvent) {
        var to string = this.UserInterface.GetSecondaryEmail(); _ = to
        var email *BasicEmail = NewBasicEmail(emailEvent.GetEmailInfo(), to); _ = email
        EmailQueueFactory_GetInstance().Offer(any(email.GetEmail()).(EmailInterface))
}


// throws Exception
func (this *UserEmailEventListenerModule) OnEmailSendRequest(emailEvent *EmailEvent) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF("Add Email To Que For Sending: " + fmt.Sprint(emailEvent), this, "onEmailSendRequest")
        }
        var to string = this.UserInterface.GetMainEmail(); _ = to
        var email *BasicEmail = NewBasicEmail(emailEvent.GetEmailInfo(), to); _ = email
        EmailQueueFactory_GetInstance().Offer(any(email.GetEmail()).(EmailInterface))
}

