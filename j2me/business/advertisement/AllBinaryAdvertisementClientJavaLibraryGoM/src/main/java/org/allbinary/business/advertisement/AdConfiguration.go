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

type AdConfiguration struct {
    INNERACTIVE string
    LEADBOLT string
    ADMOB string
}

// The Java class AdConfiguration has subclasses so its variables are this interface
type IAdConfiguration interface {
    AdConfiguration_() *AdConfiguration
    GetValidAdSpots() []*Direction
    Process(gameState *GameState)
    ProcessDemo(state int)
    SetShowAds()
    SetToggleAds(showAds bool)
    IsShowAds() bool
}

func (this *AdConfiguration) AdConfiguration_() *AdConfiguration {
    return this
}


func (this *AdConfiguration) initFields() {
    
    
    
}


func NewAdConfiguration(advertIdArray []any) *AdConfiguration {
    this := &AdConfiguration{}
    this.initFields()
        this.INNERACTIVE = any(advertIdArray[0]).(string)
        this.ADMOB = any(advertIdArray[1]).(string)
        this.LEADBOLT = any(advertIdArray[2]).(string)
    return this
}


func (this *AdConfiguration) GetValidAdSpots() []*Direction {
        return make([]*Direction, 0)
}


func (this *AdConfiguration) Process(gameState *GameState) {
}


func (this *AdConfiguration) ProcessDemo(state int) {
}


func (this *AdConfiguration) SetShowAds() {
}


func (this *AdConfiguration) SetToggleAds(showAds bool) {
}


func (this *AdConfiguration) IsShowAds() bool {
        return true
}

