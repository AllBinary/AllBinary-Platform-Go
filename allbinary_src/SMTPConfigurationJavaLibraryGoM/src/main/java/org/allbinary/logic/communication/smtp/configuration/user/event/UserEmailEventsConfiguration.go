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

type UserEmailEventsConfiguration struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    EmailEventHashMap *HashMap
}

// Implements UserEmailEventsConfigurationInterface


func (this *UserEmailEventsConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewUserEmailEventsConfiguration() *UserEmailEventsConfiguration {
    this := &UserEmailEventsConfiguration{}
    this.initFields()
        this.Init()
    return this
}


// Java overloaded constructor: Go needs a different name
func NewUserEmailEventsConfiguration(hashMap *HashMap) *UserEmailEventsConfiguration {
    this := &UserEmailEventsConfiguration{}
    this.initFields()
        this.Init()
    return this
}


func (this *UserEmailEventsConfiguration) Init() {
        this.EmailEventHashMap = StdUtil_GetInstance().CreateHashMap()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, " Constructor")
        }
}


func (this *UserEmailEventsConfiguration) GetEventConfigurationHashMap() *HashMap {
        return this.EmailEventHashMap
}


func (this *UserEmailEventsConfiguration) AddUserEmailEventConfiguration(userEmailEventConfigurationInterface UserEmailEventConfigurationInterface) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF("Adding: " + fmt.Sprint(userEmailEventConfigurationInterface.Log()), this, "addUserEmailEventConfiguration")
        }
        this.EmailEventHashMap.Put(userEmailEventConfigurationInterface.GetName(), userEmailEventConfigurationInterface)
}


// throws Exception
func (this *UserEmailEventsConfiguration) GetEventListener(abeClientInformation AbeClientInformationInterface, userEmailEventNameData *UserEmailEventNameData, userInterface UserInterface) UserEmailEventListenerInterface {
        var userEmailEventConfigurationInterface UserEmailEventConfigurationInterface = any(this.EmailEventHashMap.Get(fmt.Sprint(userEmailEventNameData))).(UserEmailEventConfigurationInterface); _ = userEmailEventConfigurationInterface
        return UserEmailEventListenerFactory_GetInstance(abeClientInformation, userEmailEventConfigurationInterface, userInterface)
}

