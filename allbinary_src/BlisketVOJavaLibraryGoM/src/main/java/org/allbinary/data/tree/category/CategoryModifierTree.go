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

type CategoryModifierTree struct {
    CategoryPrivateTree
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Directory *Directory
}

// Implements CategoryModifierTreeInterface

// The Java class CategoryModifierTree has subclasses so its variables are this interface
type ICategoryModifierTree interface {
    ICategoryPrivateTree
    CategoryModifierTree_() *CategoryModifierTree
    Insert(parentCategoryInterface CategoryInterface, newChildCategoryInterface CategoryInterface)
    Delete(parentCategoryInterface CategoryInterface, existingChildCategoryInterface CategoryInterface)
}

func (this *CategoryModifierTree) CategoryModifierTree_() *CategoryModifierTree {
    return this
}


func (this *CategoryModifierTree) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.Directory = Directory_GetInstance()
}


func NewCategoryModifierTree(categoryFactoryInterface CategoryFactoryInterface) *CategoryModifierTree {
    this := &CategoryModifierTree{}
        this.CategoryPrivateTree = *NewCategoryPrivateTree(categoryFactoryInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
            this.LogUtil.PutF(this.CommonStrings.START, this, "CategoryModifierTree(CategoryFactoryInterface categoryFactoryInterface)")
        }
    return this
}


// synchronized
func (this *CategoryModifierTree) Insert(parentCategoryInterface CategoryInterface, newChildCategoryInterface CategoryInterface) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, "insert", e)
            }
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                this.LogUtil.PutF("Inserting", this, "insert()")
            }
            if parentCategoryInterface.IsValid() {
                parentCategoryInterface.AddChild(newChildCategoryInterface)
                var directoryToBeCreatedAbPath IAbPath = NewAbPath(fmt.Sprint(newChildCategoryInterface.GetRootFilePath()) + fmt.Sprint(newChildCategoryInterface.GetPath()), StringUtil_GetInstance().EMPTY_STRING); _ = directoryToBeCreatedAbPath
                this.Directory.Create(directoryToBeCreatedAbPath)
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().CATEGORY) {
                    this.LogUtil.PutF("Saving", this, "insert()")
                }
                this.Save(parentCategoryInterface)
                this.Save(newChildCategoryInterface)
            } else {
                panic(errors.New(fmt.Sprint("Will Not Add Null Category")))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func (this *CategoryModifierTree) Delete(parentCategoryInterface CategoryInterface, existingChildCategoryInterface CategoryInterface) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLLOGGING) {
                this.LogUtil.Put(this.CommonStrings.FAILURE, this, this.CommonStrings.Delete, e)
            }
                }
            }()
            if parentCategoryInterface.IsValid() {
                parentCategoryInterface.AddChild(existingChildCategoryInterface)
                var directoryToBeDeletedAbPath IAbPath = NewAbPath(fmt.Sprint(existingChildCategoryInterface.GetRootFilePath()) + fmt.Sprint(existingChildCategoryInterface.GetPath()), StringUtil_GetInstance().EMPTY_STRING); _ = directoryToBeDeletedAbPath
                this.Delete(existingChildCategoryInterface)
                parentCategoryInterface.RemoveChild(existingChildCategoryInterface)
                this.Save(parentCategoryInterface)
                this.Directory.Remove(directoryToBeDeletedAbPath)
            } else {
                panic(errors.New(fmt.Sprint("Will Not Add Null Category")))
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

