(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-c18fd512"],{1667:function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"dashboard-container"},[a("div",{staticClass:"app-container"},[a("el-form",{staticClass:"demo-form-inline",attrs:{inline:!0,model:t.queryForm}},[a("el-form-item",{attrs:{label:"名称"}},[a("el-input",{attrs:{placeholder:""},model:{value:t.queryForm.name,callback:function(e){t.$set(t.queryForm,"name",e)},expression:"queryForm.name"}})],1),a("el-form-item",{attrs:{label:"代码"}},[a("el-input",{attrs:{placeholder:"000001"},model:{value:t.queryForm.secucode,callback:function(e){t.$set(t.queryForm,"secucode",e)},expression:"queryForm.secucode"}})],1),a("el-form-item",{attrs:{label:"增减率"}},[a("el-input-number",{attrs:{step:2,min:-100,max:100,"step-strictly":""},model:{value:t.queryForm.totalRatio,callback:function(e){t.$set(t.queryForm,"totalRatio",e)},expression:"queryForm.totalRatio"}})],1),a("el-form-item",{attrs:{label:"排序"}},[a("el-select",{attrs:{clearable:"",placeholder:"全部"},on:{change:t.onSelectValue},model:{value:t.queryForm.sortTyp,callback:function(e){t.$set(t.queryForm,"sortTyp",e)},expression:"queryForm.sortTyp"}},t._l(t.queryForm.options,(function(t){return a("el-option",{key:t.value,attrs:{label:t.label,value:t.value}})})),1)],1),a("el-form-item",{attrs:{label:"日期"}},[a("el-date-picker",{attrs:{"value-format":"timestamp",type:"daterange","range-separator":"至","start-placeholder":"开始日期","end-placeholder":"结束日期"},on:{input:t.onSelectDate},model:{value:t.queryForm.dateRange,callback:function(e){t.$set(t.queryForm,"dateRange",e)},expression:"queryForm.dateRange"}})],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:t.onQuerySubmit}},[t._v("查询")])],1)],1)],1),a("el-divider"),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],attrs:{data:t.tableData,"element-loading-text":"Loading",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{fixed:"left",label:"名称",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:"danger",effect:"dark"}},[t._v(t._s(e.row.name))])]}}])}),a("el-table-column",{attrs:{label:"代码",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("span",[t._v(t._s(e.row.secucode))])]}}])}),a("el-table-column",{attrs:{label:"持仓人数",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:"warning",effect:"dark"}},[t._v(t._s(e.row.holderTotalNum))])]}}])}),a("el-table-column",{attrs:{prop:"totalNumRatio",label:"持仓变化率",width:"140",align:"center",sortable:""},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.totalNumRatio)+" ")]}}])}),a("el-table-column",{attrs:{label:"发布时间",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.releaseDate)+" ")]}}])}),a("el-table-column",{attrs:{label:"集中度",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.holdFocus)+" ")]}}])}),a("el-table-column",{attrs:{label:"当前价",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:"danger",effect:"light"}},[t._v(t._s(e.row.presentPrice))])]}}])}),a("el-table-column",{attrs:{label:"收盘价",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:"warning",effect:"light"}},[t._v(t._s(e.row.price))])]}}])}),a("el-table-column",{attrs:{label:"非流通股东率",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.holdRatioTotal)+" ")]}}])}),a("el-table-column",{attrs:{label:"流通股东率",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.freeholdRatioTotal)+" ")]}}])}),a("el-table-column",{attrs:{fixed:"right",label:"操作",width:"220"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),t.onRenshuDetail(e.$index,t.tableData)}}},[t._v(" 历史 ")]),a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),t.onConfirmFocus(e.$index,t.tableData)}}},[t._v(" "+t._s(e.row.focused)+" ")]),a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),t.klineChart(e.$index,t.tableData)}}},[t._v(" K线图 ")])]}}])})],1),a("div",{staticClass:"gva-pagination"},[a("el-pagination",{attrs:{"current-page":t.queryForm.pageNum,"page-size":t.queryForm.pageSize,"page-sizes":[10,30,50,100],total:t.queryForm.total,layout:"total, sizes, prev, pager, next, jumper"},on:{"current-change":t.handleCurrentChange,"size-change":t.handleSizeChange}})],1),a("div",[a("el-dialog",{attrs:{title:"K线查询",visible:t.lineChartForm.lineChartVisible},on:{"update:visible":function(e){return t.$set(t.lineChartForm,"lineChartVisible",e)}}},[a("el-tabs",{on:{"tab-click":t.selectTabClick},model:{value:t.lineChartForm.activeName,callback:function(e){t.$set(t.lineChartForm,"activeName",e)},expression:"lineChartForm.activeName"}},[a("el-tab-pane",{attrs:{label:"日线",name:"date"}},[a("el-image",{attrs:{src:t.lineChartForm.lineChartSrc}})],1),a("el-tab-pane",{attrs:{label:"周线",name:"contour"}},[a("el-image",{attrs:{src:t.lineChartForm.lineChartSrc}})],1),a("el-tab-pane",{attrs:{label:"分时线",name:"minute"}},[a("el-image",{attrs:{src:t.lineChartForm.lineChartSrc}})],1)],1)],1)],1),a("div",[a("el-dialog",{attrs:{title:"股东人数变化详情",visible:t.detailForm.dialogVisible},on:{"update:visible":function(e){return t.$set(t.detailForm,"dialogVisible",e)}}},[a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],attrs:{data:t.detailData,"element-loading-text":"Loading",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{fixed:"left",label:"名称",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:"danger",effect:"dark"}},[t._v(t._s(e.row.name))])]}}])}),a("el-table-column",{attrs:{label:"代码",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("span",[t._v(t._s(e.row.secucode))])]}}])}),a("el-table-column",{attrs:{label:"持仓人数",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-tag",{attrs:{type:"warning",effect:"dark"}},[t._v(t._s(e.row.holderTotalNum))])]}}])}),a("el-table-column",{attrs:{label:"持仓变化率",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.totalNumRatio)+" ")]}}])}),a("el-table-column",{attrs:{label:"发布时间",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.releaseDate)+" ")]}}])}),a("el-table-column",{attrs:{label:"集中度",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.holdFocus)+" ")]}}])}),a("el-table-column",{attrs:{label:"收盘价",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.price)+" ")]}}])}),a("el-table-column",{attrs:{label:"非流通股东率",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.holdRatioTotal)+" ")]}}])}),a("el-table-column",{attrs:{label:"流通股东率",width:"110",align:"center"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(e.row.freeholdRatioTotal)+" ")]}}])})],1)],1)],1)],1)},r=[],o=(a("b0c0"),a("a15b"),a("ac1f"),a("1276"),a("d43b")),l={filters:{statusFilter:function(t){var e={published:"success",draft:"gray",deleted:"danger"};return e[t]}},data:function(){return{tableData:null,detailData:null,listLoading:!0,lineChartForm:{activeName:"date",lineChartVisible:!1,lineChartSrc:"",secucode:""},detailForm:{dialogVisible:!1},queryForm:{pageNum:1,pageSize:20,total:0,name:"",secucode:"",dateRange:"",startDate:0,endDate:0,totalRatio:0,sortTyp:0,options:[{value:0,label:"全部"},{value:1,label:"减少"},{value:2,label:"增加"}]}}},mounted:function(){this.fetchData()},methods:{fetchData:function(){var t=this;console.log("==>>TODO 111: ",this.queryForm.sortTyp),this.listLoading=!0;var e={name:this.queryForm.name,secucode:this.queryForm.secucode,decrease:this.queryForm.decrease,releaseStart:this.queryForm.startDate,releaseEnd:this.queryForm.endDate,totalRatio:this.queryForm.totalRatio,limit:this.queryForm.pageSize,sortTyp:this.queryForm.sortTyp,offset:(this.queryForm.pageNum-1)*this.queryForm.pageSize};Object(o["g"])(e).then((function(e){t.tableData=e.data.items,t.listLoading=!1,t.queryForm.total=e.data.total}))},handleCurrentChange:function(t){this.queryForm.pageNum=t,this.fetchData()},handleSizeChange:function(){},onSelectValue:function(){console.log("您选择了",this.queryForm.sortTyp)},onConfirmFocus:function(t,e){var a=this,n=e[t],r={name:n.name,secucode:n.secucode,expectPrice:n.presentPrice};console.log("==>>TODO fucus 01:",r),Object(o["c"])(r).then((function(t){console.log("==>>TODO fucus 02:",t),a.fetchData()}))},onSelectDate:function(t){console.log("==>>TODO date is: ",t[0],t[1]),this.queryForm.startDate=t[0],this.queryForm.endDate=t[1]},onQuerySubmit:function(t){console.log("==>>TODO query is: ",t),this.fetchData()},onRenshuDetail:function(t,e){var a=this;this.listLoading=!0;var n=e[t];this.detailForm.dialogVisible=!0;var r={secucode:n.secucode};Object(o["f"])(r).then((function(t){a.listLoading=!1,a.detailData=t.data.items}))},klineChart:function(t,e){var a=e[t],n=a.secucode.split(".").join("").toLowerCase();this.lineChartForm.secucode=n,this.lineChartForm.lineChartVisible=!this.lineChartForm.lineChartVisible,this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/daily/n/"+n+".gif"},selectTabClick:function(t){var e=this.lineChartForm.secucode;"date"===t.name?this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/daily/n/"+e+".gif":"contour"==t.name?this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/weekly/n/"+e+".gif":this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/min/n/"+e+".gif",this.lineChartForm.activeName="date"}}},i=l,s=(a("28bc"),a("2877")),u=Object(s["a"])(i,n,r,!1,null,"44a80e63",null);e["default"]=u.exports},"28bc":function(t,e,a){"use strict";a("b09d")},a15b:function(t,e,a){"use strict";var n=a("23e7"),r=a("44ad"),o=a("fc6a"),l=a("a640"),i=[].join,s=r!=Object,u=l("join",",");n({target:"Array",proto:!0,forced:s||!u},{join:function(t){return i.call(o(this),void 0===t?",":t)}})},b09d:function(t,e,a){},d43b:function(t,e,a){"use strict";a.d(e,"m",(function(){return r})),a.d(e,"j",(function(){return o})),a.d(e,"p",(function(){return l})),a.d(e,"h",(function(){return i})),a.d(e,"l",(function(){return s})),a.d(e,"c",(function(){return u})),a.d(e,"i",(function(){return c})),a.d(e,"b",(function(){return d})),a.d(e,"n",(function(){return m})),a.d(e,"g",(function(){return f})),a.d(e,"f",(function(){return h})),a.d(e,"o",(function(){return p})),a.d(e,"a",(function(){return b})),a.d(e,"k",(function(){return g})),a.d(e,"e",(function(){return y})),a.d(e,"d",(function(){return v}));var n=a("b775");function r(t){return Object(n["a"])({url:"/api/backend/updateCNConfig",method:"post",data:t})}function o(t){return Object(n["a"])({url:"/api/stock/GetRecommend",method:"post",data:t})}function l(t){return Object(n["a"])({url:"/api/stock/UpdateRecommend",method:"post",data:t})}function i(t){return Object(n["a"])({url:"/api/stock/GetDailyList",method:"post",data:t})}function s(t){return Object(n["a"])({url:"/api/stock/ManualDecreaseList",method:"post",data:t})}function u(t){return Object(n["a"])({url:"/api/stock/ConfirmFocus",method:"post",data:t})}function c(t){return Object(n["a"])({url:"/api/stock/GetFocusList",method:"post",data:t})}function d(t){return Object(n["a"])({url:"/api/stock/CancelFocus",method:"post",data:t})}function m(t){return Object(n["a"])({url:"/api/stock/updateFocus",method:"post",data:t})}function f(t){return Object(n["a"])({url:"/api/stock/GDRenshuList",method:"post",data:t})}function h(t){return Object(n["a"])({url:"/api/stock/GDRenshuDetail",method:"post",data:t})}function p(t){return Object(n["a"])({url:"/api/stock/UpdateGPZhouQi",method:"post",data:t})}function b(t){return Object(n["a"])({url:"/api/stock/AddGPZhouQiRemark",method:"post",data:t})}function g(t){return Object(n["a"])({url:"/api/stock/GPZhouQiList",method:"post",data:t})}function y(t){return Object(n["a"])({url:"/api/stock/GDAggregationReset",method:"post",data:t})}function v(t){return Object(n["a"])({url:"/api/stock/GDAggregationList",method:"post",data:t})}}}]);