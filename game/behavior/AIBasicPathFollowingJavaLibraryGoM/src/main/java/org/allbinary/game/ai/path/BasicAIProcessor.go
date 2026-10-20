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

type BasicAIProcessor struct {
    Name string
    GeographicMapCellHistory *GeographicMapCellHistory
    GeographicMapCellPositionBasicArrayList IBasicArrayList
    GeographicMapCellPathHistoryInfo *GeographicMapCellPathHistoryInfo
    BasicAI *BasicDirectionAIHelper
    GeographicMapDirectionUtil *GeographicMapDirectionUtil
}


func (this *BasicAIProcessor) initFields() {
    
    
    
    
    
    this.GeographicMapDirectionUtil = GeographicMapDirectionUtil_GetInstance()
}


func NewBasicAIProcessor(name string, geographicMapCellHistory *GeographicMapCellHistory, geographicMapCellPathHistoryInfo *GeographicMapCellPathHistoryInfo, chosenPathList IBasicArrayList, angleInfo *AngleInfo, seed int) *BasicAIProcessor {
    this := &BasicAIProcessor{}
    this.initFields()
        this.Name = name
        this.GeographicMapCellHistory = geographicMapCellHistory
        this.GeographicMapCellPositionBasicArrayList = chosenPathList
        this.GeographicMapCellPathHistoryInfo = geographicMapCellPathHistoryInfo
        this.SetName(name)
        this.SetNewPath(geographicMapCellHistory, chosenPathList)
        this.BasicAI = NewBasicDirectionAIHelper(this.GetName(), angleInfo)
    return this
}


func (this *BasicAIProcessor) SetNewPath(geographicMapCellHistory *GeographicMapCellHistory, chosenPathList IBasicArrayList) {
        this.SetGeographicMapCellHistory(geographicMapCellHistory)
        this.GeographicMapCellPositionBasicArrayList = chosenPathList
        this.Init()
}


func (this *BasicAIProcessor) Init() {
        this.GeographicMapCellPathHistoryInfo = NewGeographicMapCellPathHistoryInfo()
        this.GeographicMapCellPathHistoryInfo.SetPreviousOnPathGeographicMapCellPosition(any(this.GeographicMapCellPositionBasicArrayList.Get(this.GeographicMapCellPositionBasicArrayList.Size() - 1)).(IGeographicMapCellPosition))
        this.GeographicMapCellPathHistoryInfo.SetNextOnPathGeographicMapCellPosition(any(this.GeographicMapCellPositionBasicArrayList.Get(this.GeographicMapCellPositionBasicArrayList.Size() - 1)).(IGeographicMapCellPosition))
        this.Update()
}


// throws Exception
func (this *BasicAIProcessor) GoTowardNextChosenOnPathPosition(currentGeographicMapCellPosition IGeographicMapCellPosition) int {
        var keyDirection int = -1; _ = keyDirection
        var goToGeographicMapCellPosition IGeographicMapCellPosition = this.GeographicMapCellPathHistoryInfo.GetNextChosenOnPathGeographicMapCellPosition(); _ = goToGeographicMapCellPosition
        var geographicMapDirectionData *Direction = this.GeographicMapDirectionUtil.GetDirectionFromCellPositionToAdjacentCellPosition(currentGeographicMapCellPosition, goToGeographicMapCellPosition); _ = geographicMapDirectionData
        if geographicMapDirectionData == DirectionFactory_GetInstance().NOT_BORDERED_WITH {
            geographicMapDirectionData = this.GeographicMapDirectionUtil.GetDirectionFromCellPositionToCellPosition(currentGeographicMapCellPosition, goToGeographicMapCellPosition)
        }
        keyDirection = this.BasicAI.GetAIKeyPressedFromDirection(geographicMapDirectionData)
        return keyDirection
}


func (this *BasicAIProcessor) Update() {
        this.GeographicMapCellPathHistoryInfo.SetNextUnvisitedOnPathGeographicMapCellPosition(this.GetGeographicMapCellHistory().GetFirstUnvisited())
        this.GeographicMapCellPathHistoryInfo.SetNextChosenOnPathGeographicMapCellPosition(this.GeographicMapCellPathHistoryInfo.GetNextUnvisitedOnPathGeographicMapCellPosition())
}


// throws Exception
func (this *BasicAIProcessor) ProcessAI(myManager IAllBinaryLayerManager, currentGeographicMapCellPosition IGeographicMapCellPosition) int {
        var keyDirection int = -1; _ = keyDirection
        if this.GeographicMapCellPositionBasicArrayList.Contains(currentGeographicMapCellPosition) {
            this.Update()
            keyDirection = this.GoTowardNextChosenOnPathPosition(currentGeographicMapCellPosition)
            this.GeographicMapCellPathHistoryInfo.SetPreviousOnPathGeographicMapCellPosition(currentGeographicMapCellPosition)
        } else {
            keyDirection = this.GoTowardNextChosenOnPathPosition(currentGeographicMapCellPosition)
        }
        return keyDirection
}


func (this *BasicAIProcessor) GetName() string {
        return this.Name
}


func (this *BasicAIProcessor) SetName(name string) {
        this.Name = name
}


func (this *BasicAIProcessor) GetGeographicMapCellPathHistoryInfo() *GeographicMapCellPathHistoryInfo {
        return this.GeographicMapCellPathHistoryInfo
}


func (this *BasicAIProcessor) SetGeographicMapCellPathHistoryInfo(geographicMapCellPathHistoryInfo *GeographicMapCellPathHistoryInfo) {
        this.GeographicMapCellPathHistoryInfo = geographicMapCellPathHistoryInfo
}


func (this *BasicAIProcessor) GetGeographicMapCellHistory() *GeographicMapCellHistory {
        return this.GeographicMapCellHistory
}


func (this *BasicAIProcessor) SetGeographicMapCellHistory(geographicMapCellHistory *GeographicMapCellHistory) {
        this.GeographicMapCellHistory = geographicMapCellHistory
}

