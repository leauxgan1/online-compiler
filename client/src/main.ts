import './style.css'
//import { setupCounter } from './counter.ts'
import {setupEditor} from './editor.ts'

//setupCounter(document.querySelector<HTMLButtonElement>("#counter")!)
setupEditor(document.querySelector<HTMLTextAreaElement>("#editor")!)
