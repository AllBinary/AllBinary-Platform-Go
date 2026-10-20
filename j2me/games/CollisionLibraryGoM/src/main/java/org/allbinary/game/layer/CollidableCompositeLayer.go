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

type CollidableCompositeLayer struct {
    AllBinaryGameLayer
    CollidableInferface ICollidableBaseBehavior
}

// Implements CollidableInterfaceCompositeInterface

// The Java class CollidableCompositeLayer has subclasses so its variables are this interface
type ICollidableCompositeLayer interface {
    IAllBinaryGameLayer
    CollidableCompositeLayer_() *CollidableCompositeLayer
    GetArtificialIntelligenceInterface() ArtificialIntelligenceInterface
    GetCollidableInferface() ICollidableBaseBehavior
    SetCollidableInferface(collidableInferface ICollidableBaseBehavior)
    ImplmentsCollidableInterface() bool
    ToStringAppend(stringBuffer *StringMaker)
    ToString() string
}

func (this *CollidableCompositeLayer) CollidableCompositeLayer_() *CollidableCompositeLayer {
    return this
}

var CollidableCompositeLayer_NULL_COLLIDABLE_COMPOSITE_LAYER any = NullUtil_GetInstance().NULL_OBJECT

func (this *CollidableCompositeLayer) initFields() {
    this.CollidableInferface = CollidableNeverCollideBehaviorFactory_GetInstance().CreateBehavior()
}


func CollidableCompositeLayer_GetNullInstance() ICollidableCompositeLayer {
        if CollidableCompositeLayer_NULL_COLLIDABLE_COMPOSITE_LAYER == NullUtil_GetInstance().NULL_OBJECT {
            CollidableCompositeLayer_NULL_COLLIDABLE_COMPOSITE_LAYER = NewCollidableCompositeLayer(StringUtil_GetInstance().EMPTY_STRING, RectangleFactory_SINGLETON, ViewPositionBase_NULL_VIEW_POSITION, CollidableNeverCollideBehaviorFactory_GetInstance())
        }
        return any(CollidableCompositeLayer_NULL_COLLIDABLE_COMPOSITE_LAYER).(ICollidableCompositeLayer)
}


func NewCollidableCompositeLayer(name string, layerInfo *Rectangle, viewPosition IViewPositionBase, collidableBaseBehaviorFactory ICollidableBaseBehaviorFactory) *CollidableCompositeLayer {
    this := &CollidableCompositeLayer{}
        this.AllBinaryGameLayer = *NewAllBinaryGameLayer(name, layerInfo, viewPosition)
        this.initFields()
        if collidableBaseBehaviorFactory != CollidableNeverCollideBehaviorFactory_GetInstance() {
            this.SetCollidableInferface(collidableBaseBehaviorFactory.CreateBehavior())
        }
    return this
}


func (this *CollidableCompositeLayer) GetArtificialIntelligenceInterface() ArtificialIntelligenceInterface {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, this)
        return ArtificialIntelligence_GetInstance()
}


func (this *CollidableCompositeLayer) GetCollidableInferface() ICollidableBaseBehavior {
        return this.CollidableInferface
}


func (this *CollidableCompositeLayer) SetCollidableInferface(collidableInferface ICollidableBaseBehavior) {
        this.CollidableInferface = collidableInferface
}


func (this *CollidableCompositeLayer) ImplmentsCollidableInterface() bool {
        return true
}


func (this *CollidableCompositeLayer) ToStringAppend(stringBuffer *StringMaker) {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        this.AllBinaryGameLayer.ToStringAppend(stringBuffer)
        stringBuffer.Append(commonSeps.NEW_LINE)
        stringBuffer.Append(fmt.Sprint(this.GetCollidableInferface()))
}


func (this *CollidableCompositeLayer) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        this.ToStringAppend(stringBuffer)
        return fmt.Sprint(stringBuffer)
}

func (this *CollidableCompositeLayer) String() string {
    return this.ToString()
}

