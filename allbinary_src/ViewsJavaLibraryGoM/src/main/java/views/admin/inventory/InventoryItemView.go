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

type InventoryItemView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    Request *HttpServletRequest
    ImageFileName string
    MediaData *MediaData
    ItemInterface ItemInterface
    DownloadableItemVector IBasicArrayList
    RequestHashMap *HashMap
}

// Implements RequestMapInterface

// The Java class InventoryItemView has subclasses so its variables are this interface
type IInventoryItemView interface {
    IHttpStoreComponentView
    InventoryItemView_() *InventoryItemView
    GetTypeId() int
    GetFormData()
    AddDomNodeInterfaces()
    GetItemInterface() ItemInterface
    View() string
    ProcessImageFiles()
    SetRequestHashMap(requestHashMap *HashMap)
    GetRequestHashMap() *HashMap
}

func (this *InventoryItemView) InventoryItemView_() *InventoryItemView {
    return this
}

var InventoryItemView_TYPE_ID int = 10

func (this *InventoryItemView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
    
    
    
    
}




// Java overloads of new InventoryItemView
func NewInventoryItemView(args ...any) *InventoryItemView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[TransformInfoInterface](args[0], pass) {
            return NewInventoryItemView_1(Java_Arg[TransformInfoInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return NewInventoryItemView_2(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of new InventoryItemView for the arguments")
}
func NewInventoryItemView_1(transformInfoInterface TransformInfoInterface) *InventoryItemView {
    this := &InventoryItemView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.Request = any(this.GetPageContext().GetRequest()).(*HttpServletRequest)
        this.GetFormData()
    return this
}


func NewInventoryItemView_2(transformInfoInterface TransformInfoInterface, empty string) *InventoryItemView {
    this := &InventoryItemView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.Request = any(this.GetPageContext().GetRequest()).(*HttpServletRequest)
    return this
}


func (this *InventoryItemView) GetTypeId() int {
        return InventoryItemView_TYPE_ID
}


// throws Exception
func (this *InventoryItemView) GetFormData() {
        this.SetRequestHashMap(NewMultipartRequestParams(this.Request).ToHashMap())
        var imageFileItemObject any = this.GetRequestHashMap().Get(this.BasicItemData.IMAGE); _ = imageFileItemObject
        if HttpFileUploadUtil_GetInstance().IsValid(imageFileItemObject) {
            var fileItem *FileItem = any(imageFileItemObject).(*FileItem); _ = fileItem
            if fileItem != nil && fileItem.GetSize() > 1 {
                this.ImageFileName = HttpRequestUtil_GetInstance().GenerateFileName(fileItem.GetName())
                var pathData *AbPathData = AbPathData_GetInstance(); _ = pathData
                var pathUtil *PathUtil = PathUtil_GetInstance(); _ = pathUtil
                this.MediaData = MediaData_Get(pathData.GetExtension(this.ImageFileName))
                this.ImageFileName = pathUtil.GetWithoutExtension(this.ImageFileName)
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("Uploaded File Data: ")
                    stringBuffer.Append(this.ImageFileName)
                    stringBuffer.Append(" Extension: ")
                    stringBuffer.Append(this.MediaData.GetName())
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "getFormData()")
                }
            }
        }
        this.ItemInterface = any(NewBasicItem(this.GetRequestHashMap())).(ItemInterface)
}


func (this *InventoryItemView) AddDomNodeInterfaces() {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        var downloadableItem *DownloadableItem; _ = downloadableItem
        var size int = this.DownloadableItemVector.Size(); _ = size
        for index := int(0); index < size; index++ {
            downloadableItem = any(this.DownloadableItemVector.Get(index)).(*DownloadableItem)
            vector.Add(NewDownloadableItemView(downloadableItem))
        }
        this.AddDomNodeInterface(NewBasicItemView(this.ItemInterface, vector))
}


func (this *InventoryItemView) GetItemInterface() ItemInterface {
        return this.ItemInterface
}


// throws Exception
func (this *InventoryItemView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            this.AddDomNodeInterfaces()
            tryControl1, tryResult1 = 1, this.HttpStoreComponentView.View(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *InventoryItemView) ProcessImageFiles() {
        var set Set = this.GetRequestHashMap().KeySet(); _ = set
        var fieldNameArray []any = set.ToArray(); _ = fieldNameArray
        var size int = len(fieldNameArray); _ = size
        for index := int(0); index < size; index++ {
            var fieldName string = any(fieldNameArray[index]).(string); _ = fieldName
            if strings.Compare(fieldName, this.BasicItemData.IMAGE) == 0 {
                var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetWeblisketSession().GetStoreName()); _ = storeFrontInterface
                var inventoryUploadMediaUtil *InventoryUploadMediaUtil = NewInventoryUploadMediaUtil(storeFrontInterface, this.ItemInterface); _ = inventoryUploadMediaUtil
                var fileItem *FileItem = any(this.GetRequestHashMap().Get(this.BasicItemData.IMAGE)).(*FileItem); _ = fileItem
                this.ItemInterface = inventoryUploadMediaUtil.SaveFiles(fileItem.Get(), this.ImageFileName, this.MediaData)
            }
        }
}


func (this *InventoryItemView) SetRequestHashMap(requestHashMap *HashMap) {
        this.RequestHashMap = requestHashMap
}


func (this *InventoryItemView) GetRequestHashMap() *HashMap {
        return this.RequestHashMap
}

