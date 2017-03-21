import {Order, Product} from 'mixin'

const testJson = {
  "Email": "test@test",
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
    "Price": 5.99,
    "Currency": ""
  }]
} as Order

localStorage.setObject('test', testJson)

export function defineClass(tagname: string) {
  return function <T extends { new (...args: any[]): HTMLElement }>(constructor: T) {
    console.log("Define: " + constructor.name)
    window.customElements.define(tagname, constructor)
    return class extends constructor {
      state = testJson
    }
  }
}
