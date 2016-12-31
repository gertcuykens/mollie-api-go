/// <amd-module name="list"/>
export default class MyList extends HTMLElement {

  items(s: ShadowRoot) {
    let items: HTMLElement[] = []
    // Object.keys(localStorage).forEach(key => console.log(localStorage.getObject(key)));
    const test:any = localStorage.getObject('test')
    Object.keys(test).forEach(i => {
      const item = document.createElement('my-item')
      item.innerHTML = `
        - <div slot="name">${test[i].n}</div> <br/> 
        <span slot="description">${test[i].d}</span>
        <span slot="amount">${test[i].a}</span>
      `
      s.appendChild(item)
    });
  }

  constructor() {
    super()
  }

  connectedCallback() {
    const shadowRoot = this.attachShadow({ mode: 'open' })
    fetch('list.html')
      .then( response => response.text() )
      .then( text => {
        shadowRoot.innerHTML = text
        this.items(shadowRoot)
      })
  }
  disconnectedCallback() { }
  attributeChangedCallback(name: string, oldValue: string, newValue: string) { }
  adoptedCallback() { }
}
