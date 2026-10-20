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

type UserEmailEventHandlerSingletons struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    UserEmailEventHandlerHashMap *HashMap
}

var UserEmailEventHandlerSingletons_Instance *UserEmailEventHandlerSingletons = NewUserEmailEventHandlerSingletons()

func (this *UserEmailEventHandlerSingletons) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.UserEmailEventHandlerHashMap = StdUtil_GetInstance().CreateHashMap()
}


func UserEmailEventHandlerSingletons_GetInstance() *UserEmailEventHandlerSingletons {
        return UserEmailEventHandlerSingletons_Instance
}


func NewUserEmailEventHandlerSingletons() *UserEmailEventHandlerSingletons {
    this := &UserEmailEventHandlerSingletons{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


// throws Exception
func (this *UserEmailEventHandlerSingletons) GetInstance(abeClientInformation AbeClientInformationInterface, userEmailEventNameData *UserEmailEventNameData, userInterface UserInterface) *UserEmailEventHandler {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.GET_INSTANCE)
        }
        var userEmailEventHandler *UserEmailEventHandler = any(this.UserEmailEventHandlerHashMap.Get(userEmailEventNameData)).(*UserEmailEventHandler); _ = userEmailEventHandler
        if userEmailEventHandler == nil {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.PutF("Creating New Named UserEmailEventHandler", this, this.CommonStrings.GET_INSTANCE)
            }
            var newUserEmailEventHandler *UserEmailEventHandler = NewUserEmailEventHandler(); _ = newUserEmailEventHandler
            var vector IBasicArrayList = EmailEventHandlerUtil_GetUserEmailEventListenerVector(abeClientInformation, userEmailEventNameData, userInterface); _ = vector
            newUserEmailEventHandler.AddListener(vector)
            newUserEmailEventHandler.AddListener(NewLogUserEmailEventListenerModule())
            this.UserEmailEventHandlerHashMap.Put(userEmailEventNameData, newUserEmailEventHandler)
            return newUserEmailEventHandler
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
                this.LogUtil.PutF("Returning existing UserEmailEventHandler", this, this.CommonStrings.GET_INSTANCE)
            }
            return userEmailEventHandler
        }
}

