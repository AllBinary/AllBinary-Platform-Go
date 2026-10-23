/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

//ArkTs does not have this.  So this is for the ArkTs build.
type EnumerationUtil struct {
}

var EnumerationUtil_Instance *EnumerationUtil = NewEnumerationUtil()

func NewEnumerationUtil() *EnumerationUtil {
    this := &EnumerationUtil{}
    return this
}


func EnumerationUtil_GetInstance() *EnumerationUtil {
        return EnumerationUtil_Instance
}


func (this *EnumerationUtil) HasMoreElements(enumeration Enumeration) bool {
        return enumeration.HasMoreElements()
}


func (this *EnumerationUtil) NextElement(enumeration Enumeration) any {
        return enumeration.NextElement()
}


func (this *EnumerationUtil) GetAsArray(enumeration Enumeration) []any {
        var basicArrayList IBasicArrayList = NewBasicArrayListD(); _ = basicArrayList
        for this.HasMoreElements(enumeration) {
            basicArrayList.Add(this.NextElement(enumeration))
        }
        return basicArrayList.ToArray()
}

