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

type GraphicItemFactory struct {
    GraphicItems *HashMap
    DEFAULT GraphicsItemInterfaceFactoryInterface
}

var GraphicItemFactory_Instance *GraphicItemFactory = NewGraphicItemFactory()

func (this *GraphicItemFactory) initFields() {
    
    this.DEFAULT = NewSelectionToolFactory()
}


func GraphicItemFactory_GetInstance() *GraphicItemFactory {
        return GraphicItemFactory_Instance
}


func NewGraphicItemFactory() *GraphicItemFactory {
    this := &GraphicItemFactory{}
    this.initFields()
        this.GraphicItems = StdUtil_GetInstance().CreateHashMap()
        this.GraphicItems.Put(LinesGraphicItem_GetStaticName(), NewLinesGraphicItemFactory())
    return this
}


func (this *GraphicItemFactory) GetInstance_1(itemName string) GraphicsItemInterfaceFactoryInterface {
        var graphicsItemInterfaceFactoryInterface GraphicsItemInterfaceFactoryInterface = any(this.GraphicItems.Get(itemName)).(GraphicsItemInterfaceFactoryInterface); _ = graphicsItemInterfaceFactoryInterface
        if graphicsItemInterfaceFactoryInterface == nil {
            graphicsItemInterfaceFactoryInterface = this.DEFAULT
        }
        return graphicsItemInterfaceFactoryInterface
}

// Java overloads of GraphicItemFactory.getInstance
func (this *GraphicItemFactory) GetInstance(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GetInstance_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[IBasicArrayList](args[0], pass) {
            return this.GetInstance_2(Java_Arg[IBasicArrayList](args[0]))
        }
    }
    panic("No Java overload of GraphicItemFactory.getInstance for the arguments")
}


// throws Exception
func (this *GraphicItemFactory) GetInstance_2(graphicItemNodeList IBasicArrayList) *HashMap {
        var numberOfItems int = graphicItemNodeList.Size(); _ = numberOfItems
        var graphicItemHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = graphicItemHashMap
        if LOGGING.Contains(LOGGING.MENUEVENT) {
        }
        for index := int(0); index < numberOfItems; index++ {
            var graphicItem GraphicItemInterface = GraphicItemBuilder_GetInstance(any(graphicItemNodeList.Get(index)).(*Node)); _ = graphicItem
            graphicItemHashMap.Put(graphicItem.GetTreeNode(), graphicItem)
        }
        return graphicItemHashMap
}

