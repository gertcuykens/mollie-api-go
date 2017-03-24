/// <amd-module name="formelement"/>
import {Payment, Order, Product, payment, render} from 'mixin'
import {defineClass} from 'decorator'

@defineClass('form-ts')
export default class FormElement extends payment(render(HTMLElement)) {

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
      "Price": 5.99,
      "Currency": ""
    }]
  }

  constructor() {
    super()
    // this.state = {
    //   "Email": "",
    //   "Method": "",
    //   "Issuer": "",
    //   "Product": []
    // } as Order
    // this.state = localStorage.getObject('test') as Order
  }

  value(name: string): string {
    if (!this.shadowRoot) return ''
    const input = this.shadowRoot.querySelector(name) as HTMLInputElement
    return input.value
  } 

  add() {
    const product: Product = {
      'Name': this.value('[name="Name"]'),
      'Description': this.value('[name="Description"]'),
      'Quantity': Number(this.value('[name="Quantity"]')),
      'Price': Number(this.value('[name="Price"]')),
      'Currency': this.value('[name="Currency"]')
    }
    this.state['Product'].unshift(product)

    if (!this.shadowRoot) return
    const t = this.shadowRoot.querySelector('template')
    if (!t) return
    const td = t.content.querySelectorAll('td')
    td[0].innerText = this.value('[name="Name"]')
    td[1].innerText = this.value('[name="Description"]')
    td[2].innerText = this.value('[name="Quantity"]')
    td[3].innerText = this.value('[name="Price"]')
    const n = Number(this.value('[name="Price"]')) * Number(this.value('[name="Quantity"]'))
    td[4].innerText = String( n.round2f() )

    const table = this.shadowRoot.querySelectorAll('table')
    if (1 > table.length) return
    const tr = table[1].querySelectorAll('tr')
    if(5 > tr.length) return
    const parent = tr[4].parentNode
    if(!parent) return
    const clone = document.importNode(t.content, true)
    parent.insertBefore(clone, tr[4])

    const btn = table[1].querySelectorAll('button')
    if (1 > btn.length) return
    btn[1].addEventListener('click', ()=>{ this.rm(btn[1]) } )

    const span = table[1].querySelector('span')
    if (!span) return
    span.innerText = this.total()
  }

  rm(btn:HTMLButtonElement) {
    if (!btn.parentNode) return
    const tr = btn.parentNode.parentNode as HTMLTableRowElement
    let name = ""
    let td = tr.querySelectorAll('td')
    if (td) name = td[0].innerText
    this.state['Product'].forEach((p:Product, i:number) => {
      if (p.Name == name) this.state['Product'].splice(i, 1)
    })
    const table = tr.parentNode as HTMLTableElement
    if (!table) return
    table.removeChild(tr)
    
    const span = table.querySelector('span')
    if (!span) return
    span.innerText = this.total()
  }

  submit() {
    this.state['Email'] = this.value('[name="Email"]')
    this.state['Method'] = this.value('[name="Method"]')
    this.state['Issuer'] = this.value('[name="Issuer"]')
    this.payment(this.state)
      .then(json => {
        const p = json as Payment
        localStorage.setItem('payment', p.links.paymentUrl)
        localStorage.setItem('transaction', p.id)
        const menu = document.querySelector('menu-ts')
        if (menu) {
          const event = new CustomEvent('transaction', { 'detail': null })
          menu.dispatchEvent(event)
        }
      })
      .catch(err => console.log(err))
  }

  total(){
    return this.state.Product.reduce(function(acc, val) {
        return acc + (val.Price * val.Quantity)
      }, 0).round2f()
  }

  totalChanged() {
    if (!this.shadowRoot) return
    const t = this.shadowRoot.querySelector('#total')
    const n = Number(this.value('[name="Price"]')) * Number(this.value('[name="Quantity"]'))
    if (t) t.innerHTML = String( n.round2f() )
  }

  renderCallback(root:ShadowRoot) {
    let input = root.querySelector('[name="Price"]') as HTMLInputElement
    input.addEventListener('change', this.totalChanged.bind(this))

    input = root.querySelector('[name="Quantity"]') as HTMLInputElement
    input.addEventListener('change', this.totalChanged.bind(this))
    
    const btn = root.querySelectorAll('button')
    if (2 > btn.length) return
    btn[0].addEventListener('click', this.add.bind(this))
    btn[btn.length-1].addEventListener('click', this.submit.bind(this))
    for (let i=1; i < btn.length-1; ++i ) {
      btn[i].addEventListener('click', ()=>{ this.rm(btn[i]) })
    }
  }

  connectedCallback() {
    const uri = encodeURIComponent(JSON.stringify(this.state))
    this.render('form.md?json='+uri)
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
