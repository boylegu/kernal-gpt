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
