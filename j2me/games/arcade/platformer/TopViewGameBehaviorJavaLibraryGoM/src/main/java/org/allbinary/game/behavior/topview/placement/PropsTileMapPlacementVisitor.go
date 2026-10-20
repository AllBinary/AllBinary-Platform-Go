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

type PropsTileMapPlacementVisitor struct {
    TileMapPlacementVisitor
    LogUtil *LogUtil
    RandomFactory *MyRandomFactory
}


func (this *PropsTileMapPlacementVisitor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.RandomFactory = MyRandomFactory_GetInstance()
}

func NewPropsTileMapPlacementVisitor() *PropsTileMapPlacementVisitor {
    this := &PropsTileMapPlacementVisitor{}
    this.TileMapPlacementVisitor = *NewTileMapPlacementVisitor()
    this.initFields()
    return this
}


func (this *PropsTileMapPlacementVisitor) Visit(lastMap *TiledMap, mapData [][]int) {
        var layerIndex int = 0; _ = layerIndex
        var tileLayer *TileLayer = (any(lastMap.GetLayer(layerIndex)).(*TileLayer)); _ = tileLayer
        var mapArray [][]int = tileLayer.GetMapArray(); _ = mapArray
        var size int = len(mapData); _ = size
        var size2 int = len(mapData[0]); _ = size2
        for index := int(0); index < size; index++ {
            for index2 := int(0); index2 < size2; index2++ {
                if mapArray[index][index2] == 14 {
                } else {
                    this.Visit2(mapArray, mapData, index, index2)
                }
                if index % 7 == 0 {
                    if mapArray[index][index2] == 19 {
                        var randomInt int = this.RandomFactory.GetAbsoluteNextIntAllowZero(3); _ = randomInt
                        if randomInt == 0 {
                            mapData[index][index2] = 1
                        } else if randomInt == 1 {
                            mapData[index][index2] = 81
                        } else if randomInt == 2 {
                            mapData[index][index2] = 97
                        }
                    } else if mapArray[index][index2] == 17 {
                        var randomInt int = this.RandomFactory.GetAbsoluteNextIntAllowZero(3); _ = randomInt
                        if randomInt == 0 {
                            mapData[index][index2] = 17
                        } else if randomInt == 1 {
                            mapData[index][index2] = 81
                        } else if randomInt == 2 {
                            mapData[index][index2] = 97
                        }
                    }
                }
                if index2 % 7 == 0 {
                    if mapArray[index][index2] == 34 {
                        var randomInt int = this.RandomFactory.GetAbsoluteNextIntAllowZero(4); _ = randomInt
                        if randomInt == 0 {
                            mapData[index][index2] = 33
                        } else if randomInt == 1 {
                            mapData[index][index2] = 65
                        } else if randomInt == 2 {
                            mapData[index][index2] = 81
                        } else if randomInt == 3 {
                            mapData[index][index2] = 97
                        }
                    }
                }
            }
        }
}


func (this *PropsTileMapPlacementVisitor) Visit2(mapArray [][]int, mapData [][]int, index int, index2 int) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var countX int = 0; _ = countX
        var countY int = 0; _ = countY
        var index3 int = index - 1; _ = index3
        for index4 := int(index2 - 1); index4 > 0; index4-- {
            index3--
            if index3 < 0 || index4 < 0 {
                return
            }
            if mapArray[index3][index4] != 14 || mapData[index3][index4] != 0 {
                if countX > 10 && countY > 10 {
                    for index5 := int(index3 + 1); index5 < index - 1; index5++ {
                        for index6 := int(index4 + 1); index6 < index2 - 1; index6++ {
                            if mapArray[index5][index6] != 14 || mapData[index5][index6] != 0 {
                                return
                            }
                        }
                    }
                    var x int = index - (countX / 2); _ = x
                    var y int = index2 - (countY / 2); _ = y
                    if x >= 0 && y >= 0 && x < len(mapData) && y < len(mapData[0]) {
                    }
                }
                return
            }
            countY++
            countX++
        }
}

