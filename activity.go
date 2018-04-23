package azure

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"io/ioutil"
	"net/http"
	"strings"
)

var log = logger.GetLogger("activity-queue")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval
	MethodPut := "PUT"
	StatusCreated := 201
	queuename := strings.ToLower(context.GetInput("QueueName").(string))
	sastoken := context.GetInput("SAS_token").(string)
	storagename := context.GetInput("StorageName").(string)

	log.Debugf("The values set to queuename:[%s] and sas_token:[%s]", queuename, sastoken)

	url := "https://" + storagename + ".queue.core.windows.net/" + queuename + sastoken

	request, err := http.NewRequest(MethodPut, url, strings.NewReader("<golang>really</golang>"))
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Error(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		fmt.Println("   ", response.StatusCode)
		if response.StatusCode == StatusCreated {
//			fmt.Println("QueueCreted Successfully!!!")
			context.SetOutput("Status", "QueueCreated Successfully!!!")
		} else {
//			fmt.Println("Queue Not created.Either QueueName already exist")
			context.SetOutput("Status", "Queue Not created.Either QueueName already exist")
		}
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		}
//		fmt.Println(contents)

	}
	return true, nil
}