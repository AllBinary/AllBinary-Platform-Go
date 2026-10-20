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

type InventorySearchProductUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    ITEM_NOT_FOUND string
}

var InventorySearchProductUtil_Instance *InventorySearchProductUtil = NewInventorySearchProductUtil()

func (this *InventorySearchProductUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.ITEM_NOT_FOUND = "Item Not Found."
}

func NewInventorySearchProductUtil() *InventorySearchProductUtil {
    this := &InventorySearchProductUtil{}
    this.initFields()
    return this
}


func InventorySearchProductUtil_GetInstance() *InventorySearchProductUtil {
        return InventorySearchProductUtil_Instance
}


// throws Exception
func (this *InventorySearchProductUtil) GetProduct(abeClientInformation AbeClientInformationInterface, searchRequest *SearchRequest, product string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRODUCTSEARCHLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "getProduct", e)
            }
            panic(errors.New(fmt.Sprint("Failed to getProduct")))
                }
            }()
            var inventoryEntityInterface *InventoryEntity = InventoryEntityFactory_GetInstance().GetInventoryEntityInstance(); _ = inventoryEntityInterface
            var itemInterface ItemInterface = inventoryEntityInterface.GetItem(product); _ = itemInterface
            if itemInterface != nil {
                var viewDocumentInterface TransformDocumentInterface = TransformStoreDocumentFactory_GetInstance(searchRequest); _ = viewDocumentInterface
                viewDocumentInterface.GetBaseNode().AppendChild(NewBasicItemView(itemInterface, NewBasicArrayListD()).ToXmlNode(viewDocumentInterface.GetDoc()))
                var success string = XmlDocumentHelper_ToString(viewDocumentInterface.GetDoc()); _ = success
                var outputStr string = NewStoreTransformer(abeClientInformation, any(NewTransformInfoHttpSearch(searchRequest)).(TransformInfoInterface)).Translate(success); _ = outputStr
                tryControl1, tryResult1 = 1, outputStr; return
            } else {
                tryControl1, tryResult1 = 1, this.ITEM_NOT_FOUND; return
            }
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

