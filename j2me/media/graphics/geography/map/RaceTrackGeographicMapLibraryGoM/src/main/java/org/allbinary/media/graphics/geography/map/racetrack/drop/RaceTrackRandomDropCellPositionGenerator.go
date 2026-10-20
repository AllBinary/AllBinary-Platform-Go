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

type RaceTrackRandomDropCellPositionGenerator struct {
    RaceTrackDropCellPositionGenerator
    MyRandomFactory *MyRandomFactory
    CellsPerRowOrColumn int
    TotalCells int
    RowArray []int
    ColumnArray []int
    CellWidth int
    CellHeight int
}

var RaceTrackRandomDropCellPositionGenerator_SINGLETON *RaceTrackRandomDropCellPositionGenerator = NewRaceTrackRandomDropCellPositionGenerator()

func (this *RaceTrackRandomDropCellPositionGenerator) initFields() {
    this.MyRandomFactory = MyRandomFactory_GetInstance()
    this.CellsPerRowOrColumn = 3
    this.TotalCells = this.CellsPerRowOrColumn * this.CellsPerRowOrColumn
    this.RowArray = []int{0, 1, 2, 0, 1, 2, 0, 1, 2}
    this.ColumnArray = []int{0, 0, 0, 1, 1, 1, 2, 2, 2}
    
    
}


func RaceTrackRandomDropCellPositionGenerator_GetInstance() *RaceTrackRandomDropCellPositionGenerator {
        return RaceTrackRandomDropCellPositionGenerator_SINGLETON
}


func NewRaceTrackRandomDropCellPositionGenerator() *RaceTrackRandomDropCellPositionGenerator {
    this := &RaceTrackRandomDropCellPositionGenerator{}
    this.RaceTrackDropCellPositionGenerator = *NewRaceTrackDropCellPositionGenerator()
    this.initFields()
    return this
}


// throws Exception
func (this *RaceTrackRandomDropCellPositionGenerator) Update(allBinaryGameLayerManager IAllBinaryGameLayerManager, geographicMapInterface IBasicGeographicMap) {
        this.RaceTrackDropCellPositionGenerator.Update(allBinaryGameLayerManager, geographicMapInterface)
        var tiledLayer IAllBinaryTiledLayer = geographicMapInterface.GetAllBinaryTiledLayer(); _ = tiledLayer
        this.CellWidth = tiledLayer.GetCellWidth() / this.CellsPerRowOrColumn
        this.CellHeight = tiledLayer.GetCellHeight() / this.CellsPerRowOrColumn
}


// throws Exception
func (this *RaceTrackRandomDropCellPositionGenerator) Drop(allBinaryLayerManager IAllBinaryLayerManager, index int) {
        var geographicMapCellPosition IGeographicMapCellPosition = any(this.List.Get(index)).(IGeographicMapCellPosition); _ = geographicMapCellPosition
        var point IGPoint = geographicMapCellPosition.GetPoint(); _ = point
        var randomCell int = this.MyRandomFactory.GetAbsoluteNextInt(this.TotalCells); _ = randomCell
        var row int = this.RowArray[randomCell]; _ = row
        var column int = this.ColumnArray[randomCell]; _ = column
        var x int = point.GetX() + (row * this.CellWidth); _ = x
        var y int = point.GetY() + (column * this.CellHeight); _ = y
        var pickedUpLayerInterfaceFactory PickedUpLayerInterfaceFactoryInterface = DropLayerFactory_GetInstance().GetRandomInstance(); _ = pickedUpLayerInterfaceFactory
        var tiledLayer IAllBinaryTiledLayer = this.RaceTrackGeographicMap.GetAllBinaryTiledLayer(); _ = tiledLayer
        var pickupLayer *RaceTrackPickupLayer = any(PickupLayerCircularStaticPool_GetInstance().GetInstanceXYZ(pickedUpLayerInterfaceFactory, x, y, tiledLayer.GetZP() + 3)).(*RaceTrackPickupLayer); _ = pickupLayer
        pickupLayer.SetTiledLayer(tiledLayer)
        allBinaryLayerManager.Append(pickupLayer)
}

