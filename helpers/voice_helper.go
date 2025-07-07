package helpers

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	vosk "github.com/alphacep/vosk-api/go"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func FormatVoiceToText(ctx *th.Context, voice *telego.File) string {
	//fileURL := ctx.Bot().FileDownloadURL() + "file/" + voice.FileID

	//return voice.FilePath // voice/file_1.oga

	const model = "./models/vosk-model-small-ru-0.22.zip"

	Recognize(model, convertOgaToWav(voice))

	return ""

	/*resp, err := http.Get(voice.FilePath)
	if err != nil {
		log.Println("Ошибка загрузки файла:", err)
		return ""
	}

	fmt.Println(resp.Request.URL)

	defer resp.Body.Close()

	return resp.Request.RequestURI*/

	/*fileBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения тела ответа:", err)
		return ""
	}

	return ""*/
}

/*
не работает fatal error: 'vosk_api.h' file not found

	7 |  #include <vosk_api.h>
	  |           ^~~~~~~~~~~~
*/

// in terminal cd $HOME/go
// open .
// github/aplhacep and add file
func Recognize(modelPath string, voiceFile string) {
	var filename string
	flag.StringVar(&filename, "f", "", voiceFile)
	flag.Parse()

	model, err := vosk.NewModel(modelPath)
	if err != nil {
		log.Fatal(err)
	}

	sampleRate := 16000.0
	rec, err := vosk.NewRecognizer(model, sampleRate)
	if err != nil {
		log.Fatal(err)
	}
	rec.SetWords(1)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 4096)

	for {
		_, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}

		if rec.AcceptWaveform(buf) != 0 {
			fmt.Println(rec.Result())
		}
	}

	var jres map[string]interface{}
	json.Unmarshal([]byte(rec.FinalResult()), &jres)
	fmt.Println(jres["text"])
}

func convertOgaToWav(voice *telego.File) string {
	inputPath := voice.FilePath
	outputPath := strings.Replace(voice.FilePath, ".oga", ".wav", 1)

	cmd := exec.Command("ffmpeg", "-i", inputPath, outputPath)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error converting OGA to WAV: %v", err)
	}

	return outputPath
}
