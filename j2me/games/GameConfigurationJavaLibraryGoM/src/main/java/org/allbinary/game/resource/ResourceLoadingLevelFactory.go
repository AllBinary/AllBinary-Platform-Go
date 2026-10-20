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

type ResourceLoadingLevelFactory struct {
    LOAD_EARLY *ResourceLoadingLevel
    LOAD_TOUCH *ResourceLoadingLevel
    LOAD_GAME *ResourceLoadingLevel
    LEVEL *ResourceLoadingLevel
    MAX_LEVEL *ResourceLoadingLevel
    LOAD_ALL *ResourceLoadingLevel
    RESOURCE_LOADING_LEVEL_ARRAY []*ResourceLoadingLevel
}

var ResourceLoadingLevelFactory_Instance *ResourceLoadingLevelFactory = NewResourceLoadingLevelFactory()

func (this *ResourceLoadingLevelFactory) initFields() {
    this.LOAD_EARLY = NewResourceLoadingLevel("Load Early", Integer_MAX_VALUE)
    this.LOAD_TOUCH = NewResourceLoadingLevel("Load Touch", Integer_MAX_VALUE - 1)
    this.LOAD_GAME = NewResourceLoadingLevel("Load Game", Integer_MAX_VALUE - 2)
    this.LEVEL = NewResourceLoadingLevel("Level", Integer_MAX_VALUE - 99)
    this.MAX_LEVEL = NewResourceLoadingLevel("Max Level", Integer_MAX_VALUE - 100)
    this.LOAD_ALL = NewResourceLoadingLevel("Load All", -1)
    this.RESOURCE_LOADING_LEVEL_ARRAY = []*ResourceLoadingLevel{this.LOAD_ALL, this.MAX_LEVEL, this.LEVEL, this.LOAD_GAME, this.LOAD_TOUCH, this.LOAD_EARLY}
}

func NewResourceLoadingLevelFactory() *ResourceLoadingLevelFactory {
    this := &ResourceLoadingLevelFactory{}
    this.initFields()
    return this
}


func ResourceLoadingLevelFactory_GetInstance() *ResourceLoadingLevelFactory {
        return ResourceLoadingLevelFactory_Instance
}


func (this *ResourceLoadingLevelFactory) GetLevelString(level int) string {
        for index := int(0); index < len(this.RESOURCE_LOADING_LEVEL_ARRAY); index++ {
            if level == this.RESOURCE_LOADING_LEVEL_ARRAY[index].GetLevel() {
                return this.RESOURCE_LOADING_LEVEL_ARRAY[index].GetName()
            }
        }
        return fmt.Sprint(level)
}

