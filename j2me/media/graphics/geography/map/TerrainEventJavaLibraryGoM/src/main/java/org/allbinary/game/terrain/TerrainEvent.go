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

type TerrainEvent struct {
    AllBinaryEventObject
    BasicTerrainInfo *BasicTerrainInfo
}


func (this *TerrainEvent) initFields() {
    this.BasicTerrainInfo = BasicTerrainInfo_NULL_BASIC_TERRAIN_INFO
}


func NewTerrainEvent(basicTerrainInfo any) *TerrainEvent {
    this := &TerrainEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(basicTerrainInfo)
        this.initFields()
        this.SetBasicTerrainInfo(this.BasicTerrainInfo)
    return this
}


func (this *TerrainEvent) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("TerrainEvent: \n")
        stringBuffer.Append("LayerInterface: ")
        stringBuffer.Append(fmt.Sprint(this.BasicTerrainInfo))
        return fmt.Sprint(stringBuffer)
}

func (this *TerrainEvent) String() string {
    return this.ToString()
}


func (this *TerrainEvent) GetBasicTerrainInfo() *BasicTerrainInfo {
        return this.BasicTerrainInfo
}


func (this *TerrainEvent) SetBasicTerrainInfo(basicTerrainInfo *BasicTerrainInfo) {
        this.BasicTerrainInfo = basicTerrainInfo
}


func (this *TerrainEvent) SetBasicTerrainInfoForCircularStaticPool(basicTerrainInfo *BasicTerrainInfo) {
        this.BasicTerrainInfo = basicTerrainInfo
}

