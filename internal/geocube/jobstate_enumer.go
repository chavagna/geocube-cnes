// Code generated by "enumer -json -sql -type JobState -trimprefix JobState"; DO NOT EDIT.

//
package geocube

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _JobStateName = "NEWCREATEDCONSOLIDATIONINPROGRESSCONSOLIDATIONDONECONSOLIDATIONINDEXEDCONSOLIDATIONEFFECTIVECONSOLIDATIONFAILEDCONSOLIDATIONRETRYINGCONSOLIDATIONCANCELLINGDELETIONINPROGRESSDELETIONEFFECTIVEDELETIONFAILEDDONEFAILEDINITIALISATIONFAILEDCANCELLATIONFAILEDABORTEDROLLBACKFAILEDDONEBUTUNTIDY"

var _JobStateIndex = [...]uint16{0, 3, 10, 33, 50, 70, 92, 111, 132, 155, 173, 190, 204, 208, 214, 234, 252, 259, 273, 286}

func (i JobState) String() string {
	if i < 0 || i >= JobState(len(_JobStateIndex)-1) {
		return fmt.Sprintf("JobState(%d)", i)
	}
	return _JobStateName[_JobStateIndex[i]:_JobStateIndex[i+1]]
}

var _JobStateValues = []JobState{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

var _JobStateNameToValueMap = map[string]JobState{
	_JobStateName[0:3]:     0,
	_JobStateName[3:10]:    1,
	_JobStateName[10:33]:   2,
	_JobStateName[33:50]:   3,
	_JobStateName[50:70]:   4,
	_JobStateName[70:92]:   5,
	_JobStateName[92:111]:  6,
	_JobStateName[111:132]: 7,
	_JobStateName[132:155]: 8,
	_JobStateName[155:173]: 9,
	_JobStateName[173:190]: 10,
	_JobStateName[190:204]: 11,
	_JobStateName[204:208]: 12,
	_JobStateName[208:214]: 13,
	_JobStateName[214:234]: 14,
	_JobStateName[234:252]: 15,
	_JobStateName[252:259]: 16,
	_JobStateName[259:273]: 17,
	_JobStateName[273:286]: 18,
}

// JobStateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func JobStateString(s string) (JobState, error) {
	if val, ok := _JobStateNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to JobState values", s)
}

// JobStateValues returns all values of the enum
func JobStateValues() []JobState {
	return _JobStateValues
}

// IsAJobState returns "true" if the value is listed in the enum definition. "false" otherwise
func (i JobState) IsAJobState() bool {
	for _, v := range _JobStateValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for JobState
func (i JobState) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for JobState
func (i *JobState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("JobState should be a string, got %s", data)
	}

	var err error
	*i, err = JobStateString(s)
	return err
}

func (i JobState) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *JobState) Scan(value interface{}) error {
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

	val, err := JobStateString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
