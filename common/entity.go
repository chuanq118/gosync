package common

// RunMode is a string type to store run mode.
type RunMode string

// Watch :the program will watch the file system using configuration defined in config home.
const Watch RunMode = "watch"

// Batch :the program will run in batch mode. Only sync the specific directory.
const Batch RunMode = "batch"

// CloudObject is a struct that describes a file object stored on a remote cloud server.
type CloudObject struct {
	Key   string `json:"key"`
	Bytes int64  `json:"bytes"`
	MD5   string `json:"md5"`
}
