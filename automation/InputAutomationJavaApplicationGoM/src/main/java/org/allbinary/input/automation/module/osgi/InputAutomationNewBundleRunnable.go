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

type InputAutomationNewBundleRunnable struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    InputAutomationBundleActivator *InputAutomationBundleActivator
    Running bool
    FileBasicArrayList IBasicArrayList
}

// Implements RunnableInterface

var InputAutomationNewBundleRunnable_FILE string = "file:"
var InputAutomationNewBundleRunnable_JAR_DIR_PROP string = "org.knopflerfish.gosg.jars"
var InputAutomationNewBundleRunnable_INPUT_AUTMATION_MODULE_BUNDLE_JAR_PATH string = "ia/auto"
var InputAutomationNewBundleRunnable_MODULES_START_LEVEL int = 8

func (this *InputAutomationNewBundleRunnable) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}


func NewInputAutomationNewBundleRunnable(inputAutomationBundleActivator *InputAutomationBundleActivator) *InputAutomationNewBundleRunnable {
    this := &InputAutomationNewBundleRunnable{}
    this.initFields()
        this.InputAutomationBundleActivator = inputAutomationBundleActivator
        this.FileBasicArrayList = NewBasicArrayListD()
    return this
}


// throws Exception
func (this *InputAutomationNewBundleRunnable) SetThread(thread IThread) {
}


// synchronized
func (this *InputAutomationNewBundleRunnable) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *InputAutomationNewBundleRunnable) SetRunning(running bool) {
        this.Running = running
}


// throws Exception
func (this *InputAutomationNewBundleRunnable) UpdateModules() {
        this.LogUtil.PutF(this.CommonStrings.START, this, "updateModules")
        var list IBasicArrayList = this.FindNewModules(); _ = list
        var size int = list.Size(); _ = size
        var bundle Bundle; _ = bundle
        for index := int(0); index < size; index++ {
            bundle = this.Install(any(list.Get(index)).(*URL))
            if bundle != nil {
                bundle.Start(InputAutomationNewBundleRunnable_MODULES_START_LEVEL)
            }
        }
}


// throws Exception
func (this *InputAutomationNewBundleRunnable) GetAllJarSymbolicNameHashMap() *HashMap {
        this.LogUtil.PutF(this.CommonStrings.START, this, "getAllJarSymbolicNameHashMap")
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        var jarFileBasicArrayList IBasicArrayList = this.GetJarModuleFileBasicArrayList(); _ = jarFileBasicArrayList
        this.LogUtil.PutF("Jar Module Files: " + fmt.Sprint(jarFileBasicArrayList), this, "getAllJarSymbolicNameHashMap")
        var size int = jarFileBasicArrayList.Size(); _ = size
        var file *File; _ = file
        for index := int(0); index < size; index++ {
            file = any(jarFileBasicArrayList.Get(index)).(*File)
            if !File_IsDirectory() {
                var fileInputStream *FileInputStream = NewFileInputStream(file); _ = fileInputStream
                var jarInputStream *JarInputStream = NewJarInputStream(fileInputStream); _ = jarInputStream
                var manifest *Manifest = jarInputStream.GetManifest(); _ = manifest
                if manifest == nil {
                } else {
                    var symbolicName string = manifest.GetMainAttributes().GetValue(Constants_BUNDLE_SYMBOLICNAME); _ = symbolicName
                    if symbolicName != "" {
                        hashMap.Put(symbolicName, NewURL(InputAutomationNewBundleRunnable_FILE + file.GetAbsolutePath()))
                    }
                }
            }
        }
        return hashMap
}


func (this *InputAutomationNewBundleRunnable) GetJarModuleFileBasicArrayList() IBasicArrayList {
        this.LogUtil.PutF(this.CommonStrings.START, this, "getJarModuleFileBasicArrayList")
        var baseJarPath string = System_GetProperty(InputAutomationNewBundleRunnable_JAR_DIR_PROP); _ = baseJarPath
        if strings.HasPrefix(baseJarPath, InputAutomationNewBundleRunnable_FILE) {
            baseJarPath = baseJarPath[len(InputAutomationNewBundleRunnable_FILE):]
        }
        var jarFileFilter *FileFilter = BasicFileFilterUtil_GetInstance(".jar"); _ = jarFileFilter
        var path string = baseJarPath + fmt.Sprint(InputAutomationNewBundleRunnable_INPUT_AUTMATION_MODULE_BUNDLE_JAR_PATH); _ = path
        this.LogUtil.PutF("Path: " + path, this, "getJarModuleFileBasicArrayList")
        var file *File = NewFile(path); _ = file
        this.LogUtil.PutF("File: " + fmt.Sprint(file.GetAbsolutePath()) + " isDirectory: " + fmt.Sprint(File_IsDirectory()), this, "getJarModuleFileBasicArrayList")
        return SubDirectory_GetInstance().Search(jarFileFilter, FileWrapperUtil_WrapFile(file))
}


// throws Exception
func (this *InputAutomationNewBundleRunnable) GetInstalledJarSymbolicNameBasicArrayList() IBasicArrayList {
        this.LogUtil.PutF(this.CommonStrings.START, this, "getInstalledJarSymbolicNameBasicArrayList")
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        var bundleContext BundleContext = InputAutomationBundleActivator_GetBundleContext(); _ = bundleContext
        var bundleArray []Bundle = bundleContext.GetBundles(); _ = bundleArray
        if bundleArray != nil {
            this.LogUtil.PutF("bundleArray: " + fmt.Sprint(len(bundleArray)), this, "getInputAutomationModuleServices")
            for index := int(0); index < len(bundleArray); index++ {
                var bundle Bundle = bundleArray[index]; _ = bundle
                vector.Add(bundle.GetSymbolicName())
            }
        }
        return vector
}


// throws Exception
func (this *InputAutomationNewBundleRunnable) IsInstalled(symbolicName string) bool {
        this.LogUtil.PutF(fmt.Sprint(CommonLabels_GetInstance().START) + symbolicName, this, "isInstalled")
        var list IBasicArrayList = this.GetInstalledJarSymbolicNameBasicArrayList(); _ = list
        var size int = list.Size(); _ = size
        var nextSymbolicName string; _ = nextSymbolicName
        for index := int(0); index < size; index++ {
            nextSymbolicName = any(list.Get(index)).(string)
            if strings.Compare(nextSymbolicName, symbolicName) == 0 {
                return true
            }
        }
        return false
}


// throws Exception
func (this *InputAutomationNewBundleRunnable) FindNewModules() IBasicArrayList {
        this.LogUtil.PutF(this.CommonStrings.START, this, "findNewModules")
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        var hashMap *HashMap = this.GetAllJarSymbolicNameHashMap(); _ = hashMap
        this.LogUtil.PutF("All: " + fmt.Sprint(hashMap), this, "findNewModules")
        var set Set = hashMap.KeySet(); _ = set
        var symbolicNameArray []any = set.ToArray(); _ = symbolicNameArray
        var size int = len(symbolicNameArray); _ = size
        for index := int(0); index < size; index++ {
            var symbolicName string = any(symbolicNameArray[index]).(string); _ = symbolicName
            if !this.IsInstalled(symbolicName) {
                vector.Add(hashMap.Get(symbolicName))
            }
        }
        return vector
}


// throws Exception
func (this *InputAutomationNewBundleRunnable) Install(url *URL) Bundle {
        this.LogUtil.PutF(CommonLabels_GetInstance().START + url, this, "install")
        var bundleContext BundleContext = InputAutomationBundleActivator_GetBundleContext(); _ = bundleContext
        return bundleContext.InstallBundle(fmt.Sprint(url))
}


func (this *InputAutomationNewBundleRunnable) Run() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.RUN, e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.RUN)
            this.SetRunning(true)
            var timeHelper ITimeDelayHelper = NewTimeDelayHelper(1000); _ = timeHelper
            for this.IsRunning() {
                timeHelper.SetStartTimeTNT()
                this.LogUtil.PutF(CommonLabels_GetInstance().ELAPSED + timeHelper.GetElapsedTNT(), this, this.CommonStrings.RUN)
                this.UpdateModules()
                break
            }
            this.LogUtil.PutF(this.CommonStrings.END, this, this.CommonStrings.RUN)
            return
        }(); tryControl1 == 1 {
            return
        }
}

