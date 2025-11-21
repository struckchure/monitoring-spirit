# Monitoring Spirit

Monitoring Spirit is a CLI tool designed to generate work reports based on your git commit history using AI. It leverages Large Language Models (LLMs) to summarize your contributions and present them in a readable format.

## Features

- **AI-Powered Summaries**: Uses AI to analyze and summarize git commits.
- **Multiple Providers**: Supports Ollama and LLM Studio as AI providers.
- **Customizable Prompts**: Choose between default, technical, or non-technical prompt styles.
- **Flexible Filtering**: Filter commits by date range, specific commit hashes, or author email.

## Installation

Ensure you have Go installed on your machine.

```bash
go install github.com/struckchure/monitoring-spirit@latest
```

Or build from source:

```bash
git clone https://github.com/struckchure/monitoring-spirit.git
cd monitoring-spirit
go build -o ms cmd/*.go
```

## Usage

The basic command structure is:

```bash
ms [command] [flags]
```

### Global Flags

| Flag | Shorthand | Description | Default |
|------|-----------|-------------|---------|
| `--model` | `-m` | The AI model to use (e.g., `llama3`, `mistral`) | |
| `--api-url` | `-u` | The API URL for the provider | |
| `--api-key` | `-k` | API Key (if required) | |
| `--api-provider` | `-p` | AI Provider (`ollama`, `llmstudio`) | `ollama` |
| `--prompt-type` | | Prompt style (`default`, `technical`, `non-technical`) | `default` |

### Commands

#### `report`

Generates a report based on git commits.

```bash
ms report [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--from` | Start commit hash |
| `--to` | End commit hash |
| `--since` | Start date (e.g., "2023-01-01") |
| `--until` | End date (e.g., "2023-01-31") |
| `--email` | Filter commits by author email |

### Examples

**Generate a report using Ollama (default):**

```bash
ms report --model llama3 --since "2023-10-01" --until "2023-10-07"
```

**Generate a report for a specific author:**

```bash
ms report --email "user@example.com" --since "yesterday"
```

**Use LLM Studio with a technical prompt:**

```bash
ms report --api-provider llmstudio --api-url "http://localhost:1234/v1" --prompt-type technical
```

## Configuration

You can set the `GIT_DIR` environment variable to specify the path to the git repository if you are not running the command from the root of the repo.

```bash
export GIT_DIR=/path/to/your/repo
ms report ...
```
