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

© 2024, [ernestmicklei.com](http://ernestmicklei.com).  MIT License. Contributions welcome.