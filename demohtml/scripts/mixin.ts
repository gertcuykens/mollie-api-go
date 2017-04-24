type Constructor<T> = new (...args: any[]) => T;

export function render < T extends Constructor<HTMLElement> > (Base: T) {
  return class extends Base {
    render(url: string):Promise<string> {
      return fetch(url, {
        headers: {
          'Accept': 'text/plain, text/html',
          'Accept-Encoding': 'gzip',
          'Content-Type': 'application/json'
        },
        method: 'GET',
        mode: 'cors'
      }).then(response => response.text())
    }
  }
}

export type Order = {
    "Email": string
    "Method": string
    "Issuer": string
    "Product": Array<Product>
}

export type Product = {
    "Name": string
    "Description": string
    "Quantity": number
    "Price": number
    "Currency": string
}

export type Payment = {
	"id": string
	"links": {
		"paymentUrl": string
		"redirectUrl":string
	}
}

export function payment < T extends Constructor<HTMLElement> > (Base: T) {
  return class extends Base {
    payment(state: Order):Promise<{}> {
      return fetch('demo/payment.json', {
        headers: {
          'Accept': 'application/json',
          'Accept-Encoding': 'gzip',
          'Content-Type': 'application/json',
          'Cache-control': 'no-cache'
        },
        method: 'POST',
        mode: 'cors',
        body: JSON.stringify(state)
      }).then( response => response.json() )
    }
  }
}

export type Transaction = {
  "Id": string
  "Mode": string
  "CreatedDatetime": string
  "Status": string
  "PaidDatetime": string
  "Amount": string
  "Description": string
  "Method": string
  "Metadata": object
  "Locale": string
  "Links":{
    "WebhookUrl": string
    "RedirectUrl": string
  }
}

export function transaction < T extends Constructor<HTMLElement> > (Base: T) {
  return class extends Base {
    transaction(id: string):Promise<{}> {
      return fetch('demo/transaction.json?id='+id, {
        headers: {
          'Accept': 'application/json',
          'Accept-Encoding': 'gzip',
          'Cache-control': 'no-cache'
        },
        method: 'GET',
        mode: 'cors'
      }).then( response => response.json() )
    }
  }
}

export type Method = {
  resource: string
  id: string
  description: string
  amount: object
  image: object
}

export function method < T extends Constructor<HTMLElement> > (Base: T) {
  return class extends Base {
    method():Promise<{}> {
      return fetch('demo/method.json', {
        headers: {
          'Accept': 'application/json',
          'Accept-Encoding': 'gzip',
          'Cache-control': 'no-cache'
        },
        method: 'GET',
        mode: 'cors'
      }).then( response => response.json() )
    }
  }
}

export type Issuer = {
  resource: string
  id: string
  name: string
  method: string
}

export function issuer < T extends Constructor<HTMLElement> > (Base: T) {
  return class extends Base {
    issuer():Promise<{}> {
      return fetch('demo/issuer.json', {
        headers: {
          'Accept': 'application/json',
          'Accept-Encoding': 'gzip',
          'Cache-control': 'no-cache'
        },
        method: 'GET',
        mode: 'cors'
      }).then( response => response.json() )
    }
  }
}
