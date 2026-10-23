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

type StoreFileGenerator struct {
    TransformInfoObjectConfigGenerator
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Output string
    FileAbPath IAbPath
    File *AbFile
}


func (this *StoreFileGenerator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}


func NewStoreFileGenerator(transformInfoInterface TransformInfoInterface) *StoreFileGenerator {
    this := &StoreFileGenerator{}
    this.TransformInfoObjectConfigGenerator = *NewTransformInfoObjectConfigGenerator()
    this.initFields()
        var httpTransformInfoInterface TransformInfoHttpInterface = any(transformInfoInterface).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
        var weblisketSession *WeblisketSession = NewWeblisketSession(httpTransformInfoInterface.GetPropertiesHashMap(), httpTransformInfoInterface.GetPageContext()); _ = weblisketSession
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(TransformInfoObjectConfigGenerator_URLGLOBALS.GetWebappPath())
        stringBuffer.Append(weblisketSession.GetStoreName())
        stringBuffer.Append(AbPathData_GetInstance().SEPARATOR)
        var abPath IAbPath = NewAbPath(fmt.Sprint(stringBuffer), StringUtil_GetInstance().EMPTY_STRING); _ = abPath
        var fileName string = any(httpTransformInfoInterface.GetPropertiesHashMap().Get(InputOutputTypeData_GetInstance().FILE)).(string); _ = fileName
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        if stringValidationUtil.IsEmpty(fileName) {
            var transformInfoHttpStoreInterface TransformInfoHttpStoreInterface = any(transformInfoInterface).(TransformInfoHttpStoreInterface); _ = transformInfoHttpStoreInterface
            fileName = transformInfoInterface.GetName()[len(transformInfoHttpStoreInterface.GetStoreName()) + 1:]
            var outputMappingHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = outputMappingHashMap
            outputMappingHashMap.Put(TransformInfosData_GetInstance().PREVIEW, "Preview")
            outputMappingHashMap.Put(TransformInfosData_GetInstance().SMALL_PREVIEW, "SmallPreview")
            outputMappingHashMap.Put(CommonSeps_GetInstance().SPACE, StringUtil_GetInstance().EMPTY_STRING)
            var replace *Replace = NewReplace(outputMappingHashMap); _ = replace
            fileName = replace.All(fileName)
        }
        if stringValidationUtil.IsEmpty(fileName) {
            panic(errors.New(fmt.Sprint("TransformInfoObjectConfigGenerator FileName Not Specified")))
        }
        var tempOutput string = any(httpTransformInfoInterface.GetPropertiesHashMap().Get(InputOutputTypeData_GetInstance().NAME)).(string); _ = tempOutput
        if stringValidationUtil.IsEmpty(tempOutput) {
            tempOutput = InputOutputTypeData_GetInstance().DEFAULT
        }
        this.Output = tempOutput
        if stringValidationUtil.IsEmpty(this.Output) {
            panic(errors.New(fmt.Sprint("TransformInfoObjectConfigGenerator Output=extension Not Specified")))
        }
        stringBuffer = NewStringMaker()
        stringBuffer.Append(fileName)
        stringBuffer.Append(AbPathData_GetInstance().EXTENSION_SEP)
        stringBuffer.Append(this.Output)
        abPath = NewAbPath(fmt.Sprint(abPath), fmt.Sprint(stringBuffer))
        this.FileAbPath = abPath
        this.File = AbFile_CreateAbFileFromAbPath(this.FileAbPath)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("File: " + fmt.Sprint(this.FileAbPath), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


// throws Exception
func (this *StoreFileGenerator) Process(input string) string {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Processing", this, this.CommonStrings.PROCESS)
        }
        if !this.File.Exists() {
            this.File.CreateNewFile()
        }
        if this.File.Exists() {
            var inputStream IInputStream = NewByteArrayInputStream(Java_GetBytes(input)); _ = inputStream
            FileUtil_GetInstance().Write(inputStream, this.File)
        } else {
            panic(errors.New(fmt.Sprint("Could Not Create: " + fmt.Sprint(this.FileAbPath))))
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Processed", this, this.CommonStrings.PROCESS)
        }
        return StringUtil_GetInstance().EMPTY_STRING
}

