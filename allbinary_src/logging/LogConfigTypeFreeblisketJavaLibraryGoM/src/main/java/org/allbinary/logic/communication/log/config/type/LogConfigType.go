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

type LogConfigType struct {
    LogUtil *LogUtil
    Name string
    Description string
}

var LogConfigType_AvailableLogConfigTypes *BasicArrayList = NewBasicArrayListD()

func (this *LogConfigType) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewLogConfigType(name string, description string) *LogConfigType {
    this := &LogConfigType{}
    this.initFields()
        this.Name = name
        this.Description = description
        LogConfigType_AvailableLogConfigTypes.Add(this)
    return this
}


func (this *LogConfigType) GetName() string {
        return this.Name
}


func (this *LogConfigType) GetDescription() string {
        return this.Description
}


func (this *LogConfigType) SetName(value string) {
        this.Name = value
}


func (this *LogConfigType) SetDescription(value string) {
        this.Description = value
}

