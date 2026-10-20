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

type LogUserEmailEventListenerModule struct {
    LogUtil *LogUtil
}

// Implements UserEmailEventListenerInterface


func (this *LogUserEmailEventListenerModule) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewLogUserEmailEventListenerModule() *LogUserEmailEventListenerModule {
    this := &LogUserEmailEventListenerModule{}
    this.initFields()
    return this
}


// throws Exception
func (this *LogUserEmailEventListenerModule) OnEmailSendRequest(userEmailInfoEvent *EmailEvent) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            var message string = "Logging EmailEvent - Temporary Listener for debugging.\n\n" + fmt.Sprint(userEmailInfoEvent); _ = message
            this.LogUtil.PutF(message, this, "onEmailSendRequest")
        }
}

