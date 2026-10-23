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

type UserEmailEventConfigurationMapping struct {
    UserEmailEventConfigurationInterface UserEmailEventConfigurationInterface
}

// Implements TableMappingInterface


func (this *UserEmailEventConfigurationMapping) initFields() {
    
}


func NewUserEmailEventConfigurationMapping(userEmailEventConfigurationInterface UserEmailEventConfigurationInterface) *UserEmailEventConfigurationMapping {
    this := &UserEmailEventConfigurationMapping{}
    this.initFields()
        this.UserEmailEventConfigurationInterface = userEmailEventConfigurationInterface
    return this
}


func (this *UserEmailEventConfigurationMapping) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(UserEmailEventConfigurationData_NAME, this.UserEmailEventConfigurationInterface.GetName())
        hashMap.Put(UserEmailEventConfigurationData_LISTENER_CLASSPATH, this.UserEmailEventConfigurationInterface.GetEventListenerClassPath())
        return hashMap
}


// throws Exception
func (this *UserEmailEventConfigurationMapping) GetKey() any {
        return UserEmailEventConfigurationData_NAME
}


// throws Exception
func (this *UserEmailEventConfigurationMapping) ToVector() IBasicArrayList {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        vector.Add(this.UserEmailEventConfigurationInterface.GetName())
        vector.Add(this.UserEmailEventConfigurationInterface.GetEventListenerClassPath())
        return vector
}

