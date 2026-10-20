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

type WeaponLayerArrayLayerCircularStaticPool struct {
    MAX int
    CircularIndexUtil *CircularIndexUtil
    ALL_WEAPONLAYER_ARRAY [][][]any
}

var WeaponLayerArrayLayerCircularStaticPool_Instance *WeaponLayerArrayLayerCircularStaticPool = NewWeaponLayerArrayLayerCircularStaticPool()

func (this *WeaponLayerArrayLayerCircularStaticPool) initFields() {
    this.MAX = 5
    this.CircularIndexUtil = CircularIndexUtil_CreateInstance(this.MAX)
    this.ALL_WEAPONLAYER_ARRAY = make([][][]any, 4)
}

func NewWeaponLayerArrayLayerCircularStaticPool() *WeaponLayerArrayLayerCircularStaticPool {
    this := &WeaponLayerArrayLayerCircularStaticPool{}
    this.initFields()
    return this
}


func WeaponLayerArrayLayerCircularStaticPool_GetInstance() *WeaponLayerArrayLayerCircularStaticPool {
        return WeaponLayerArrayLayerCircularStaticPool_Instance
}


func (this *WeaponLayerArrayLayerCircularStaticPool) Init() {
        var WEAPONLAYER_ARRAY [][]any = make([][]any, this.MAX * 2); _ = WEAPONLAYER_ARRAY
        var TWO_WEAPONLAYER_ARRAY [][]any = make([][]any, this.MAX); _ = TWO_WEAPONLAYER_ARRAY
        var THREE_WEAPONLAYER_ARRAY [][]any = make([][]any, this.MAX); _ = THREE_WEAPONLAYER_ARRAY
        for index := int(0); index < this.MAX; index++ {
            WEAPONLAYER_ARRAY[index] = make([]any, 1)
            TWO_WEAPONLAYER_ARRAY[index] = make([]any, 2)
            THREE_WEAPONLAYER_ARRAY[index] = make([]any, 3)
        }
        this.ALL_WEAPONLAYER_ARRAY[1] = WEAPONLAYER_ARRAY
        this.ALL_WEAPONLAYER_ARRAY[2] = TWO_WEAPONLAYER_ARRAY
        this.ALL_WEAPONLAYER_ARRAY[3] = THREE_WEAPONLAYER_ARRAY
}


// throws Exception
// synchronized
func (this *WeaponLayerArrayLayerCircularStaticPool) GetInstanceArray(size int) []IWeaponLayer {
        var weaponLayerArray []IWeaponLayer = any(this.ALL_WEAPONLAYER_ARRAY[size][this.CircularIndexUtil.GetIndex()]).([]IWeaponLayer); _ = weaponLayerArray
        this.CircularIndexUtil.Next()
        return weaponLayerArray
}

