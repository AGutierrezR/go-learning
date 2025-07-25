package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Item struct represents a ToDo item
// first letter is lower case to make "private" to the package
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of ToDo items. It is a slice of item structs
// The first letter is upper case to make it "public" to the package
type List []item
