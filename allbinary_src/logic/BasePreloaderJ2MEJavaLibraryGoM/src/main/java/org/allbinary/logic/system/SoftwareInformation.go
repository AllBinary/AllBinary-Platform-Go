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

type SoftwareInformation struct {
    Name string
    ShortName string
    Version string
}


func (this *SoftwareInformation) initFields() {
    
    
    
}


func NewSoftwareInformation(name string, version string, shortName string) *SoftwareInformation {
    this := &SoftwareInformation{}
    this.initFields()
        this.Name = name
        this.Version = version
        this.ShortName = fmt.Sprint(NewStringMaker().Append(shortName).Append(KotlinUtil_GetShortName()))
    return this
}


func (this *SoftwareInformation) GetName() string {
        return this.Name
}


func (this *SoftwareInformation) GetVersion() string {
        return this.Version
}


func (this *SoftwareInformation) ToShortString() string {
        return this.ShortName
}


func (this *SoftwareInformation) ToString() string {
        var NAME string = "Name: "; _ = NAME
        var VERSION string = "Version: "; _ = VERSION
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(NAME)
        stringBuffer.Append(this.GetName())
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(VERSION)
        stringBuffer.Append(this.GetVersion())
        return fmt.Sprint(stringBuffer)
}

func (this *SoftwareInformation) String() string {
    return this.ToString()
}

