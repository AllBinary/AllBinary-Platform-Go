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

type RandomRotationFactory struct {
    MyRandomFactory *MyRandomFactory
    AngleFactory *AngleFactory
}

var RandomRotationFactory_Instance *RandomRotationFactory = NewRandomRotationFactory()

func (this *RandomRotationFactory) initFields() {
    this.MyRandomFactory = MyRandomFactory_GetInstance()
    this.AngleFactory = AngleFactory_GetInstance()
}

func NewRandomRotationFactory() *RandomRotationFactory {
    this := &RandomRotationFactory{}
    this.initFields()
    return this
}


func RandomRotationFactory_GetInstance() *RandomRotationFactory {
        return RandomRotationFactory_Instance
}


// throws Exception
func (this *RandomRotationFactory) GetInstanceList(vector IBasicArrayList, howMuch int) IBasicArrayList {
        var graphicsPipe IBasicGraphicsPipeline = NewBasicGraphicsPipeline(vector); _ = graphicsPipe
        graphicsPipe.CreateMatrix()
        graphicsPipe.RotateToAngle(this.AngleFactory.GetAt(this.GetNextRandomAngle(howMuch)))
        return graphicsPipe.GetMatrix()
}


func (this *RandomRotationFactory) GetNextRandomAngle(howMuch int) int {
        return this.MyRandomFactory.GetNextInt(howMuch * 10)
}

