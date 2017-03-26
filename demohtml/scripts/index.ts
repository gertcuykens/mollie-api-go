import './prototype.js'
import './menuelement.js'

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
    if ('SCRIPT' == elms[i].tagName) continue;
    elms[i].setAttribute('hidden', 'true')
  }
}

;(function(){
  const menu = document.querySelector('menu-ts')
  menu!.addEventListener('transaction', () => { script('./scripts/transactionelement.js').then( () => show('menu-ts', 'transaction-ts') ) })      
  menu!.addEventListener('form', () => { script('./scripts/formelement.js', './scripts/issuerelement.js', './scripts/methodelement.js').then( () => show('menu-ts', 'form-ts') ) })
  menu!.addEventListener('test', () => { script('./scripts/formelement_test.js').then( () => show('menu-ts', 'form-test') ) })
  menu!.addEventListener('online', () => { online(); show('menu-ts') })
  menu!.addEventListener('offline', () => { offline(); show('menu-ts') })
  const page = window.location.search.substr(1)
  let event = new CustomEvent('form', { 'detail': null })
  if (page) event = new CustomEvent(page, { 'detail': null })
  menu!.dispatchEvent(event)
})()

function script(...href:string[]):Promise<{}> {
  const p:Promise<{}>[] = []
  href.forEach( v => p.push( new Promise( (resolve, reject) => {
    const script = document.createElement('script')
    script.type = 'module'
    script.src = href[0]
    script.onload = resolve
    script.onerror = reject
    script.setAttribute('async', '')
    document.body.appendChild(script)
  })))
  return Promise.all([p])
}

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
