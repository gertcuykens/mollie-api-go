function loadScript(src: string) {
  return new Promise(function (resolve, reject) {
    const script = document.createElement('script');
    script.async = true;
    script.src = src;
    script.onload = resolve;
    script.onerror = reject;
    document.head.appendChild(script);
  });
}

;(function() {
  loadScript('font.js')
  const p: Promise<{}>[] = [];
  p.push(loadScript('/bower_components/requirejs/require.js'));
  p.push(loadScript('storage.js'));
  if (!('customElements' in window)) {
    p.push(loadScript('/bower_components/custom-elements/custom-elements.min.js'));
  }
  // if (!('fetch' in window)) {
  //   p.push(loadScript('/bower_components/fetch/fetch.min.js'));
  // }
  // let shady = false;
  // if (!HTMLElement.prototype.attachShadow) {
  //   shady = true
  //   p.push(loadScript('/bower_components/shadydom/shadydom.min.js'));
  //   p.push(loadScript('/bower_components/shadycss/shadycss.min.js'));
  // }
  Promise.all(p)
    .then(e => loadScript('test.js'))
    .then(e => { 
      require(['test'], function(test:any){ test.main() }) 
    })
})()
