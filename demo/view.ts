/// <amd-module name="view"/>
export default class MyView extends HTMLElement {

  private encodeParams(params: Object) {
    return Object.keys(params).map(function (k) {
      return encodeURIComponent(k) + '=' + encodeURIComponent((<any>params)[k])
    }).join('&')
  }

  private decodeParams(paramString: string) {
    var params = {};
    paramString = (paramString || '').replace(/\+/g, '%20');
    var paramList = paramString.split('&');
    for (var i = 0; i < paramList.length; i++) {
      var param = paramList[i].split('=');
      if (param[0]) {
        (<any>params)[decodeURIComponent(param[0])] = decodeURIComponent(param[1] || '');
      }
    }
    return params;
  }

  public query(): any {
    return this.decodeParams(location.search.slice(1))
  }

  public view(v: string) {
    const d = this.query()
    d.view = v
    window.dispatchEvent(new CustomEvent('view', { 'detail': d }))
  }

  constructor() {
    super()
    window.addEventListener('view', (e: any) => {
      window.dispatchEvent(new CustomEvent(e.detail.view, { 'detail': e.detail }))
      history.pushState(e.detail, e.detail.view, location.pathname + "?" + this.encodeParams(e.detail))
    }, false)
    window.addEventListener('popstate', (e: any) => {
      if (!e.state) return
      window.dispatchEvent(new CustomEvent(e.state.view, { 'detail': e.state }))
    })
  }

  connectedCallback() { }
  disconnectedCallback() { }
  attributeChangedCallback(name: string, oldValue: string, newValue: string) { }
  adoptedCallback() { }
}

// var encodedParams = [];
// for (var key in params) {
//   var value = params[key];
//   if (value === '') {
//     encodedParams.push(encodeURIComponent(key));
//   } else if (value) {
//     encodedParams.push(
//       encodeURIComponent(key) +
//       '=' +
//       encodeURIComponent(value.toString())
//     );
//   }
// }
// return encodedParams.join('&');
