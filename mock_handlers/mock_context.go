// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SebastianCoetzee/blog-order-service-example/handlers (interfaces: Context)

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	io "io"
	multipart "mime/multipart"
	reflect "reflect"
	time "time"

	gin "github.com/gin-gonic/gin"
	binding "github.com/gin-gonic/gin/binding"
	render "github.com/gin-gonic/gin/render"
	gomock "github.com/golang/mock/gomock"
)

// MockContext is a mock of Context interface
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Abort mocks base method
func (m *MockContext) Abort() {
	m.ctrl.Call(m, "Abort")
}

// Abort indicates an expected call of Abort
func (mr *MockContextMockRecorder) Abort() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockContext)(nil).Abort))
}

// AbortWithError mocks base method
func (m *MockContext) AbortWithError(arg0 int, arg1 error) *gin.Error {
	ret := m.ctrl.Call(m, "AbortWithError", arg0, arg1)
	ret0, _ := ret[0].(*gin.Error)
	return ret0
}

// AbortWithError indicates an expected call of AbortWithError
func (mr *MockContextMockRecorder) AbortWithError(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortWithError", reflect.TypeOf((*MockContext)(nil).AbortWithError), arg0, arg1)
}

// AbortWithStatus mocks base method
func (m *MockContext) AbortWithStatus(arg0 int) {
	m.ctrl.Call(m, "AbortWithStatus", arg0)
}

// AbortWithStatus indicates an expected call of AbortWithStatus
func (mr *MockContextMockRecorder) AbortWithStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortWithStatus", reflect.TypeOf((*MockContext)(nil).AbortWithStatus), arg0)
}

// AbortWithStatusJSON mocks base method
func (m *MockContext) AbortWithStatusJSON(arg0 int, arg1 interface{}) {
	m.ctrl.Call(m, "AbortWithStatusJSON", arg0, arg1)
}

// AbortWithStatusJSON indicates an expected call of AbortWithStatusJSON
func (mr *MockContextMockRecorder) AbortWithStatusJSON(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortWithStatusJSON", reflect.TypeOf((*MockContext)(nil).AbortWithStatusJSON), arg0, arg1)
}

// AsciiJSON mocks base method
func (m *MockContext) AsciiJSON(arg0 int, arg1 interface{}) {
	m.ctrl.Call(m, "AsciiJSON", arg0, arg1)
}

// AsciiJSON indicates an expected call of AsciiJSON
func (mr *MockContextMockRecorder) AsciiJSON(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsciiJSON", reflect.TypeOf((*MockContext)(nil).AsciiJSON), arg0, arg1)
}

// Bind mocks base method
func (m *MockContext) Bind(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "Bind", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind
func (mr *MockContextMockRecorder) Bind(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockContext)(nil).Bind), arg0)
}

// BindJSON mocks base method
func (m *MockContext) BindJSON(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "BindJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindJSON indicates an expected call of BindJSON
func (mr *MockContextMockRecorder) BindJSON(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindJSON", reflect.TypeOf((*MockContext)(nil).BindJSON), arg0)
}

// BindQuery mocks base method
func (m *MockContext) BindQuery(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "BindQuery", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindQuery indicates an expected call of BindQuery
func (mr *MockContextMockRecorder) BindQuery(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindQuery", reflect.TypeOf((*MockContext)(nil).BindQuery), arg0)
}

// ClientIP mocks base method
func (m *MockContext) ClientIP() string {
	ret := m.ctrl.Call(m, "ClientIP")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientIP indicates an expected call of ClientIP
func (mr *MockContextMockRecorder) ClientIP() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientIP", reflect.TypeOf((*MockContext)(nil).ClientIP))
}

// ContentType mocks base method
func (m *MockContext) ContentType() string {
	ret := m.ctrl.Call(m, "ContentType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContentType indicates an expected call of ContentType
func (mr *MockContextMockRecorder) ContentType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentType", reflect.TypeOf((*MockContext)(nil).ContentType))
}

// Cookie mocks base method
func (m *MockContext) Cookie(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "Cookie", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cookie indicates an expected call of Cookie
func (mr *MockContextMockRecorder) Cookie(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cookie", reflect.TypeOf((*MockContext)(nil).Cookie), arg0)
}

// Copy mocks base method
func (m *MockContext) Copy() *gin.Context {
	ret := m.ctrl.Call(m, "Copy")
	ret0, _ := ret[0].(*gin.Context)
	return ret0
}

// Copy indicates an expected call of Copy
func (mr *MockContextMockRecorder) Copy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockContext)(nil).Copy))
}

// Data mocks base method
func (m *MockContext) Data(arg0 int, arg1 string, arg2 []byte) {
	m.ctrl.Call(m, "Data", arg0, arg1, arg2)
}

// Data indicates an expected call of Data
func (mr *MockContextMockRecorder) Data(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockContext)(nil).Data), arg0, arg1, arg2)
}

// DataFromReader mocks base method
func (m *MockContext) DataFromReader(arg0 int, arg1 int64, arg2 string, arg3 io.Reader, arg4 map[string]string) {
	m.ctrl.Call(m, "DataFromReader", arg0, arg1, arg2, arg3, arg4)
}

// DataFromReader indicates an expected call of DataFromReader
func (mr *MockContextMockRecorder) DataFromReader(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataFromReader", reflect.TypeOf((*MockContext)(nil).DataFromReader), arg0, arg1, arg2, arg3, arg4)
}

// Deadline mocks base method
func (m *MockContext) Deadline() (time.Time, bool) {
	ret := m.ctrl.Call(m, "Deadline")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Deadline indicates an expected call of Deadline
func (mr *MockContextMockRecorder) Deadline() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deadline", reflect.TypeOf((*MockContext)(nil).Deadline))
}

// DefaultPostForm mocks base method
func (m *MockContext) DefaultPostForm(arg0, arg1 string) string {
	ret := m.ctrl.Call(m, "DefaultPostForm", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// DefaultPostForm indicates an expected call of DefaultPostForm
func (mr *MockContextMockRecorder) DefaultPostForm(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultPostForm", reflect.TypeOf((*MockContext)(nil).DefaultPostForm), arg0, arg1)
}

// DefaultQuery mocks base method
func (m *MockContext) DefaultQuery(arg0, arg1 string) string {
	ret := m.ctrl.Call(m, "DefaultQuery", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// DefaultQuery indicates an expected call of DefaultQuery
func (mr *MockContextMockRecorder) DefaultQuery(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultQuery", reflect.TypeOf((*MockContext)(nil).DefaultQuery), arg0, arg1)
}

// Done mocks base method
func (m *MockContext) Done() <-chan struct{} {
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Done indicates an expected call of Done
func (mr *MockContextMockRecorder) Done() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockContext)(nil).Done))
}

// Err mocks base method
func (m *MockContext) Err() error {
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockContextMockRecorder) Err() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockContext)(nil).Err))
}

// Error mocks base method
func (m *MockContext) Error(arg0 error) *gin.Error {
	ret := m.ctrl.Call(m, "Error", arg0)
	ret0, _ := ret[0].(*gin.Error)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockContextMockRecorder) Error(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockContext)(nil).Error), arg0)
}

// File mocks base method
func (m *MockContext) File(arg0 string) {
	m.ctrl.Call(m, "File", arg0)
}

// File indicates an expected call of File
func (mr *MockContextMockRecorder) File(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "File", reflect.TypeOf((*MockContext)(nil).File), arg0)
}

// FormFile mocks base method
func (m *MockContext) FormFile(arg0 string) (*multipart.FileHeader, error) {
	ret := m.ctrl.Call(m, "FormFile", arg0)
	ret0, _ := ret[0].(*multipart.FileHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormFile indicates an expected call of FormFile
func (mr *MockContextMockRecorder) FormFile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormFile", reflect.TypeOf((*MockContext)(nil).FormFile), arg0)
}

// Get mocks base method
func (m *MockContext) Get(arg0 string) (interface{}, bool) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockContextMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContext)(nil).Get), arg0)
}

// GetBool mocks base method
func (m *MockContext) GetBool(arg0 string) bool {
	ret := m.ctrl.Call(m, "GetBool", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetBool indicates an expected call of GetBool
func (mr *MockContextMockRecorder) GetBool(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBool", reflect.TypeOf((*MockContext)(nil).GetBool), arg0)
}

// GetDuration mocks base method
func (m *MockContext) GetDuration(arg0 string) time.Duration {
	ret := m.ctrl.Call(m, "GetDuration", arg0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetDuration indicates an expected call of GetDuration
func (mr *MockContextMockRecorder) GetDuration(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDuration", reflect.TypeOf((*MockContext)(nil).GetDuration), arg0)
}

// GetFloat64 mocks base method
func (m *MockContext) GetFloat64(arg0 string) float64 {
	ret := m.ctrl.Call(m, "GetFloat64", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetFloat64 indicates an expected call of GetFloat64
func (mr *MockContextMockRecorder) GetFloat64(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloat64", reflect.TypeOf((*MockContext)(nil).GetFloat64), arg0)
}

// GetHeader mocks base method
func (m *MockContext) GetHeader(arg0 string) string {
	ret := m.ctrl.Call(m, "GetHeader", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHeader indicates an expected call of GetHeader
func (mr *MockContextMockRecorder) GetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockContext)(nil).GetHeader), arg0)
}

// GetInt mocks base method
func (m *MockContext) GetInt(arg0 string) int {
	ret := m.ctrl.Call(m, "GetInt", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetInt indicates an expected call of GetInt
func (mr *MockContextMockRecorder) GetInt(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt", reflect.TypeOf((*MockContext)(nil).GetInt), arg0)
}

// GetInt64 mocks base method
func (m *MockContext) GetInt64(arg0 string) int64 {
	ret := m.ctrl.Call(m, "GetInt64", arg0)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetInt64 indicates an expected call of GetInt64
func (mr *MockContextMockRecorder) GetInt64(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInt64", reflect.TypeOf((*MockContext)(nil).GetInt64), arg0)
}

// GetPostForm mocks base method
func (m *MockContext) GetPostForm(arg0 string) (string, bool) {
	ret := m.ctrl.Call(m, "GetPostForm", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPostForm indicates an expected call of GetPostForm
func (mr *MockContextMockRecorder) GetPostForm(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostForm", reflect.TypeOf((*MockContext)(nil).GetPostForm), arg0)
}

// GetPostFormArray mocks base method
func (m *MockContext) GetPostFormArray(arg0 string) ([]string, bool) {
	ret := m.ctrl.Call(m, "GetPostFormArray", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPostFormArray indicates an expected call of GetPostFormArray
func (mr *MockContextMockRecorder) GetPostFormArray(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostFormArray", reflect.TypeOf((*MockContext)(nil).GetPostFormArray), arg0)
}

// GetPostFormMap mocks base method
func (m *MockContext) GetPostFormMap(arg0 string) (map[string]string, bool) {
	ret := m.ctrl.Call(m, "GetPostFormMap", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetPostFormMap indicates an expected call of GetPostFormMap
func (mr *MockContextMockRecorder) GetPostFormMap(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostFormMap", reflect.TypeOf((*MockContext)(nil).GetPostFormMap), arg0)
}

// GetQuery mocks base method
func (m *MockContext) GetQuery(arg0 string) (string, bool) {
	ret := m.ctrl.Call(m, "GetQuery", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetQuery indicates an expected call of GetQuery
func (mr *MockContextMockRecorder) GetQuery(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuery", reflect.TypeOf((*MockContext)(nil).GetQuery), arg0)
}

// GetQueryArray mocks base method
func (m *MockContext) GetQueryArray(arg0 string) ([]string, bool) {
	ret := m.ctrl.Call(m, "GetQueryArray", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetQueryArray indicates an expected call of GetQueryArray
func (mr *MockContextMockRecorder) GetQueryArray(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryArray", reflect.TypeOf((*MockContext)(nil).GetQueryArray), arg0)
}

// GetQueryMap mocks base method
func (m *MockContext) GetQueryMap(arg0 string) (map[string]string, bool) {
	ret := m.ctrl.Call(m, "GetQueryMap", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetQueryMap indicates an expected call of GetQueryMap
func (mr *MockContextMockRecorder) GetQueryMap(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryMap", reflect.TypeOf((*MockContext)(nil).GetQueryMap), arg0)
}

// GetRawData mocks base method
func (m *MockContext) GetRawData() ([]byte, error) {
	ret := m.ctrl.Call(m, "GetRawData")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawData indicates an expected call of GetRawData
func (mr *MockContextMockRecorder) GetRawData() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawData", reflect.TypeOf((*MockContext)(nil).GetRawData))
}

// GetString mocks base method
func (m *MockContext) GetString(arg0 string) string {
	ret := m.ctrl.Call(m, "GetString", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetString indicates an expected call of GetString
func (mr *MockContextMockRecorder) GetString(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetString", reflect.TypeOf((*MockContext)(nil).GetString), arg0)
}

// GetStringMap mocks base method
func (m *MockContext) GetStringMap(arg0 string) map[string]interface{} {
	ret := m.ctrl.Call(m, "GetStringMap", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// GetStringMap indicates an expected call of GetStringMap
func (mr *MockContextMockRecorder) GetStringMap(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringMap", reflect.TypeOf((*MockContext)(nil).GetStringMap), arg0)
}

// GetStringMapString mocks base method
func (m *MockContext) GetStringMapString(arg0 string) map[string]string {
	ret := m.ctrl.Call(m, "GetStringMapString", arg0)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetStringMapString indicates an expected call of GetStringMapString
func (mr *MockContextMockRecorder) GetStringMapString(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringMapString", reflect.TypeOf((*MockContext)(nil).GetStringMapString), arg0)
}

// GetStringMapStringSlice mocks base method
func (m *MockContext) GetStringMapStringSlice(arg0 string) map[string][]string {
	ret := m.ctrl.Call(m, "GetStringMapStringSlice", arg0)
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// GetStringMapStringSlice indicates an expected call of GetStringMapStringSlice
func (mr *MockContextMockRecorder) GetStringMapStringSlice(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringMapStringSlice", reflect.TypeOf((*MockContext)(nil).GetStringMapStringSlice), arg0)
}

// GetStringSlice mocks base method
func (m *MockContext) GetStringSlice(arg0 string) []string {
	ret := m.ctrl.Call(m, "GetStringSlice", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetStringSlice indicates an expected call of GetStringSlice
func (mr *MockContextMockRecorder) GetStringSlice(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringSlice", reflect.TypeOf((*MockContext)(nil).GetStringSlice), arg0)
}

// GetTime mocks base method
func (m *MockContext) GetTime(arg0 string) time.Time {
	ret := m.ctrl.Call(m, "GetTime", arg0)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTime indicates an expected call of GetTime
func (mr *MockContextMockRecorder) GetTime(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTime", reflect.TypeOf((*MockContext)(nil).GetTime), arg0)
}

// HTML mocks base method
func (m *MockContext) HTML(arg0 int, arg1 string, arg2 interface{}) {
	m.ctrl.Call(m, "HTML", arg0, arg1, arg2)
}

// HTML indicates an expected call of HTML
func (mr *MockContextMockRecorder) HTML(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTML", reflect.TypeOf((*MockContext)(nil).HTML), arg0, arg1, arg2)
}

// Handler mocks base method
func (m *MockContext) Handler() gin.HandlerFunc {
	ret := m.ctrl.Call(m, "Handler")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// Handler indicates an expected call of Handler
func (mr *MockContextMockRecorder) Handler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockContext)(nil).Handler))
}

// HandlerName mocks base method
func (m *MockContext) HandlerName() string {
	ret := m.ctrl.Call(m, "HandlerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// HandlerName indicates an expected call of HandlerName
func (mr *MockContextMockRecorder) HandlerName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlerName", reflect.TypeOf((*MockContext)(nil).HandlerName))
}

// Header mocks base method
func (m *MockContext) Header(arg0, arg1 string) {
	m.ctrl.Call(m, "Header", arg0, arg1)
}

// Header indicates an expected call of Header
func (mr *MockContextMockRecorder) Header(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockContext)(nil).Header), arg0, arg1)
}

// IndentedJSON mocks base method
func (m *MockContext) IndentedJSON(arg0 int, arg1 interface{}) {
	m.ctrl.Call(m, "IndentedJSON", arg0, arg1)
}

// IndentedJSON indicates an expected call of IndentedJSON
func (mr *MockContextMockRecorder) IndentedJSON(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndentedJSON", reflect.TypeOf((*MockContext)(nil).IndentedJSON), arg0, arg1)
}

// IsAborted mocks base method
func (m *MockContext) IsAborted() bool {
	ret := m.ctrl.Call(m, "IsAborted")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAborted indicates an expected call of IsAborted
func (mr *MockContextMockRecorder) IsAborted() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAborted", reflect.TypeOf((*MockContext)(nil).IsAborted))
}

// IsWebsocket mocks base method
func (m *MockContext) IsWebsocket() bool {
	ret := m.ctrl.Call(m, "IsWebsocket")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsWebsocket indicates an expected call of IsWebsocket
func (mr *MockContextMockRecorder) IsWebsocket() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWebsocket", reflect.TypeOf((*MockContext)(nil).IsWebsocket))
}

// JSON mocks base method
func (m *MockContext) JSON(arg0 int, arg1 interface{}) {
	m.ctrl.Call(m, "JSON", arg0, arg1)
}

// JSON indicates an expected call of JSON
func (mr *MockContextMockRecorder) JSON(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockContext)(nil).JSON), arg0, arg1)
}

// JSONP mocks base method
func (m *MockContext) JSONP(arg0 int, arg1 interface{}) {
	m.ctrl.Call(m, "JSONP", arg0, arg1)
}

// JSONP indicates an expected call of JSONP
func (mr *MockContextMockRecorder) JSONP(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSONP", reflect.TypeOf((*MockContext)(nil).JSONP), arg0, arg1)
}

// MultipartForm mocks base method
func (m *MockContext) MultipartForm() (*multipart.Form, error) {
	ret := m.ctrl.Call(m, "MultipartForm")
	ret0, _ := ret[0].(*multipart.Form)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultipartForm indicates an expected call of MultipartForm
func (mr *MockContextMockRecorder) MultipartForm() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipartForm", reflect.TypeOf((*MockContext)(nil).MultipartForm))
}

// MustBindWith mocks base method
func (m *MockContext) MustBindWith(arg0 interface{}, arg1 binding.Binding) error {
	ret := m.ctrl.Call(m, "MustBindWith", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MustBindWith indicates an expected call of MustBindWith
func (mr *MockContextMockRecorder) MustBindWith(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustBindWith", reflect.TypeOf((*MockContext)(nil).MustBindWith), arg0, arg1)
}

// MustGet mocks base method
func (m *MockContext) MustGet(arg0 string) interface{} {
	ret := m.ctrl.Call(m, "MustGet", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// MustGet indicates an expected call of MustGet
func (mr *MockContextMockRecorder) MustGet(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustGet", reflect.TypeOf((*MockContext)(nil).MustGet), arg0)
}

// Negotiate mocks base method
func (m *MockContext) Negotiate(arg0 int, arg1 gin.Negotiate) {
	m.ctrl.Call(m, "Negotiate", arg0, arg1)
}

// Negotiate indicates an expected call of Negotiate
func (mr *MockContextMockRecorder) Negotiate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Negotiate", reflect.TypeOf((*MockContext)(nil).Negotiate), arg0, arg1)
}

// NegotiateFormat mocks base method
func (m *MockContext) NegotiateFormat(arg0 ...string) string {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NegotiateFormat", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// NegotiateFormat indicates an expected call of NegotiateFormat
func (mr *MockContextMockRecorder) NegotiateFormat(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NegotiateFormat", reflect.TypeOf((*MockContext)(nil).NegotiateFormat), arg0...)
}

// Next mocks base method
func (m *MockContext) Next() {
	m.ctrl.Call(m, "Next")
}

// Next indicates an expected call of Next
func (mr *MockContextMockRecorder) Next() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockContext)(nil).Next))
}

// Param mocks base method
func (m *MockContext) Param(arg0 string) string {
	ret := m.ctrl.Call(m, "Param", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param
func (mr *MockContextMockRecorder) Param(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockContext)(nil).Param), arg0)
}

// PostForm mocks base method
func (m *MockContext) PostForm(arg0 string) string {
	ret := m.ctrl.Call(m, "PostForm", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// PostForm indicates an expected call of PostForm
func (mr *MockContextMockRecorder) PostForm(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostForm", reflect.TypeOf((*MockContext)(nil).PostForm), arg0)
}

// PostFormArray mocks base method
func (m *MockContext) PostFormArray(arg0 string) []string {
	ret := m.ctrl.Call(m, "PostFormArray", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// PostFormArray indicates an expected call of PostFormArray
func (mr *MockContextMockRecorder) PostFormArray(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostFormArray", reflect.TypeOf((*MockContext)(nil).PostFormArray), arg0)
}

// PostFormMap mocks base method
func (m *MockContext) PostFormMap(arg0 string) map[string]string {
	ret := m.ctrl.Call(m, "PostFormMap", arg0)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// PostFormMap indicates an expected call of PostFormMap
func (mr *MockContextMockRecorder) PostFormMap(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostFormMap", reflect.TypeOf((*MockContext)(nil).PostFormMap), arg0)
}

// Query mocks base method
func (m *MockContext) Query(arg0 string) string {
	ret := m.ctrl.Call(m, "Query", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Query indicates an expected call of Query
func (mr *MockContextMockRecorder) Query(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockContext)(nil).Query), arg0)
}

// QueryArray mocks base method
func (m *MockContext) QueryArray(arg0 string) []string {
	ret := m.ctrl.Call(m, "QueryArray", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// QueryArray indicates an expected call of QueryArray
func (mr *MockContextMockRecorder) QueryArray(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryArray", reflect.TypeOf((*MockContext)(nil).QueryArray), arg0)
}

// QueryMap mocks base method
func (m *MockContext) QueryMap(arg0 string) map[string]string {
	ret := m.ctrl.Call(m, "QueryMap", arg0)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// QueryMap indicates an expected call of QueryMap
func (mr *MockContextMockRecorder) QueryMap(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryMap", reflect.TypeOf((*MockContext)(nil).QueryMap), arg0)
}

// Redirect mocks base method
func (m *MockContext) Redirect(arg0 int, arg1 string) {
	m.ctrl.Call(m, "Redirect", arg0, arg1)
}

// Redirect indicates an expected call of Redirect
func (mr *MockContextMockRecorder) Redirect(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Redirect", reflect.TypeOf((*MockContext)(nil).Redirect), arg0, arg1)
}

// Render mocks base method
func (m *MockContext) Render(arg0 int, arg1 render.Render) {
	m.ctrl.Call(m, "Render", arg0, arg1)
}

// Render indicates an expected call of Render
func (mr *MockContextMockRecorder) Render(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockContext)(nil).Render), arg0, arg1)
}

// SSEvent mocks base method
func (m *MockContext) SSEvent(arg0 string, arg1 interface{}) {
	m.ctrl.Call(m, "SSEvent", arg0, arg1)
}

// SSEvent indicates an expected call of SSEvent
func (mr *MockContextMockRecorder) SSEvent(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SSEvent", reflect.TypeOf((*MockContext)(nil).SSEvent), arg0, arg1)
}

// SaveUploadedFile mocks base method
func (m *MockContext) SaveUploadedFile(arg0 *multipart.FileHeader, arg1 string) error {
	ret := m.ctrl.Call(m, "SaveUploadedFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUploadedFile indicates an expected call of SaveUploadedFile
func (mr *MockContextMockRecorder) SaveUploadedFile(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUploadedFile", reflect.TypeOf((*MockContext)(nil).SaveUploadedFile), arg0, arg1)
}

// SecureJSON mocks base method
func (m *MockContext) SecureJSON(arg0 int, arg1 interface{}) {
	m.ctrl.Call(m, "SecureJSON", arg0, arg1)
}

// SecureJSON indicates an expected call of SecureJSON
func (mr *MockContextMockRecorder) SecureJSON(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecureJSON", reflect.TypeOf((*MockContext)(nil).SecureJSON), arg0, arg1)
}

// Set mocks base method
func (m *MockContext) Set(arg0 string, arg1 interface{}) {
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set
func (mr *MockContextMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockContext)(nil).Set), arg0, arg1)
}

// SetAccepted mocks base method
func (m *MockContext) SetAccepted(arg0 ...string) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetAccepted", varargs...)
}

// SetAccepted indicates an expected call of SetAccepted
func (mr *MockContextMockRecorder) SetAccepted(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccepted", reflect.TypeOf((*MockContext)(nil).SetAccepted), arg0...)
}

// SetCookie mocks base method
func (m *MockContext) SetCookie(arg0, arg1 string, arg2 int, arg3, arg4 string, arg5, arg6 bool) {
	m.ctrl.Call(m, "SetCookie", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// SetCookie indicates an expected call of SetCookie
func (mr *MockContextMockRecorder) SetCookie(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCookie", reflect.TypeOf((*MockContext)(nil).SetCookie), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ShouldBind mocks base method
func (m *MockContext) ShouldBind(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "ShouldBind", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShouldBind indicates an expected call of ShouldBind
func (mr *MockContextMockRecorder) ShouldBind(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldBind", reflect.TypeOf((*MockContext)(nil).ShouldBind), arg0)
}

// ShouldBindBodyWith mocks base method
func (m *MockContext) ShouldBindBodyWith(arg0 interface{}, arg1 binding.BindingBody) error {
	ret := m.ctrl.Call(m, "ShouldBindBodyWith", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShouldBindBodyWith indicates an expected call of ShouldBindBodyWith
func (mr *MockContextMockRecorder) ShouldBindBodyWith(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldBindBodyWith", reflect.TypeOf((*MockContext)(nil).ShouldBindBodyWith), arg0, arg1)
}

// ShouldBindJSON mocks base method
func (m *MockContext) ShouldBindJSON(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "ShouldBindJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShouldBindJSON indicates an expected call of ShouldBindJSON
func (mr *MockContextMockRecorder) ShouldBindJSON(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldBindJSON", reflect.TypeOf((*MockContext)(nil).ShouldBindJSON), arg0)
}

// ShouldBindQuery mocks base method
func (m *MockContext) ShouldBindQuery(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "ShouldBindQuery", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShouldBindQuery indicates an expected call of ShouldBindQuery
func (mr *MockContextMockRecorder) ShouldBindQuery(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldBindQuery", reflect.TypeOf((*MockContext)(nil).ShouldBindQuery), arg0)
}

// ShouldBindWith mocks base method
func (m *MockContext) ShouldBindWith(arg0 interface{}, arg1 binding.Binding) error {
	ret := m.ctrl.Call(m, "ShouldBindWith", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShouldBindWith indicates an expected call of ShouldBindWith
func (mr *MockContextMockRecorder) ShouldBindWith(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldBindWith", reflect.TypeOf((*MockContext)(nil).ShouldBindWith), arg0, arg1)
}

// Status mocks base method
func (m *MockContext) Status(arg0 int) {
	m.ctrl.Call(m, "Status", arg0)
}

// Status indicates an expected call of Status
func (mr *MockContextMockRecorder) Status(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockContext)(nil).Status), arg0)
}

// Stream mocks base method
func (m *MockContext) Stream(arg0 func(io.Writer) bool) {
	m.ctrl.Call(m, "Stream", arg0)
}

// Stream indicates an expected call of Stream
func (mr *MockContextMockRecorder) Stream(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockContext)(nil).Stream), arg0)
}

// String mocks base method
func (m *MockContext) String(arg0 int, arg1 string, arg2 ...interface{}) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "String", varargs...)
}

// String indicates an expected call of String
func (mr *MockContextMockRecorder) String(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockContext)(nil).String), varargs...)
}

// Value mocks base method
func (m *MockContext) Value(arg0 interface{}) interface{} {
	ret := m.ctrl.Call(m, "Value", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockContextMockRecorder) Value(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockContext)(nil).Value), arg0)
}

// XML mocks base method
func (m *MockContext) XML(arg0 int, arg1 interface{}) {
	m.ctrl.Call(m, "XML", arg0, arg1)
}

// XML indicates an expected call of XML
func (mr *MockContextMockRecorder) XML(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XML", reflect.TypeOf((*MockContext)(nil).XML), arg0, arg1)
}

// YAML mocks base method
func (m *MockContext) YAML(arg0 int, arg1 interface{}) {
	m.ctrl.Call(m, "YAML", arg0, arg1)
}

// YAML indicates an expected call of YAML
func (mr *MockContextMockRecorder) YAML(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "YAML", reflect.TypeOf((*MockContext)(nil).YAML), arg0, arg1)
}