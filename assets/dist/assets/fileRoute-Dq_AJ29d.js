import{g as r,h as i,i as a,j as u,k as n,b as h}from"./index-CARAERQ_.js";class c{constructor(e){this.useMatch=t=>r({select:t==null?void 0:t.select,from:this.options.id}),this.useRouteContext=t=>r({from:this.options.id,select:o=>t!=null&&t.select?t.select(o.context):o.context}),this.useSearch=t=>i({...t,from:this.options.id}),this.useParams=t=>a({...t,from:this.options.id}),this.useLoaderDeps=t=>u({...t,from:this.options.id}),this.useLoaderData=t=>n({...t,from:this.options.id}),this.useNavigate=()=>h({from:this.options.id}),this.options=e,this.$$typeof=Symbol.for("react.memo")}}function d(s){return e=>new c({id:s,...e})}export{d as c};