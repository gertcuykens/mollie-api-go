import {render} from './mixin.js'
import {defineClass} from './decorator.js'

@defineClass('menu-ts')
export default class MenuElement extends render(HTMLElement) {
  constructor() {
    super()
  }

  renderCallback(root:ShadowRoot){
    const click1 = new CustomEvent('form', {'detail': null})
    const click2 = new CustomEvent('transaction', {'detail': null})
    const test = new CustomEvent('test', {'detail': null})    
    const li = root.querySelectorAll('li')
    li[0].addEventListener('click', ()=>{ this.dispatchEvent(click1) })
    li[1].addEventListener('click', ()=>{ this.dispatchEvent(click2) })
    li[2].addEventListener('click', ()=>{ document.location.href = 'payment.csv' }) // http://localhost:8081/
    li[3].addEventListener('click', ()=>{ this.dispatchEvent(test) })
    this.style.padding = '20px 0'
  }

  connectedCallback() { 
    this.render('menu.md')
      .then(text => {
        const root = this.attachShadow({mode: 'open'});
        root.innerHTML = text;
        this.renderCallback(root)
      })
      .catch(err => console.log(err))
  }

  disconnectedCallback() { }
  attributeChangedCallback(name: string, oldValue: string, newValue: string) { }
  adoptedCallback() { }
}
