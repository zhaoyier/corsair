(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-57015196"],{"0ccb":function(e,t,a){var r=a("50c4"),n=a("1148"),o=a("1d80"),i=Math.ceil,l=function(e){return function(t,a,l){var s,c,u=String(o(t)),m=u.length,d=void 0===l?" ":String(l),p=r(a);return p<=m||""==d?u:(s=p-m,c=n.call(d,i(s/d.length)),c.length>s&&(c=c.slice(0,s)),e?u+c:c+u)}};e.exports={start:l(!1),end:l(!0)}},1014:function(e,t,a){},1148:function(e,t,a){"use strict";var r=a("a691"),n=a("1d80");e.exports="".repeat||function(e){var t=String(n(this)),a="",o=r(e);if(o<0||o==1/0)throw RangeError("Wrong number of repetitions");for(;o>0;(o>>>=1)&&(t+=t))1&o&&(a+=t);return a}},"4d90":function(e,t,a){"use strict";var r=a("23e7"),n=a("0ccb").start,o=a("9a0c");r({target:"String",proto:!0,forced:o},{padStart:function(e){return n(this,e,arguments.length>1?arguments[1]:void 0)}})},"9a0c":function(e,t,a){var r=a("342f");e.exports=/Version\/10\.\d+(\.\d+)?( Mobile\/\w+)? Safari\//.test(r)},a15b:function(e,t,a){"use strict";var r=a("23e7"),n=a("44ad"),o=a("fc6a"),i=a("a640"),l=[].join,s=n!=Object,c=i("join",",");r({target:"Array",proto:!0,forced:s||!c},{join:function(e){return l.call(o(this),void 0===e?",":e)}})},c7d5:function(e,t,a){"use strict";a("1014")},d43b:function(e,t,a){"use strict";a.d(t,"n",(function(){return n})),a.d(t,"k",(function(){return o})),a.d(t,"q",(function(){return i})),a.d(t,"h",(function(){return l})),a.d(t,"m",(function(){return s})),a.d(t,"c",(function(){return c})),a.d(t,"i",(function(){return u})),a.d(t,"b",(function(){return m})),a.d(t,"o",(function(){return d})),a.d(t,"g",(function(){return p})),a.d(t,"f",(function(){return f})),a.d(t,"p",(function(){return h})),a.d(t,"a",(function(){return b})),a.d(t,"l",(function(){return F})),a.d(t,"e",(function(){return g})),a.d(t,"d",(function(){return v})),a.d(t,"j",(function(){return k}));var r=a("b775");function n(e){return Object(r["a"])({url:"/api/backend/updateCNConfig",method:"post",data:e})}function o(e){return Object(r["a"])({url:"/api/stock/GetRecommend",method:"post",data:e})}function i(e){return Object(r["a"])({url:"/api/stock/UpdateRecommend",method:"post",data:e})}function l(e){return Object(r["a"])({url:"/api/stock/GetDailyList",method:"post",data:e})}function s(e){return Object(r["a"])({url:"/api/stock/ManualDecreaseList",method:"post",data:e})}function c(e){return Object(r["a"])({url:"/api/stock/ConfirmFocus",method:"post",data:e})}function u(e){return Object(r["a"])({url:"/api/stock/GetFocusList",method:"post",data:e})}function m(e){return Object(r["a"])({url:"/api/stock/CancelFocus",method:"post",data:e})}function d(e){return Object(r["a"])({url:"/api/stock/updateFocus",method:"post",data:e})}function p(e){return Object(r["a"])({url:"/api/stock/GDRenshuList",method:"post",data:e})}function f(e){return Object(r["a"])({url:"/api/stock/GDRenshuDetail",method:"post",data:e})}function h(e){return Object(r["a"])({url:"/api/stock/UpdateGPZhouQi",method:"post",data:e})}function b(e){return Object(r["a"])({url:"/api/stock/AddGPZhouQiRemark",method:"post",data:e})}function F(e){return Object(r["a"])({url:"/api/stock/GPZhouQiList",method:"post",data:e})}function g(e){return Object(r["a"])({url:"/api/stock/GDAggregationReset",method:"post",data:e})}function v(e){return Object(r["a"])({url:"/api/stock/GDAggregationList",method:"post",data:e})}function k(e){return Object(r["a"])({url:"/api/stock/GetFundFlowList",method:"post",data:e})}},ed08:function(e,t,a){"use strict";a.d(t,"a",(function(){return n}));a("a4d3"),a("e01a"),a("d3b7"),a("d28b"),a("3ca3"),a("ddb0");function r(e){return r="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e},r(e)}a("ac1f"),a("00b4"),a("5319"),a("4d63"),a("2c3e"),a("25f0"),a("4d90"),a("1276"),a("159b");function n(e,t){if(0===arguments.length||!e)return null;var a,n=t||"{y}-{m}-{d} {h}:{i}:{s}";"object"===r(e)?a=e:("string"===typeof e&&(e=/^[0-9]+$/.test(e)?parseInt(e):e.replace(new RegExp(/-/gm),"/")),"number"===typeof e&&10===e.toString().length&&(e*=1e3),a=new Date(e));var o={y:a.getFullYear(),m:a.getMonth()+1,d:a.getDate(),h:a.getHours(),i:a.getMinutes(),s:a.getSeconds(),a:a.getDay()},i=n.replace(/{([ymdhisa])+}/g,(function(e,t){var a=o[t];return"a"===t?["日","一","二","三","四","五","六"][a]:a.toString().padStart(2,"0")}));return i}},f305:function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"dashboard-container"},[a("div",{staticClass:"app-container"},[a("el-form",{staticClass:"demo-form-inline",staticStyle:{float:"left"},attrs:{inline:!0,model:e.queryForm}},[a("el-form-item",{attrs:{label:"名称"}},[a("el-input",{attrs:{placeholder:"茅台股份"},model:{value:e.queryForm.name,callback:function(t){e.$set(e.queryForm,"name",t)},expression:"queryForm.name"}})],1),a("el-form-item",{attrs:{label:"代码"}},[a("el-input",{attrs:{placeholder:"SZ.000001"},model:{value:e.queryForm.secucode,callback:function(t){e.$set(e.queryForm,"secucode",t)},expression:"queryForm.secucode"}})],1),a("el-form-item",{attrs:{label:"关注"}},[a("el-select",{attrs:{clearable:"",placeholder:"全部"},model:{value:e.queryForm.disabled,callback:function(t){e.$set(e.queryForm,"disabled",t)},expression:"queryForm.disabled"}},e._l(e.queryForm.focusOpts,(function(e){return a("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1),a("el-form-item",{attrs:{label:"状态"}},[a("el-select",{attrs:{clearable:"",placeholder:"全部"},model:{value:e.queryForm.state,callback:function(t){e.$set(e.queryForm,"state",t)},expression:"queryForm.state"}},e._l(e.queryForm.options,(function(e){return a("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:e.onQuerySubmit}},[e._v("查询")])],1)],1),a("el-button",{staticStyle:{float:"right"},attrs:{type:"primary"},on:{click:e.onCreateSubmit}},[e._v("新建周期")])],1),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],attrs:{data:e.tableData,"element-loading-text":"Loading",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{fixed:"left",label:"名称",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"danger",effect:"dark"}},[e._v(e._s(t.row.name))])]}}])}),a("el-table-column",{attrs:{label:"代码",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.secucode))])]}}])}),a("el-table-column",{attrs:{label:"状态",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:e._f("statusFilter")(t.row.state),effect:"dark"}},[e._v(e._s(t.row.state))])]}}])}),a("el-table-column",{attrs:{label:"当前价格",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.presentPrice)+" ")]}}])}),a("el-table-column",{attrs:{label:"期望最低价",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"danger",effect:"light"}},[e._v(e._s(t.row.expectMin))])]}}])}),a("el-table-column",{attrs:{label:"期望最高价",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{effect:"light"}},[e._v(e._s(t.row.expectMax))])]}}])}),a("el-table-column",{attrs:{label:"期望开始时间",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"warning",effect:"light"}},[e._v(e._s(e._f("dateFilter")(t.row.expectStart)))])]}}])}),a("el-table-column",{attrs:{label:"期望结束时间",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"warning",effect:"light"}},[e._v(e._s(e._f("dateFilter")(t.row.expectEnd)))])]}}])}),a("el-table-column",{attrs:{label:"最近备注",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.remark)+" ")]}}])}),a("el-table-column",{attrs:{label:"主营业务",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"warning",effect:"light"}},[e._v(e._s(t.row.mainBusiness))])]}}])}),a("el-table-column",{attrs:{label:"更新日期",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.updateDate)+" ")]}}])}),a("el-table-column",{attrs:{fixed:"right",label:"操作",width:"220"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.onUpdateZhouQi(t.$index,e.tableData)}}},[e._v(" 修改 ")]),a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.onRemarkZhouQi(t.$index,e.tableData)}}},[e._v(" 备注 ")]),a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.onKlineChart(t.$index,e.tableData)}}},[e._v(" K线图 ")])]}}])})],1),a("div",{staticClass:"gva-pagination"},[a("el-pagination",{attrs:{"current-page":e.paginationForm.pageNum,"page-size":e.paginationForm.pageSize,"page-sizes":[10,30,50,100],total:e.paginationForm.total,layout:"total, sizes, prev, pager, next, jumper"},on:{"current-change":e.handleCurrentChange,"size-change":e.handleSizeChange}})],1),a("div",[a("el-dialog",{attrs:{title:"K线查询",visible:e.lineChartForm.lineChartVisible},on:{"update:visible":function(t){return e.$set(e.lineChartForm,"lineChartVisible",t)}}},[a("el-tabs",{on:{"tab-click":e.onSelectTabClick},model:{value:e.lineChartForm.activeName,callback:function(t){e.$set(e.lineChartForm,"activeName",t)},expression:"lineChartForm.activeName"}},[a("el-tab-pane",{attrs:{label:"日线",name:"date"}},[a("el-image",{attrs:{src:e.lineChartForm.lineChartSrc}})],1),a("el-tab-pane",{attrs:{label:"周线",name:"contour"}},[a("el-image",{attrs:{src:e.lineChartForm.lineChartSrc}})],1),a("el-tab-pane",{attrs:{label:"分时线",name:"minute"}},[a("el-image",{attrs:{src:e.lineChartForm.lineChartSrc}})],1)],1)],1)],1),a("div",[a("el-drawer",{attrs:{title:"更新周期信息",visible:e.updateForm.updateZhouQiDrawer,"with-header":!0},on:{"update:visible":function(t){return e.$set(e.updateForm,"updateZhouQiDrawer",t)}}},[a("el-form",{ref:"updateForm",staticClass:"demo-ruleForm",attrs:{model:e.updateForm,"label-width":"100px"}},[a("el-form-item",{attrs:{label:"代码",prop:"secucode",rules:[{required:!0,message:"代码不能为空"},{type:"string",message:"代码必须为字符串"}]}},[a("el-input",{attrs:{type:"secucode",autocomplete:"off"},model:{value:e.updateForm.secucode,callback:function(t){e.$set(e.updateForm,"secucode",e._n(t))},expression:"updateForm.secucode"}})],1),a("el-form-item",{attrs:{label:"预期最低价",prop:"expectMin",rules:[{required:!0,message:"价格不能为空"},{type:"number",message:"价格必须为数字值"}]}},[a("el-input-number",{attrs:{type:"expectMin",precision:2,step:.5,max:1e4},model:{value:e.updateForm.expectMin,callback:function(t){e.$set(e.updateForm,"expectMin",t)},expression:"updateForm.expectMin"}})],1),a("el-form-item",{attrs:{label:"预期最高价",prop:"expectMax",rules:[{required:!0,message:"价格不能为空"},{type:"number",message:"价格必须为数字值"}]}},[a("el-input-number",{attrs:{type:"expectMax",precision:2,step:.5,max:1e4},model:{value:e.updateForm.expectMax,callback:function(t){e.$set(e.updateForm,"expectMax",t)},expression:"updateForm.expectMax"}})],1),a("el-form-item",{attrs:{label:"取消关注"}},[a("el-switch",{attrs:{"active-color":"#13ce66","active-value":1,"inactive-value":0},model:{value:e.updateForm.disabled,callback:function(t){e.$set(e.updateForm,"disabled",t)},expression:"updateForm.disabled"}})],1),a("el-form-item",{attrs:{label:"主营业务"}},[a("el-input",{attrs:{type:"textarea"},model:{value:e.updateForm.mainBusiness,callback:function(t){e.$set(e.updateForm,"mainBusiness",t)},expression:"updateForm.mainBusiness"}})],1),a("el-form-item",{attrs:{label:"预期时间"}},[a("el-date-picker",{attrs:{"value-format":"timestamp",type:"daterange","range-separator":"至","start-placeholder":"开始日期","end-placeholder":"结束日期"},on:{input:e.onSelectDate},model:{value:e.updateForm.expectDate,callback:function(t){e.$set(e.updateForm,"expectDate",t)},expression:"updateForm.expectDate"}})],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.onSubmitForm("updateForm")}}},[e._v("提交")]),a("el-button",{on:{click:function(t){return e.onResetForm("updateForm")}}},[e._v("重置")])],1)],1)],1)],1),a("div",[a("el-drawer",{attrs:{title:"备注信息",visible:e.remarkForm.remarkZhouQiDrawer,"with-header":!0},on:{"update:visible":function(t){return e.$set(e.remarkForm,"remarkZhouQiDrawer",t)}}},[a("el-row",{attrs:{gutter:20}},[a("el-col",{attrs:{span:18,offset:1}},[a("el-tabs",{on:{"tab-click":e.onSelectRemarkTab},model:{value:e.remarkForm.activeName,callback:function(t){e.$set(e.remarkForm,"activeName",t)},expression:"remarkForm.activeName"}},[a("el-tab-pane",{attrs:{label:"备注列表",name:"list"}},[a("el-table",{staticStyle:{width:"100%"},attrs:{data:e.remarkForm.remarkList,stripe:""}},[a("el-table-column",{attrs:{prop:"remark",label:"备注",width:"180"}}),a("el-table-column",{attrs:{prop:"createDate",label:"日期",width:"180"}})],1)],1),a("el-tab-pane",{attrs:{label:"新增备注",name:"create"}},[a("el-form",{attrs:{"label-position":e.remarkForm.labelPosition,"label-width":"80px",model:e.remarkForm}},[a("el-form-item",{attrs:{label:"名称"}},[a("el-input",{model:{value:e.remarkForm.content,callback:function(t){e.$set(e.remarkForm,"content",t)},expression:"remarkForm.content"}})],1),a("el-form-item",{attrs:{label:"活动区域"}},[a("el-date-picker",{attrs:{"value-format":"timestamp",type:"date",placeholder:"选择日期"},model:{value:e.remarkForm.createDate,callback:function(t){e.$set(e.remarkForm,"createDate",t)},expression:"remarkForm.createDate"}})],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.onCreateRemark("remarkForm")}}},[e._v("提交")]),a("el-button",{on:{click:function(t){return e.onResetRemark("remarkForm")}}},[e._v("重置")])],1)],1)],1)],1)],1)],1)],1)],1)],1)},n=[],o=(a("b0c0"),a("a15b"),a("ac1f"),a("1276"),a("d43b")),i=a("ed08"),l={filters:{statusFilter:function(e){var t={"已达时间":"danger","已达价格":"warning","待定":"success"};return t[e]},dateFilter:function(e){return Object(i["a"])(e)}},data:function(){return{activeName:"second",tableData:null,timer:null,listLoading:!0,lineChartForm:{activeName:"date",lineChartVisible:!1,lineChartSrc:"",secucode:""},queryForm:{name:"",secucode:"",state:0,disabled:0,options:[{value:0,label:"全部"},{value:1,label:"已达时间"},{value:2,label:"已达价格"}],focusOpts:[{value:0,label:"全部"},{value:1,label:"已关注"},{value:2,label:"取消关注"}]},updateForm:{secucode:"",expectMin:0,expectMax:0,expectStart:0,expectEnd:0,expectDate:[],disabled:0,remark:"",updateZhouQiDrawer:!1,mainBusiness:""},remarkForm:{content:"",createDate:"",labelPosition:"right",remarkZhouQiDrawer:!1,remarkList:null,activeName:"list"},paginationForm:{pageNum:1,pageSize:20,total:0}}},mounted:function(){this.fetchData()},methods:{fetchData:function(){var e=this;this.listLoading=!0;var t={name:this.queryForm.name,secucode:this.queryForm.secucode,state:this.queryForm.state,disabled:this.queryForm.disabled,limit:this.paginationForm.pageSize,offset:(this.paginationForm.pageNum-1)*this.paginationForm.pageSize};Object(o["l"])(t).then((function(t){e.tableData=t.data.items,e.listLoading=!1,e.paginationForm.total=t.data.total}))},handleCurrentChange:function(e){this.paginationForm.pageNum=e},handleSizeChange:function(e){this.paginationForm.pageSize=e},onQuerySubmit:function(e){this.fetchData()},onKlineChart:function(e,t){var a=t[e],r=a.secucode.split(".").join("").toLowerCase();this.lineChartForm.secucode=r,this.lineChartForm.activeName="date",this.lineChartForm.lineChartVisible=!this.lineChartForm.lineChartVisible,this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/daily/n/"+r+".gif"},onSelectTabClick:function(e,t){var a=this.lineChartForm.secucode;"date"===e.name?this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/daily/n/"+a+".gif":"contour"==e.name?this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/weekly/n/"+a+".gif":this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/min/n/"+a+".gif",this.lineChartForm.activeName="date"},onUpdateZhouQi:function(e,t){var a=t[e];this.updateForm.secucode=a.secucode,this.updateForm.expectMin=a.expectMin,this.updateForm.expectMax=a.expectMax,this.updateForm.disabled=a.disabled?1:0,this.updateForm.updateZhouQiDrawer=!0,this.updateForm.expectStart=1e3*a.expectStart,this.updateForm.expectEnd=1e3*a.expectEnd,0==this.updateForm.expectStart&&(this.updateForm.expectStart=Date.parse(new Date)),0==this.updateForm.expectEnd&&(this.updateForm.expectEnd=Date.parse(new Date)),console.log("==>>TODO 221: ",this.updateForm.expectStart),this.updateForm.expectDate=[new Date(this.updateForm.expectStart),this.updateForm.expectEnd]},onRemarkZhouQi:function(e,t){var a=t[e];this.remarkForm.secucode=a.secucode,this.remarkForm.remarkList=a.remarks,this.remarkForm.remarkZhouQiDrawer=!0},onResetForm:function(e){this.$refs[e].resetFields()},onSubmitForm:function(e){var t=this;this.$refs[e].validate((function(e){e||t.$message({message:"提交失败",type:"warning"});var a={secucode:t.updateForm.secucode,expectMin:t.updateForm.expectMin,expectMax:t.updateForm.expectMax,expectStart:t.updateForm.expectStart,expectEnd:t.updateForm.expectEnd,disabled:1==t.updateForm.disabled,mainBusiness:t.updateForm.mainBusiness};Object(o["p"])(a).then((function(e){2e4==e.code&&t.$message({message:"编辑成功",type:"success"}),t.fetchData()}))}))},onCreateSubmit:function(){this.updateForm.updateZhouQiDrawer=!0},onSelectDate:function(e){e?(this.updateForm.expectStart=e[0],this.updateForm.expectEnd=e[1]):(this.updateForm.expectStart=0,this.updateForm.expectEnd=0)},onSelectRemarkTab:function(e){"list"===e.name?this.remarkForm.activeName="list":this.remarkForm.activeName="create"},onCreateRemark:function(){console.log("==>>551: ",this.remarkForm.secucode,this.remarkForm.createDate),console.log("==>>551: ",this.remarkForm.content,this.remarkForm.createDate);var e={secucode:this.remarkForm.secucode,content:this.remarkForm.content,createDate:this.remarkForm.createDate};Object(o["a"])(e).then((function(e){}))},onResetRemark:function(){}}},s=l,c=(a("c7d5"),a("2877")),u=Object(c["a"])(s,r,n,!1,null,"7bb5f531",null);t["default"]=u.exports}}]);