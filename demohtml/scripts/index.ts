import 'prototype'

function offline() {
  navigator.serviceWorker.getRegistrations()
    .then(function(registrations:any) {
      for(let registration of registrations) registration.unregister()
    })
}

function online() {
  if ('serviceWorker' in navigator) navigator.serviceWorker.register('service.js')
}

function show(...args:string[]) {
  const elms = document.querySelectorAll('body > *')
  for (let i = 0; i < elms.length; ++i) {
    let next = false
    args.forEach( (v) => {if (v.toUpperCase() == elms[i].tagName) next = true; })
    if (next) {elms[i].removeAttribute('hidden'); continue;}
    if ('STYLE' == elms[i].tagName) continue; 
    elms[i].setAttribute('hidden', 'true')
  }
}

;(function(){
  requirejs(['menuelement'], () => {
    const menu = document.querySelector('menu-ts')
    menu!.addEventListener('transaction', () => { requirejs(['transactionelement'], () => { show('menu-ts', 'transaction-ts') }) })      
    menu!.addEventListener('form', () => { requirejs(['formelement', 'issuerelement', 'methodelement'], () => { show('menu-ts', 'form-ts') }) })
    menu!.addEventListener('test', () => { requirejs(['formelement_test'], () => { show('menu-ts', 'form-test') }) })
    menu!.addEventListener('online', () => { online(); show('menu-ts') })
    menu!.addEventListener('offline', () => { offline(); show('menu-ts') })
    const page = window.location.search.substr(1)
    let event = new CustomEvent('form', { 'detail': null })
    if (page) event = new CustomEvent(page, { 'detail': null })
    menu!.dispatchEvent(event)
  })
})()

function link(href:string):Promise<{}> {
  return new Promise((resolve, reject) => {
    const link = document.createElement('link')
    link.rel = 'import'
    link.href = href
    link.onload = resolve
    link.onerror = reject
    link.setAttribute('async', '')
    document.head.appendChild(link)
  })
}
