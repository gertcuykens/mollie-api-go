import {Order} from 'mixin'

const testJson = {
  "Email": "test@decorator",
  "Method": "",
  "Issuer": "",
  "Product": [{
    "Name": "decorator",
    "Description": "",
    "Quantity": 3.03,
    "Price": 7.96,
    "Currency": ""
  },
  {
    "Name": "test2",
    "Description": "",
    "Quantity": 2.09,
    "Price": 0.99,
    "Currency": ""
  }]
} as Order

// localStorage.setObject('test', testJson)

export function defineClass(tagname: string) {
  return function <T extends { new (...args: any[]): HTMLElement }>(constructor: T) {
    console.log("Define: " + constructor.name)
    window.customElements.define(tagname, constructor)
    return class extends constructor {
      state = testJson
    }
  }
}
