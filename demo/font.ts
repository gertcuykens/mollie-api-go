// interface HTMLLinkElement {
//   property:string
// }

;(function(){
  const l = document.createElement('link')
  l.href = 'https://fonts.googleapis.com/css?family=Open+Sans|Roboto|Source+Code+Pro'
  l.rel = 'stylesheet'
  document.body.appendChild(l)
})()

// <link href="https://fonts.googleapis.com/css?family=Open+Sans|Roboto|Source+Code+Pro" rel="stylesheet">

// <style>
//   @import url('https://fonts.googleapis.com/css?family=Open+Sans|Roboto|Source+Code+Pro');
// </style>
