package golang

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	_ "embed"
)

//go:embed prompt.txt
var defaultPrompt string

func explainGoPackageIn(dir string, excluded []string, out string, promptFilename string) error {
	ctx := context.Background()

	modelName := os.Getenv("GEMINI_MODEL")
	if modelName == "" {
		modelName = "gemini-1.5-pro"
	}

	prompt := defaultPrompt
	if promptFilename != "" {
		log.Println("reading prompt from", promptFilename)
		data, err := os.ReadFile(promptFilename)
		if err != nil {
			return err
		}
		prompt = string(data)
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return err
	}
	defer client.Close()

	ctxParts := []genai.Part{genai.Text("you are an expert Go programmer")}
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), "_test.go") {
			log.Println("excluding test files (*_test.go)", info.Name())
			return nil
		}
		if strings.Contains(path, "test") {
			log.Println("excluding files with 'test' in path", path)
			return nil
		}
		if strings.Contains(path, "/example") {
			log.Println("excluding files with '/example' in path", path)
			return nil
		}
		if strings.Contains(path, "/demo") {
			log.Println("excluding files with '/demo' in path", path)
			return nil
		}
		// allow /cmd
		// check excludes
		for _, each := range excluded {
			rel, _ := filepath.Rel(dir, path)
			if strings.HasPrefix(rel, each) {
				return nil
			}
		}

		if strings.HasSuffix(info.Name(), ".go") {
			log.Println("adding source to the context", path)
			src, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file:%w", err)
			}
			bodyless, err := sourceWithoutFunctionBodies(path, string(src))
			if err != nil {
				return fmt.Errorf("failed to parse file:%w", err)
			}
			ctxParts = append(ctxParts, genai.Text(bodyless))
		}
		return nil
	})

	model := client.GenerativeModel(modelName)
	model.SystemInstruction = &genai.Content{
		Parts: ctxParts,
	}
	log.Println("requesting AI to describe the software using", modelName)
	if prompt == "" {
		return errors.New("missing prompt")
	}
	resp, err := model.GenerateContent(ctx,
		genai.Text(prompt),
	)
	if err != nil {
		return err
	}
	return printResponse(out, resp)

}

func printResponse(output string, resp *genai.GenerateContentResponse) error {
	log.Println("writing LLM response to", output)
	w, err := os.Create(output)
	if err != nil {
		return err
	}
	defer w.Close()
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Fprintln(w, part)
			}
		}
	}
	log.Printf("usage: %#v\n", resp.UsageMetadata)
	return nil
}
