package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testModel struct {
}

func TestGenerateUUID(t *testing.T) {
	oldNewUUID := newUUID
	defer func() { newUUID = oldNewUUID }()
	newUUID = func() uuid.UUID {
		var b = [16]byte{}
		copy(b[:], "ABCDEFGHIJKLMNOP")
		return b
	}

	uuid := GenerateUUID()
	assert.Equal(t, "41424344-4546-4748-494a-4b4c4d4e4f50", uuid)
}

func TestGetTypeName(t *testing.T) {
	model := testModel{}
	var models []testModel
	n := GetTypeName(model)
	assert.Equal(t, "testModel", n)
	n = GetTypeName(&model)
	assert.Equal(t, "testModel", n)
	n = GetTypeName(models)
	assert.Equal(t, "testModel", n)
	n = GetTypeName(&models)
	assert.Equal(t, "testModel", n)

}

func TestTraceExecutionTime(t *testing.T) {
	c := &gin.Context{}
	TraceExecutionTime(c, "My_Method", GetCurrentTimeStampInMiliSec())
	oldLogInfo := logInfo
	defer func() { logInfo = oldLogInfo }()
	logInfo = func(ctx *gin.Context, traceCode string, traceData map[string]interface{}) {
		assert.Equal(t, "My_Method", traceCode)
		assert.Equal(t, c, ctx)
		assert.Equal(t, 0, traceData["execution_time"])
	}
}

func TestGetMapKeys(t *testing.T) {
	m := make(map[string]interface{})
	s := GetMapKeys(m)
	assert.Equal(t, 0, len(s))
	m["A"] = "B"
	m["X"] = "Y"
	s = GetMapKeys(m)
	assert.Contains(t, s, "X")
	assert.Contains(t, s, "A")
}

func TestStringSliceUnique(t *testing.T) {
	s1 := []string{"A", "B", "C", "A", "A", "C"}
	s2 := StringSliceUnique(s1)
	assert.Equal(t, 3, len(s2))
	assert.Contains(t, s2, "A")
	assert.Contains(t, s2, "B")
	assert.Contains(t, s2, "C")
}

func TestStringSliceDiff(t *testing.T) {
	s1 := []string{"A", "B", "C", "A", "A", "C"}
	s2 := []string{"A", "B", "C", "A", "A", "C"}
	s3 := StringSliceDiff(s1, s2)
	assert.Equal(t, 0, len(s3))

	s2 = []string{"A", "C", "A", "A", "C"}
	s3 = StringSliceDiff(s1, s2)
	assert.Equal(t, 1, len(s3))
	assert.Equal(t, "B", s3[0])

}

func TestConvertToInt(t *testing.T) {
	s1 := "10"
	n, err := ConvertToInt(s1)
	assert.Equal(t, int64(10), n)
	assert.Equal(t, nil, err)

	s2 := []string{"10"}
	n, err = ConvertToInt(s2)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, int64(0), n)

	n, err = ConvertToInt(nil)
	assert.Equal(t, int64(0), n)
	assert.Equal(t, nil, err)

}

func TestIsArray(t *testing.T) {
	s2 := []string{"10"}
	r := IsArray(s2)
	assert.Equal(t, true, r)

	s5 := []string{}
	r = IsArray(s5)
	assert.Equal(t, true, r)

	s1 := 10
	r = IsArray(s1)
	assert.Equal(t, false, r)

	s3 := "10"
	r = IsArray(s3)
	assert.Equal(t, false, r)

	s4 := make(map[string]string)
	r = IsArray(s4)
	assert.Equal(t, false, r)

}

func TestIsMap(t *testing.T) {

	s2 := []string{"10"}
	r := IsMap(s2)
	assert.Equal(t, false, r)

	s5 := []string{}
	r = IsMap(s5)
	assert.Equal(t, false, r)

	s1 := 10
	r = IsMap(s1)
	assert.Equal(t, false, r)

	s3 := "10"
	r = IsMap(s3)
	assert.Equal(t, false, r)

	s4 := make(map[string]string)
	r = IsMap(s4)
	assert.Equal(t, true, r)

	s6 := make(map[string]string)
	s6["A"] = "B"
	r = IsMap(s6)
	assert.Equal(t, true, r)

}

func TestIsEmpty(t *testing.T) {

	s2 := []string{"10"}
	r := IsEmpty(s2)
	assert.Equal(t, false, r)

	s5 := []string{}
	r = IsEmpty(s5)
	assert.Equal(t, true, r)

	s1 := 10
	r = IsEmpty(s1)
	assert.Equal(t, false, r)

	s3 := "10"
	r = IsEmpty(s3)
	assert.Equal(t, false, r)

	s4 := make(map[string]string)
	r = IsEmpty(s4)
	assert.Equal(t, false, r)

	s6 := make(map[string]string)
	s6["A"] = "B"
	r = IsEmpty(s6)
	assert.Equal(t, false, r)

	s7 := int64(10)
	r = IsEmpty(s7)
	assert.Equal(t, false, r)

	r = IsEmpty(nil)
	assert.Equal(t, true, r)

	s8 := ""
	r = IsEmpty(s8)
	assert.Equal(t, true, r)

}

func TestGetCurrentTimeStamp(t *testing.T) {
	t1 := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	oldCurrentTime := currentTime
	defer func() { currentTime = oldCurrentTime }()
	currentTime = func() time.Time {
		return t1
	}

	t2 := GetCurrentTimeStamp()
	assert.Equal(t, t1, t2)
}

func TestGetCurrentTimeStampInMiliSec(t *testing.T) {
	t1 := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	oldCurrentTime := currentTime
	defer func() { currentTime = oldCurrentTime }()
	currentTime = func() time.Time {
		return t1
	}

	t2 := GetCurrentTimeStampInMiliSec()
	assert.Equal(t, t1.UnixNano()/1000000, t2)
}

func TestIsValidUUID(t *testing.T) {
	r := IsValidUUID("732b4666-9a49-45b5-ba52-bb806dec8dd4")
	assert.Equal(t, true, r)
	r = IsValidUUID("41424344-4546-4748-494a-4b4c4d4e4f50")
	assert.Equal(t, false, r)
}

func TestDeserializeJSON(t *testing.T) {
	c := &gin.Context{}
	var m interface{}
	jsonStr := `{"id":1,"name":"bowser","breed":"husky"}`
	r := DeserializeJSON(c, jsonStr, &m)
	msgMap := m.(map[string]interface{})
	assert.Equal(t, "bowser", msgMap["name"])
	assert.Nil(t, r)

	jsonInvalidStr := `{"id":1,"name":"bowser}`
	r2 := DeserializeJSON(c, jsonInvalidStr, &m)
	assert.NotNil(t, r2)

	r3 := DeserializeJSON(c, nil, &m)
	assert.NotNil(t, r3)
}

func TestContains(t *testing.T) {
	var input = []string{"X", "Y", "Z"}
	contains := Contains(input, "X")
	assert.True(t, contains)

	contains2 := Contains(input, "A")
	assert.False(t, contains2)
}
