/// <amd-module name="item"/>
export default class MyItem extends HTMLElement {
  constructor() {
    super()
    // if (shady) ShadyCSS.prepareTemplate(this.template, 'hello-world-b');
  }
  connectedCallback() {
    const shadowRoot = this.attachShadow({ mode: 'open' })
    fetch('item.html')
      .then( response => response.text() )
      .then( text => shadowRoot.innerHTML = text )
    // if (shady) { ShadyCSS.applyStyle(this) }
  }
  disconnectedCallback() { }
  attributeChangedCallback(name: string, oldValue: string, newValue: string) { }
  adoptedCallback() { }
}
