(this.webpackJsonpgoking=this.webpackJsonpgoking||[]).push([[0],{12:function(t,n,e){},13:function(t,n,e){},16:function(t,n,e){"use strict";e.r(n);var c=e(1),a=e.n(c),r=e(4),s=e.n(r),i=(e(12),e(7)),o=(e(13),e(6)),l=e(3),u=e.n(l),b=e(5),h=e(0);WebAssembly.instantiateStreaming||(WebAssembly.instantiateStreaming=function(){var t=Object(b.a)(u.a.mark((function t(n,e){var c;return u.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,n;case 2:return t.next=4,t.sent.arrayBuffer();case 4:return c=t.sent,t.next=7,WebAssembly.instantiate(c,e);case 7:return t.abrupt("return",t.sent);case 8:case"end":return t.stop()}}),t)})));return function(n,e){return t.apply(this,arguments)}}()),function(t){var n=new window.Go;return new Promise((function(e,c){console.log(n),WebAssembly.instantiateStreaming(fetch(t),n.importObject).then((function(t){n.run(t.instance),e(t.instance)})).catch((function(t){c(t)}))}))}("./main.wasm").then((function(t){console.log(t)})).catch((function(t){console.log("ouch",t)}));var d,m=Object(o.hot)((function(){var t,n=Object(c.useState)(null),e=Object(i.a)(n,2),a=e[0],r=e[1];return window.getStr&&(null===(t=window.getStr())||void 0===t||t.then(r)),Object(h.jsxs)("div",{className:"App",children:[Object(h.jsx)("p",{style:{color:"white",fontSize:"11px",width:"100vw",height:"300px",position:"absolute",top:15,left:15,display:"flex",flexWrap:"wrap"},children:Object.keys(window).map((function(t){return Object(h.jsx)("span",{style:{margin:3},children:t},t)}))}),Object(h.jsx)("header",{className:"App-header",children:Object(h.jsxs)("div",{className:"btn-row",children:[a&&Object(h.jsx)("button",{className:"btn",children:a}),Object(h.jsx)("button",{className:"btn",children:10*Math.random()}),Object(h.jsx)("button",{className:"btn",children:10*Math.random()}),Object(h.jsx)("button",{className:"btn",children:10*Math.random()}),Object(h.jsx)("button",{className:"btn",children:10*Math.random()})]})})]})})),j=function(t){t&&t instanceof Function&&e.e(3).then(e.bind(null,17)).then((function(n){var e=n.getCLS,c=n.getFID,a=n.getFCP,r=n.getLCP,s=n.getTTFB;e(t),c(t),a(t),r(t),s(t)}))};d=m,s.a.render(Object(h.jsx)(a.a.StrictMode,{children:Object(h.jsx)(d,{})}),document.getElementById("root")),j(),window.forceClearCache=function(){for(var t=document.getElementsByTagName("script"),n=["myscript.js","myscript2.js"],e=0;e<t.length;e++)for(var c=0;c<n.length;c++)t[e].src&&t[e].src.indexOf(n[c])>-1&&(t[e].src=t[e].src.replace(n[c],"".concat(n[c],"k=").concat(1)));window.location.reload(!0)}}},[[16,1,2]]]);
//# sourceMappingURL=main.ee856552.chunk.js.map