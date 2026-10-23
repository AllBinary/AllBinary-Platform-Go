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

type UserConfiguration struct {
    UserEmailConfigurationInterface UserEmailConfigurationInterface
    TimeCreated *TimeCreated
    TimeLastModified *TimeLastModified
}

// Implements UserConfigurationInterface


func (this *UserConfiguration) initFields() {
    
    
    
}




// Java overloads of new UserConfiguration
func NewUserConfiguration(args ...any) *UserConfiguration {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewUserConfiguration_1()
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewUserConfiguration_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new UserConfiguration for the arguments")
}
func NewUserConfiguration_1() *UserConfiguration {
    this := &UserConfiguration{}
    this.initFields()
        this.UserEmailConfigurationInterface = any(NewUserEmailConfiguration()).(UserEmailConfigurationInterface)
        this.TimeCreated = NewTimeCreated(0)
        this.TimeLastModified = NewTimeLastModified(0)
    return this
}


func NewUserConfiguration_2(hashMap *HashMap) *UserConfiguration {
    this := &UserConfiguration{}
    this.initFields()
        this.UserEmailConfigurationInterface = any(NewUserEmailConfiguration(hashMap)).(UserEmailConfigurationInterface)
        this.TimeCreated = NewTimeCreated(any(hashMap.Get(EntryData_GetInstance().TIMECREATED)).(string))
        this.TimeLastModified = NewTimeLastModified(any(hashMap.Get(EntryData_GetInstance().LASTMODIFIED)).(string))
    return this
}


func (this *UserConfiguration) GetUserEmailConfigurationInterface() UserEmailConfigurationInterface {
        return this.UserEmailConfigurationInterface
}


func (this *UserConfiguration) SetUserEmailConfigurationInterface(userEmailConfigurationInterface UserEmailConfigurationInterface) {
        this.UserEmailConfigurationInterface = userEmailConfigurationInterface
}


func (this *UserConfiguration) GetTimeCreated() *TimeCreated {
        return this.TimeCreated
}


func (this *UserConfiguration) GetTimeLastModified() *TimeLastModified {
        return this.TimeLastModified
}


func (this *UserConfiguration) SetTimeCreated(timeCreated *TimeCreated) {
        this.TimeCreated = timeCreated
}


func (this *UserConfiguration) SetTimeLastModified(timeLastModified *TimeLastModified) {
        this.TimeLastModified = timeLastModified
}

