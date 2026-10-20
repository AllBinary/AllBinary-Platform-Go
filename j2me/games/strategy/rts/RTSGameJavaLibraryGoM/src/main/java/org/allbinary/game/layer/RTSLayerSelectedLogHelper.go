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

type RTSLayerSelectedLogHelper struct {
    RTSLayerLogHelper
    LogUtil *LogUtil
    CommonSeps *CommonSeps
    PositionStrings *PositionStrings
    TRACKTO_TURNTO string
}

var RTSLayerSelectedLogHelper_InstanceC *RTSLayerSelectedLogHelper = NewRTSLayerSelectedLogHelper()

func (this *RTSLayerSelectedLogHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonSeps = CommonSeps_GetInstance()
    this.PositionStrings = PositionStrings_GetInstance()
    this.TRACKTO_TURNTO = "trackTo:turnTo"
}

func NewRTSLayerSelectedLogHelper() *RTSLayerSelectedLogHelper {
    this := &RTSLayerSelectedLogHelper{}
    this.RTSLayerLogHelper = *NewRTSLayerLogHelper()
    this.initFields()
    return this
}


func RTSLayerSelectedLogHelper_GetInstance() *RTSLayerSelectedLogHelper {
        return RTSLayerSelectedLogHelper_InstanceC
}


func (this *RTSLayerSelectedLogHelper) SetClosestGeographicMapCellHistory(associatedAdvancedRTSGameLayer PathFindingLayerInterface, pathsList IBasicArrayList) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(CommonLabels_GetInstance().START)
        stringBuffer.Appendint(pathsList.Size())
        stringBuffer.Append(" -> ")
        stringBuffer.Append(StringUtil_GetInstance().ToString(pathsList))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: setClosestGeographicMapCellHistory")
}


func (this *RTSLayerSelectedLogHelper) TrackTo(associatedAdvancedRTSGameLayer PathFindingLayerInterface, nextUnvisitedPathGeographicMapCellPosition IGeographicMapCellPosition, dx int, dy int, reason string) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(StringUtil_GetInstance().ToString(nextUnvisitedPathGeographicMapCellPosition))
        stringBuffer.Append(this.PositionStrings.DX_LABEL)
        stringBuffer.Appendint(dx)
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.PositionStrings.DY_LABEL)
        stringBuffer.Appendint(dy)
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(reason)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "selected: trackTo")
}


func (this *RTSLayerSelectedLogHelper) TurnTo(associatedAdvancedRTSGameLayer PathFindingLayerInterface, dx int, dy int, angleInfo *AngleInfo, angle int, movementAngle *NamedAngle, evading bool, targetAngle int) {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(associatedAdvancedRTSGameLayer.GetName())
        stringBuffer.Append(" steering - ")
        stringBuffer.Append(this.PositionStrings.DX_LABEL)
        stringBuffer.Appendint(dx)
        stringBuffer.Append(this.CommonSeps.SPACE)
        stringBuffer.Append(this.PositionStrings.DY_LABEL)
        stringBuffer.Appendint(dy)
        stringBuffer.Append(" angle: ")
        stringBuffer.Appendint(angle)
        stringBuffer.Append(" movementAngle: ")
        stringBuffer.Append(movementAngle.Name)
        stringBuffer.Append(this.CommonSeps.FORWARD_SLASH)
        stringBuffer.Appendshort(movementAngle.GetValue())
        if angleInfo != nil {
            var angleIncrement int16 = angleInfo.GetAngleIncrementInfo().GetAngleIncrement(); _ = angleIncrement
            stringBuffer.Append(" angleIncrement: ")
            stringBuffer.Appendshort(angleIncrement)
        }
        stringBuffer.Append(" Evading: ")
        stringBuffer.Appendboolean(evading)
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) DoneMoving(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - done moving 0,0")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) CloseEnough(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - close enough")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) MovingLeft(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - allow movingLeft")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) MovingRight(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - allow movingRight")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) MovingUp(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - allow movingUp")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) MovingDown(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - allow movingDown")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) CurrentMoveEnded(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - correct angle, but movement not needed for current deltas so reseting movement angle")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) Evade(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - evade")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) RotateLeft(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - rotating left")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) RotateRight(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - rotating right")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) Handle(associatedAdvancedRTSGameLayer PathFindingLayerInterface, movementAngle *NamedAngle) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - movementAngle: ").Append(movementAngle.Name).Append(this.CommonSeps.FORWARD_SLASH).Appendshort(movementAngle.GetValue())), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) NoRotation(associatedAdvancedRTSGameLayer PathFindingLayerInterface) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(" steering - Do not rotate")), this, this.TRACKTO_TURNTO)
}


func (this *RTSLayerSelectedLogHelper) NotOnPath(associatedAdvancedRTSGameLayer PathFindingLayerInterface, geographicMapCellHistory *GeographicMapCellHistory, currentGeographicMapCellPosition IGeographicMapCellPosition, pathList IBasicArrayList) {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(associatedAdvancedRTSGameLayer.GetName()).Append(CommonSeps_GetInstance().SPACE).Appendint(geographicMapCellHistory.GetTotalVisited()).Append(CommonSeps_GetInstance().SPACE).Append(stringUtil.ToString(currentGeographicMapCellPosition)).Append(" - trying to move but not on path: ").Append(stringUtil.ToString(pathList))), this, "turnTo")
}

