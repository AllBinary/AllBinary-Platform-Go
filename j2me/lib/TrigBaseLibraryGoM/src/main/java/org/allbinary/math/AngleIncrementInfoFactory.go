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

type AngleIncrementInfoFactory struct {
    AngleIncrementInfo []*AngleIncrementInfo
}

var AngleIncrementInfoFactory_Instance *AngleIncrementInfoFactory = NewAngleIncrementInfoFactory()

func (this *AngleIncrementInfoFactory) initFields() {
    this.AngleIncrementInfo = make([]*AngleIncrementInfo, int(AngleFactory_GetInstance().TOTAL_ANGLE))
}

func NewAngleIncrementInfoFactory() *AngleIncrementInfoFactory {
    this := &AngleIncrementInfoFactory{}
    this.initFields()
    return this
}


func AngleIncrementInfoFactory_GetInstance() *AngleIncrementInfoFactory {
        return AngleIncrementInfoFactory_Instance
}


func (this *AngleIncrementInfoFactory) GetAt(angleIncrement int16) *AngleIncrementInfo {
        var halfAngleIncrement int = (int(angleIncrement) >> 1); _ = halfAngleIncrement
        var angleIncrementInfo *AngleIncrementInfo; _ = angleIncrementInfo
        var angleIncrementInfoCanBeNull *AngleIncrementInfo = this.AngleIncrementInfo[halfAngleIncrement]; _ = angleIncrementInfoCanBeNull
        if angleIncrementInfoCanBeNull == nil {
            angleIncrementInfo = NewAngleIncrementInfo(angleIncrement)
            this.AngleIncrementInfo[halfAngleIncrement] = angleIncrementInfo
            return angleIncrementInfo
        } else {
            angleIncrementInfo = angleIncrementInfoCanBeNull
        }
        return angleIncrementInfo
}
