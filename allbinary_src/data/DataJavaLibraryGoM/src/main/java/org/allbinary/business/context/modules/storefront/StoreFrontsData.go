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

type StoreFrontsData struct {
    TYPE string
    RETAIL string
    SUBSCRIPTION string
    AUCTION string
    NEWS string
    BLOG string
    FORUM string
}

var StoreFrontsData_Instance *StoreFrontsData = NewStoreFrontsData()

func (this *StoreFrontsData) initFields() {
    this.TYPE = "STOREFRONTS_TYPE"
    this.RETAIL = "STOREFRONTS_RETAIL"
    this.SUBSCRIPTION = "STOREFRONTS_SUBSCRIPTION"
    this.AUCTION = "STOREFRONTS_AUCTION"
    this.NEWS = "STOREFRONTS_NEWS"
    this.BLOG = "STOREFRONTS_BLOG"
    this.FORUM = "STOREFRONTS_FORUM"
}

func NewStoreFrontsData() *StoreFrontsData {
    this := &StoreFrontsData{}
    this.initFields()
    return this
}


func StoreFrontsData_GetInstance() *StoreFrontsData {
        return StoreFrontsData_Instance
}

