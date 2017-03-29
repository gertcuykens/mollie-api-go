const VERSION = 'v1'
const RUNTIME = 'runtime-'+VERSION
const PRECACHE = 'precache-'+VERSION
const PRECACHE_URLS = [
  '/',
  'manifest.json',
  'scripts/alameda.js',
  'scripts/index.js'
]

interface WorkerGlobalScope {
  skipWaiting():any
  clients:{
    claim:any
  }
}

self.addEventListener('install', (event:any) => {
  console.log('SW: install')
  event.waitUntil(
    caches.open(PRECACHE)
      .then(cache => cache.addAll(PRECACHE_URLS))
      .then(self.skipWaiting())
      .catch(err => console.log(err))
  )
})

self.addEventListener('activate', (event:any) => {
  console.log('SW: activate')
  const currentCaches = [PRECACHE, RUNTIME];
  event.waitUntil(
    caches.keys()
      .then( (cacheNames:string[]) => {
        return cacheNames.filter( (cacheName:string) => !currentCaches.includes(cacheName) ) })
      .then( (cachesToDelete:string[]) => {
        return Promise.all( cachesToDelete.map( (cacheToDelete:string) => { return caches.delete(cacheToDelete) }) )})
      .then(() => self.clients.claim())
      .catch((err:any) => console.log(err))
  )
})

self.addEventListener('fetch', (event:any) => {
  const request = event.request
  if (request.method !== 'GET') return
  event.respondWith( caches.match(request)
    .then( cachedResponse => {
      if (cachedResponse) return cachedResponse
      return caches.open(RUNTIME)
        .then(cache => fetch(request)
            .then(response => cache.put(request, response.clone())
                .then(() => response) ))  
  }))
})

// console.log('SW :', request)
// if (request.url.endsWith('.json') &&
//     !request.url.includes('manifest.json'))
// {
//   if (self.location.origin.includes('localhost:8080')) request = new Request(request.url.replace(':8080',':8081'))
//   event.respondWith( fetch(request) )
//   return
// }

// if (request.headers.get('Cache-control') === 'no-cache') return
// console.log('SW :', request.headers.get('Cache-control'))
