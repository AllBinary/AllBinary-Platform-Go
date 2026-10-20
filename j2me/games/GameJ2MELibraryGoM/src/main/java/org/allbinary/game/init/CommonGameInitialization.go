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

type CommonGameInitialization struct {
    BaseGameInitialization
    CommonStrings *CommonStrings
}


func (this *CommonGameInitialization) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewCommonGameInitialization(resourceInitializationArray []*ResourceInitialization, portion int) *CommonGameInitialization {
    this := &CommonGameInitialization{}
        this.BaseGameInitialization = *NewBaseGameInitialization(resourceInitializationArray, portion)
        this.initFields()
    return this
}


// throws Exception
func (this *CommonGameInitialization) Init(abeClientInformation AbeClientInformationInterface, commandListener CommandListener, level int) {
        this.BaseGameInitialization.Init(abeClientInformation, commandListener, level)
        var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
        if !this.IsGameInitialized() && level == resourceLoadingLevelFactory.LOAD_ALL.GetLevel() {
            this.SetGameInitialized(true)
            var progressCanvas IProgressCanvas = ProgressCanvasFactory_GetInstance(); _ = progressCanvas
            SmallIntegerSingletonFactory_GetInstance().Init()
            progressCanvas.AddNormalPortion(50, "Integers")
            PointFactory_GetInstance().Init()
            progressCanvas.AddNormalPortion(50, "Points")
            AngleFactory_GetInstance()
            progressCanvas.AddNormalPortion(50, "Angles")
            this.InitGame()
            this.ResourceInitializationArray[this.GAME_RESOURCES].Init()
        }
        this.BaseGameInitialization.ResourceInitialization(level)
}


// throws Exception
func (this *CommonGameInitialization) InitGame() {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}

