(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-70087557"],{"9fad":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"dashboard-container"},[a("div",{staticClass:"app-container"},[a("el-form",{staticClass:"demo-form-inline",attrs:{inline:!0,model:t.dailyForm}},[a("el-form-item",{attrs:{label:"名称"}},[a("el-input",{attrs:{placeholder:"茅台股份"},model:{value:t.dailyForm.name,callback:function(e){t.$set(t.dailyForm,"name",e)},expression:"dailyForm.name"}})],1),a("el-form-item",{attrs:{label:"代码"}},[a("el-input",{attrs:{placeholder:"000001"},model:{value:t.dailyForm.secucode,callback:function(e){t.$set(t.dailyForm,"secucode",e)},expression:"dailyForm.secucode"}})],1),a("el-form-item",{attrs:{label:"跌幅"}},[a("el-input-number",{attrs:{step:2,min:2,max:20,"step-strictly":""},model:{value:t.dailyForm.decrease,callback:function(e){t.$set(t.dailyForm,"decrease",e)},expression:"dailyForm.decrease"}})],1),a("el-form-item",{attrs:{label:"日期"}},[a("el-date-picker",{attrs:{"value-format":"timestamp",type:"daterange","range-separator":"至","start-placeholder":"开始日期","end-placeholder":"结束日期"},on:{input:t.selectDate},model:{value:t.dailyForm.dateRange,callback:function(e){t.$set(t.dailyForm,"dateRange",e)},expression:"dailyForm.dateRange"}})],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:t.onQuerySubmit}},[t._v("查询")])],1)],1)],1),a("el-divider"),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],attrs:{data:t.list,"element-loading-text":"Loading",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{fixed:"left",label:"名称",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:"danger",effect:"dark"}},[t._v(t._s(e.row.name))])]}}])}),a("el-table-column",{attrs:{label:"名称",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("span",[t._v(t._s(e.row.secucode))])]}}])}),a("el-table-column",{attrs:{label:"涨幅",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:"warning",effect:"dark"}},[t._v(t._s(e.row.prise))])]}}])}),a("el-table-column",{attrs:{label:"量比",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.liangbi)+" ")]}}])}),a("el-table-column",{attrs:{label:"开盘价",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.opening)+" ")]}}])}),a("el-table-column",{attrs:{label:"收盘价",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.closing)+" ")]}}])}),a("el-table-column",{attrs:{label:"最高价",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.maxPrice)+" ")]}}])}),a("el-table-column",{attrs:{label:"最低价",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.minPrice)+" ")]}}])}),a("el-table-column",{attrs:{label:"流通市值(亿)",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.traded)+" ")]}}])}),a("el-table-column",{attrs:{label:"创建日期",width:"210",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.createDate)+" ")]}}])}),a("el-table-column",{attrs:{fixed:"right",label:"操作",width:"220"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),t.updateDaily(e.$index,t.tableData)}}},[t._v(" 编辑 ")])]}}])})],1),a("div",{staticClass:"gva-pagination"},[a("el-pagination",{attrs:{"current-page":t.dailyForm.pageNum,"page-size":t.dailyForm.pageSize,"page-sizes":[10,30,50,100],total:t.dailyForm.total,layout:"total, sizes, prev, pager, next, jumper"},on:{"current-change":t.handleCurrentChange,"size-change":t.handleSizeChange}})],1)],1)},r=[],i=(a("b0c0"),a("d43b")),o={filters:{statusFilter:function(t){var e={published:"success",draft:"gray",deleted:"danger"};return e[t]}},data:function(){return{list:null,timer:null,listLoading:!0,dailyForm:{pageNum:1,pageSize:20,total:0,lineChartVisible:!1,lineChartSrc:"",name:"",secucode:"",dateRange:"",startDate:0,endDate:0}}},mounted:function(){this.fetchData()},methods:{fetchData:function(){var t=this;this.listLoading=!0;var e={name:this.dailyForm.name,secucode:this.dailyForm.secucode,decrease:this.dailyForm.decrease,startDate:this.dailyForm.startDate,endDate:this.dailyForm.endDate,limit:this.dailyForm.pageSize,offset:(this.dailyForm.pageNum-1)*this.dailyForm.pageSize};Object(i["e"])(e).then((function(e){t.list=e.data.items,t.listLoading=!1,t.dailyForm.total=e.data.total,t.dailyForm.startDate=0,t.dailyForm.endDate=0}))},handleCurrentChange:function(t){this.dailyForm.pageNum=t,this.fetchData()},handleSizeChange:function(){},updateDaily:function(t,e){var a=e[t],n=a.secucode.toLowerCase();console.log("==>>TODO secucode is 02: ",n)},selectDate:function(t){console.log("==>>TODO date is: ",t[0],t[1]),this.dailyForm.startDate=t[0],this.dailyForm.endDate=t[1]},onQuerySubmit:function(t){console.log("==>>TODO query is: ",t),this.fetchData()}}},l=o,s=(a("d41b"),a("2877")),u=Object(s["a"])(l,n,r,!1,null,"85856e98",null);e["default"]=u.exports},d41b:function(t,e,a){"use strict";a("dba4")},d43b:function(t,e,a){"use strict";a.d(e,"j",(function(){return r})),a.d(e,"g",(function(){return i})),a.d(e,"m",(function(){return o})),a.d(e,"e",(function(){return l})),a.d(e,"i",(function(){return s})),a.d(e,"b",(function(){return u})),a.d(e,"f",(function(){return c})),a.d(e,"a",(function(){return d})),a.d(e,"k",(function(){return m})),a.d(e,"d",(function(){return f})),a.d(e,"c",(function(){return p})),a.d(e,"l",(function(){return h})),a.d(e,"h",(function(){return b}));var n=a("b775");function r(t){return Object(n["a"])({url:"/api/backend/updateCNConfig",method:"post",data:t})}function i(t){return Object(n["a"])({url:"/api/stock/GetRecommend",method:"post",data:t})}function o(t){return Object(n["a"])({url:"/api/stock/UpdateRecommend",method:"post",data:t})}function l(t){return Object(n["a"])({url:"/api/stock/GetDailyList",method:"post",data:t})}function s(t){return Object(n["a"])({url:"/api/stock/ManualDecreaseList",method:"post",data:t})}function u(t){return Object(n["a"])({url:"/api/stock/ConfirmFocus",method:"post",data:t})}function c(t){return Object(n["a"])({url:"/api/stock/GetFocusList",method:"post",data:t})}function d(t){return Object(n["a"])({url:"/api/stock/CancelFocus",method:"post",data:t})}function m(t){return Object(n["a"])({url:"/api/stock/updateFocus",method:"post",data:t})}function f(t){return Object(n["a"])({url:"/api/stock/GDRenshuList",method:"post",data:t})}function p(t){return Object(n["a"])({url:"/api/stock/GDRenshuDetail",method:"post",data:t})}function h(t){return Object(n["a"])({url:"/api/stock/UpdateGPZhouQi",method:"post",data:t})}function b(t){return Object(n["a"])({url:"/api/stock/GPZhouQiList",method:"post",data:t})}},dba4:function(t,e,a){}}]);