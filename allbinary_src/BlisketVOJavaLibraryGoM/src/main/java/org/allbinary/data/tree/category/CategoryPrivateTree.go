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

type CategoryPrivateTree struct {
    LogUtil *LogUtil
    CategoryFactoryInterface CategoryFactoryInterface
}

// The Java class CategoryPrivateTree has subclasses so its variables are this interface
type ICategoryPrivateTree interface {
    CategoryPrivateTree_() *CategoryPrivateTree
    Save(categoryInterface CategoryInterface)
    Delete(categoryInterface CategoryInterface)
}

func (this *CategoryPrivateTree) CategoryPrivateTree_() *CategoryPrivateTree {
    return this
}


func (this *CategoryPrivateTree) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewCategoryPrivateTree(categoryFactoryInterface CategoryFactoryInterface) *CategoryPrivateTree {
    this := &CategoryPrivateTree{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.START, this, "CategoryPrivateTree(CategoryFactoryInterface categoryFactoryInterface)")
        }
        this.CategoryFactoryInterface = categoryFactoryInterface
    return this
}


// throws Exception
// synchronized
func (this *CategoryPrivateTree) Save(categoryInterface CategoryInterface) {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
            this.LogUtil.PutF("Saving Category File: " + fmt.Sprint(categoryInterface.GetFilePath()), this, "save()")
        }
        var document *Document = NewCategoryComponent(categoryInterface).ToXmlDoc(); _ = document
        var cryptFileWriter *CryptFileWriter = NewCryptFileWriter(CategoryData_GetInstance().UNCRYPTED_EXTENSION, CategoryData_GetInstance().ENCRYPTED_EXTENSION); _ = cryptFileWriter
        cryptFileWriter.Write(categoryInterface.GetFilePath(), document)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
            this.LogUtil.PutF("end", this, "save()")
        }
}


// throws Exception
// synchronized
func (this *CategoryPrivateTree) Delete(categoryInterface CategoryInterface) {
        if categoryInterface.IsLeaf() {
            var categoryFile *AbFile = AbFile_CreateAbFileFromAbPath(categoryInterface.GetFilePath()); _ = categoryFile
            if categoryFile.Delete() {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                    this.LogUtil.PutF("Removed: " + fmt.Sprint(categoryInterface.GetFilePath()), this, "delete(CategoryInterface)")
                }
            }
        } else {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.PutF("Not a leaf unable to remove", this, "delete(CategoryInterface)")
            }
        }
}

