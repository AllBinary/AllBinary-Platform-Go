/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2007 AllBinary 
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

type CacheType struct {
    Vendor string
    Type string
}


func (this *CacheType) initFields() {
    
    
}


func NewCacheType(vendor string, type_ string) *CacheType {
    this := &CacheType{}
    this.initFields()
        this.Vendor = vendor
        this.Type = type_
    return this
}


func (this *CacheType) GetVendor() string {
        return this.Vendor
}


func (this *CacheType) GetType() string {
        return this.Type
}


func (this *CacheType) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("CacheType: Vendor: ")
        stringBuffer.Append(this.GetVendor())
        stringBuffer.Append(" Type: ")
        stringBuffer.Append(this.GetType())
        return fmt.Sprint(stringBuffer)
}

func (this *CacheType) String() string {
    return this.ToString()
}

