interface Storage {
    setObject<T>(key:string, value:T):void;
    getObject<T>(key:string):T;
}

Storage.prototype.setObject = function(this:Storage, key:string, value:any) {
    this.setItem(key, JSON.stringify(value));
}

Storage.prototype.getObject = function(this:Storage, key:string) {
    const value = this.getItem(key);
    return value && JSON.parse(value);
}
