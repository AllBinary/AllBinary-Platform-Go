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

type UserEmailConfiguration struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    UserEmailEventsConfigurationInterface UserEmailEventsConfigurationInterface
}

// Implements UserEmailConfigurationInterface


func (this *UserEmailConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewUserEmailConfiguration() *UserEmailConfiguration {
    this := &UserEmailConfiguration{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, this.CommonStrings.CONSTRUCTOR)
        }
        this.SetUserEmailEventsConfigurationInterface(NewUserEmailEventsConfiguration())
    return this
}


// Java overloaded constructor: Go needs a different name
func NewUserEmailConfiguration(hashMap *HashMap) *UserEmailConfiguration {
    this := &UserEmailConfiguration{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, this.CommonStrings.CONSTRUCTOR)
        }
        this.SetUserEmailEventsConfigurationInterface(NewUserEmailEventsConfiguration(hashMap))
    return this
}


func (this *UserEmailConfiguration) GetUserEmailEventsConfigurationInterface() UserEmailEventsConfigurationInterface {
        return this.UserEmailEventsConfigurationInterface
}


func (this *UserEmailConfiguration) SetUserEmailEventsConfigurationInterface(userEmailEventsConfigurationInterface UserEmailEventsConfigurationInterface) {
        this.UserEmailEventsConfigurationInterface = userEmailEventsConfigurationInterface
}

