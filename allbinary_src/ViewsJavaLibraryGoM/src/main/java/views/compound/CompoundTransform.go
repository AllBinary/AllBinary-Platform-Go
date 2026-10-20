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

type CompoundTransform struct {
    AbTransformer
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
}


func (this *CompoundTransform) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
}


func NewCompoundTransform(transformInfoInterface TransformInfoInterface) *CompoundTransform {
    this := &CompoundTransform{}
        this.AbTransformer = *NewAbTransformer(transformInfoInterface)
        this.initFields()
        this.SetURIResolver(any(NewStoreUriResolver(this.GetTransformInfoInterface(), any(this.GetURIResolver()).(*BasicUriResolver))).(*URIResolver))
        var objectConfig *CompoundContextTransformInfoObjectConfig = NewCompoundContextTransformInfoObjectConfig(this.AbeClientInformation, this.GetTransformInfoInterface(), this.GetTransformInfoInterface().GetObjectConfigInterface().ToXmlDoc()); _ = objectConfig
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("\nObjectConfig: \n" + fmt.Sprint(objectConfig), this, this.CommonStrings.CONSTRUCTOR)
        }
        var templateInputStream *InputStream = objectConfig.CreateInputStream(); _ = templateInputStream
        this.SetInputStream(templateInputStream)
    return this
}

