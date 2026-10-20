/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type GameAdStateBase struct {
    AdConfigurationP IAdConfiguration
    AdvertisementsP *Advertisements
    OkayToShowAdsP bool
}

// The Java class GameAdStateBase has subclasses so its variables are this interface
type IGameAdStateBase interface {
    GameAdStateBase_() *GameAdStateBase
    InitViewArray(viewArray []any)
    OnStart(anyType any)
    Init()
    IsShowingAtLocation(location int) bool
    IsShowingAt() bool
    Show()
    Hide()
    OnEvent(eventObject IAllBinaryEventObject)
    ProcessAdState(gameState *GameState, gameType *GameType)
    ProcessPageAdState()
    PlayingAdState()
    PlayingAdStateWH(width int, height int)
    NotPlayingAdState()
    SetGameIsReady(gameIsReady bool)
    IsGameIsReady() bool
    GetAdvertisements() *Advertisements
    GetAdConfiguration() IAdConfiguration
    IsOkayToShowPageAd() bool
    SetOkayToShowAds(okayToShowAds bool)
    IsOkayToShowAds() bool
    IsEnabled() bool
}

func (this *GameAdStateBase) GameAdStateBase_() *GameAdStateBase {
    return this
}


func (this *GameAdStateBase) initFields() {
    
    this.AdvertisementsP = NewAdvertisements()
    
}


func NewGameAdStateBase(adConfiguration IAdConfiguration) *GameAdStateBase {
    this := &GameAdStateBase{}
    this.initFields()
        this.AdConfigurationP = adConfiguration
    return this
}


func (this *GameAdStateBase) InitViewArray(viewArray []any) {
}


func (this *GameAdStateBase) OnStart(anyType any) {
        this.Init()
}


func (this *GameAdStateBase) Init() {
}


func (this *GameAdStateBase) IsShowingAtLocation(location int) bool {
        return false
}


func (this *GameAdStateBase) IsShowingAt() bool {
        return false
}


func (this *GameAdStateBase) Show() {
}


func (this *GameAdStateBase) Hide() {
}


func (this *GameAdStateBase) OnEvent(eventObject IAllBinaryEventObject) {
        ForcedLogUtil_Log(EventStrings_GetInstance().PERFORMANCE_MESSAGE, this)
}


func (this *GameAdStateBase) ProcessAdState(gameState *GameState, gameType *GameType) {
}


func (this *GameAdStateBase) ProcessPageAdState() {
}


// throws Exception
func (this *GameAdStateBase) PlayingAdState() {
}


// throws Exception
func (this *GameAdStateBase) PlayingAdStateWH(width int, height int) {
}


// throws Exception
func (this *GameAdStateBase) NotPlayingAdState() {
}


func (this *GameAdStateBase) SetGameIsReady(gameIsReady bool) {
}


func (this *GameAdStateBase) IsGameIsReady() bool {
        return false
}


func (this *GameAdStateBase) GetAdvertisements() *Advertisements {
        return this.AdvertisementsP
}


func (this *GameAdStateBase) GetAdConfiguration() IAdConfiguration {
        return this.AdConfigurationP
}


func (this *GameAdStateBase) IsOkayToShowPageAd() bool {
        return false
}


func (this *GameAdStateBase) SetOkayToShowAds(okayToShowAds bool) {
        this.OkayToShowAdsP = okayToShowAds
}


func (this *GameAdStateBase) IsOkayToShowAds() bool {
        return this.OkayToShowAdsP
}


func (this *GameAdStateBase) IsEnabled() bool {
        return false
}

