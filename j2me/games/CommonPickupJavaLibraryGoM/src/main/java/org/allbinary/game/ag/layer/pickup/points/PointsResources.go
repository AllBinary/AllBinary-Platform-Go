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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type PointsResources struct {
    RESOURCE string
    RESOURCE_2 string
    RESOURCE_3 string
    RESOURCE_4 string
}

var PointsResources_SINGLETON *PointsResources = NewPointsResources()

func (this *PointsResources) initFields() {
    this.RESOURCE = "points_drop_20_by_20.png"
    this.RESOURCE_2 = "points_2_drop_20_by_20.png"
    this.RESOURCE_3 = "points_3_drop_20_by_20.png"
    this.RESOURCE_4 = "points_4_drop_20_by_20.png"
}


func NewPointsResources() *PointsResources {
    this := &PointsResources{}
    this.initFields()
    return this
}


func PointsResources_GetInstance() *PointsResources {
        return PointsResources_SINGLETON
}

