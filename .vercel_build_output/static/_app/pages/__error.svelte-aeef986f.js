import{S as s,i as t,s as e,k as a,e as r,t as o,I as n,d as i,n as u,c,a as d,g as l,f as h,E as m,h as f,J as p}from"../chunks/vendor-0839ef83.js";function g(s){let t,e,g,E,v,y,$,k;return document.title=t=s[0],{c(){e=a(),g=r("h1"),E=o("Ooopsy... ...this is my standard error page."),v=a(),y=r("h2"),$=o("Error message: "),k=o(s[0])},l(t){n('[data-svelte="svelte-1uo06u1"]',document.head).forEach(i),e=u(t),g=c(t,"H1",{});var a=d(g);E=l(a,"Ooopsy... ...this is my standard error page."),a.forEach(i),v=u(t),y=c(t,"H2",{});var r=d(y);$=l(r,"Error message: "),k=l(r,s[0]),r.forEach(i)},m(s,t){h(s,e,t),h(s,g,t),m(g,E),h(s,v,t),h(s,y,t),m(y,$),m(y,k)},p(s,[e]){1&e&&t!==(t=s[0])&&(document.title=t),1&e&&f(k,s[0])},i:p,o:p,d(s){s&&i(e),s&&i(g),s&&i(v),s&&i(y)}}}function E({error:s,status:t}){return{props:{title:`${t}: ${s.message}`}}}function v(s,t,e){let{title:a}=t;return s.$$set=s=>{"title"in s&&e(0,a=s.title)},[a]}class y extends s{constructor(s){super(),t(this,s,v,g,e,{title:0})}}export{y as default,E as load};
