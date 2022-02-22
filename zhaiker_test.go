package zhaiker

import (
	"context"
	"encoding/json"
	"os"
	"testing"
)

func Test(t *testing.T) {
	accountName, accountPass := os.Getenv("ZHAIKER_NAME"), os.Getenv("ZHAIKER_PASS")
	if accountName == "" || accountPass == "" {
		t.Fatal("Please set ZHAIKER_NAME and ZHAIKER_PASS environment variables.")
	}
	debug = os.Getenv("DEBUG") == "1"

	ctx := context.Background()
	client := NewClient(MustGetKey(ctx, accountName, accountPass))
	var gyms []Gym
	client.MustRequest(ctx, API_GET_GYMS, Params("pageSize", 100), &gyms, "DATA.*")
	b, _ := json.Marshal(gyms)
	t.Log("gyms:", string(b))

	if len(gyms) == 0 {
		t.Log("no gym to test")
		return
	}

	var count int
	client.MustRequest(ctx, API_GET_EXAMS_COUNT, Params(), &count, "DATA")
	t.Log("total exams:", count)

	var exams []Exam
	client.MustRequest(ctx, API_GET_EXAMS, Params("gymId", gyms[0].Id), &exams, "DATA.*")
	if len(exams) == 0 {
		t.Log("no exam to test")
		return
	}
	index := 0
	for i, exam := range exams {
		if exam.BodyDetect != nil {
			index = i
			break
		}
	}
	b, _ = json.Marshal(exams[index])
	t.Log("example exam:", string(b))

	for i, exam := range exams {
		if exam.BodyDetect != nil && exam.Age < 18 {
			index = i
			break
		}
	}

	var exam ExamWithStandards
	client.MustRequest(ctx, API_GET_EXAM, Params("gymId", gyms[0].Id, "id", exams[index].Id), &exam)
	b, _ = json.Marshal(exam)
	t.Log("example exam:", string(b))

	var device Device
	client.MustRequest(ctx, API_GET_DEVICE, Params("gymId", gyms[0].Id, "device", exams[index].DeviceFrom), &device, "DATA")
	b, _ = json.Marshal(device)
	t.Log("device:", string(b))

	if image := exams[index].BodyImage; image != "" {
		url := MustGetFileURL(ctx, image)
		t.Log("url:", url)
	} else {
		t.Log("no image to test")
	}
}
