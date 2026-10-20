/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2003 AllBinary 
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

type GeographicMapCellPositionArea struct {
    GeographicMapCellPositionAreaBase
    CellPositionsUtil *CellPositionsUtil
    LayerCoveringCellPositionsUtil *LayerCoveringCellPositionsUtil
    ReusableOccupyingGeographicMapCellPositionList IBasicArrayList
    ReusableSurroundingGeographicMapCellPositionList IBasicArrayList
    SurroundingCircularIndexUtil *CircularIndexUtil
    LayerInterface IAllBinaryLayer
    LIST IBasicArrayList
    OccupyingGeographicMapCellPositionList IBasicArrayList
    SurroundingGeographicMapCellPositionList IBasicArrayList
}


func (this *GeographicMapCellPositionArea) initFields() {
    this.CellPositionsUtil = CellPositionsUtil_GetInstance()
    this.LayerCoveringCellPositionsUtil = LayerCoveringCellPositionsUtil_GetInstance()
    this.ReusableOccupyingGeographicMapCellPositionList = NewBasicArrayListS(4)
    this.ReusableSurroundingGeographicMapCellPositionList = NewBasicArrayListS(12)
    this.SurroundingCircularIndexUtil = CircularIndexUtil_CreateInstance(0)
    
    this.LIST = BasicArrayListUtil_GetInstance().GetImmutableInstance()
    this.OccupyingGeographicMapCellPositionList = this.LIST
    this.SurroundingGeographicMapCellPositionList = this.LIST
}


func NewGeographicMapCellPositionArea(layerInterface IAllBinaryLayer) *GeographicMapCellPositionArea {
    this := &GeographicMapCellPositionArea{}
    this.GeographicMapCellPositionAreaBase = *NewGeographicMapCellPositionAreaBase()
    this.initFields()
        this.LayerInterface = layerInterface
    return this
}


// throws Exception
func (this *GeographicMapCellPositionArea) Update(geographicMapInterface IBasicGeographicMap) {
        this.OccupyingGeographicMapCellPositionList = this.LayerCoveringCellPositionsUtil.GetAllXY(geographicMapInterface, this.LayerInterface, this.LayerInterface.GetXP(), this.LayerInterface.GetYP(), this.ReusableOccupyingGeographicMapCellPositionList)
        this.SurroundingGeographicMapCellPositionList = this.CellPositionsUtil.GetAllSurrounding(geographicMapInterface, this.OccupyingGeographicMapCellPositionList, this.ReusableSurroundingGeographicMapCellPositionList)
        this.SurroundingCircularIndexUtil.SetSize(this.SurroundingGeographicMapCellPositionList.Size())
}


func (this *GeographicMapCellPositionArea) GetOccupyingGeographicMapCellPositionList() IBasicArrayList {
        return this.OccupyingGeographicMapCellPositionList
}


func (this *GeographicMapCellPositionArea) GetSurroundingGeographicMapCellPositionList() IBasicArrayList {
        return this.SurroundingGeographicMapCellPositionList
}


func (this *GeographicMapCellPositionArea) GetNextSurroundingGeographicMapCellPosition() IGeographicMapCellPosition {
        var geographicMapCellPosition IGeographicMapCellPosition = any(this.SurroundingGeographicMapCellPositionList.Get(this.SurroundingCircularIndexUtil.GetIndex())).(IGeographicMapCellPosition); _ = geographicMapCellPosition
        this.SurroundingCircularIndexUtil.Next()
        return geographicMapCellPosition
}

