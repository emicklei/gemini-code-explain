# gemini-code-explain

Use Google Gemini API to produce a design description from a Go package.

## install

    go install github.com/emicklei/gemini-code-explain/cmd/gemini-code-explain@latest

## usage

    GEMINI_API_KEY=YOUR_TOKEN gemini-code-explain -gopkg github.com/emicklei/dot@v1.6.2

Optionally, you can override the model used, currently it is set to:

    GEMINI_MODEL=gemini-1.5-pro

Here is how to [get a Gemini API Key](https://ai.google.dev/gemini-api/docs/api-key).

## output examples

- [dot package](examples/dot@v1.6.2.md)
- [mox application](examples/mox@v0.0.11.md)
- [structexplorer tool](examples/structexplorer@v0.1.0.md)
- [proto package](examples/proto@v1.13.2.md)

## how it works

After downloading the package sources, I will create an LLM context window with all sources:
- skipping all *_test.go files
- stripping the bodies for each declared function found in the remaining *.go files
- asking the LLM model to `describe the design of this Go software` (prompt).

## custom prompt

    GEMINI_API_KEY=YOUR_TOKEN gemini-code-explain -prompt myprompt.txt -gopkg github.com/emicklei/dot@v1.6.2

© 2024, [ernestmicklei.com](http://ernestmicklei.com).  MIT License. Contributions welcome.