(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-08109582"],{"34c4":function(e,t,a){"use strict";a("76f4")},"76f4":function(e,t,a){},"9fad":function(e,t,a){"use strict";a.r(t);var n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"dashboard-container"},[a("div",{staticClass:"app-container"},[a("el-form",{staticClass:"demo-form-inline",attrs:{inline:!0,model:e.dailyForm}},[a("el-form-item",{attrs:{label:"名称"}},[a("el-input",{attrs:{placeholder:"茅台股份"},model:{value:e.dailyForm.name,callback:function(t){e.$set(e.dailyForm,"name",t)},expression:"dailyForm.name"}})],1),a("el-form-item",{attrs:{label:"代码"}},[a("el-input",{attrs:{placeholder:"000001"},model:{value:e.dailyForm.secucode,callback:function(t){e.$set(e.dailyForm,"secucode",t)},expression:"dailyForm.secucode"}})],1),a("el-form-item",{attrs:{label:"跌幅"}},[a("el-input-number",{attrs:{step:2,min:2,max:20,"step-strictly":""},model:{value:e.dailyForm.decrease,callback:function(t){e.$set(e.dailyForm,"decrease",t)},expression:"dailyForm.decrease"}})],1),a("el-form-item",{attrs:{label:"日期"}},[a("el-date-picker",{attrs:{"value-format":"timestamp",type:"daterange","range-separator":"至","start-placeholder":"开始日期","end-placeholder":"结束日期"},on:{input:e.selectDate},model:{value:e.dailyForm.dateRange,callback:function(t){e.$set(e.dailyForm,"dateRange",t)},expression:"dailyForm.dateRange"}})],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:e.onQuerySubmit}},[e._v("查询")])],1)],1)],1),a("el-divider"),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],attrs:{data:e.list,"element-loading-text":"Loading",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{fixed:"left",label:"名称",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"danger",effect:"dark"}},[e._v(e._s(t.row.name))])]}}])}),a("el-table-column",{attrs:{label:"名称",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.secucode))])]}}])}),a("el-table-column",{attrs:{label:"涨幅",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"warning",effect:"dark"}},[e._v(e._s(t.row.prise))])]}}])}),a("el-table-column",{attrs:{label:"量比",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.liangbi)+" ")]}}])}),a("el-table-column",{attrs:{label:"开盘价",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.opening)+" ")]}}])}),a("el-table-column",{attrs:{label:"收盘价",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.closing)+" ")]}}])}),a("el-table-column",{attrs:{label:"最高价",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.maxPrice)+" ")]}}])}),a("el-table-column",{attrs:{label:"最低价",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.minPrice)+" ")]}}])}),a("el-table-column",{attrs:{label:"流通市值",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.traded)+" ")]}}])}),a("el-table-column",{attrs:{label:"创建日期",width:"210",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.createDate)+" ")]}}])}),a("el-table-column",{attrs:{fixed:"right",label:"操作",width:"220"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.updateDaily(t.$index,e.tableData)}}},[e._v(" 编辑 ")])]}}])})],1),a("div",{staticClass:"gva-pagination"},[a("el-pagination",{attrs:{"current-page":e.dailyForm.pageNum,"page-size":e.dailyForm.pageSize,"page-sizes":[10,30,50,100],total:e.dailyForm.total,layout:"total, sizes, prev, pager, next, jumper"},on:{"current-change":e.handleCurrentChange,"size-change":e.handleSizeChange}})],1)],1)},l=[],r=(a("b0c0"),a("d43b")),i={filters:{statusFilter:function(e){var t={published:"success",draft:"gray",deleted:"danger"};return t[e]}},data:function(){return{list:null,timer:null,listLoading:!0,dailyForm:{pageNum:1,pageSize:20,total:0,lineChartVisible:!1,lineChartSrc:"",name:"",secucode:"",dateRange:"",startDate:0,endDate:0}}},mounted:function(){this.fetchData()},methods:{fetchData:function(){var e=this;this.listLoading=!0;var t={name:this.dailyForm.name,secucode:this.dailyForm.secucode,decrease:this.dailyForm.decrease,startDate:this.dailyForm.startDate,endDate:this.dailyForm.endDate,limit:this.dailyForm.pageSize,offset:(this.dailyForm.pageNum-1)*this.dailyForm.pageSize};Object(r["b"])(t).then((function(t){e.list=t.data.items,e.listLoading=!1,e.dailyForm.total=t.data.total,e.dailyForm.name="",e.dailyForm.secucode="",e.dailyForm.decrease=0,e.dailyForm.startDate=0,e.dailyForm.endDate=0,e.dailyForm.pageNum=1}))},handleCurrentChange:function(){},handleSizeChange:function(){},updateDaily:function(e,t){var a=t[e],n=a.secucode.toLowerCase();console.log("==>>TODO secucode is: ",n)},selectDate:function(e){console.log("==>>TODO date is: ",e[0],e[1]),this.dailyForm.startDate=e[0],this.dailyForm.endDate=e[1]},onQuerySubmit:function(e){console.log("==>>TODO query is: ",e),this.fetchData()}}},o=i,s=(a("34c4"),a("2877")),c=Object(s["a"])(o,n,l,!1,null,"ab504de2",null);t["default"]=c.exports},d43b:function(e,t,a){"use strict";a.d(t,"e",(function(){return l})),a.d(t,"h",(function(){return r})),a.d(t,"g",(function(){return i})),a.d(t,"d",(function(){return o})),a.d(t,"b",(function(){return s})),a.d(t,"f",(function(){return c})),a.d(t,"a",(function(){return d})),a.d(t,"c",(function(){return u}));var n=a("b775");function l(e){return Object(n["a"])({url:"/api/stock/GetRecommend",method:"post",data:e})}function r(e){return Object(n["a"])({url:"/api/stock/UpdateRecommend",method:"post",data:e})}function i(e){return Object(n["a"])({url:"/api/stock/PromptBuyList",method:"post",data:e})}function o(e){return Object(n["a"])({url:"/api/stock/GetLongLineList",method:"post",data:e})}function s(e){return Object(n["a"])({url:"/api/stock/GetDailyList",method:"post",data:e})}function c(e){return Object(n["a"])({url:"/api/stock/ManualDecreaseList",method:"post",data:e})}function d(e){return Object(n["a"])({url:"/api/stock/FocusConfirm",method:"post",data:e})}function u(e){return Object(n["a"])({url:"/api/stock/GetFocusList",method:"post",data:e})}}}]);