package gpt

import "fmt"

func ConstructRunningPrompt(input string) string {
	examples := Retriever(input)
	return fmt.Sprintf(`
	As a supportive assistant to a Linux system administrator,
		your role involves leveraging bpftrace to generate eBPF code that aids
	in problem-solving, as well as responding to queries.
		Note that you may not always need to call the bpftrace tool function.
		Here are some pertinent examples that align with the user's requests:

	%s

	Now, you have received the following request from a user: %s
	Please utilize your capabilities to the fullest extent to accomplish this task.
	`, examples, input)

}

var entityPrompt = `Please determine whether the following user input corresponds to a scenario that can be converted into a Linux operating system command or an eBPF scenario. If it's an operating system command scenario, return "oscmd"; if it's an eBPF scenario, return "ebpf".

Scenarios that can be converted into operating system commands, such as:

"List all files in the current directory"
"Check which process is consuming the most CPU resources"
"View the local network IP address", etc.
Scenarios that can be converted into eBPF, such as:

"Monitor the establishment and closure of all TCP connections"
"Filter out all network packets from a specific IP"
"Track CPU usage for each core", etc.
Operating System Commands: Actions typically executed in user space, involving user-space tools or commands (such as ls, ps, ifconfig, etc.)

eBPF Scenarios: Actions typically executed in kernel space, involving kernel event processing or data capture (such as network packet filtering, system call tracing, etc.).

User Intent

Operating system commands: The user's intent is to perform specific system administration tasks.
eBPF scenarios: The user's intent is to monitor, analyze, or optimize system performance.
User input: {{.user_input}}`

var osCmdPrompt = `Welcome to the natural language command execution tool, rather than script-based execution. Please describe the operating system command or task you want to perform.
For example, you can input: "List all files in the current directory" or "Create a folder named 'test'."
If there are subdirectories, please consider them as well.
Please note that I can only help you execute commands related to the operating system, such as file operations or system information queries.
Do not input commands that may harm the system or change data, such as deleting, moving, copying, or killing processes.
"Please return the raw operating system command that performs the following task, without any code block formatting or extra characters. For example, if I ask to list files in a directory, return only the command itself like ls /tmp/*."

Please determine whether the following user input is a valid operating system command. If it is, return the command; if it could potentially harm the system or change data, return "danger-command"; if it cannot be converted into an operating system command, return "non-command."

User input: {{.user_input}}`
