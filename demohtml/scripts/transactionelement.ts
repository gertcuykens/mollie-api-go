import {Transaction, transaction} from './mixin.js'
import {defineClass} from './decorator.js'

@defineClass('transaction-ts')
export default class TransactionElement extends transaction(HTMLElement) {
  constructor() {
    super()
  }
  render(json:{}) {
    //while (this.hasChildNodes()) this.removeChild(this.lastChild)
    this.innerHTML = 'Payment screen: '
    const t:Transaction = json as Transaction
    const a = document.createElement('a')
    a.href = localStorage.getItem('payment') || ''
    a.innerHTML = localStorage.getItem('payment') || ''
    const pre = document.createElement('pre')
    pre.innerHTML = JSON.stringify(t, null, '\t')
    this.appendChild(a)
    this.appendChild(pre)  
  }
  connectedCallback() {
    this.transaction(localStorage.getItem('transaction')||'')
      .then(json => this.render(json))
      .catch(err => console.log(err))
  }
  disconnectedCallback() { }
  attributeChangedCallback(name: string, oldValue: string, newValue: string) { }
  adoptedCallback() { }
}
