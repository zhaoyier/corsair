(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-06fa210e"],{9036:function(t,e,n){"use strict";n.r(e);var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app-container"},[n("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],attrs:{data:t.list,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":""}},[n("el-table-column",{attrs:{align:"center",label:"ID",width:"95"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.$index)+" ")]}}])}),n("el-table-column",{attrs:{label:"Title"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.title)+" ")]}}])}),n("el-table-column",{attrs:{label:"Author",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("span",[t._v(t._s(e.row.author))])]}}])}),n("el-table-column",{attrs:{label:"Pageviews",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.pageviews)+" ")]}}])}),n("el-table-column",{attrs:{"class-name":"status-col",label:"Status",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("el-tag",{attrs:{type:t._f("statusFilter")(e.row.status)}},[t._v(t._s(e.row.status))])]}}])}),n("el-table-column",{attrs:{align:"center",prop:"created_at",label:"Display_time",width:"200"},scopedSlots:t._u([{key:"default",fn:function(e){return[n("i",{staticClass:"el-icon-time"}),n("span",[t._v(t._s(e.row.display_time))])]}}])})],1)],1)},l=[],s=n("d43b"),i={filters:{statusFilter:function(t){var e={published:"success",draft:"gray",deleted:"danger"};return e[t]}},data:function(){return{list:null,listLoading:!0}},created:function(){this.fetchData()},methods:{fetchData:function(){var t=this;this.listLoading=!0,Object(s["getList"])().then((function(e){t.list=e.data.items,t.listLoading=!1}))}}},r=i,o=n("2877"),u=Object(o["a"])(r,a,l,!1,null,null,null);e["default"]=u.exports},d43b:function(t,e,n){"use strict";n.d(e,"b",(function(){return l})),n.d(e,"d",(function(){return s})),n.d(e,"c",(function(){return i}));var a=n("b775");function l(t){return Object(a["a"])({url:"/api/stock/GetRecommend",method:"post",data:t})}function s(t){return Object(a["a"])({url:"/api/stock/UpdateRecommend",method:"post",data:t})}function i(t){return Object(a["a"])({url:"/api/stock/PromptBuyList",method:"post",data:t})}}}]);