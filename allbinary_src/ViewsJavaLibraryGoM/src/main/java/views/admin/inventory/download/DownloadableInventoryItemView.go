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

type DownloadableInventoryItemView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    BasicItemData *BasicItemData
    Request *HttpServletRequest
    Id string
    ItemInterface ItemInterface
    DownloadableItem *DownloadableItem
    RequestHashMap *HashMap
}

// Implements RequestMapInterface

// The Java class DownloadableInventoryItemView has subclasses so its variables are this interface
type IDownloadableInventoryItemView interface {
    IHttpStoreComponentView
    DownloadableInventoryItemView_() *DownloadableInventoryItemView
    GetTypeId() int
    GetFormData()
    AddDomNodeInterfaces()
    View() string
    ProcessFile(fileItem *FileItem)
    Unzip(fileItem *FileItem)
    GetItemFilePath() string
    SetRequestHashMap(requestHashMap *HashMap)
    GetRequestHashMap() *HashMap
    GetDownloadableItem() *DownloadableItem
}

func (this *DownloadableInventoryItemView) DownloadableInventoryItemView_() *DownloadableInventoryItemView {
    return this
}

var DownloadableInventoryItemView_TYPE_ID int = 11

func (this *DownloadableInventoryItemView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.BasicItemData = BasicItemData_GetInstance()
    
    
    
    
    
}




// Java overloads of new DownloadableInventoryItemView
func NewDownloadableInventoryItemView(args ...any) *DownloadableInventoryItemView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[TransformInfoInterface](args[0], pass) {
            return NewDownloadableInventoryItemView_1(Java_Arg[TransformInfoInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return NewDownloadableInventoryItemView_2(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of new DownloadableInventoryItemView for the arguments")
}
func NewDownloadableInventoryItemView_1(transformInfoInterface TransformInfoInterface) *DownloadableInventoryItemView {
    this := &DownloadableInventoryItemView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.Request = any(this.GetPageContext().GetRequest()).(*HttpServletRequest)
        this.GetFormData()
    return this
}


func NewDownloadableInventoryItemView_2(transformInfoInterface TransformInfoInterface, empty string) *DownloadableInventoryItemView {
    this := &DownloadableInventoryItemView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        this.Request = any(this.GetPageContext().GetRequest()).(*HttpServletRequest)
    return this
}


func (this *DownloadableInventoryItemView) GetTypeId() int {
        return DownloadableInventoryItemView_TYPE_ID
}


// throws Exception
func (this *DownloadableInventoryItemView) GetFormData() {
        this.SetRequestHashMap(NewMultipartRequestParams(this.Request).ToHashMap())
        this.Id = any(this.GetRequestHashMap().Get(this.BasicItemData.ID)).(string)
}


func (this *DownloadableInventoryItemView) AddDomNodeInterfaces() {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        if this.DownloadableItem != nil {
            vector.Add(NewDownloadableItemView(this.DownloadableItem))
        }
        this.AddDomNodeInterface(NewBasicItemView(this.ItemInterface, vector))
}


// throws Exception
func (this *DownloadableInventoryItemView) View() string {
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
func (this *DownloadableInventoryItemView) ProcessFile(fileItem *FileItem) {
        var fileName string = fileItem.GetName(); _ = fileName
        var fullPath string = this.GetItemFilePath(); _ = fullPath
        var itemResourceFile *AbFile = AbFile_CreateAbFile(fullPath); _ = itemResourceFile
        itemResourceFile.Mkdir()
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(fullPath)
        stringBuffer.Append(fileName)
        this.LogUtil.PutF("FileName: " + fileName, this, "processFile()")
        var file *AbFile = AbFile_CreateAbFile(fmt.Sprint(stringBuffer)); _ = file
        this.LogUtil.PutF(file.GetPath(), this, "processFiles()")
        file.CreateNewFile()
        var byteArray []int8 = fileItem.Get(); _ = byteArray
        FileUtil_GetInstance().Write(NewByteArrayInputStream(byteArray), file)
}


// throws Exception
func (this *DownloadableInventoryItemView) Unzip(fileItem *FileItem) {
        var fileName string = fileItem.GetName(); _ = fileName
        var fullPath string = this.GetItemFilePath(); _ = fullPath
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(fullPath)
        stringBuffer.Append(fileName)
        this.LogUtil.PutF("FileName: " + fileName, this, "unzip()")
        var file *AbFile = AbFile_CreateAbFile(fmt.Sprint(stringBuffer)); _ = file
        this.LogUtil.PutF(file.GetPath(), this, "unzip()")
        ZipFileUtil_GetInstance().Unzip(fullPath, file, fileName)
}


func (this *DownloadableInventoryItemView) GetItemFilePath() string {
        var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(this.GetWeblisketSession().GetStoreName()); _ = storeFrontInterface
        var filePathData *FilePathData = FilePathData_GetInstance(); _ = filePathData
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(HttpStoreComponentView_URLGLOBALS.GetWebappPath())
        stringBuffer.Append(storeFrontInterface.GetCurrentHostNamePath())
        stringBuffer.Append(this.ItemInterface.GetCategory())
        stringBuffer.Appendchar(filePathData.SEPARATORCHAR)
        stringBuffer.Append(this.ItemInterface.GetId())
        stringBuffer.Appendchar(filePathData.SEPARATORCHAR)
        stringBuffer.Append(this.DownloadableItem.GetId())
        stringBuffer.Appendchar(filePathData.SEPARATORCHAR)
        return fmt.Sprint(stringBuffer)
}


func (this *DownloadableInventoryItemView) SetRequestHashMap(requestHashMap *HashMap) {
        this.RequestHashMap = requestHashMap
}


func (this *DownloadableInventoryItemView) GetRequestHashMap() *HashMap {
        return this.RequestHashMap
}


func (this *DownloadableInventoryItemView) GetDownloadableItem() *DownloadableItem {
        return this.DownloadableItem
}

