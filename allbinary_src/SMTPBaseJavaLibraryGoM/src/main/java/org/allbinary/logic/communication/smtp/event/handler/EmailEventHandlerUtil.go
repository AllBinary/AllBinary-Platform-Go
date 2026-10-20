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

type EmailEventHandlerUtil struct {
    LogUtil *LogUtil
}

var EmailEventHandlerUtil_Instance *EmailEventHandlerUtil = NewEmailEventHandlerUtil()

func (this *EmailEventHandlerUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func EmailEventHandlerUtil_GetInstance() *EmailEventHandlerUtil {
        return EmailEventHandlerUtil_Instance
}


func NewEmailEventHandlerUtil() *EmailEventHandlerUtil {
    this := &EmailEventHandlerUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *EmailEventHandlerUtil) GetEventHandler(abeClientInformation AbeClientInformationInterface, userEmailEventNameData *UserEmailEventNameData, userVector *BasicArrayList) *UserEmailEventHandler {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Returning UserEmailEventHandler for ")
            stringBuffer.Appendint(userVector.Size())
            stringBuffer.Append(" users.")
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "getEventHandler")
        }
        var userEmailEventHandler *UserEmailEventHandler = NewUserEmailEventHandler(); _ = userEmailEventHandler
        userEmailEventHandler.AddListener(NewLogUserEmailEventListenerModule())
        var size int = userVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var userInterface UserInterface = any(userVector.Get(index)).(UserInterface); _ = userInterface
            var vector *BasicArrayList = EmailEventHandlerUtil_GetUserEmailEventListenerVector(abeClientInformation, userEmailEventNameData, userInterface); _ = vector
            userEmailEventHandler.AddListener(vector)
        }
        return userEmailEventHandler
}


// throws Exception
func EmailEventHandlerUtil_GetUserEmailEventListenerVector(abeClientInformation AbeClientInformationInterface, userEmailEventNameData *UserEmailEventNameData, userInterface UserInterface) *BasicArrayList {
        var userConfigurationInterface UserConfigurationInterface = userInterface.GetUserConfigurationInterface(); _ = userConfigurationInterface
        var userEmailConfigurationInterface UserEmailConfigurationInterface = userConfigurationInterface.GetUserEmailConfigurationInterface(); _ = userEmailConfigurationInterface
        var userEmailEventsConfigurationInterface UserEmailEventsConfigurationInterface = userEmailConfigurationInterface.GetUserEmailEventsConfigurationInterface(); _ = userEmailEventsConfigurationInterface
        var userEmailEventListenerInterface UserEmailEventListenerInterface = userEmailEventsConfigurationInterface.GetEventListener(abeClientInformation, userEmailEventNameData, userInterface); _ = userEmailEventListenerInterface
        var vector *BasicArrayList = NewBasicArrayListD(); _ = vector
        vector.Add(userEmailEventListenerInterface)
        return vector
}

