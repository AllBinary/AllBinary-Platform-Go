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

type CategoryRequestHelper struct {
    ModifyTable
    LogUtil *LogUtil
    PageContext *PageContext
    HashMap *HashMap
    Request *HttpServletRequest
    CategoryLoaderInterface CategoryLoaderInterface
    CategoryInterface CategoryInterface
    ChildCategoryInterface CategoryInterface
    TransformInfoInterface TransformInfoInterface
}

var CategoryRequestHelper_CategoryRequest string = fmt.Sprint(NewStringMaker().Append("<").Append(CategoryData_GetInstance().REQUEST).Append(">"))

func (this *CategoryRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
}


func NewCategoryRequestHelper(hashMap *HashMap, pageContext *PageContext) *CategoryRequestHelper {
    this := &CategoryRequestHelper{}
    this.ModifyTable = *NewModifyTable()
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "CategoryRequestHelper()", e)
            }
                }
            }()
            this.Request = any(pageContext.GetRequest()).(*HttpServletRequest)
            this.PageContext = pageContext
            this.HashMap = hashMap
            this.GetXmlData()
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func (this *CategoryRequestHelper) GetXmlData() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getXmlData()", e)
            }
                }
            }()
            var map_ Map = this.Request.GetParameterMap(); _ = map_
            var categoryData *CategoryData = CategoryData_GetInstance(); _ = categoryData
            var keys Set = map_.KeySet(); _ = keys
            var keyArray []any = keys.ToArray(); _ = keyArray
            var size int = len(keyArray); _ = size
            if size > 0 {
                var xmlRequest string = any(keyArray[0]).(string); _ = xmlRequest
                var index int = 1; _ = index
                for index < size && !strings.HasPrefix(xmlRequest, CategoryRequestHelper_CategoryRequest) {
                    xmlRequest = any(keyArray[index]).(string)
                    if strings.HasPrefix(xmlRequest, CategoryRequestHelper_CategoryRequest) {
                        break
                    }
                }
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                    this.LogUtil.PutF(xmlRequest, this, "getXmlData()")
                }
                var document *Document = DomDocumentHelper_Create(xmlRequest); _ = document
                var requestNode *Node = DomSearchHelper_GetNode(categoryData.REQUEST, document.GetChildNodes()); _ = requestNode
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                    this.LogUtil.PutF(XmlDocumentHelper_ToString(document), this, "getXmlData()")
                }
                this.SetCategoryLoader(requestNode)
                var parentCategoryNode *Node = DomSearchHelper_GetNode(categoryData.PARENT, requestNode.GetChildNodes()); _ = parentCategoryNode
                var categoryNode *Node = DomSearchHelper_GetNode(categoryData.NAME, parentCategoryNode.GetChildNodes()); _ = categoryNode
                this.CategoryInterface = any(NewStoreCategoryFactory(this.TransformInfoInterface).GetRootInstanceFromNode(categoryNode)).(CategoryInterface)
                this.CategoryInterface = this.CategoryLoaderInterface.Get(this.CategoryInterface)
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                    this.CategoryInterface.Log()
                    this.LogUtil.PutF("Loaded Parent Category", this, "getXmlData()")
                }
                var childCategoryNode *Node = DomSearchHelper_GetNodeNoThrow(categoryData.NAME, requestNode.GetChildNodes()); _ = childCategoryNode
                if childCategoryNode != nil {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                        this.LogUtil.PutF("Loading Child Category", this, "getXmlData()")
                    }
                    this.ChildCategoryInterface = any(NewStoreCategoryFactory(this.TransformInfoInterface).GetInstance(childCategoryNode)).(CategoryInterface)
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                        this.ChildCategoryInterface.Log()
                        this.LogUtil.PutF("Loaded Child Category", this, "getXmlData()")
                    }
                } else {
                    this.ChildCategoryInterface = this.CategoryInterface
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *CategoryRequestHelper) SetCategoryLoader(requestNode *Node) {
        var storeNameNode *Node = DomSearchHelper_GetNode(StoreFrontData_GetInstance().NAME, requestNode.GetChildNodes()); _ = storeNameNode
        if storeNameNode != nil {
            var storeName string = DomNodeHelper_GetTextNodeValue(storeNameNode); _ = storeName
            if storeName != "" {
                var storeFrontInterface StoreFrontInterface = StoreFrontFactory_GetInstance(storeName); _ = storeFrontInterface
                this.TransformInfoInterface = any(NewTransformInfoBasic(storeFrontInterface, this.HashMap, this.PageContext)).(TransformInfoInterface)
                if this.TransformInfoInterface == nil {
                    panic(errors.New(fmt.Sprint("TransformInfo null")))
                }
                var storeCategoryFactory *StoreCategoryFactory = NewStoreCategoryFactory(this.TransformInfoInterface); _ = storeCategoryFactory
                this.CategoryLoaderInterface = CategoryLoaderFactory_GetInstance(storeCategoryFactory)
            } else {
                panic(errors.New(fmt.Sprint("Store Name Error: " + storeName)))
            }
        }
}


func (this *CategoryRequestHelper) Insert() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to add item to Category"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "insert()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Successfully Added the following to the Category table"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("inserting", this, "insert()")
            }
            this.CategoryLoaderInterface.Insert(this.CategoryInterface, this.ChildCategoryInterface)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "insert()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *CategoryRequestHelper) Delete() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e2 := Java_ToError(recovered)
                        _ = e2
                var error_ string = "Failed to remove category and show the path of the failed category"; _ = error_
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                    this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "delete()", e2)
                }
                tryControl2, tryResult2 = 1, error_; return
                    }
                }()
                var error_ string = "Failed to remove category: " + fmt.Sprint(this.ChildCategoryInterface.GetPath()); _ = error_
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                    this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "delete()", e)
                }
                tryControl2, tryResult2 = 1, error_; return
                return
            }(); tryControl2 == 1 {
                tryControl1, tryResult1 = 1, tryResult2; return
            }
                }
            }()
            var success string = "Successfully Removed " + fmt.Sprint(CategoryData_GetInstance().NAME) + "=" + fmt.Sprint(this.ChildCategoryInterface.GetPath()); _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF("Deleting", this, "delete()")
            }
            this.CategoryLoaderInterface.Delete(this.CategoryInterface, this.ChildCategoryInterface)
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "delete()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *CategoryRequestHelper) ViewCategory() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to get Category"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "viewCategory()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var document *Document = NewCategoryComponent(this.ChildCategoryInterface).ToXmlDoc(); _ = document
            var xmlString string = XmlDocumentHelper_ToString(document); _ = xmlString
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(xmlString, this, "viewCategory()")
            }
            tryControl1, tryResult1 = 1, xmlString; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *CategoryRequestHelper) ViewCategories() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to view Categories table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "viewCategories()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = StringUtil_GetInstance().EMPTY_STRING; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "viewCategories()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *CategoryRequestHelper) Update() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to add Item"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "update()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "New Item Successfully added"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "update()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

