/// <amd-module name="formelement_test"/>
import FormElement from 'formelement'
import {defineClass} from 'decorator'
import {Payment, Order, Product, payment, render} from 'mixin'

@defineClass('form-test')
export default class FormElement_test extends FormElement {
  state: Order = {
    "Email": "test@test -,_.!'~*()",
    "Method": "",
    "Issuer": "",
    "Product": [{
      "Name": "test",
      "Description": "",
      "Quantity": 2.01,
      "Price": 1.99,
      "Currency": ""
    },
    {
      "Name": "test2",
      "Description": "",
      "Quantity": 4.01,
      "Price": 55.99,
      "Currency": ""
    }]
  }
  constructor() {
    super()
  }
  connectedCallback() {
    const pre = document.createElement('pre') 
    pre.innerHTML = JSON.stringify(this.state, null, '\t')
    const p = document.createElement('p') 
    p.innerHTML = 'Total: '+ this.total()
    this.appendChild(pre)    
    this.appendChild(p)
  }
  disconnectedCallback() { }
  attributeChangedCallback(name: string, oldValue: string, newValue: string) { }
  adoptedCallback() { }
}