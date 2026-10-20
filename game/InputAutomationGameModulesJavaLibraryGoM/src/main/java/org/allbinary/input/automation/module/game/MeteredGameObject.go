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

type MeteredGameObject struct {
    Ratios []float64
    Good bool
    Dropping bool
    MinX int
    MaxX int
    Y int
    Size int
    RATIOS_LABEL string
    IS_GOOD_LABEL string
    IS_DROPPPING_LABEL string
}

// Implements MeteredGameObjectInterface

// The Java class MeteredGameObject has subclasses so its variables are this interface
type IMeteredGameObject interface {
    MeteredGameObject_() *MeteredGameObject
    IsGood() bool
    SetGood(good bool)
    IsDropping() bool
    SetDropping(dropping bool)
    GetRatios() []float64
    ToString() string
    SetRatios(ratios []float64)
    GetMinX() int
    SetMinX(minX int)
    GetMaxX() int
    SetMaxX(maxX int)
    GetY() int
    SetY(y int)
    GetSize() int
    SetSize(size int)
}

func (this *MeteredGameObject) MeteredGameObject_() *MeteredGameObject {
    return this
}


func (this *MeteredGameObject) initFields() {
    
    this.Good = false
    this.Dropping = false
    
    
    
    
    this.RATIOS_LABEL = "Ratios: "
    this.IS_GOOD_LABEL = "isGood: "
    this.IS_DROPPPING_LABEL = "isDropping: "
}


func NewMeteredGameObject(bufferedImage []*BufferedImage, minX int, maxX int, y int) *MeteredGameObject {
    this := &MeteredGameObject{}
    this.initFields()
        this.SetRatios(make([]float64, len(bufferedImage)))
        this.SetMinX(minX)
        this.SetMaxX(maxX)
        this.SetY(y)
        this.SetSize(maxX - minX)
        for index := int(0); index < len(bufferedImage); index++ {
            var nominator float64 = GraphicsAnalysisUtil_GetNominator(bufferedImage[index], this.GetMinX(), this.GetMaxX(), this.GetY()); _ = nominator
            this.GetRatios()[index] = float64(nominator / this.GetSize())
        }
        if float64(this.GetRatios()[len(bufferedImage) - 1]) > float64(0.93) {
            this.SetGood(true)
        } else {
            this.SetGood(false)
        }
        if float64(this.GetRatios()[len(bufferedImage) - 1]) < float64(this.GetRatios()[0]) {
            this.SetDropping(true)
        } else {
            this.SetDropping(false)
        }
    return this
}


func (this *MeteredGameObject) IsGood() bool {
        return this.Good
}


func (this *MeteredGameObject) SetGood(good bool) {
        this.Good = good
}


func (this *MeteredGameObject) IsDropping() bool {
        return this.Dropping
}


func (this *MeteredGameObject) SetDropping(dropping bool) {
        this.Dropping = dropping
}


func (this *MeteredGameObject) GetRatios() []float64 {
        return this.Ratios
}


func (this *MeteredGameObject) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(this.RATIOS_LABEL)
        for index := int(0); index < len(this.GetRatios()); index++ {
            stringBuffer.Append(fmt.Sprint(this.GetRatios()[index]))
            if index < len(this.GetRatios()) - 1 {
                stringBuffer.Append(commonSeps.COMMA_SEP)
            }
        }
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(this.IS_GOOD_LABEL)
        stringBuffer.Appendboolean(this.IsGood())
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(this.IS_DROPPPING_LABEL)
        stringBuffer.Appendboolean(this.IsDropping())
        return fmt.Sprint(stringBuffer)
}

func (this *MeteredGameObject) String() string {
    return this.ToString()
}


func (this *MeteredGameObject) SetRatios(ratios []float64) {
        this.Ratios = ratios
}


func (this *MeteredGameObject) GetMinX() int {
        return this.MinX
}


func (this *MeteredGameObject) SetMinX(minX int) {
        this.MinX = minX
}


func (this *MeteredGameObject) GetMaxX() int {
        return this.MaxX
}


func (this *MeteredGameObject) SetMaxX(maxX int) {
        this.MaxX = maxX
}


func (this *MeteredGameObject) GetY() int {
        return this.Y
}


func (this *MeteredGameObject) SetY(y int) {
        this.Y = y
}


func (this *MeteredGameObject) GetSize() int {
        return this.Size
}


func (this *MeteredGameObject) SetSize(size int) {
        this.Size = size
}

