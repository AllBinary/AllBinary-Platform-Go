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

type UserEmailEventConfiguration struct {
    LogUtil *LogUtil
    Name string
    EventListenerClassPath string
}

// Implements UserEmailEventConfigurationInterface


func (this *UserEmailEventConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}




// Java overloads of new UserEmailEventConfiguration
func NewUserEmailEventConfiguration(args ...any) *UserEmailEventConfiguration {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewUserEmailEventConfiguration_1()
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewUserEmailEventConfiguration_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new UserEmailEventConfiguration for the arguments")
}
func NewUserEmailEventConfiguration_1() *UserEmailEventConfiguration {
    this := &UserEmailEventConfiguration{}
    this.initFields()
        this.SetName(StringUtil_GetInstance().EMPTY_STRING)
        this.SetEventListenerClassPath(StringUtil_GetInstance().EMPTY_STRING)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF("New", this, " Constructor")
        }
    return this
}


func NewUserEmailEventConfiguration_2(hashMap *HashMap) *UserEmailEventConfiguration {
    this := &UserEmailEventConfiguration{}
    this.initFields()
        this.SetName(any(hashMap.Get(UserEmailEventConfigurationData_NAME)).(string))
        this.SetEventListenerClassPath(any(hashMap.Get(UserEmailEventConfigurationData_LISTENER_CLASSPATH)).(string))
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF("Log: " + fmt.Sprint(this.Log()), this, " Constructor")
        }
    return this
}


func (this *UserEmailEventConfiguration) GetName() string {
        return this.Name
}


func (this *UserEmailEventConfiguration) SetName(name string) {
        this.Name = name
}


func (this *UserEmailEventConfiguration) GetEventListenerClassPath() string {
        return this.EventListenerClassPath
}


func (this *UserEmailEventConfiguration) SetEventListenerClassPath(eventListenerClassPath string) {
        this.EventListenerClassPath = eventListenerClassPath
}


func (this *UserEmailEventConfiguration) Log() string {
        return "\nUser Email Event Log Info:" + "\nName: " + this.GetName() + "\nEventListenerClassPath: " + fmt.Sprint(this.GetEventListenerClassPath())
}

