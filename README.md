<p align="center"><img src="tgpt.svg"></p>

# Terminal GPT (tgpt) 🚀

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/aandrew-me/tgpt)](https://github.com/aandrew-me/tgpt)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/aandrew-me/tgpt)](https://github.com/aandrew-me/tgpt/releases/latest)
![Arch Linux package](https://img.shields.io/archlinux/v/extra/x86_64/tgpt)
![Chocolatey Version](https://img.shields.io/chocolatey/v/tgpt)

tgpt is a cross-platform command-line interface (CLI) tool that allows you to use AI chatbot in your Terminal without requiring API keys. 

### Currently available providers: 
- [Blackbox AI](https://www.blackbox.ai/) (Blackbox model)
- [Duckduckgo](https://duckduckgo.com/aichat) (Supports several models)
- [Groq](https://groq.com/) (Requires a free API Key. LLaMA2-70b & Mixtral-8x7b)
- [KoboldAI](https://koboldai-koboldcpp-tiefighter.hf.space/)  (koboldcpp/HF_SPACE_Tiefighter-13B)
- [Ollama](https://www.ollama.com/) (Supports many models)
- [OpenAI](https://platform.openai.com/docs/guides/text-generation/chat-completions-api) (All models, Requires API Key, supports custom endpoints)
- [Phind](https://www.phind.com/agent) (Phind Model)

**Image Generation Model**: BlackBoxAi

## Usage 

```
Usage: tgpt [Flags] [Prompt]

Flags:
-s, --shell                                        Generate and Execute shell commands. (Experimental) 
-c, --code                                         Generate Code. (Experimental)
-q, --quiet                                        Gives response back without loading animation
-w, --whole                                        Gives response back as a whole text
-img, --image                                      Generate images from text
--provider                                         Set Provider. Detailed information has been provided below. (Env: AI_PROVIDER)

Some additional options can be set. However not all options are supported by all providers. Not supported options will just be ignored.
--model                                            Set Model
--key                                              Set API Key
--url                                              Set OpenAI API endpoint url
--temperature                                      Set temperature
--top_p                                            Set top_p
--max_length                                       Set max response length
--log                                              Set filepath to log conversation to (For interactive modes)
--preprompt                                        Set preprompt
-y                                                 Execute shell command without confirmation

Options:
-v, --version                                      Print version 
-h, --help                                         Print help message 
-i, --interactive                                  Start normal interactive mode 
-m, --multiline                                    Start multi-line interactive mode 
-cl, --changelog                                   See changelog of versions 
-u, --update                                       Update program 

Providers:
The default provider is phind. The AI_PROVIDER environment variable can be used to specify a different provider.
Available providers to use: blackboxai, duckduckgo, groq, koboldai, ollama, openai and phind

Provider: blackboxai
Uses BlackBox model. Great for developers

Provider: duckduckgo
Available models: gpt-4o-mini (default), meta-llama/Meta-Llama-3.1-70B-Instruct-Turbo, mistralai/Mixtral-8x7B-Instruct-v0.1, claude-3-haiku-20240307

Provider: groq
Requires a free API Key. Supports LLaMA2-70b & Mixtral-8x7b

Provider: koboldai
Uses koboldcpp/HF_SPACE_Tiefighter-13B only, answers from novels

Provider: ollama
Needs to be run locally. Supports many models

Provider: openai
Needs API key to work and supports various models. Recognizes the OPENAI_API_KEY and OPENAI_MODEL environment variables. Supports custom urls with --url

Provider: phind
Uses Phind Model. Great for developers

Examples:
tgpt "What is internet?"
tgpt -m
tgpt -s "How to update my system?"
tgpt --provider duckduckgo "What is 1+1"
tgpt --provider openai --key "sk-xxxx" --model "gpt-3.5-turbo" "What is 1+1"
cat install.sh | tgpt "Explain the code"
```

### New Flag Parsing Function

The new flag parsing function `parseFlags` is used to parse the command-line flags and returns the parsed values. This function improves the readability and maintainability of the code by separating the flag parsing logic from the main function.

#### Usage

To use the new flag parsing function, simply call it in the main function and use the returned values as needed.

```go
func main() {
    isQuiet, isWhole, isCode, isShell, isImage, isInteractive, isMultiline, isVersion, isHelp, isUpdate, isChangelog, prompt, logFile, shouldExecuteCommand := parseFlags()

    // Use the parsed flag values as needed
    if isVersion {
        fmt.Println("tgpt", localVersion)
    }
    // ...
}
```

### Examples and Use Cases

Here are some examples and use cases to help you understand how to use the application effectively:

1. **Basic Usage**: To get a response from the AI chatbot, simply provide a prompt as an argument.
    ```bash
    tgpt "What is the capital of France?"
    ```

2. **Interactive Mode**: Start the interactive mode to have a continuous conversation with the AI chatbot.
    ```bash
    tgpt -i
    ```

3. **Multi-line Interactive Mode**: Start the multi-line interactive mode to input longer prompts.
    ```bash
    tgpt -m
    ```

4. **Generate and Execute Shell Commands**: Use the `-s` flag to generate and execute shell commands.
    ```bash
    tgpt -s "How to list all files in a directory?"
    ```

5. **Generate Code**: Use the `-c` flag to generate code snippets.
    ```bash
    tgpt -c "Hello world in Python"
    ```

6. **Generate Images**: Use the `-img` flag to generate images from text.
    ```bash
    tgpt -img "A beautiful sunset over the mountains"
    ```

7. **Quiet Mode**: Use the `-q` flag to get a response without the loading animation.
    ```bash
    tgpt -q "What is the meaning of life?"
    ```

8. **Whole Text Mode**: Use the `-w` flag to get the response as a whole text.
    ```bash
    tgpt -w "Explain the theory of relativity"
    ```

9. **Set Provider**: Use the `--provider` flag to specify a different provider.
    ```bash
    tgpt --provider openai --key "sk-xxxx" --model "gpt-3.5-turbo" "What is 1+1"
    ```

10. **Log Conversation**: Use the `--log` flag to log the conversation to a file.
    ```bash
    tgpt --log "conversation.log" "What is the weather today?"
    ```

![demo](https://user-images.githubusercontent.com/66430340/233759296-c4cf8cf2-0cab-48aa-9e84-40765b823282.gif)

## Installation ⏬

### Download for GNU/Linux 🐧 or MacOS 🍎

The default download location is `/usr/local/bin`, but you can change it in the command to use a different location. However, make sure the location is added to your PATH environment variable for easy accessibility.

You can download it with the following command:

```bash
curl -sSL https://raw.githubusercontent.com/aandrew-me/tgpt/main/install | bash -s /usr/local/bin
```

If you are using Arch Linux, you can install with pacman:

```bash
pacman -S tgpt
```


### FreeBSD 😈 

To install the [port](https://www.freshports.org/www/tgpt):
```
cd /usr/ports/www/tgpt/ && make install clean
```
To install the package, run one of these commands:
```
pkg install www/tgpt
pkg install tgpt
```

### Install with Go
You need to [add the Go install directory to your system's shell path](https://go.dev/doc/tutorial/compile-install). 

```bash
go install github.com/aandrew-me/tgpt/v2@latest
```

### Windows 🪟

-   **Scoop:** Package installation with [Scoop](https://scoop.sh/) can be done using the following command:

    ```bash
    scoop install https://raw.githubusercontent.com/aandrew-me/tgpt/main/tgpt.json
    ```
- **Chocolatey** 
    ```bash
    choco install tgpt
    ```    

## Updating ⬆️
If you installed the program with the installation script, you may update it with
```bash
tgpt -u
```
**It may require admin privileges.**
### Proxy

Support:

1. environment variable

http_proxy or HTTP_PROXY with following available formats:

- Http Proxy [ `http://ip:port` ]
- Http Auth [ `http://user:pass@ip:port` ]
- Socks5 Proxy [ `socks5://ip:port ]`
- Socks5 Auth [ `socks5://user:pass@ip:port` ]

2. configuration file

file location in the following order:

- ./proxy.txt (in the same directory from where you are executing)
- ~/.config/tgpt/proxy.txt

Example:

```bash
http://127.0.0.1:8080
```

### From Release

You can download the executable for your operating system, rename it to `tgpt` (or any other desired name), and then execute it by typing `./tgpt` while in that directory. Alternatively, you can add it to your PATH environmental variable and then execute it by simply typing `tgpt`.


## Uninstalling
If you installed with the install script, you can execute the following command to remove the tgpt executable
```
sudo rm $(which tgpt)
```
Configuration file is usually located in `~/.config/tgpt` on GNU/Linux Systems and in `"Library/Application Support/tgpt"` on MacOS
