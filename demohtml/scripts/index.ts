import 'prototype'

function show(...args:string[]) {
  const elms = document.querySelectorAll('body > *')
  for (let i = 0; i < elms.length; ++i) {
    let next = false
    args.forEach( (v) => {if (v.toUpperCase() == elms[i].tagName) next = true; })
    if (next) {elms[i].removeAttribute('hidden'); continue;}
    elms[i].setAttribute('hidden', 'true')
  }
}

;(function(){
  requirejs(['menuelement'], () => {
    let menu = document.querySelector('menu-ts')
    if (menu) {
      menu.addEventListener('transaction', (e) => { requirejs(['transactionelement'], () => { show('menu-ts', 'transaction-ts') }) })      
      menu.addEventListener('form', (e) => { requirejs(['formelement', 'issuerelement', 'methodelement'], () => { show('menu-ts', 'form-ts') }) })
      const page = window.location.search.substr(1)
      let event = new CustomEvent('form', { 'detail': null })
      if (page) event = new CustomEvent(page, { 'detail': null })
      menu.dispatchEvent(event)
    }
  })
})()
