/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type LogConfigTypeFactory struct {
    NO_DESCRIPTION string
    OS *LogConfigType
    FACTORYERROR *LogConfigType
}

var LogConfigTypeFactory_Instance *LogConfigTypeFactory = NewLogConfigTypeFactory()

func (this *LogConfigTypeFactory) initFields() {
    this.NO_DESCRIPTION = "No Description"
    this.OS = NewLogConfigType("Operating System", this.NO_DESCRIPTION)
    this.FACTORYERROR = NewLogConfigType("Factory Error", this.NO_DESCRIPTION)
}


func LogConfigTypeFactory_GetInstance() *LogConfigTypeFactory {
        return LogConfigTypeFactory_Instance
}


func NewLogConfigTypeFactory() *LogConfigTypeFactory {
    this := &LogConfigTypeFactory{}
    this.initFields()
        LogConfigTypes_LOGGING.Add(this.OS)
        LogConfigTypes_LOGGING.Add(this.FACTORYERROR)
    return this
}

