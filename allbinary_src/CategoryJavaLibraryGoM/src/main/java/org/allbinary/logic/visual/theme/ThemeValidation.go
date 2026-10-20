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

type ThemeValidation struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TransformInfoInterface TransformInfoInterface
    StyleValidationInterface ValidationInterface
    ThemeName string
    ThemePreviewImageName string
    WebAppAbPath IAbPath
    CategoryAbPath IAbPath
    FileAbPath IAbPath
}

// Implements ThemeInterface, ValidationInterface, DomNodeInterface, CompositeTransformInfoInterface


func (this *ThemeValidation) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    
}




// Java overloads of new ThemeValidation
func NewThemeValidation(args ...any) *ThemeValidation {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 3 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[ICssStyleValidation](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return NewThemeValidation_1(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[ICssStyleValidation](args[1]), Java_Arg[string](args[2]))
        }
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[*HashMap](args[1], pass) {
            return NewThemeValidation_2(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[*HashMap](args[1]))
        }
        if len(args) == 2 && Java_IsArg[StoreThemeCategoryInterface](args[0], pass) && Java_IsArg[*Node](args[1], pass) {
            return NewThemeValidation_3(Java_Arg[StoreThemeCategoryInterface](args[0]), Java_Arg[*Node](args[1]))
        }
    }
    panic("No Java overload of new ThemeValidation for the arguments")
}
func NewThemeValidation_1(transformInfoInterface TransformInfoInterface, cssStyleValidation ICssStyleValidation, categoryThemePath string) *ThemeValidation {
    this := &ThemeValidation{}
    this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        var pathData *AbPathData = AbPathData_GetInstance(); _ = pathData
        var pathUtil *PathUtil = PathUtil_GetInstance(); _ = pathUtil
        this.CategoryAbPath = pathUtil.RemoveNameFromPath(categoryThemePath)
        this.ThemeName = pathData.GetNameFromPath(categoryThemePath)
        this.StyleValidationInterface = cssStyleValidation
        this.Init()
    return this
}


func NewThemeValidation_2(transformInfoInterface TransformInfoInterface, hashMap *HashMap) *ThemeValidation {
    this := &ThemeValidation{}
    this.initFields()
        this.TransformInfoInterface = transformInfoInterface
        var categoryThemeAbPath IAbPath = NewAbPath(any(hashMap.Get(ThemeData_GetInstance().PATH)).(string), StringUtil_GetInstance().EMPTY_STRING); _ = categoryThemeAbPath
        var pathData *AbPathData = AbPathData_GetInstance(); _ = pathData
        var pathUtil *PathUtil = PathUtil_GetInstance(); _ = pathUtil
        this.CategoryAbPath = pathUtil.RemoveNameFromPath(fmt.Sprint(categoryThemeAbPath))
        this.ThemeName = pathData.GetNameFromPath(fmt.Sprint(categoryThemeAbPath))
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("CategoryThemePath: ")
            stringBuffer.Append(fmt.Sprint(categoryThemeAbPath))
            stringBuffer.Append(" & CategoryPath: ")
            stringBuffer.Append(fmt.Sprint(this.CategoryAbPath))
            stringBuffer.Append(" & ThemeName: ")
            stringBuffer.Append(this.ThemeName)
            this.LogUtil.PutF("Http Request Constructor", this, fmt.Sprint(stringBuffer))
        }
        this.Init()
    return this
}


func NewThemeValidation_3(storeThemeCategoryInterface StoreThemeCategoryInterface, node *Node) *ThemeValidation {
    this := &ThemeValidation{}
    this.initFields()
        this.TransformInfoInterface = storeThemeCategoryInterface.GetTransformInfoInterface()
        this.CategoryAbPath = storeThemeCategoryInterface.GetPath()
        var themeNameNode *Node = DomSearchHelper_GetNode(ThemeData_GetInstance().NAME, node.GetChildNodes()); _ = themeNameNode
        var valueNode *Node = DomSearchHelper_GetNode(DomData_VALUE, themeNameNode.GetChildNodes()); _ = valueNode
        this.ThemeName = DomNodeHelper_GetTextNodeValue(valueNode)
        var previewImageNameNode *Node = DomSearchHelper_GetNode(ThemeData_GetInstance().PREVIEW_IMAGE_NAME, node.GetChildNodes()); _ = previewImageNameNode
        var previewImageNameValueNode *Node = DomSearchHelper_GetNode(DomData_VALUE, previewImageNameNode.GetChildNodes()); _ = previewImageNameValueNode
        this.ThemePreviewImageName = DomNodeHelper_GetTextNodeValue(previewImageNameValueNode)
        this.WebAppAbPath = storeThemeCategoryInterface.GetWebAppPath()
        this.Init(storeThemeCategoryInterface)
    return this
}


// throws Exception
func (this *ThemeValidation) Init_1() {
        var categoryFactoryInterface CategoryFactoryInterface = NewStoreThemeCategoryFactory(this.GetTransformInfoInterface()); _ = categoryFactoryInterface
        var categoryLoaderInterface CategoryLoaderInterface = CategoryLoaderFactory_GetInstance(categoryFactoryInterface); _ = categoryLoaderInterface
        var rootStoreThemeCategoryInterface StoreThemeCategoryInterface = any(categoryFactoryInterface.GetRootInstance()).(StoreThemeCategoryInterface); _ = rootStoreThemeCategoryInterface
        this.Init(rootStoreThemeCategoryInterface)
}

// Java overloads of ThemeValidation.init
func (this *ThemeValidation) Init(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            this.Init_1()
            return
        }
        if len(args) == 1 && Java_IsArg[StoreThemeCategoryInterface](args[0], pass) {
            this.Init_2(Java_Arg[StoreThemeCategoryInterface](args[0]))
            return
        }
    }
    panic("No Java overload of ThemeValidation.init for the arguments")
}


// throws Exception
func (this *ThemeValidation) Init_2(storeThemeCategoryInterface StoreThemeCategoryInterface) {
        this.FileAbPath = NewAbPath(fmt.Sprint(storeThemeCategoryInterface.GetRootFilePath()) + fmt.Sprint(this.CategoryAbPath), StringUtil_GetInstance().EMPTY_STRING)
        this.WebAppAbPath = storeThemeCategoryInterface.GetWebAppPath()
}


// throws Exception
func (this *ThemeValidation) GetTransformInfoInterface() TransformInfoInterface {
        return this.TransformInfoInterface
}


func (this *ThemeValidation) GetName() string {
        return this.ThemeName
}


func (this *ThemeValidation) GetPreviewImageName() string {
        return this.ThemePreviewImageName
}


func (this *ThemeValidation) GetPreviewImagePath() string {
        return fmt.Sprint(this.WebAppAbPath) + fmt.Sprint(this.CategoryAbPath)
}


func (this *ThemeValidation) GetPath() string {
        return fmt.Sprint(this.CategoryAbPath)
}


// throws Exception
func (this *ThemeValidation) GetCssStyleValidation() ICssStyleValidation {
        if this.StyleValidationInterface == nil {
            var cssStyleFileAbPath IAbPath = NewAbPath(fmt.Sprint(this.FileAbPath), this.ThemeName + fmt.Sprint(AbPathData_GetInstance().EXTENSION_SEP) + fmt.Sprint(CategoryData_GetInstance().UNCRYPTED_EXTENSION)); _ = cssStyleFileAbPath
            var cryptFileReader *CryptFileReader = NewCryptFileReader(CategoryData_GetInstance().UNCRYPTED_EXTENSION, CategoryData_GetInstance().ENCRYPTED_EXTENSION); _ = cryptFileReader
            var document *Document = DomDocumentHelper_Create(cryptFileReader.Get(cssStyleFileAbPath)); _ = document
            this.StyleValidationInterface = NewCssStyleValidation(document)
        }
        return any(this.StyleValidationInterface).(ICssStyleValidation)
}


func (this *ThemeValidation) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to validate form", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var isValid bool = true; _ = isValid
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.IS_VALID)
            }
            if !StringValidationUtil_GetInstance().IsValidRequired(fmt.Sprint(this.CategoryAbPath), AbSqlData_MINSTRING, AbSqlData_MAXSTRING) {
                isValid = false
            }
            if !StringValidationUtil_GetInstance().IsValidRequired(this.ThemeName, AbSqlData_MINSTRING, AbSqlData_MAXSTRING) {
                isValid = false
            }
            if !this.GetCssStyleValidation().IsValid() {
                isValid = false
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("End: " + fmt.Sprint(isValid), this, this.CommonStrings.IS_VALID)
            }
            tryControl1, tryResult1 = 1, isValid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *ThemeValidation) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Validating Form"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append("Theme Validation Error")
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *ThemeValidation) ToValidationInfoDoc() *Document {
        return nil
}


func (this *ThemeValidation) ToValidationInfoNode(document *Document) *Node {
        return nil
}


func (this *ThemeValidation) GetKey() any {
        return any(this.GetName()).(any)
}


func (this *ThemeValidation) ToVector() IBasicArrayList {
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        vector.Add(this.GetName())
        vector.Add(this.GetPath())
        vector.Add(this.GetPreviewImageName())
        vector.Add(this.GetPreviewImagePath())
        return vector
}


func (this *ThemeValidation) ToHashMap() *HashMap {
        var themeData *ThemeData = ThemeData_GetInstance(); _ = themeData
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(themeData.NAME, this.GetName())
        hashMap.Put(themeData.PATH, this.GetPath())
        hashMap.Put(themeData.PREVIEW_IMAGE_NAME, this.GetPreviewImageName())
        hashMap.Put(themeData.PREVIEW_IMAGE_PATH, this.GetPreviewImagePath())
        return hashMap
}


// throws Exception
func (this *ThemeValidation) ToXmlNode(document *Document) *Node {
        var node *Node = ModDomHelper_CreateNameValueNodes(document, ThemeData_GetInstance().NAME, this.ToHashMap()); _ = node
        var domNodeInterface DomNodeInterface = any(this.GetCssStyleValidation()).(DomNodeInterface); _ = domNodeInterface
        node.AppendChild(domNodeInterface.ToXmlNode(document))
        return node
}

