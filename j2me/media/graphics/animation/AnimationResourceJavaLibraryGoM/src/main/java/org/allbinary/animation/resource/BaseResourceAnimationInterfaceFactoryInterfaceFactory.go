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

type BaseResourceAnimationInterfaceFactoryInterfaceFactory struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Hashtable *ABHashtable[any, any]
    RectangleHashtable *ABHashtable[any, any]
    RectangleArrayOfArraysHashtable *ABHashtable[any, any]
    Name string
    Initialized bool
}

// Implements FeatureResourceAnimationInterfaceFactoryInterface

// The Java class BaseResourceAnimationInterfaceFactoryInterfaceFactory has subclasses so its variables are this interface
type IBaseResourceAnimationInterfaceFactoryInterfaceFactory interface {
    BaseResourceAnimationInterfaceFactoryInterfaceFactory_() *BaseResourceAnimationInterfaceFactoryInterfaceFactory
    GetName() string
    Init(level int)
    InitImageCache(imageCache IImageCache, level int)
    Add(resource string, animationInterfaceFactoryInterface BasicAnimationInterfaceFactoryInterface)
    GetBasicAnimationInterfaceFactoryInstance(resource string) BasicAnimationInterfaceFactoryInterface
    GetRectangle(resource string) *Rectangle
    AddRectangle(resource string, rectangle *Rectangle)
    GetRectangleArrayOfArrays(resource string) [][]*Rectangle
    AddRectangleArrayOfArrays(resource string, rectangleArrayOfArrays [][]*Rectangle)
    IsFeature() bool
    IsLoadingLevel(level int) bool
    ToString() string
    GetHashtable() *ABHashtable[any, any]
    GetRectangleHashtable() *ABHashtable[any, any]
    GetRectangleArrayOfArraysHashtable() *ABHashtable[any, any]
    SetInitialized(initialized bool)
    IsInitialized() bool
}

func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) BaseResourceAnimationInterfaceFactoryInterfaceFactory_() *BaseResourceAnimationInterfaceFactoryInterfaceFactory {
    return this
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
}


func NewBaseResourceAnimationInterfaceFactoryInterfaceFactory(name string, hashtable *ABHashtable[any, any], rectangleHashtable *ABHashtable[any, any], rectangleArrayHashtable *ABHashtable[any, any]) *BaseResourceAnimationInterfaceFactoryInterfaceFactory {
    this := &BaseResourceAnimationInterfaceFactoryInterfaceFactory{}
    this.initFields()
        this.Hashtable = hashtable
        this.RectangleHashtable = rectangleHashtable
        this.RectangleArrayOfArraysHashtable = rectangleArrayHashtable
        this.Name = name
    return this
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) GetName() string {
        return this.Name
}


// throws Exception
func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) Init(level int) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Available List of Animations: ").Append(fmt.Sprint(this.Hashtable))), this, commonStrings.INIT)
        this.SetInitialized(true)
}


// throws Exception
func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) InitImageCache(imageCache IImageCache, level int) {
}


// throws Exception
func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) Add(resource string, animationInterfaceFactoryInterface BasicAnimationInterfaceFactoryInterface) {
        if this.Hashtable.ContainsKey(any(resource).(any)) {
            panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("Resource Already Created: ").Append(resource)))))
        }
        this.Hashtable.Put(resource, animationInterfaceFactoryInterface)
}


// throws Exception
func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) GetBasicAnimationInterfaceFactoryInstance(resource string) BasicAnimationInterfaceFactoryInterface {
        var basicAnimationInterfaceFactoryInterfaceCanBeNull any = this.Hashtable.Get(resource); _ = basicAnimationInterfaceFactoryInterfaceCanBeNull
        if basicAnimationInterfaceFactoryInterfaceCanBeNull == nil {
            return NullAnimationFactory_NULL_NOT_FOR_USE_ANIMATION_FACTORY
        }
        return any(basicAnimationInterfaceFactoryInterfaceCanBeNull).(BasicAnimationInterfaceFactoryInterface)
}


// throws Exception
func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) GetRectangle(resource string) *Rectangle {
        var rectangleCanBeNull any = this.RectangleHashtable.Get(resource); _ = rectangleCanBeNull
        if rectangleCanBeNull == nil {
            return RectangleFactory_SINGLETON
        }
        return any(rectangleCanBeNull).(*Rectangle)
}


// throws Exception
func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) AddRectangle(resource string, rectangle *Rectangle) {
        this.RectangleHashtable.Put(resource, rectangle)
}


// throws Exception
func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) GetRectangleArrayOfArrays(resource string) [][]*Rectangle {
        return any(this.RectangleArrayOfArraysHashtable.Get(resource)).([][]*Rectangle)
}


// throws Exception
func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) AddRectangleArrayOfArrays(resource string, rectangleArrayOfArrays [][]*Rectangle) {
        this.RectangleArrayOfArraysHashtable.Put(resource, rectangleArrayOfArrays)
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) IsFeature() bool {
        return false
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) IsLoadingLevel(level int) bool {
        var resourceLoadingLevelFactory *ResourceLoadingLevelFactory = ResourceLoadingLevelFactory_GetInstance(); _ = resourceLoadingLevelFactory
        if level == resourceLoadingLevelFactory.LOAD_ALL.GetLevel() {
            return true
        } else {
            return false
        }
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(reflect.TypeOf(this).String()).Append(CommonSeps_GetInstance().SEMICOLON).Append(CommonSeps_GetInstance().SPACE).Append(this.Name))
}

func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) String() string {
    return this.ToString()
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) GetHashtable() *ABHashtable[any, any] {
        return this.Hashtable
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) GetRectangleHashtable() *ABHashtable[any, any] {
        return this.RectangleHashtable
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) GetRectangleArrayOfArraysHashtable() *ABHashtable[any, any] {
        return this.RectangleArrayOfArraysHashtable
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) SetInitialized(initialized bool) {
        this.Initialized = initialized
}


func (this *BaseResourceAnimationInterfaceFactoryInterfaceFactory) IsInitialized() bool {
        return this.Initialized
}

