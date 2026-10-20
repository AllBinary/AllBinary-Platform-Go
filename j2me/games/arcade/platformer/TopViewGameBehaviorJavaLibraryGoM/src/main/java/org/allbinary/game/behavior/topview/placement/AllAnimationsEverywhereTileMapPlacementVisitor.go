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

type AllAnimationsEverywhereTileMapPlacementVisitor struct {
    TileMapPlacementVisitor
}


func NewAllAnimationsEverywhereTileMapPlacementVisitor() *AllAnimationsEverywhereTileMapPlacementVisitor {
    this := &AllAnimationsEverywhereTileMapPlacementVisitor{}
    this.TileMapPlacementVisitor = *NewTileMapPlacementVisitor()
    return this
}


func (this *AllAnimationsEverywhereTileMapPlacementVisitor) Visit(lastMap *TiledMap, mapData [][]int) {
        var size int = len(mapData); _ = size
        var size2 int = len(mapData[0]) - 7; _ = size2
        for index := int(0); index < size; index++ {
            for index2 := int(0); index2 < size2; index2 += 7 {
                mapData[index][index2] = 1
                mapData[index][index2 + 1] = 17
                mapData[index][index2 + 2] = 33
                mapData[index][index2 + 3] = 49
                mapData[index][index2 + 4] = 65
                mapData[index][index2 + 5] = 81
                mapData[index][index2 + 6] = 97
            }
        }
}

