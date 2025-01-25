package providers

import (
	"fmt"
	"os"

	"github.com/aandrew-me/tgpt/v2/providers/blackboxai"
	"github.com/aandrew-me/tgpt/v2/providers/duckduckgo"
	"github.com/aandrew-me/tgpt/v2/providers/groq"
	"github.com/aandrew-me/tgpt/v2/providers/isou"
	"github.com/aandrew-me/tgpt/v2/providers/koboldai"
	"github.com/aandrew-me/tgpt/v2/providers/ollama"
	"github.com/aandrew-me/tgpt/v2/providers/openai"
	"github.com/aandrew-me/tgpt/v2/providers/phind"
	"github.com/aandrew-me/tgpt/v2/providers/pollinations"
	"github.com/aandrew-me/tgpt/v2/structs"
	http "github.com/bogdanfinn/fhttp"
)

var availableProviders = []string{
	"", "blackboxai", "duckduckgo", "isou", "groq", "koboldai", "ollama", "openai", "phind", "pollinations",
}

// GetMainText extracts the main text content from the response line based on the provider.
func GetMainText(line string, provider string, input string) string {
	switch provider {
	case "blackboxai":
		return blackboxai.GetMainText(line)
	case "duckduckgo":
		return duckduckgo.GetMainText(line)
	case "isou":
		return isou.GetMainText((line))
	case "groq":
		return groq.GetMainText(line)
	case "koboldai":
		return koboldai.GetMainText(line)
	case "ollama":
		return ollama.GetMainText(line)
	case "openai":
		return openai.GetMainText(line)
	case "pollinations":
		return pollinations.GetMainText(line)
	default:
		return phind.GetMainText(line)
	}
}

// NewRequest creates a new request to the specified provider with the given parameters and extra options.
// It validates the provider and then delegates the request creation to the appropriate provider-specific function.
func NewRequest(input string, params structs.Params, extraOptions structs.ExtraOptions) (*http.Response, error) {
	validProvider := false
	for _, str := range availableProviders {
		if str == params.Provider {
			validProvider = true
			break
		}
	}
	if !validProvider {
		fmt.Fprintln(os.Stderr, "Invalid provider")
		os.Exit(1)
	}

	switch params.Provider {
	case "blackboxai":
		return blackboxai.NewRequest(input, params)
	case "duckduckgo":
		return duckduckgo.NewRequest(input, params, params.PrevMessages)
	case "groq":
		return groq.NewRequest(input, params)
	case "isou":
		return isou.NewRequest(input, params)
	case "koboldai":
		return koboldai.NewRequest(input, params)
	case "ollama":
		return ollama.NewRequest(input, params)
	case "openai":
		return openai.NewRequest(input, params)
	case "pollinations":
		return pollinations.NewRequest(input, params)
	default:
		return phind.NewRequest(input, params)
	}

}
