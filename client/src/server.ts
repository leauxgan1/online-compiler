
export async function compileCode() {
	const code = (document.getElementById('code') as HTMLTextAreaElement).value;
	const response = await fetch('/compile', {
			method: 'POST',
			headers: {
					'Content-Type': 'application/json'
			},
			body: JSON.stringify({ code: code })
	});

	const result = await response.json();
	const outputElement = document.getElementById('output');
	if (outputElement) {
			outputElement.innerText = result.output + result.error;
	}
}

export async function runCode() {
	console.log("PUSHED RUN");
}

(window as any).compileCode = compileCode;
(window as any).runCode = runCode;

