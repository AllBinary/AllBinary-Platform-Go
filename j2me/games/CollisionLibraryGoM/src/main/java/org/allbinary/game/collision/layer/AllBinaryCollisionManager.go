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

type AllBinaryCollisionManager struct {
}

var AllBinaryCollisionManager_AllBinaryCollisionManager *AllBinaryCollisionManager = NewAllBinaryCollisionManager()


func NewAllBinaryCollisionManager() *AllBinaryCollisionManager {
    this := &AllBinaryCollisionManager{}
    return this
}


func AllBinaryCollisionManager_GetInstance() *AllBinaryCollisionManager {
        return AllBinaryCollisionManager_AllBinaryCollisionManager
}


// throws Exception
func (this *AllBinaryCollisionManager) Process(layerManager ILayerManager, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface, startIndex int) {
        var size int = layerManager.GetSize(); _ = size
        for index := int(startIndex + 1); index < size; index++ {
            var collidableInterfaceCompositeInterface2 CollidableInterfaceCompositeInterface = any(layerManager.GetLayerAt(index)).(CollidableInterfaceCompositeInterface); _ = collidableInterfaceCompositeInterface2
            this.Collide(layerManager, collidableInterfaceCompositeInterface, collidableInterfaceCompositeInterface2)
        }
}


// throws Exception
func (this *AllBinaryCollisionManager) Collide(layerManager ILayerManager, collidableInterfaceCompositeInterface CollidableInterfaceCompositeInterface, collidableInterfaceCompositeInterface2 CollidableInterfaceCompositeInterface) {
        var collidableBase2 ICollidableBaseBehavior = collidableInterfaceCompositeInterface2.GetCollidableInferface(); _ = collidableBase2
        var collidableBase ICollidableBaseBehavior = collidableInterfaceCompositeInterface.GetCollidableInferface(); _ = collidableBase
        if collidableBase2.IsCollidable(any(collidableInterfaceCompositeInterface).(ICollidableCompositeLayer)) {
            if collidableBase.IsCollisionInterface(any(collidableInterfaceCompositeInterface).(ICollidableCompositeLayer), collidableInterfaceCompositeInterface2) {
                collidableBase.CollideInterface(any(collidableInterfaceCompositeInterface).(ICollidableCompositeLayer), collidableInterfaceCompositeInterface2)
                collidableBase2.CollideInterface(any(collidableInterfaceCompositeInterface).(ICollidableCompositeLayer), collidableInterfaceCompositeInterface)
            }
        }
}

