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

type SmallBasicColorCacheFactory struct {
    SIZE int
    INDEX_TO_COLOR []int
    BASIC_COLOR_ARRAY []*BasicColor
    ColorIndex int
}

var SmallBasicColorCacheFactory_Instance *SmallBasicColorCacheFactory = NewSmallBasicColorCacheFactory()

func (this *SmallBasicColorCacheFactory) initFields() {
    this.SIZE = 255
    this.INDEX_TO_COLOR = make([]int, this.SIZE)
    this.BASIC_COLOR_ARRAY = make([]*BasicColor, this.SIZE)
    this.ColorIndex = 0
}


func SmallBasicColorCacheFactory_GetInstance() *SmallBasicColorCacheFactory {
        return SmallBasicColorCacheFactory_Instance
}


func NewSmallBasicColorCacheFactory() *SmallBasicColorCacheFactory {
    this := &SmallBasicColorCacheFactory{}
    this.initFields()
    return this
}


func (this *SmallBasicColorCacheFactory) Add(basicDefaultColor *BasicColor) {
        this.BASIC_COLOR_ARRAY[this.ColorIndex] = basicDefaultColor
        this.INDEX_TO_COLOR[this.ColorIndex] = basicDefaultColor.IntValue()
        this.ColorIndex++
}


func (this *SmallBasicColorCacheFactory) GetAndOrCreate(colorAsInt int) *BasicColor {
        var basicColor *BasicColor; _ = basicColor
        for index := int(0); index < this.SIZE; index++ {
            if this.INDEX_TO_COLOR[index] == colorAsInt {
                return this.BASIC_COLOR_ARRAY[index]
            }
        }
        var ALPHA_MASK int = 0xFF000000; _ = ALPHA_MASK
        basicColor = BasicColorFactory_GetInstance().CreateInstanceAN(colorAsInt & ALPHA_MASK, colorAsInt & 0x00FFFFFF, StringUtil_GetInstance().EMPTY_STRING)
        this.Add(basicColor)
        return basicColor
}
