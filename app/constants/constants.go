package constants

const (
	// Input
	Input = "input"

	// Base path
	BasePath = "base_path"

	// Default base path
	DefaultBasePath = "."

	// Environment
	Env = "env"

	// Dev environment
	Development = "dev"

	// API auth field name
	APIAuthUserFieldName = "API"

	// ContainerID: k8s container id
	ContainerID = "container_id"

	// CommitID - git commit hash
	CommitID = "commit_id"

	// GitCommitHash - git commit hash env key
	GitCommitHash = "GIT_COMMIT_HASH"

	// Hostname: k8s container id env key
	Hostname = "HOSTNAME"

	// Logger: holds entry in context
	LOGGER = "logger"

	// RequestID: holds the unique request identifier for the request
	RequestID = "request_id"

	// TaskID: is the key to hold the task id of the process
	TaskID = "task_id"

	//Service: map which contains all the system/pod information
	Service = "service"

	//Context: map contains tracedata
	Context = "context"

	// Response: keys where response has to be written
	Request = "request"

	// AppMode: used it identify the application to run on debug mode
	AppMode = "APP_MODE"

	// Mode : cli flag to specify env for migrations
	Mode = "mode"

	// TaskID: is the key to hold the task id of the process
	X_RAZORPAY_TASK_ID = "X-Razorpay-TaskId"

	// Response: keys where response has to be written
	Response = "response"

	// Headers: key for headers
	Headers = "headers"

	// Message: message key used in response
	Message = "message"

	// Message: debug key used in response
	DEBUG_MESSAGE = "debug"

	// Code: key to hold the status code
	Code = "code"

	CreateSecret = "add_secret"
	GetSecret = "get_secret"
)

// PerformanceMetricsActions: List of actions which need performance metrics
var PerformanceMetricsActions = []string{CreateSecret,GetSecret}
