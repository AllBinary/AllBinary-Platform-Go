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

type RaceTrackInfo struct {
    Id int
    Name string
    RaceTrackFrictionProperties *RaceTrackFrictionProperties
    ForegroundBasicColor *BasicColor
    BackgroundBasicColor *BasicColor
    QualifyTime int
    WorstContinuePosition int
    TotalLaps int
}


func (this *RaceTrackInfo) initFields() {
    
    
    
    
    
    
    
    
}


func NewRaceTrackInfo(id int, name string, raceTrackFrictionProperties *RaceTrackFrictionProperties, foregroundBasicColor *BasicColor, backgroundBasicColor *BasicColor, qualifyTime int, worstContinuePosition int, totalLaps int) *RaceTrackInfo {
    this := &RaceTrackInfo{}
    this.initFields()
        this.Id = id
        this.Name = name
        this.RaceTrackFrictionProperties = raceTrackFrictionProperties
        this.BackgroundBasicColor = backgroundBasicColor
        this.ForegroundBasicColor = foregroundBasicColor
        this.QualifyTime = qualifyTime
        this.WorstContinuePosition = worstContinuePosition
        this.TotalLaps = totalLaps
    return this
}


func (this *RaceTrackInfo) GetId() int {
        return this.Id
}


func (this *RaceTrackInfo) SetId(id int) {
        this.Id = id
}


func (this *RaceTrackInfo) GetName() string {
        return this.Name
}


func (this *RaceTrackInfo) SetName(name string) {
        this.Name = name
}


func (this *RaceTrackInfo) GetQualifyTime() int {
        return this.QualifyTime
}


func (this *RaceTrackInfo) SetQualifyTime(qualifyTime int) {
        this.QualifyTime = qualifyTime
}


func (this *RaceTrackInfo) GetWorstContinuePosition() int {
        return this.WorstContinuePosition
}


func (this *RaceTrackInfo) SetWorstContinuePosition(worstContinuePosition int) {
        this.WorstContinuePosition = worstContinuePosition
}


func (this *RaceTrackInfo) GetTotalLaps() int {
        return this.TotalLaps
}


func (this *RaceTrackInfo) SetTotalLaps(totalLaps int) {
        this.TotalLaps = totalLaps
}


func (this *RaceTrackInfo) GetForegroundBasicColor() *BasicColor {
        return this.ForegroundBasicColor
}


func (this *RaceTrackInfo) SetForegroundBasicColor(foregroundBasicColor *BasicColor) {
        this.ForegroundBasicColor = foregroundBasicColor
}


func (this *RaceTrackInfo) GetBackgroundBasicColor() *BasicColor {
        return this.BackgroundBasicColor
}


func (this *RaceTrackInfo) SetBackgroundBasicColor(backgroundBasicColor *BasicColor) {
        this.BackgroundBasicColor = backgroundBasicColor
}


func (this *RaceTrackInfo) GetRaceTrackFrictionProperties() *RaceTrackFrictionProperties {
        return this.RaceTrackFrictionProperties
}


func (this *RaceTrackInfo) SetRaceTrackFrictionProperties(raceTrackFrictionProperties *RaceTrackFrictionProperties) {
        this.RaceTrackFrictionProperties = raceTrackFrictionProperties
}

