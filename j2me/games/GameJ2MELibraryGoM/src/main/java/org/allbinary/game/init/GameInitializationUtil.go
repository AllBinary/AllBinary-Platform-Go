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

type GameInitializationUtil struct {
    EVENT *GameInitializedEvent
}

var GameInitializationUtil_Instance *GameInitializationUtil = NewGameInitializationUtil()

func (this *GameInitializationUtil) initFields() {
    this.EVENT = NewGameInitializedEvent(this)
}

func NewGameInitializationUtil() *GameInitializationUtil {
    this := &GameInitializationUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *GameInitializationUtil) InitDemo(abeClientInformation AbeClientInformationInterface, canvas IMyCanvas, gameInitializationInterfaceFactoryInterface *BasicBuildGameInitializerFactory) {
        var mainFeatureFactory *MainFeatureFactory = MainFeatureFactory_GetInstance(); _ = mainFeatureFactory
        if ChangedGameFeatureListener_GetInstance().IsChangedFeature(mainFeatureFactory.STATIC) {
            var gameInitializationInterface GameInitializationInterface = gameInitializationInterfaceFactoryInterface.GetInstance(); _ = gameInitializationInterface
            var features *Features = Features_GetInstance(); _ = features
            if features.IsFeature(mainFeatureFactory.LOAD_ALL) {
                this.InitGame(abeClientInformation, canvas, gameInitializationInterfaceFactoryInterface)
            } else if features.IsFeature(mainFeatureFactory.LOAD_ONDEMAND) {
                var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
                gameInitializationInterface.Init(abeClientInformation, canvas.GetCustomCommandListener(), resourceLoadingLevelFactory.LOAD_EARLY.GetLevel())
                this.EVENT.SetResourceLoadingLevel(resourceLoadingLevelFactory.LOAD_EARLY)
            } else {
                panic(errors.New(fmt.Sprint("No Loading Feature Available")))
            }
            GameInitializedEventHandler_GetInstance().FireEvent(this.EVENT)
        }
}


// throws Exception
func (this *GameInitializationUtil) InitGame(abeClientInformation AbeClientInformationInterface, canvas IMyCanvas, gameInitializationInterfaceFactoryInterface *BasicBuildGameInitializerFactory) {
        if ChangedGameFeatureListener_GetInstance().IsChangedFeature(MainFeatureFactory_GetInstance().STATIC) {
            var gameInitializationInterface GameInitializationInterface = gameInitializationInterfaceFactoryInterface.GetInstance(); _ = gameInitializationInterface
            var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
            gameInitializationInterface.Init(abeClientInformation, canvas.GetCustomCommandListener(), resourceLoadingLevelFactory.LOAD_ALL.GetLevel())
            this.EVENT.SetResourceLoadingLevel(resourceLoadingLevelFactory.LOAD_ALL)
            GameInitializedEventHandler_GetInstance().FireEvent(this.EVENT)
            ChangedGameFeatureListener_GetInstance().Remove(MainFeatureFactory_GetInstance().STATIC)
        }
}


func GameInitializationUtil_GetInstance() *GameInitializationUtil {
        return GameInitializationUtil_Instance
}

