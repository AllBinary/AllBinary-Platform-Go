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

type PointsLayer struct {
    AllBinaryLayer
    Points int
}


func (this *PointsLayer) initFields() {
    
}


func NewPointsLayer(points int) *PointsLayer {
    this := &PointsLayer{}
        this.AllBinaryLayer = *NewAllBinaryLayer(StringUtil_GetInstance().EMPTY_STRING, NewRectangle(PointFactory_GetInstance().ZERO_ZERO, 0, 0), ViewPosition_GetInstanceD())
        this.initFields()
        this.Points = points
    return this
}


func (this *PointsLayer) Paint(graphics IGraphics) {
}


func (this *PointsLayer) Visit(scoreableInterface ScoreableInterface) {
        scoreableInterface.AddPoints(this.Points)
}

