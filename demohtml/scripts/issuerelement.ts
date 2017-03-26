import {defineClass} from './decorator.js'
import {Issuer, issuer} from './mixin.js'

@defineClass('issuer-ts')
export default class IssuerElement extends issuer(HTMLElement) {
  constructor() {
    super()
  }

  selector(select:HTMLSelectElement, json:any){
    const data:Issuer[] = json.data
    data.forEach((v)=>{
      const option = document.createElement('option')
      option.value = v.id
      option.innerText = v.name
      select.appendChild(option)
    })
  }

  connectedCallback() { 
    const root = this.attachShadow({mode: 'open'})
    root.innerHTML = `<slot name="input"></slot>`
    const select = this.querySelector('select')
    if (!select) return
    this.issuer()
      .then(json => this.selector(select, json))
      .catch(err => console.log(err))
  }

  disconnectedCallback() { }
  attributeChangedCallback(name: string, oldValue: string, newValue: string) { }
  adoptedCallback() { }
}
