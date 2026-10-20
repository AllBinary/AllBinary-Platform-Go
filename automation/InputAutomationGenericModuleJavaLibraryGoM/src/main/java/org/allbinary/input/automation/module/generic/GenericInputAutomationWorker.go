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

type GenericInputAutomationWorker struct {
    AbstractInputAutomationWorker
    LogUtil *LogUtil
    PROCESS_DATA_WORKER_RESULTS string
    LastFrame int64
    Frame int64
    GenericProfile *GenericProfile
}


func (this *GenericInputAutomationWorker) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.PROCESS_DATA_WORKER_RESULTS = "processDataWorkerResults"
    this.LastFrame = int64(-1)
    
    
}


func NewGenericInputAutomationWorker(inputAutomationActionInterface InputAutomationActionInterface, genericProfile *GenericProfile, motionRectangleConstraintsInterface MotionRectangleConstraintsInterface, imageComparatorConstraintsInterface ImageComparatorConstraintsInterface) *GenericInputAutomationWorker {
    this := &GenericInputAutomationWorker{}
        this.AbstractInputAutomationWorker = *NewAbstractInputAutomationWorker(inputAutomationActionInterface)
        this.initFields()
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        this.SetCaptureWorker(GenericProfileCaptureWorkerFactory_GetInstance(genericProfile))
        this.SetInputAutomationActionInterface(inputAutomationActionInterface)
        this.SetImageComparisonWorker(NewImageComparisonWorker(imageComparatorConstraintsInterface))
        this.SetMotionRectanglesWorker(NewMotionRectanglesWorker(motionRectangleConstraintsInterface))
        this.SetGenericProfile(genericProfile)
        var vector IBasicArrayList = this.GetGenericProfile().GetGenericProfileDataWorkerTypeVector(); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var genericProfileDataWorkerType IGenericProfileDataWorkerType = any(vector.Get(index)).(IGenericProfileDataWorkerType); _ = genericProfileDataWorkerType
            this.LogUtil.PutF("Adding Listener: " + fmt.Sprint(genericProfileDataWorkerType), this, this.CommonStrings.CONSTRUCTOR)
            if genericProfileDataWorkerType == GenericProfileDataWorkerType_COMPARISON {
                this.GetCaptureWorker().AddListenerInterface(this.GetImageComparisonWorker())
            } else if genericProfileDataWorkerType == GenericProfileDataWorkerType_MOTION {
                this.GetImageComparisonWorker().AddListenerInterface(this.GetMotionRectanglesWorker())
            }
        }
    return this
}


// throws Exception
func (this *GenericInputAutomationWorker) ProcessDataWorkerResults() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.PROCESS_DATA_WORKER_RESULTS)
        this.WaitForDataWorkers()
        var cacheInterface J2SECacheInterface = any(CapturedBufferedImagesCacheSingleton_GetInstance()).(J2SECacheInterface); _ = cacheInterface
        var keyArray []any = cacheInterface.KeySet().ToArray(); _ = keyArray
        if len(keyArray) > 0 {
            this.LogUtil.PutF("Image Available", this, this.PROCESS_DATA_WORKER_RESULTS)
            this.SetFrame(any(keyArray[len(keyArray) - 1]).(int64))
            if this.GetFrame() > this.LastFrame {
                this.LogUtil.PutF("Processing new frame: " + fmt.Sprint(this.GetFrame()), this, this.PROCESS_DATA_WORKER_RESULTS)
                var hashMap *HashMap = this.GetGenericProfile().GetGenericProfileActions().GetHashMap(); _ = hashMap
                var set Set = hashMap.KeySet(); _ = set
                this.LogUtil.PutF("Processing " + fmt.Sprint(set.Size()) + "Actions", this, this.PROCESS_DATA_WORKER_RESULTS)
                var actionNameArray []any = set.ToArray(); _ = actionNameArray
                var size int = len(actionNameArray); _ = size
                for index := int(0); index < size; index++ {
                    var actionNameString string = any(actionNameArray[index]).(string); _ = actionNameString
                    this.LogUtil.PutF("Processing Action: " + actionNameString, this, this.PROCESS_DATA_WORKER_RESULTS)
                    var genericProfileAction *GenericProfileAction = any(hashMap.Get(actionNameString)).(*GenericProfileAction); _ = genericProfileAction
                    var genericProfileActionScript *GenericProfileActionScript = genericProfileAction.GetGenericProfileActionScript(); _ = genericProfileActionScript
                    var vector IBasicArrayList = genericProfileActionScript.GetProfileActionConditionInterfaceVector(); _ = vector
                    CaptureWorkerUtil_ProcessProfileActionConditions(vector, this.GetFrame())
                }
                this.LastFrame = this.GetFrame()
            }
        } else {
            this.LogUtil.PutF("Image Not Available", this, this.PROCESS_DATA_WORKER_RESULTS)
        }
}


// throws Exception
func (this *GenericInputAutomationWorker) Process() {
        this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.PROCESS)
        this.StartDataWorkers()
        this.ProcessDataWorkerResults()
}


func (this *GenericInputAutomationWorker) GetGenericProfile() *GenericProfile {
        return this.GenericProfile
}


func (this *GenericInputAutomationWorker) SetGenericProfile(genericProfile *GenericProfile) {
        this.GenericProfile = genericProfile
}


func (this *GenericInputAutomationWorker) GetFrame() int64 {
        return this.Frame
}


func (this *GenericInputAutomationWorker) SetFrame(frame int64) {
        this.Frame = frame
}

