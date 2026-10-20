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

type PtsDamageFloaters struct {
    DamageFloaters
    BasicSetColorUtil IBasicColorSetUtil
    LayerInterface IAllBinaryLayer
    LastDamage []int
    LastDamageString [][]rune
    LastDamageStringSizeArray []int
    CircularIndexUtil *CircularIndexUtil
    PrimitiveLongUtil *PrimitiveLongUtil
    MyRandomFactory *MyRandomFactory
}

var PtsDamageFloaters_BasicColorArrayIndexer *BasicColorArrayIndexer = NewBasicColorArrayIndexer([]*BasicColor{BasicColorFactory_GetInstance().WHITE, BasicColorFactory_GetInstance().RED, BasicColorFactory_GetInstance().YELLOW, BasicColorFactory_GetInstance().BLUE})

func (this *PtsDamageFloaters) initFields() {
    this.BasicSetColorUtil = BasicColorSetUtil_GetInstance()
    
    this.LastDamage = make([]int, 5)
    this.LastDamageString = func() [][]rune { array := make([][]rune, 5); for index := range array { array[index] = make([]rune, 0) }; return array }()
    this.LastDamageStringSizeArray = make([]int, 5)
    this.CircularIndexUtil = CircularIndexUtil_CreateInstance(5)
    
    this.MyRandomFactory = MyRandomFactory_GetInstance()
}


func NewPtsDamageFloaters(layerInterface IAllBinaryLayer) *PtsDamageFloaters {
    this := &PtsDamageFloaters{}
    this.DamageFloaters = *NewDamageFloaters()
    this.initFields()
        this.PrimitiveLongUtil = PrimitiveLongUtil_CreatePowerOfTen(100000)
        this.LayerInterface = layerInterface
        var CHAR_ARRAY []rune = CharArrayFactory_GetInstance().GetZeroCharArray(); _ = CHAR_ARRAY
        for index := int(len(this.LastDamageString) - 1); index >= 0; index-- {
            this.LastDamageString[index] = CHAR_ARRAY
        }
    return this
}


func (this *PtsDamageFloaters) Add(damage int) {
        var index int = this.CircularIndexUtil.GetIndex(); _ = index
        this.LastDamage[index] = damage
        this.LastDamageString[index] = this.PrimitiveLongUtil.GetCharArray(this.LastDamage[index])
        this.LastDamageStringSizeArray[index] = this.PrimitiveLongUtil.GetCurrentTotalDigits()
        this.CircularIndexUtil.Next()
}


func (this *PtsDamageFloaters) Paint(graphics IGraphics) {
        var viewPosition IViewPositionBase = this.LayerInterface.GetViewPosition(); _ = viewPosition
        var x int = viewPosition.GetX(); _ = x
        var y int = viewPosition.GetY(); _ = y
        for index := int(len(this.LastDamage) - 1); index >= 0; index-- {
            if this.LastDamage[index] != 0 {
                var delta int = 10 * (index + 2); _ = delta
                this.BasicSetColorUtil.SetBasicColorP(graphics, PtsDamageFloaters_BasicColorArrayIndexer.Get())
                graphics.DrawChars(this.LastDamageString[index], 0, this.LastDamageStringSizeArray[index], x - this.MyRandomFactory.GetNextInt(delta), y - delta, 0)
                PtsDamageFloaters_BasicColorArrayIndexer.Next()
            }
            this.LastDamage[index] = 0
        }
}

