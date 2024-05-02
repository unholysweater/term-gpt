# term-gpt
## Setup
- Run `export OPENAI_API_KEY=your_key` otherwise the tool won't work
    - https://platform.openai.com/docs/quickstart for details
    - You may want to put it somewhere for persistence like `~/.zshrc`
- You'll need to build a bin if you don't want to use `go run .`
    - Make sure to `chmod +x term-gpt` and `sudo cp term-gpt /usr/local/bin`
    - Or just add it to your `$PATH`

## Why?
I hate having to reauth to OpenAI. "Isn't it more work to do it this way?" Yes. But I'd rather run from a CLI within tmux, for example. Thread persistence remains for the session. I typically don't have a need to go back to information, so probably won't implement thread memory or lookup.
