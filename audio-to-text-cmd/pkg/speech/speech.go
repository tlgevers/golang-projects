package speech

import (
	speech "cloud.google.com/go/speech/apiv1"
	"context"
	"fmt"
	"google.golang.org/api/option"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
	"os"
)

func client() (client *speech.Client, err error) {
	ctx := context.Background()
	client, err = speech.NewClient(ctx, option.WithCredentialsFile("service-account.json"))
	if err != nil {
		return
	}
	return
}

func Convert(filename string, output string) {
	ctx := context.Background()
	client, err := client()
	if err != nil {
		panic(err)
	}
	req := &speechpb.LongRunningRecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_FLAC,
			SampleRateHertz: 16000,
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: filename},
		},
	}
	op, err := client.LongRunningRecognize(ctx, req)
	if err != nil {
		panic(err)
	}
	resp, err := op.Wait(ctx)
	if err != nil {
		panic(err)
	}
	// Write the results.
	f, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
			f.WriteString(alt.Transcript + "\n")
		}
	}
}
