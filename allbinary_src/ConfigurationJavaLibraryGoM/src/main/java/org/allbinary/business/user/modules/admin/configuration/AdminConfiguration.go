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

type AdminConfiguration struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ContextConfigurationInterface ContextConfigurationInterface
}

// Implements AdminConfigurationInterface


func (this *AdminConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}




// Java overloads of new AdminConfiguration
func NewAdminConfiguration(args ...any) *AdminConfiguration {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewAdminConfiguration_1()
        }
        if len(args) == 1 && Java_IsArg[*HttpServletRequest](args[0], pass) {
            return NewAdminConfiguration_2(Java_Arg[*HttpServletRequest](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewAdminConfiguration_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new AdminConfiguration for the arguments")
}
func NewAdminConfiguration_1() *AdminConfiguration {
    this := &AdminConfiguration{}
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.ContextConfigurationInterface = ContextConfigurationInterfaceFactory_GetInstance().GetInstance(AdminConfigurationData_CONTEXTNAME)
    return this
}


func NewAdminConfiguration_2(request *HttpServletRequest) *AdminConfiguration {
    this := &AdminConfiguration{}
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.GetFormData(NewRequestParams(request).ToHashMap())
    return this
}


func NewAdminConfiguration_3(storeHashMap *HashMap) *AdminConfiguration {
    this := &AdminConfiguration{}
    this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.GetFormData(storeHashMap)
    return this
}


// throws Exception
func (this *AdminConfiguration) GetFormData(storeHashMap *HashMap) {
        this.LogUtil.PutF(this.CommonStrings.START, this, "getFormData")
        this.SetContextConfigurationInterface(any(NewContextConfiguration(storeHashMap)).(ContextConfigurationInterface))
}


// throws Exception
// synchronized
func (this *AdminConfiguration) Write() {
        var contextConfigurationDomDocumentMapping *ContextConfigurationDomDocumentMapping = NewContextConfigurationDomDocumentMapping(this.GetContextConfigurationInterface()); _ = contextConfigurationDomDocumentMapping
        var cryptFileWriter *CryptFileWriter = NewCryptFileWriter(ContextConfigurationData_GetInstance().UNCRYPTED_EXTENSION, ContextConfigurationData_GetInstance().ENCRYPTED_EXTENSION); _ = cryptFileWriter
        cryptFileWriter.Write(ContextConfigurationPathUtil_GetAbPath(AdminConfigurationData_CONTEXTNAME), contextConfigurationDomDocumentMapping.ToXmlDoc())
}


func (this *AdminConfiguration) GetContextConfigurationInterface() ContextConfigurationInterface {
        return this.ContextConfigurationInterface
}


func (this *AdminConfiguration) SetContextConfigurationInterface(contextConfigurationInterface ContextConfigurationInterface) {
        this.ContextConfigurationInterface = contextConfigurationInterface
}

