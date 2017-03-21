Storage.prototype.setObject = function(this:Storage, key:string, value:any) {
  this.setItem(key, JSON.stringify(value))
}

Storage.prototype.getObject = function(this:Storage, key:string) {
  const value = this.getItem(key)
  return value && JSON.parse(value)
}

interface String {
  isRegistered: (this:string) => boolean | undefined
  loadScript: (this:string) => Promise < (resolve: Function, reject: Function) => void >
}

String.prototype.isRegistered = function(this:string) { 
  switch(document.createElement(this).constructor) {
    case HTMLElement: return false; 
    case HTMLUnknownElement: return undefined; 
  }
  return true;
}

String.prototype.loadScript = function (this:string) {
  const src = this
  return new Promise(function (resolve, reject) {
    const script = document.createElement('script')
    script.async = true
    script.src = src
    script.onload = resolve
    script.onerror = reject
    document.head.appendChild(script)
  });
}

interface Number {
  round2: (this:number) => number
  round2f: (this:number) => string
  
}

Number.prototype.round2 = function (this:Number) {  
    const s = Number(String(this) + "e+2")
    const r = Math.round(s)
    return Number(r + "e-2")
}

Number.prototype.round2f = function (this:Number) {  
    const s = Number(String(this) + "e+2")
    const r = Math.round(s)
    return Number(r + "e-2").toFixed(2)
}
