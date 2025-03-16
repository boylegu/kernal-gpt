<div align="center">

<p align="center">
<img src="https://cdn.jsdelivr.net/gh/boylegu/kernal-gpt/assets/kernal_gpt.png" width="360" height="300">
</p>

<h1 style="border-bottom: none">
    <b>Kernal-GPT</b><br />
</h1>

<p>
An AI agent based on the Ollama large model, capable of executing Linux commands through natural language and invoking kernel hooks to delve into the underlying system.
</p>

[![go](https://img.shields.io/badge/Go-1.24+-66C9D6)]()
[![ver](https://img.shields.io/badge/version-0.3.0.dev-E940AF)]()
[![go](https://img.shields.io/badge/license-MIT-E940AF)]()
</div>

## Abilities & Possibilities

[✔] Support all large models locally in Ollama

[✔] langchain and langgraph integration

[✔] multi-modal LLM's support 

[✔] only tool calling support

[✔] redis-vector and RAG integration

[✔] eBPF integration (Only supports Linux kernel 5.10 and above.)

[✔] dangerous command detection

[×] support memory and context （Coming soon）

[×] It will not run directly at the moment.（Currently, it is an experimental version and not suitable for direct execution.）

## Mechanism

<p align="center">
<img src="https://cdn.jsdelivr.net/gh/boylegu/kernal-gpt/assets/black.png">
</p>

## Usage example

- Operate Linux commands using natural language.

<p align="center">
<img src="https://cdn.jsdelivr.net/gh/boylegu/kernal-gpt/assets/oscmd_en_eg01.gif">
</p>

dangerous command detection

<p align="center">
<img src="https://cdn.jsdelivr.net/gh/boylegu/kernal-gpt/assets/oscmd_en_eg02.gif">
</p>

- Operate the Linux kernel using natural language, including automatically generating eBPF scripts.

<p align="center">
<img src="https://cdn.jsdelivr.net/gh/boylegu/kernal-gpt/assets/bpf_en_eg01.gif">
</p>

We use examples from bpftrace tools to create vector store and search.


<p align="center">
<img src="https://cdn.jsdelivr.net/gh/boylegu/kernal-gpt/assets/bpf_en_eg02.gif">
</p>

## How to run

1. This project uses Redis vector as the vector database, so Redis must be started first.

```
docker run -p 6379:6379 docker.io/redislabs/redisearch:latest 
```

2. Set environment variables

```bash

export KPT_MODEL="" # "please select the llm
export KPT_OLLAMA_URL="http://127.0.0.1:11434"
export KPT_REDIS_URL="redis://127.0.0.1:6379"
```


3. Compile the code

```
make
```

>> The artifacts are in ./dist.

## the original intention of this project

This project is mainly to validate the feasibility of AI combined with eBPF for impacting the underlying infrastructure and its future potential. You can also learn how to use Golang for AI Agent related development. The above examples run on qwen2.5:1.5b. I believe the actual results for you will be even better than mine. I also welcome discussions with you. If you're interested, feel free to submit a PR.