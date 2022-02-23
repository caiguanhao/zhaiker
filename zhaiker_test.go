package zhaiker

import (
	"bytes"
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

	managerId, apiKey := MustGetKey(ctx, accountName, accountPass)
	t.Log("managerId:", managerId)
	t.Log("apiKey:", apiKey)

	client := NewClient(managerId, apiKey)
	var gyms []Gym
	client.MustRequest(ctx, API_GET_GYMS, Params("pageSize", 100), &gyms, "DATA.*")
	gymsBytes, err := json.Marshal(gyms)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("gyms:", string(gymsBytes))

	var gymsDup []Gym
	if err := json.Unmarshal(gymsBytes, &gymsDup); err != nil {
		t.Fatal(err)
	}
	gymsBytesDup, err := json.Marshal(gymsDup)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(gymsBytes, gymsBytesDup) {
		t.Log("equal")
	} else {
		t.Error("not equal")
	}

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
	examBytes, err := json.Marshal(exams[index])
	if err != nil {
		t.Fatal(err)
	}
	t.Log("example exam:", string(examBytes))

	var examDup Exam
	if err := json.Unmarshal(examBytes, &examDup); err != nil {
		t.Fatal(err)
	}
	examBytesDup, err := json.Marshal(examDup)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(examBytes, examBytesDup) {
		t.Log("equal")
	} else {
		t.Error("not equal")
	}

	for i, exam := range exams {
		if exam.BodyDetect != nil && exam.Age < 18 {
			index = i
			break
		}
	}

	var ews ExamWithStandards
	client.MustRequest(ctx, API_GET_EXAM, Params("gymId", gyms[0].Id, "id", exams[index].Id), &ews)
	ewsBytes, err := json.Marshal(ews)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("example exam:", string(ewsBytes))

	var ewsDup ExamWithStandards
	if err := json.Unmarshal(ewsBytes, &ewsDup); err != nil {
		t.Fatal(err)
	}
	ewsBytesDup, err := json.Marshal(ewsDup)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(ewsBytes, ewsBytesDup) {
		t.Log("equal")
	} else {
		t.Error("not equal")
	}

	var device Device
	client.MustRequest(ctx, API_GET_DEVICE, Params("gymId", gyms[0].Id, "device", exams[index].DeviceFrom), &device, "DATA")
	deviceBytes, err := json.Marshal(device)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("device:", string(deviceBytes))

	var deviceDup Device
	if err := json.Unmarshal(deviceBytes, &deviceDup); err != nil {
		t.Fatal(err)
	}
	deviceBytesDup, err := json.Marshal(deviceDup)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(deviceBytes, deviceBytesDup) {
		t.Log("equal")
	} else {
		t.Error("not equal")
	}

	if image := exams[index].BodyImage; image != "" {
		url := MustGetFileURL(ctx, image)
		t.Log("url:", url)
	} else {
		t.Log("no image to test")
	}
}
