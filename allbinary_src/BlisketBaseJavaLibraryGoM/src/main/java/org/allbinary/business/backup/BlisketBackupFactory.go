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

type BlisketBackupFactory struct {
    LogUtil *LogUtil
    StringUtil *StringUtil
    CommonStrings *CommonStrings
    Directory *Directory
}

var BlisketBackupFactory_Instance *BlisketBackupFactory = NewBlisketBackupFactory()

func (this *BlisketBackupFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.Directory = Directory_GetInstance()
}

func NewBlisketBackupFactory() *BlisketBackupFactory {
    this := &BlisketBackupFactory{}
    this.initFields()
    return this
}


func BlisketBackupFactory_GetInstance() *BlisketBackupFactory {
        return BlisketBackupFactory_Instance
}


// throws Exception
func (this *BlisketBackupFactory) GetFileBasicArrayList(pathString string) IBasicArrayList {
        var path IAbPath = NewAbPath(pathString, this.StringUtil.EMPTY_STRING); _ = path
        var file *AbFile = AbFile_CreateAbFileFromAbPath(path); _ = file
        return this.Directory.Search(file, true)
}


func (this *BlisketBackupFactory) Backup_1() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "backup()", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "backup()")
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(URLGLOBALS.GetWebappPath())
            stringBuffer.Append(PATH_GLOBALS.GetInstance().BACKUP_PATH)
            var backupPath string = fmt.Sprint(stringBuffer); _ = backupPath
            var path IAbPath = NewAbPath(backupPath, this.StringUtil.EMPTY_STRING); _ = path
            var fileBasicArrayList IBasicArrayList = this.GetFileBasicArrayList(backupPath); _ = fileBasicArrayList
            this.Backup(fileBasicArrayList, fmt.Sprint(path.ToFileSystemString()) + "backup.zip")
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of BlisketBackupFactory.backup
func (this *BlisketBackupFactory) Backup(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            this.Backup_1()
            return
        }
        if len(args) == 2 && Java_IsArg[IBasicArrayList](args[0], pass) && Java_IsArg[string](args[1], pass) {
            this.Backup_2(Java_Arg[IBasicArrayList](args[0]), Java_Arg[string](args[1]))
            return
        }
    }
    panic("No Java overload of BlisketBackupFactory.backup for the arguments")
}


func (this *BlisketBackupFactory) BackupViews() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "backupViews()", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "backupViews()")
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(URLGLOBALS.GetWebappPath())
            stringBuffer.Append(PATH_GLOBALS.GetInstance().BACKUP_PATH)
            var backupPath string = fmt.Sprint(stringBuffer); _ = backupPath
            var path IAbPath = NewAbPath(backupPath, this.StringUtil.EMPTY_STRING); _ = path
            var fileBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = fileBasicArrayList
            var storeFrontsEntity *StoreFrontsEntity = StoreFrontsEntityFactory_GetInstance().GetStoreFrontsEntityInstance(); _ = storeFrontsEntity
            var storeFrontNamesBasicArrayList IBasicArrayList = storeFrontsEntity.GetStoreFrontNames(); _ = storeFrontNamesBasicArrayList
            var size int = storeFrontNamesBasicArrayList.Size(); _ = size
            var nextStore string; _ = nextStore
            var viewsPath string; _ = viewsPath
            for index := int(0); index < size; index++ {
                nextStore = any(storeFrontNamesBasicArrayList.Get(index)).(string)
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(URLGLOBALS.GetWebappPath())
                stringBuffer.Append(PATH_GLOBALS.GetInstance().VIEWS_PATH)
                stringBuffer.Append(nextStore)
                viewsPath = fmt.Sprint(stringBuffer)
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append("Backup Store Views: ")
                stringBuffer.Append(nextStore)
                stringBuffer.Append(" from: ")
                stringBuffer.Append(viewsPath)
                PreLogUtil_Put(fmt.Sprint(stringBuffer), this, "backupViews()")
                fileBasicArrayList.AddAllList(this.GetFileBasicArrayList(viewsPath))
            }
            this.Backup(fileBasicArrayList, fmt.Sprint(path.ToFileSystemString()) + "backupViews.zip")
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *BlisketBackupFactory) BackupResources() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "backupResources()", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "backupResources()")
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(URLGLOBALS.GetWebappPath())
            stringBuffer.Append(PATH_GLOBALS.GetInstance().BACKUP_PATH)
            var backupPath string = fmt.Sprint(stringBuffer); _ = backupPath
            var path IAbPath = NewAbPath(backupPath, this.StringUtil.EMPTY_STRING); _ = path
            var fileBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = fileBasicArrayList
            var storeFrontsEntity *StoreFrontsEntity = StoreFrontsEntityFactory_GetInstance().GetStoreFrontsEntityInstance(); _ = storeFrontsEntity
            var storeFrontNamesBasicArrayList IBasicArrayList = storeFrontsEntity.GetStoreFrontNames(); _ = storeFrontNamesBasicArrayList
            var size int = storeFrontNamesBasicArrayList.Size(); _ = size
            var nextStore string; _ = nextStore
            var storeFrontInterface StoreFrontInterface; _ = storeFrontInterface
            var resourcesPath string; _ = resourcesPath
            for index := int(0); index < size; index++ {
                nextStore = any(storeFrontNamesBasicArrayList.Get(index)).(string)
                storeFrontInterface = storeFrontsEntity.GetStoreFrontInterface(nextStore)
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(URLGLOBALS.GetWebappPath())
                stringBuffer.Append(storeFrontInterface.GetCurrentHostNamePath())
                stringBuffer.Append(storeFrontInterface.GetCategoryPath())
                resourcesPath = fmt.Sprint(stringBuffer)
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append("Backup Store Resrouces: ")
                stringBuffer.Append(nextStore)
                stringBuffer.Append(" from: ")
                stringBuffer.Append(resourcesPath)
                PreLogUtil_Put(fmt.Sprint(stringBuffer), this, "backupResources()")
                fileBasicArrayList.AddAllList(this.GetFileBasicArrayList(resourcesPath))
            }
            this.Backup(fileBasicArrayList, fmt.Sprint(path.ToFileSystemString()) + "backupResources.zip")
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *BlisketBackupFactory) BackupJsps() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "backupJsps()", e)
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "backupJsps()")
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(URLGLOBALS.GetWebappPath())
            stringBuffer.Append(PATH_GLOBALS.GetInstance().BACKUP_PATH)
            var backupPath string = fmt.Sprint(stringBuffer); _ = backupPath
            var path IAbPath = NewAbPath(backupPath, this.StringUtil.EMPTY_STRING); _ = path
            var fileBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = fileBasicArrayList
            var storeFrontsEntity *StoreFrontsEntity = StoreFrontsEntityFactory_GetInstance().GetStoreFrontsEntityInstance(); _ = storeFrontsEntity
            var storeFrontNamesBasicArrayList IBasicArrayList = storeFrontsEntity.GetStoreFrontNames(); _ = storeFrontNamesBasicArrayList
            var size int = storeFrontNamesBasicArrayList.Size(); _ = size
            var nextStore string; _ = nextStore
            for index := int(0); index < size; index++ {
                nextStore = any(storeFrontNamesBasicArrayList.Get(index)).(string)
                PreLogUtil_Put("Backup Store Jsps: " + nextStore, this, "backupJsps()")
                stringBuffer.Delete(0, stringBuffer.Length())
                stringBuffer.Append(URLGLOBALS.GetWebappPath())
                stringBuffer.Append(nextStore)
                fileBasicArrayList.AddAllList(this.GetFileBasicArrayList(fmt.Sprint(stringBuffer)))
            }
            this.Backup(fileBasicArrayList, fmt.Sprint(path.ToFileSystemString()) + "backupJsps.zip")
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *BlisketBackupFactory) Backup_2(fileBasicArrayList IBasicArrayList, zipFile string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "backup()", e)
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Delete(0, stringBuffer.Length())
            stringBuffer.Append("ZipFile: ")
            stringBuffer.Append(zipFile)
            stringBuffer.Append(" BasicArrayList: ")
            stringBuffer.Appendint(fileBasicArrayList.Size())
            PreLogUtil_Put("Creating Backup Zip File: " + fmt.Sprint(stringBuffer), this, "backup()")
            ZipFileUtil_GetInstance().Create(zipFile, fileBasicArrayList)
            this.LogUtil.PutF("Created Backup Zip File", this, "backup()")
            return
        }(); tryControl1 == 1 {
            return
        }
}

