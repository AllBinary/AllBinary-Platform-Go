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

type BaseGameInitialization struct {
    ResourceInitializationArray []*ResourceInitialization
    Portion int
    Initialized bool
    EARLY_RESOURCES int
    GAME_RESOURCES int
    EARLY_CHANGABLE_RESOURCES int
    GAME_CHANGABLE_RESOURCES int
    GameInitialized bool
    AllLoaded bool
}

// Implements GameInitializationInterface

// The Java class BaseGameInitialization has subclasses so its variables are this interface
type IBaseGameInitialization interface {
    BaseGameInitialization_() *BaseGameInitialization
    InitKey(portion int)
    InitKeyMapping(abeClientInformation AbeClientInformationInterface, portion int)
    Init(abeClientInformation AbeClientInformationInterface, commandListener CommandListener, level int)
    ResourceInitialization(level int)
    ResourceAnimationChange() bool
    ClearResources()
    SetGameInitialized(gameInitialized bool)
    IsGameInitialized() bool
    GetPortion() int
    SetInitialized(initialized bool)
    IsInitialized() bool
}

func (this *BaseGameInitialization) BaseGameInitialization_() *BaseGameInitialization {
    return this
}

var BaseGameInitialization_NULL_BASE_GAME_INITIALIZATION IBaseGameInitialization = NewBaseGameInitialization(make([]*ResourceInitialization, 0), 0)

func (this *BaseGameInitialization) initFields() {
    
    
    
    this.EARLY_RESOURCES = 0
    this.GAME_RESOURCES = 1
    this.EARLY_CHANGABLE_RESOURCES = 2
    this.GAME_CHANGABLE_RESOURCES = 3
    
    
}


func NewBaseGameInitialization(resourceInitializationArray []*ResourceInitialization, portion int) *BaseGameInitialization {
    this := &BaseGameInitialization{}
    this.initFields()
        this.ResourceInitializationArray = resourceInitializationArray
        this.Portion = portion
    return this
}


// throws Exception
func (this *BaseGameInitialization) InitKey(portion int) {
}


// throws Exception
func (this *BaseGameInitialization) InitKeyMapping(abeClientInformation AbeClientInformationInterface, portion int) {
        if ChangedGameFeatureListener_GetInstance().IsChangedFeature(InputFeatureFactory_GetInstance().INPUT_MAPPING) {
            PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance().Init(abeClientInformation)
            ProgressCanvasFactory_GetInstance().AddNormalPortion(50, "Game Keys")
            ChangedGameFeatureListener_GetInstance().Remove(InputFeatureFactory_GetInstance().INPUT_MAPPING)
        }
}


// throws Exception
func (this *BaseGameInitialization) Init(abeClientInformation AbeClientInformationInterface, commandListener CommandListener, level int) {
        var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
        if !this.IsInitialized() && (level == resourceLoadingLevelFactory.LOAD_ALL.GetLevel() || level == resourceLoadingLevelFactory.LOAD_EARLY.GetLevel()) {
            var localPortion int = 40; _ = localPortion
            if level == resourceLoadingLevelFactory.LOAD_EARLY.GetLevel() {
                localPortion = 8
            }
            this.SetInitialized(true)
            this.InitKey(this.GetPortion())
            this.InitKeyMapping(abeClientInformation, this.GetPortion())
            GameKeyEventFactory_GetInstance().Init()
            var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
            progressCanvas.AddNormalPortion(localPortion, "Game Key Events")
            BasicTouchInputFactory_GetInstance().Init(PlatformInputMappingFactory_GetInstance().GetPersistentInputMappingInstance().GetInputMapping())
            CancelTouchButtonInputFactory_GetInstance()
            progressCanvas.AddNormalPortion(localPortion, "Touch Input")
            CompleteMotionGestureInputToGameMotionGestureInput_GetInstance().Init()
            progressCanvas.AddNormalPortion(localPortion, "Motion Input")
            GameFeatureImageCacheFactory_Init()
            progressCanvas.AddNormalPortion(localPortion, "Image Cache")
            this.ResourceInitializationArray[this.EARLY_RESOURCES].Init()
        }
}


// throws Exception
func (this *BaseGameInitialization) ResourceInitialization(level int) {
        var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
        if this.ResourceAnimationChange() && level == resourceLoadingLevelFactory.LOAD_EARLY.GetLevel() {
            this.ClearResources()
            this.ResourceInitializationArray[this.EARLY_CHANGABLE_RESOURCES].Init()
        }
        if this.ResourceAnimationChange() && level == resourceLoadingLevelFactory.LOAD_ALL.GetLevel() {
            if this.AllLoaded {
                this.ClearResources()
            }
            if FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance().GetList().Size() == 0 {
                this.ResourceInitializationArray[this.EARLY_CHANGABLE_RESOURCES].Init()
            }
            this.ResourceInitializationArray[this.GAME_CHANGABLE_RESOURCES].Init()
            this.AllLoaded = true
        }
}


func (this *BaseGameInitialization) ResourceAnimationChange() bool {
        var changedGameFeatureListener *ChangedGameFeatureListener = ChangedGameFeatureListener_GetInstance(); _ = changedGameFeatureListener
        var gameConfigurationCentral *GameConfigurationCentral = GameConfigurationCentral_GetInstance(); _ = gameConfigurationCentral
        if changedGameFeatureListener.IsChangedGameConfiguration(gameConfigurationCentral.SCALE) {
            return true
        }
        if ChangedGameFeatureListener_GetInstance().IsChangedFeature(MainFeatureFactory_GetInstance().STATIC) {
            return true
        }
        return false
}


func (this *BaseGameInitialization) ClearResources() {
        this.AllLoaded = false
        GameFeatureImageCacheFactory_ReleaseAll()
        var featuredAnimationInterfaceFactoryInterfaceFactory *FeaturedAnimationInterfaceFactoryInterfaceFactory = FeaturedAnimationInterfaceFactoryInterfaceFactory_GetInstance(); _ = featuredAnimationInterfaceFactoryInterfaceFactory
        featuredAnimationInterfaceFactoryInterfaceFactory.Clear()
        FeaturedResourceRelativeRelationshipFactory_GetInstance().Clear()
}


func (this *BaseGameInitialization) SetGameInitialized(gameInitialized bool) {
        this.GameInitialized = gameInitialized
}


func (this *BaseGameInitialization) IsGameInitialized() bool {
        return this.GameInitialized
}


func (this *BaseGameInitialization) GetPortion() int {
        return this.Portion
}


func (this *BaseGameInitialization) SetInitialized(initialized bool) {
        this.Initialized = initialized
}


func (this *BaseGameInitialization) IsInitialized() bool {
        return this.Initialized
}

