import {Order} from './mixin.js'

const testOrder = {
  "Email": "test@test -,_.!'~*()",
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

const emptyOrder = {
  "Email": "",
  "Method": "",
  "Issuer": "",
  "Product": []
} as Order

// localStorage.setObject('test', testJson)
// const storageOrder = localStorage.getObject('test') as Order

export function defineClass(tagname: string) {
  return function <T extends { new (...args: any[]): HTMLElement }>(constructor: T) {
    console.log("Define: " + constructor.name)
    const generated = class extends constructor { }
    customElements.define(tagname, generated)
    return generated
  }
}

export function state(select:string) {
  return (proto: any, propName: string) : any => {
    switch (select) {
      case 'test': proto.state = testOrder; break
      default: proto.state = emptyOrder; break
    }
  }
}

export function query(select:string) {
  return function (this:any, proto:any, propName:string, descriptor:PropertyDescriptor):any {
    let originalMethod = descriptor.value;
    descriptor.value = function(this:any, ...args:any[]) {
      if (!this.shadowRoot) return
      const elm = this.shadowRoot.querySelector(select)
      return originalMethod.apply(this, [elm])
    }
    return descriptor
  }
}
