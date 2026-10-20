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

type ItemInterface interface {
    TableMappingInterface
    SetId(itemId string)
    SetNumber(number string)
    SetInBaskets(value string)
    SetWeight(value string)
    SetEnabled(value string)
    SetNewOrUsed(value string)
    SetSummary(value string)
    SetDistributor(value string)
    SetIdUsedByDistributor(value string)
    SetProducedBy(value string)
    SetProductionDate(value string)
    SetStartProductionDate(value string)
    SetDescription(value string)
    SetKeywords(value string)
    SetCategory(value string)
    SetType(value string)
    SetSmallImage(value string)
    SetMediumImage(value string)
    SetLargeImage(value string)
    SetTimeEntered(value string)
    SetLastModified(value string)
    SetPrice(value *Money)
    SetComment(value string)
    SetCustoms(value string)
    SetDownloads(value string)
    SetGroups(value string)
    SetOptions(value string)
    SetPermissions(value string)
    SetSpecials(value string)
    GetId() string
    GetNumber() string
    GetInBaskets() string
    GetWeight() string
    GetEnabled() string
    GetNewOrUsed() string
    GetSummary() string
    GetDistributor() string
    GetIdUsedByDistributor() string
    GetProducedBy() string
    GetProductionDate() string
    GetStartProductionDate() string
    GetDescription() string
    GetKeywords() string
    GetCategory() string
    GetType() string
    GetSmallImage() string
    GetMediumImage() string
    GetLargeImage() string
    GetTimeEntered() string
    GetLastModified() string
    GetPrice() *Money
    GetComment() string
    GetCustoms() string
    GetDownloads() string
    GetGroups() string
    GetOptions() string
    GetPermissions() string
    GetSpecials() string
    GetTotal() *Money
}


