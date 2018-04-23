package azure

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("QueueName", "queue2")
	tc.SetInput("StorageName", "gauravkayande")
	tc.SetInput("SAS_token", "?sv=2017-07-29&ss=bfqt&srt=sco&sp=rwdlacup&se=2018-05-10T15:33:37Z&st=2018-04-23T07:33:37Z&spr=https,http&sig=lwwnU%2B9tvVhwhhvxwnxEUaH4Bgq8O75jNmPinB3T0ak%3D")

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("Status")
	assert.Equal(t, result, "QueueCreated Successfully!!!")
}