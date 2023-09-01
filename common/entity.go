package common

// RunMode is a string type to store run mode.
type RunMode string

// Watch :the program will watch the file system using configuration defined in config home.
const Watch RunMode = "watch"

// Batch :the program will run in batch mode. Only sync the specific directory.
const Batch RunMode = "batch"
