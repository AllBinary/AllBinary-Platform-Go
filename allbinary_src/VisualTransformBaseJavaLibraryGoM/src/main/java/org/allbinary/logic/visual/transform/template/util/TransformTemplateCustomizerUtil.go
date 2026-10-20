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

type TransformTemplateCustomizerUtil struct {
    LogUtil *LogUtil
    CUSTOMIZER_NAME string
    BODY_VIEWNAMEKEY string
    TITLE_VIEWNAMEKEY string
    GLOBALS_NEW string
    GLOBALS_INSERT string
    GLOBALS_DELETE string
    GLOBALS_EDIT string
    GLOBALS_UPDATE string
}

var TransformTemplateCustomizerUtil_Instance *TransformTemplateCustomizerUtil = NewTransformTemplateCustomizerUtil()

func (this *TransformTemplateCustomizerUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    
}


func TransformTemplateCustomizerUtil_GetInstance() *TransformTemplateCustomizerUtil {
        return TransformTemplateCustomizerUtil_Instance
}


func NewTransformTemplateCustomizerUtil() *TransformTemplateCustomizerUtil {
    this := &TransformTemplateCustomizerUtil{}
    this.initFields()
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var SPACE string = CommonSeps_GetInstance().SPACE; _ = SPACE
        this.CUSTOMIZER_NAME = SPACE + fmt.Sprint(CustomizerTransformInfoData_NAME)
        this.BODY_VIEWNAMEKEY = SPACE + fmt.Sprint(BodyData_GetInstance().VIEWNAMEKEY)
        this.TITLE_VIEWNAMEKEY = SPACE + fmt.Sprint(TitleData_GetInstance().VIEWNAMEKEY)
        this.GLOBALS_NEW = fmt.Sprint(GLOBALS2.NEW) + SPACE
        this.GLOBALS_INSERT = fmt.Sprint(commonStrings.INSERT) + SPACE
        this.GLOBALS_DELETE = fmt.Sprint(commonStrings.DELETE) + SPACE
        this.GLOBALS_EDIT = fmt.Sprint(GLOBALS2.EDIT) + SPACE
        this.GLOBALS_UPDATE = fmt.Sprint(commonStrings.UPDATE) + SPACE
    return this
}


func (this *TransformTemplateCustomizerUtil) GetPageNameHack(viewName string, storeName string) string {
        var SPACE string = CommonSeps_GetInstance().SPACE; _ = SPACE
        var ONE_EMPTY_STRING_ARRAY []string = StringUtil_GetInstance().ONE_EMPTY_STRING_ARRAY; _ = ONE_EMPTY_STRING_ARRAY
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put([]string{viewName + SPACE}, ONE_EMPTY_STRING_ARRAY)
        hashMap.Put([]string{storeName + SPACE}, ONE_EMPTY_STRING_ARRAY)
        hashMap.Put([]string{this.CUSTOMIZER_NAME}, ONE_EMPTY_STRING_ARRAY)
        hashMap.Put([]string{this.BODY_VIEWNAMEKEY}, ONE_EMPTY_STRING_ARRAY)
        hashMap.Put([]string{this.TITLE_VIEWNAMEKEY}, ONE_EMPTY_STRING_ARRAY)
        hashMap.Put([]string{this.GLOBALS_NEW}, ONE_EMPTY_STRING_ARRAY)
        hashMap.Put([]string{this.GLOBALS_INSERT}, ONE_EMPTY_STRING_ARRAY)
        hashMap.Put([]string{this.GLOBALS_DELETE}, ONE_EMPTY_STRING_ARRAY)
        hashMap.Put([]string{this.GLOBALS_EDIT}, ONE_EMPTY_STRING_ARRAY)
        hashMap.Put([]string{this.GLOBALS_UPDATE}, ONE_EMPTY_STRING_ARRAY)
        var replace *Replace = nil; _ = replace
        var rootViewName string = replace.All(viewName); _ = rootViewName
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("RootViewName for ObjectConfig: " + rootViewName, this, "getPageNameHack")
        }
        return rootViewName
}

