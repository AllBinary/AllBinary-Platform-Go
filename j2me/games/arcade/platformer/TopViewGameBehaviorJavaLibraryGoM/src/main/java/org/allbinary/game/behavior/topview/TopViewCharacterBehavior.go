/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type TopViewCharacterBehavior struct {
}

// The Java class TopViewCharacterBehavior has subclasses so its variables are this interface
type ITopViewCharacterBehavior interface {
    TopViewCharacterBehavior_() *TopViewCharacterBehavior
    TerrainEvent(layer IAllBinaryLayer, direction *Direction, x int, y int, geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellPosition IGeographicMapCellPosition)
    TerrainMove(layer IAllBinaryLayer, geographicMapInterfaceArray []IBasicGeographicMap, x int, y int)
    HasSolidBlock(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType) bool
}

func (this *TopViewCharacterBehavior) TopViewCharacterBehavior_() *TopViewCharacterBehavior {
    return this
}


func NewTopViewCharacterBehavior() *TopViewCharacterBehavior {
    this := &TopViewCharacterBehavior{}
    return this
}


// throws Exception
func (this *TopViewCharacterBehavior) TerrainEvent(layer IAllBinaryLayer, direction *Direction, x int, y int, geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellPosition IGeographicMapCellPosition) {
}


func (this *TopViewCharacterBehavior) TerrainMove(layer IAllBinaryLayer, geographicMapInterfaceArray []IBasicGeographicMap, x int, y int) {
}


func (this *TopViewCharacterBehavior) HasSolidBlock(geographicMapInterfaceArray []IBasicGeographicMap, geographicMapCellTypeArray []IGeographicMapCellType) bool {
        var size int = len(geographicMapInterfaceArray); _ = size
        var basicTopViewGeographicMapCellTypeFactory *BasicTopViewGeographicMapCellTypeFactory; _ = basicTopViewGeographicMapCellTypeFactory
        for index := int(0); index < size; index++ {
            basicTopViewGeographicMapCellTypeFactory = any(geographicMapInterfaceArray[index].GetGeographicMapCellTypeFactory()).(*BasicTopViewGeographicMapCellTypeFactory)
            if basicTopViewGeographicMapCellTypeFactory.BLOCK_CELL_TYPE.IsType(geographicMapCellTypeArray[index]) {
                return true
            }
        }
        return false
}

