<div align="center">
<h1 style="border-bottom: none">
    <b>kernal-GPT</b><br />
</h1>

<p>
An AI agent based on the Ollama large model, capable of executing Linux commands through natural language and invoking kernel hooks to delve into the underlying system.
</p>
<p align="center">
<img src="https://cdn.jsdelivr.net/gh/boylegu/kernal-gpt/assets/kernal_gpt.png" width="360" height="300">
</p>
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