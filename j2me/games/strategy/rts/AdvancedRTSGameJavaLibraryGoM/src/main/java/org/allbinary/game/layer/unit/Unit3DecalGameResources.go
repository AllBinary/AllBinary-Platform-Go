/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type Unit3DecalGameResources struct {
    BasicGameResources
}

var Unit3DecalGameResources_Instance *Unit3DecalGameResources = NewUnit3DecalGameResources()


func NewUnit3DecalGameResources() *Unit3DecalGameResources {
    this := &Unit3DecalGameResources{}
    this.BasicGameResources = *NewBasicGameResources()
        var ROOT string = "/unit_decal"; _ = ROOT
        var SMALL string = "_40_by_40.png"; _ = SMALL
        var MEDIUM string = SMALL; _ = MEDIUM
        var SIZE_FOUR string = SMALL; _ = SIZE_FOUR
        var SIZE_FIVE string = SMALL; _ = SIZE_FIVE
        var SIZE_SIX string = SMALL; _ = SIZE_SIX
        var SIZE []string = []string{SMALL, MEDIUM, SIZE_FOUR, SIZE_FIVE, SIZE_SIX}; _ = SIZE
        this.BasicGameResources.Init(ROOT, SIZE)
    return this
}


// throws Exception
func (this *Unit3DecalGameResources) Append(ROOT string, sizeString string) {
        this.BasicGameResources.Append(ROOT, sizeString)
}


// throws Exception
func (this *Unit3DecalGameResources) GetString() string {
        if AndroidUtil_IsAndroid() {
            return GameGraphicsResourceUtil_GetInstance().GetStringForFeature(GraphicsFeatureFactory_GetInstance().SPRITE_QUARTER_ROTATION_GRAPHICS)
        } else {
            return this.BasicGameResources.GetString()
        }
}


func Unit3DecalGameResources_GetInstance() IBasicGameResources {
        return Unit3DecalGameResources_Instance
}

