/// <amd-module name="test"/>
import MyList from 'list';
import MyItem from 'item';
import MyView from 'view';

const isRegistered = function(name:string) {
  return document.createElement(name).constructor !== HTMLElement;
}

export function main():void {
  localStorage.setItem('test', '[{"n":"n1", "d":"d1", "a":"a1"}, {"n":"n2", "d":"d2", "a":"a2"}]')

  customElements.define('my-view', MyView)
  const v:any = document.getElementsByTagName('my-view')[0]
  const shadowRoot = v.attachShadow({ mode: 'open' })

  window.addEventListener('pg1', (e: any) => {
    fetch('view.html')
      .then( response => response.text() )
      .then( text => {
        if (!isRegistered('my-list')) customElements.define('my-list', MyList)
        if (!isRegistered('my-item')) customElements.define('my-item', MyItem)
        shadowRoot.innerHTML = text
      })
  }, false)

  const q = v.query()
  switch (q.view) {
    case 'pg2': window.dispatchEvent(new CustomEvent('pg2', { 'detail': q })); break;
    default: window.dispatchEvent(new CustomEvent('pg1', { 'detail': q })); break;
  }

  console.log(`test done`)
}

// window.dispatchEvent(new CustomEvent('x-menu', { 'detail': { 'menu': e.detail.view } }))
// window.dispatchEvent(new CustomEvent('x-menu', { 'detail': { 'menu': e.state.view } }))
