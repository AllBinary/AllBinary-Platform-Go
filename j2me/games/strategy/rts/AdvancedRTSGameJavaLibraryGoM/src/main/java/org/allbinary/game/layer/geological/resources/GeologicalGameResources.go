/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
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

type GeologicalGameResources struct {
    BasicGameResources
}

var GeologicalGameResources_Instance *GeologicalGameResources = NewGeologicalGameResources()


func NewGeologicalGameResources() *GeologicalGameResources {
    this := &GeologicalGameResources{}
    this.BasicGameResources = *NewBasicGameResources()
        var ROOT string = "/resource"; _ = ROOT
        var SMALL string = "_20_by_20.png"; _ = SMALL
        var MEDIUM string = SMALL; _ = MEDIUM
        var SIZE_FOUR string = SMALL; _ = SIZE_FOUR
        var SIZE_FIVE string = SMALL; _ = SIZE_FIVE
        var SIZE_SIX string = SMALL; _ = SIZE_SIX
        var SIZE []string = []string{SMALL, MEDIUM, SIZE_FOUR, SIZE_FIVE, SIZE_SIX}; _ = SIZE
        this.BasicGameResources.Init(ROOT, SIZE)
    return this
}


// throws Exception
func (this *GeologicalGameResources) Append(ROOT string, sizeString string) {
        this.BasicGameResources.Append(ROOT, sizeString)
}


// throws Exception
func (this *GeologicalGameResources) GetString() string {
        if AndroidUtil_IsAndroid() {
            return GameGraphicsResourceUtil_GetInstance().GetStringForFeature(GraphicsFeatureFactory_GetInstance().SPRITE_QUARTER_ROTATION_GRAPHICS)
        } else {
            return this.BasicGameResources.GetString()
        }
}


func GeologicalGameResources_GetInstance() IBasicGameResources {
        return GeologicalGameResources_Instance
}

