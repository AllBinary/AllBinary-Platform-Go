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

type OptimizedAllBinaryCollisionLayerProcessorForCollidableLayer struct {
    LayerProcessor
}



func NewOptimizedAllBinaryCollisionLayerProcessorForCollidableLayer() *OptimizedAllBinaryCollisionLayerProcessorForCollidableLayer {
    this := &OptimizedAllBinaryCollisionLayerProcessorForCollidableLayer{}
        this.LayerProcessor = *NewLayerProcessor(NewOptimizedAllBinaryCollisionLayerManager())
    return this
}


// throws Exception
func (this *OptimizedAllBinaryCollisionLayerProcessorForCollidableLayer) ProcessAt(allBinaryLayerManager IAllBinaryLayerManager, layerInterface IAllBinaryLayer, startIndex int) {
        var collidableInterfaceCompositeInterface ICollidableCompositeLayer = any(layerInterface).(ICollidableCompositeLayer); _ = collidableInterfaceCompositeInterface
        var collidableBase ICollidableBaseBehavior = collidableInterfaceCompositeInterface.GetCollidableInferface(); _ = collidableBase
        if collidableBase.IsCollidable(collidableInterfaceCompositeInterface) {
            var layerManager ILayerManager = this.GetLayerManager(); _ = layerManager
            var size int = layerManager.GetSize(); _ = size
            var collidableInterfaceCompositeInterface2 ICollidableCompositeLayer; _ = collidableInterfaceCompositeInterface2
            var collidableBase2 ICollidableBaseBehavior; _ = collidableBase2
            for index := int(startIndex + 1); index < size; index++ {
                collidableInterfaceCompositeInterface2 = any(layerManager.GetLayerAt(index)).(ICollidableCompositeLayer)
                collidableBase2 = collidableInterfaceCompositeInterface2.GetCollidableInferface()
                if collidableBase2.IsCollidable(collidableInterfaceCompositeInterface) && collidableBase.IsCollision(collidableInterfaceCompositeInterface, collidableInterfaceCompositeInterface2) {
                    collidableBase.Collide(collidableInterfaceCompositeInterface, collidableInterfaceCompositeInterface2)
                    collidableBase2.Collide(collidableInterfaceCompositeInterface2, collidableInterfaceCompositeInterface)
                }
            }
        }
}


func (this *OptimizedAllBinaryCollisionLayerProcessorForCollidableLayer) IsProcessorLayer(layerInterface IAllBinaryLayer) bool {
        if layerInterface.ImplmentsCollidableInterface() {
            return true
        } else {
            return false
        }
}

