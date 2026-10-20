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

type GameFeatureFactory struct {
    SOUND IGameFeature
    MUSIC IGameFeature
    SCREEN_SHAKE IFeature
    POST_IMAGE_LOADING_MODIFICATION IFeature
    HEALTH_BARS IGameFeature
    DAMAGE_FLOATERS IGameFeature
    GAME_INPUT_LAYER_PROCESSOR IGameFeature
    ARTIFICIAL_INTELLEGENCE_PROCESSOR IGameFeature
    COLLIDABLE_INTERFACE_LAYER_PROCESSOR IGameFeature
    TICKABLE_LAYER_PROCESSOR IGameFeature
    CHEATING IGameFeature
    DROPPED_ITEMS IGameFeature
    DROPPED_ITEMS_FROM_DEATH IGameFeature
    TEST_DESTROYED_LAYER_PROCESSOR IGameFeature
    COLLISIONS_WITH_SOURCE IGameFeature
    COLLISIONS_FORCED_TWO_DIMENSIONAL IGameFeature
}

var GameFeatureFactory_Instance *GameFeatureFactory = NewGameFeatureFactory()

func (this *GameFeatureFactory) initFields() {
    this.SOUND = NewGameFeature("Sound")
    this.MUSIC = NewGameFeature("Music")
    this.SCREEN_SHAKE = NewFeature("Screen Shake")
    this.POST_IMAGE_LOADING_MODIFICATION = NewFeature("Post Image Loading Modification (Scaling, Rotation, or Other)")
    this.HEALTH_BARS = NewGameFeature("Health Bars")
    this.DAMAGE_FLOATERS = NewGameFeature("Damage Floaters")
    this.GAME_INPUT_LAYER_PROCESSOR = NewGameFeature("Game Input Layer Processor")
    this.ARTIFICIAL_INTELLEGENCE_PROCESSOR = NewGameFeature("Artificial Intelligence Layer Processor")
    this.COLLIDABLE_INTERFACE_LAYER_PROCESSOR = NewGameFeature("Collidable Interface Layer Processor")
    this.TICKABLE_LAYER_PROCESSOR = NewGameFeature("Tickable Layer Processor")
    this.CHEATING = NewGameFeature("Cheating")
    this.DROPPED_ITEMS = NewGameFeature("Drops")
    this.DROPPED_ITEMS_FROM_DEATH = NewGameFeature("Drops from Dead")
    this.TEST_DESTROYED_LAYER_PROCESSOR = NewGameFeature("Destroyed Layer Processor")
    this.COLLISIONS_WITH_SOURCE = NewGameFeature("Allow Collisions With Source")
    this.COLLISIONS_FORCED_TWO_DIMENSIONAL = NewGameFeature("Collisions Forced Two Dimensional")
}


func GameFeatureFactory_GetInstance() *GameFeatureFactory {
        return GameFeatureFactory_Instance
}


func NewGameFeatureFactory() *GameFeatureFactory {
    this := &GameFeatureFactory{}
    this.initFields()
    return this
}

