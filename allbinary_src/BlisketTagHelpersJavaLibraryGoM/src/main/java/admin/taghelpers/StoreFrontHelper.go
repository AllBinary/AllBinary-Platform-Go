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

type StoreFrontHelper struct {
    TagHelper
    StoreFrontInterface StoreFrontInterface
    CurrentLocation string
    CurrentHomeLocation string
}


func (this *StoreFrontHelper) initFields() {
    
    
    
}


func NewStoreFrontHelper(hashMap *HashMap, pageContext *PageContext) *StoreFrontHelper {
    this := &StoreFrontHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        var storeName string = any(hashMap.Get(StoreFrontData_GetInstance().NAME)).(string); _ = storeName
        if storeName != "" {
            this.StoreFrontInterface = StoreFrontFactory_GetInstance(storeName)
            this.CurrentLocation = NewReplace(CommonSeps_GetInstance().SPACE, "%20").All(this.StoreFrontInterface.GetCurrentHostName() + this.StoreFrontInterface.GetCurrentHostNamePath())
            this.CurrentHomeLocation = NewReplace(CommonSeps_GetInstance().SPACE, "%20").All(this.StoreFrontInterface.GetCurrentHomeHostName() + this.StoreFrontInterface.GetCurrentHomeHostNamePath())
        } else {
            var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
            this.StoreFrontInterface = NewStoreFront()
            this.CurrentLocation = stringUtil.EMPTY_STRING
            this.CurrentHomeLocation = stringUtil.EMPTY_STRING
        }
    return this
}


func (this *StoreFrontHelper) GetCurrentLocation() string {
        return this.CurrentLocation
}


func (this *StoreFrontHelper) GetCurrentHomeLocation() string {
        return this.CurrentHomeLocation
}

