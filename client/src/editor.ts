export function setupEditor(editor: HTMLTextAreaElement) {
	document.addEventListener('DOMContentLoaded', (_event: any) => {
			editor.addEventListener('keydown', (event: KeyboardEvent) => {
					if (event.key === 'Tab') {
							event.preventDefault(); // Prevent the default tab behavior
							const start = editor.selectionStart;
							const end = editor.selectionEnd;
							const value = editor.value;

							// Insert four spaces at the current cursor position
							editor.value = value.substring(0, start) + '    ' + value.substring(end);

							// Move the cursor to the correct position
							editor.selectionStart = editor.selectionEnd = start + 4;
					}
			});
	});
}
