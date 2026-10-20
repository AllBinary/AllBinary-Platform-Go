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

type ContextConfiguration struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    EmailServerConfigurationInterface EmailServerConfigurationInterface
}

// Implements ContextConfigurationInterface


func (this *ContextConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}




// Java overloads of new ContextConfiguration
func NewContextConfiguration(args ...any) *ContextConfiguration {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewContextConfiguration_1()
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewContextConfiguration_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new ContextConfiguration for the arguments")
}
func NewContextConfiguration_1() *ContextConfiguration {
    this := &ContextConfiguration{}
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, "ContextConfiguration")
        this.SetEmailServerConfigurationInterface(any(NewEmailServerConfiguration()).(EmailServerConfigurationInterface))
    return this
}


func NewContextConfiguration_2(hashMap *HashMap) *ContextConfiguration {
    this := &ContextConfiguration{}
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, "ContextConfiguration(HashMap)")
        this.SetEmailServerConfigurationInterface(any(NewEmailServerConfiguration(hashMap)).(EmailServerConfigurationInterface))
    return this
}


func (this *ContextConfiguration) GetEmailServerConfigurationInterface() EmailServerConfigurationInterface {
        return this.EmailServerConfigurationInterface
}


func (this *ContextConfiguration) SetEmailServerConfigurationInterface(emailServerConfigurationInterface EmailServerConfigurationInterface) {
        this.EmailServerConfigurationInterface = emailServerConfigurationInterface
}

