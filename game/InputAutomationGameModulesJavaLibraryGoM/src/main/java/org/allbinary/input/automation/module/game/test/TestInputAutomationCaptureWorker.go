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

type TestInputAutomationCaptureWorker struct {
    AbstractInputAutomationWorker
    LogUtil *LogUtil
    GenericProfileActions *GenericProfileActions
    MESSAGE string
    PROCESS_DATA_WORKER_RESULTS string
    IMAGE_AVAILABLE_SO_PROCESSING string
    IMAGE_NOT_AVAILABLE string
}


func (this *TestInputAutomationCaptureWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    this.MESSAGE = "GenericInputAutomationCaptureWorker"
    this.PROCESS_DATA_WORKER_RESULTS = "processDataWorkerResults"
    this.IMAGE_AVAILABLE_SO_PROCESSING = "Image Available so processing"
    this.IMAGE_NOT_AVAILABLE = "Image Not Available"
}


func NewTestInputAutomationCaptureWorker(inputAutomationActionInterface InputAutomationActionInterface, genericProfileActions *GenericProfileActions, imageComparatorConstraintsInterface ImageComparatorConstraintsInterface, motionRectangleConstraintsInterface MotionRectangleConstraintsInterface) *TestInputAutomationCaptureWorker {
    this := &TestInputAutomationCaptureWorker{}
        this.AbstractInputAutomationWorker = *NewAbstractInputAutomationWorker(inputAutomationActionInterface)
        this.initFields()
        this.LogUtil.PutF(this.MESSAGE, this, this.CommonStrings.CONSTRUCTOR)
        this.SetGenericProfileActions(genericProfileActions)
    return this
}


// throws Exception
func (this *TestInputAutomationCaptureWorker) ProcessDataWorkerResults() {
        var cacheInterface J2SECacheInterface = any(CapturedBufferedImagesCacheSingleton_GetInstance()).(J2SECacheInterface); _ = cacheInterface
        if cacheInterface.KeySet().Size() > 0 {
            this.LogUtil.PutF(this.IMAGE_AVAILABLE_SO_PROCESSING, this, this.PROCESS_DATA_WORKER_RESULTS)
            var anyType any = cacheInterface.KeySet().ToArray()[0]; _ = anyType
            var hashMap *HashMap = this.GetGenericProfileActions().GetHashMap(); _ = hashMap
            var set Set = hashMap.KeySet(); _ = set
            var actionNameArray []any = set.ToArray(); _ = actionNameArray
            var size int = len(actionNameArray); _ = size
            for index := int(0); index < size; index++ {
                var actionNameString string = any(actionNameArray[index]).(string); _ = actionNameString
                var genericProfileAction *GenericProfileAction = any(hashMap.Get(actionNameString)).(*GenericProfileAction); _ = genericProfileAction
                var genericProfileActionScript *GenericProfileActionScript = genericProfileAction.GetGenericProfileActionScript(); _ = genericProfileActionScript
                var vector IBasicArrayList = genericProfileActionScript.GetProfileActionConditionInterfaceVector(); _ = vector
            }
            cacheInterface.Get(anyType)
        } else {
            this.LogUtil.PutF(this.IMAGE_NOT_AVAILABLE, this, this.PROCESS_DATA_WORKER_RESULTS)
        }
}


// throws Exception
func (this *TestInputAutomationCaptureWorker) Process() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.PROCESS)
        this.StartDataWorkers()
        this.ProcessDataWorkerResults()
}


func (this *TestInputAutomationCaptureWorker) GetGenericProfileActions() *GenericProfileActions {
        return this.GenericProfileActions
}


func (this *TestInputAutomationCaptureWorker) SetGenericProfileActions(genericProfileActions *GenericProfileActions) {
        this.GenericProfileActions = genericProfileActions
}

