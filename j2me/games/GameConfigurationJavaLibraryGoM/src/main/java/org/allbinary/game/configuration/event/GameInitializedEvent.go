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

type GameInitializedEvent struct {
    AllBinaryEventObject
    ResourceLoadingLevel *ResourceLoadingLevel
    Level int
}


func (this *GameInitializedEvent) initFields() {
    this.ResourceLoadingLevel = ResourceLoadingLevelFactory_GetInstance().LEVEL
    
}


func NewGameInitializedEvent(anyType any) *GameInitializedEvent {
    this := &GameInitializedEvent{}
        this.AllBinaryEventObject = *NewAllBinaryEventObject(anyType)
        this.initFields()
    return this
}


func (this *GameInitializedEvent) SetLevel(level int) {
        this.Level = level
        this.ResourceLoadingLevel = ResourceLoadingLevelFactory_GetInstance().LEVEL
}


func (this *GameInitializedEvent) GetLevel() int {
        return this.Level
}


func (this *GameInitializedEvent) SetResourceLoadingLevel(resourceLoadingLevel *ResourceLoadingLevel) {
        this.ResourceLoadingLevel = resourceLoadingLevel
        this.Level = this.ResourceLoadingLevel.GetLevel()
}


func (this *GameInitializedEvent) GetResourceLoadingLevel() *ResourceLoadingLevel {
        return this.ResourceLoadingLevel
}

