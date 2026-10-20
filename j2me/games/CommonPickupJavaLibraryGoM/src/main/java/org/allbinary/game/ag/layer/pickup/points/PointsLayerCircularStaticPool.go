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

type PointsLayerCircularStaticPool struct {
    AllBinaryLayerCircularPool
}



func NewPointsLayerCircularStaticPool(allbinaryLayerFactoryInterface AllBinaryLayerFactoryInterface, total int) *PointsLayerCircularStaticPool {
    this := &PointsLayerCircularStaticPool{}
    this.AllBinaryLayerCircularPool = *NewAllBinaryLayerCircularPool()
        this.AllBinaryLayerCircularPool.InitAllBinaryLayerCircularPool(allbinaryLayerFactoryInterface, total)
    return this
}


// throws Exception
func (this *PointsLayerCircularStaticPool) Visit(sourceLayerInterface IAllBinaryLayer) {
        var pointsLayer *PointsLayer = any(this.GetNextInstance()).(*PointsLayer); _ = pointsLayer
        var scoreableInterface ScoreableInterface = any(sourceLayerInterface).(ScoreableInterface); _ = scoreableInterface
        pointsLayer.Visit(scoreableInterface)
}

