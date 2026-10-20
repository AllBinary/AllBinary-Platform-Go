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

type ContextConfigurationData struct {
    NAME string
    UNCRYPTED_EXTENSION string
    ENCRYPTED_EXTENSION string
}

var ContextConfigurationData_Instance *ContextConfigurationData = NewContextConfigurationData()

func (this *ContextConfigurationData) initFields() {
    this.NAME = "CONTEXT_CONFIGURATION_NAME"
    this.UNCRYPTED_EXTENSION = CommonDataFileStrings_GetInstance().UNCRYPTED_EXTENSION
    this.ENCRYPTED_EXTENSION = CommonDataFileStrings_GetInstance().ENCRYPTED_EXTENSION
}


func ContextConfigurationData_GetInstance() *ContextConfigurationData {
        return ContextConfigurationData_Instance
}


func NewContextConfigurationData() *ContextConfigurationData {
    this := &ContextConfigurationData{}
    this.initFields()
    return this
}

