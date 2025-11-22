# Monitoring Spirit

Monitoring Spirit is a CLI tool designed to generate work reports based on your git commit history using AI. It leverages Large Language Models (LLMs) to summarize your contributions and present them in a readable format.

## Features

- **AI-Powered Summaries**: Uses AI to analyze and summarize git commits.
- **Multiple Providers**: Supports Ollama and LLM Studio as AI providers.
- **Customizable Prompts**: Choose between default, technical, or non-technical prompt styles.
- **Flexible Filtering**: Filter commits by date range, specific commit hashes, or author email.

## How It Works

```mermaid
graph TD
    Start[Start: ms report] --> Init[Initialize Config & Flags]
    Init --> LoadPrompt[Load System Prompt]
    Init --> GitOps[Read Git Repository]
    GitOps --> Filter[Filter Commits\n(Date, Author, etc.)]
    Filter --> Context[Prepare Context\n(Commit Messages)]
    LoadPrompt --> AI[AI Processing]
    Context --> AI
    AI --> Provider{AI Provider}
    Provider -->|Ollama| Ollama[Ollama API]
    Provider -->|LLM Studio| LLM[LLM Studio API]
    Ollama --> Response[Generate Report]
    LLM --> Response
    Response --> Output[Print Markdown to Stdout]
```

## Quick Install

You can install Monitoring Spirit using the provided installation scripts.

**Linux & macOS**

```bash
curl -fsSL https://raw.githubusercontent.com/struckchure/monitoring-spirit/main/scripts/install.sh | bash
```

**Windows (PowerShell)**

```powershell
irm https://raw.githubusercontent.com/struckchure/monitoring-spirit/main/scripts/install.ps1 -OutFile install.ps1
.\install.ps1
del install.ps1
```

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
| `--prompt-type` | | Prompt style (`default`, `technical`, `non-technical`, `neutral`) | `default` |
| `--prompt-dir` | | Directory containing custom prompt files | |

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

## Prompt Customization

Monitoring Spirit allows you to customize the prompts used to generate reports. You can use the built-in prompt types or provide your own custom prompts.

### Built-in Prompts

Use the `--prompt-type` flag to select a built-in prompt:

- **`default`**: A balanced report suitable for most stakeholders. It combines high-level summaries with enough technical context to be useful for mixed audiences.
- **`technical`**: Designed for engineers and technical leads. It focuses on implementation details, code structure, refactoring, and specific technologies used.
- **`non-technical`**: Tailored for product managers and non-technical stakeholders. It translates technical work into business value, focusing on features, stability, and user impact without jargon.
- **`neutral`**: A variation of the balanced report with a strictly neutral and objective tone, avoiding any subjective language.

### Custom Prompts

To use your own prompts, you can specify a directory containing markdown files with your custom prompts using the `--prompt-dir` flag.

1. Create a directory for your prompts (e.g., `my-prompts/`).
2. Create a markdown file inside that directory (e.g., `custom.md`).
3. Run the command specifying the directory and the prompt filename (without extension) as the prompt type.

```bash
ms report --prompt-dir ./my-prompts --prompt-type custom
```

The prompt file should contain instructions for the AI on how to format and generate the report.

