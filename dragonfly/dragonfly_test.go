package dragonfly_test

import (
	//"fmt"
	"github.com/denniscollective/dragonfly.go/dragonfly"
	"testing"
)

const stubB64Job string = "W1siZmYiLCIvVXNlcnMvZGVubmlzL3dvcmtzcGFjZS96aXZpdHkvcHVibGljL2NvbnRlbnQvcGhvdG9zZXRzL29yaWdpbmFsc19hcmNoaXZlLzAwMC8wMDAwMDAvMDAwMDAwMDA3LzAwMDAwMDAwMjQtaC1vcmlnLmpwZyJdLFsicCIsInRodW1iIiwiMjB4MjAiXV0"

func TestFetch(t *testing.T) {
	job, _ := dragonfly.Decode(stubB64Job)
	file, err := job.Apply()

	if err != nil {
		t.Error("job.Apply failed")
	}

	if len(file.Name()) < 10 {
		t.Error("expected a file Object")
	}

}

func TestDecodeThingThatNeedsTwoEquals(t *testing.T) {
	jobstr := "W1siZmYiLCIvVXNlcnMvZGVubmlzL3dvcmtzcGFjZS96aXZpdHkvcHVibGljL2ltYWdlcy9pY29ucy9kZWZhdWx0XzI1Ni5qcGciXSxbInAiLCJ0aHVtYiIsIjgweDgwIyJdXQ"
	job, err := dragonfly.Decode(jobstr)

	if err != nil {
		t.Errorf("Deconde job got error %s", err)
	}

	if len(job.Steps) != 2 {
		t.Error("job should have two steps")
	}

}

func TestDecodeDragonfly(t *testing.T) {
	job, err := dragonfly.Decode(stubB64Job)

	if err != nil {
		t.Errorf("Deconde job got error %s", err)
	}

	if len(job.Steps) != 2 {
		t.Error("job should have two steps")
	}

	if job.Steps[0].Command != "ff" {
		t.Error("the first test of the stub is supposed to be fetch File")
	}

	if args := job.Steps[1].Args; args[0] != "thumb" && args[1] != "20x20" {
		t.Error("second step should be a resize to thumbnail 20x20 job")
	}
}

func TestDecodeFailse(t *testing.T) {
	t.Skip()
	job, err := dragonfly.Decode("this is y i'm hawt")
	if err == nil {
		t.Error("Decode errors aren't propagating")
	}

	if job != nil {
		t.Error("Decode should return nil when it has an error")
	}
}
