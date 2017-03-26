/// <amd-module name="methodelement"/>
import {defineClass} from 'decorator'
import {Method, method} from 'mixin'

@defineClass('method-ts')
export default class MethodElement extends method(HTMLElement) {

  state:object = {}
  constructor() {
    super()
  }

  selector(select:HTMLSelectElement, json:any){
    const data:Method[] = json.data
    data.forEach((v)=>{
      const option = document.createElement('option')
      option.value = v.id
      option.innerText = v.description
      select.appendChild(option)
    })
  }

  connectedCallback() { 
    const root = this.attachShadow({mode: 'open'});
    root.innerHTML = `<slot name="input"></slot>`
    const select = this.querySelector('select')
    this.method()
      .then(json => this.selector(select!, json))
      .catch(err => console.log(err))
  }

  disconnectedCallback() { }
  attributeChangedCallback(name: string, oldValue: string, newValue: string) { }
  adoptedCallback() { }
}
