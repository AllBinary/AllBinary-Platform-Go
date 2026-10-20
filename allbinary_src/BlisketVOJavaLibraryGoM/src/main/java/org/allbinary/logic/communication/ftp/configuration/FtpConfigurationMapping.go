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

type FtpConfigurationMapping struct {
    FtpConfigurationInterface FtpConfigurationInterface
}

// Implements TableMappingInterface


func (this *FtpConfigurationMapping) initFields() {
    
}


func NewFtpConfigurationMapping(ftpConfigurationInterface FtpConfigurationInterface) *FtpConfigurationMapping {
    this := &FtpConfigurationMapping{}
    this.initFields()
        this.FtpConfigurationInterface = ftpConfigurationInterface
    return this
}


func (this *FtpConfigurationMapping) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(FtpConfigurationData_SERVER, this.FtpConfigurationInterface.GetServer())
        hashMap.Put(FtpConfigurationData_USERNAME, this.FtpConfigurationInterface.GetUserName())
        hashMap.Put(FtpConfigurationData_PASSWORD, this.FtpConfigurationInterface.GetPassword())
        hashMap.Put(FtpConfigurationData_PATH, this.FtpConfigurationInterface.GetPath())
        return hashMap
}


// throws Exception
func (this *FtpConfigurationMapping) GetKey() any {
        return FtpConfigurationData_NAME
}


// throws Exception
func (this *FtpConfigurationMapping) ToVector() IBasicArrayList {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        vector.Add(this.FtpConfigurationInterface.GetServer())
        vector.Add(this.FtpConfigurationInterface.GetUserName())
        vector.Add(this.FtpConfigurationInterface.GetPassword())
        vector.Add(this.FtpConfigurationInterface.GetPath())
        return vector
}

