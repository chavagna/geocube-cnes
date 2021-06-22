// Code generated by "enumer -json -sql -type TaskState -trimprefix TaskState"; DO NOT EDIT.

//
package geocube

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _TaskStateName = "FAILEDCANCELLEDDONEPENDING"

var _TaskStateIndex = [...]uint8{0, 6, 15, 19, 26}

func (i TaskState) String() string {
	if i < 0 || i >= TaskState(len(_TaskStateIndex)-1) {
		return fmt.Sprintf("TaskState(%d)", i)
	}
	return _TaskStateName[_TaskStateIndex[i]:_TaskStateIndex[i+1]]
}

var _TaskStateValues = []TaskState{0, 1, 2, 3}

var _TaskStateNameToValueMap = map[string]TaskState{
	_TaskStateName[0:6]:   0,
	_TaskStateName[6:15]:  1,
	_TaskStateName[15:19]: 2,
	_TaskStateName[19:26]: 3,
}

// TaskStateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TaskStateString(s string) (TaskState, error) {
	if val, ok := _TaskStateNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TaskState values", s)
}

// TaskStateValues returns all values of the enum
func TaskStateValues() []TaskState {
	return _TaskStateValues
}

// IsATaskState returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TaskState) IsATaskState() bool {
	for _, v := range _TaskStateValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for TaskState
func (i TaskState) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for TaskState
func (i *TaskState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TaskState should be a string, got %s", data)
	}

	var err error
	*i, err = TaskStateString(s)
	return err
}

func (i TaskState) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *TaskState) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := TaskStateString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
