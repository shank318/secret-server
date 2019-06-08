package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"reflect"
	"regexp"
	"secret-server/app/constants"
	"secret-server/app/logger"
	"secret-server/app/metric"
	"secret-server/app/crerrors"
	"strconv"
	"strings"
	"time"
)

var newUUID = uuid.NewV4
var currentTime = time.Now
var logInfo = logger.Info
var regexCompile = regexp.MustCompile

// GetCurrentTimeStamp will give the current unix timestamp
func GetCurrentTimeStamp() time.Time {
	return currentTime()
}

func GetCurrentTimeStampInMiliSec() int64 {
	return currentTime().UnixNano() / 1000000
}

func Find(list []string, key string) (string, bool) {
	for _, ele := range list {
		ele = strings.ToLower(ele)
		if key == ele {
			return strings.ToUpper(ele), true
		}
	}
	return "", false
}

func GenerateUUID() string {
	return newUUID().String()
}

func GetTypeName(v interface{}) string {
	ret := ""
	t := reflect.TypeOf(v)
	getType(t, &ret)
	return ret

}

func getType(t reflect.Type, ret *string) {
	if *ret == "" {
		switch t.Kind() {
		case reflect.Array, reflect.Map, reflect.Ptr, reflect.Slice:
			getType(t.Elem(), ret)
		case reflect.Struct:
			*ret = t.Name()
		}
	}
}

func TraceExecutionTime(ctx *gin.Context, method string, startTime int64) {
	traceData := make(map[string]interface{})
	endTime := GetCurrentTimeStampInMiliSec()
	traceData[Method] = method
	traceData[StartTime] = startTime
	traceData[EndTime] = endTime
	executionTime := endTime - startTime
	traceData[ExecutionTime] = executionTime
	logInfo(ctx, method, traceData)

	if Contains(constants.PerformanceMetricsActions, method) {
		metric.TimeToProcess.Observe(float64(executionTime))
	}
}

// IsEmpty will check for given data is empty as per the go documentation
func IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}

	reflectVal := reflect.ValueOf(val)

	if val == nil {
		return true
	}

	switch reflectVal.Kind() {
	case reflect.Int:
		return val.(int) == 0

	case reflect.Int64:
		return val.(int64) == 0

	case reflect.String:
		return val.(string) == ""

	case reflect.Map:
	case reflect.Slice:
		return reflectVal.IsNil() || reflectVal.Len() == 0
	}

	return false
}

// IsArray : Check if the given value is an array or not
func IsArray(value interface{}) bool {
	rt := reflect.TypeOf(value)
	kind := rt.Kind()

	if kind != reflect.Array && kind != reflect.Slice {
		return false
	}

	return true
}

// IsMap : Check if the given value is a map
func IsMap(value interface{}) bool {
	rt := reflect.TypeOf(value)
	kind := rt.Kind()

	if kind != reflect.Map {
		return false
	}

	return true
}

// ConvertToInt : converts a value to int64 type
func ConvertToInt(value interface{}) (int64, error) {
	if IsEmpty(value) {
		return 0, nil
	}

	strValue := fmt.Sprintf("%v", value)

	i, err := strconv.ParseInt(strValue, 10, 32)

	return i, err
}

// StringSliceDiff : returns the elements in `a` that aren't in `b`
func StringSliceDiff(a, b []string) []string {
	mb := map[string]bool{}

	for _, x := range b {
		mb[x] = true
	}

	ab := make([]string, 0)

	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}

	return ab
}

// StringSliceUnique : returns a unique slice of the strings slice input
func StringSliceUnique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)

	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func GetMapKeys(mymap map[string]interface{}) []string {
	keys := make([]string, 0, len(mymap))

	for k := range mymap {
		keys = append(keys, k)
	}

	return keys
}

func IsValidUUID(uuid string) bool {
	r := regexCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// DeserializeJSON will convert the given json into map of string interface
func DeserializeJSON(ctx *gin.Context, in interface{}, out interface{}) crerrors.IError {
	var (
		ierr crerrors.IError
	)

	str, ok := in.(string)

	if !ok {
		return crerrors.NewCrError(ctx, crerrors.CodeIllegalArgument, nil).
			WithField("expected_type", "string or []byte").
			WithField("input_type", reflect.TypeOf(in)).
			Log()
	}

	input := []byte(str)

	err := json.Unmarshal(input, out)

	if err != nil {
		ierr = crerrors.NewCrError(ctx, crerrors.CodeBindingFailed, err).Log()
	}

	return ierr
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
