(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-39416102"],{"95d1":function(e,t,a){"use strict";a("b2c9")},a15b:function(e,t,a){"use strict";var n=a("23e7"),r=a("44ad"),i=a("fc6a"),o=a("a640"),l=[].join,s=r!=Object,c=o("join",",");n({target:"Array",proto:!0,forced:s||!c},{join:function(e){return l.call(i(this),void 0===e?",":e)}})},a434:function(e,t,a){"use strict";var n=a("23e7"),r=a("23cb"),i=a("a691"),o=a("50c4"),l=a("7b0b"),s=a("65f0"),c=a("8418"),u=a("1dde"),d=a("ae40"),m=u("splice"),f=d("splice",{ACCESSORS:!0,0:0,1:2}),p=Math.max,h=Math.min,b=9007199254740991,g="Maximum allowed length exceeded";n({target:"Array",proto:!0,forced:!m||!f},{splice:function(e,t){var a,n,u,d,m,f,v=l(this),F=o(v.length),_=r(e,F),w=arguments.length;if(0===w?a=n=0:1===w?(a=0,n=F-_):(a=w-2,n=h(p(i(t),0),F-_)),F+a-n>b)throw TypeError(g);for(u=s(v,n),d=0;d<n;d++)m=_+d,m in v&&c(u,d,v[m]);if(u.length=n,a<n){for(d=_;d<F-n;d++)m=d+n,f=d+a,m in v?v[f]=v[m]:delete v[f];for(d=F;d>F-n+a;d--)delete v[d-1]}else if(a>n)for(d=F-n;d>_;d--)m=d+n-1,f=d+a-1,m in v?v[f]=v[m]:delete v[f];for(d=0;d<a;d++)v[d+_]=arguments[d+2];return v.length=F-n+a,u}})},b2c9:function(e,t,a){},bdbe:function(e,t,a){"use strict";a.r(t);var n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"dashboard-container"},[a("div",{staticClass:"app-container"},[a("el-form",{staticClass:"demo-form-inline",attrs:{inline:!0,model:e.queryForm}},[a("el-form-item",{attrs:{label:"名称"}},[a("el-input",{attrs:{placeholder:"茅台股份"},model:{value:e.queryForm.name,callback:function(t){e.$set(e.queryForm,"name",t)},expression:"queryForm.name"}})],1),a("el-form-item",{attrs:{label:"代码"}},[a("el-input",{attrs:{placeholder:"000001"},model:{value:e.queryForm.secucode,callback:function(t){e.$set(e.queryForm,"secucode",t)},expression:"queryForm.secucode"}})],1),a("el-form-item",{attrs:{label:"状态"}},[a("el-select",{attrs:{clearable:"",placeholder:"全部"},model:{value:e.queryForm.disabled,callback:function(t){e.$set(e.queryForm,"disabled",t)},expression:"queryForm.disabled"}},e._l(e.queryForm.options,(function(e){return a("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:e.onQuerySubmit}},[e._v("查询")])],1)],1)],1),a("el-divider"),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],attrs:{data:e.tableData,"element-loading-text":"Loading",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{fixed:"left",label:"名称",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"danger",effect:"dark"}},[e._v(e._s(t.row.name))])]}}])}),a("el-table-column",{attrs:{label:"代码",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.secucode))])]}}])}),a("el-table-column",{attrs:{label:"价格差",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"danger",effect:"light"}},[e._v(e._s(t.row.diffPrice))])]}}])}),a("el-table-column",{attrs:{label:"关注价格",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{effect:"light"}},[e._v(e._s(t.row.focusPrice))])]}}])}),a("el-table-column",{attrs:{label:"期望价格",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"warning",effect:"light"}},[e._v(e._s(t.row.expectPrice))])]}}])}),a("el-table-column",{attrs:{label:"当前价格",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.presentPrice)+" ")]}}])}),a("el-table-column",{attrs:{label:"股东集中度",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.holdFocus)+" ")]}}])}),a("el-table-column",{attrs:{prop:"totalNumRatio",label:"股东人数",width:"110",align:"center",sortable:""},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.totalNumRatio)+" ")]}}])}),a("el-table-column",{attrs:{label:"流通市值(亿)",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.traded)+" ")]}}])}),a("el-table-column",{attrs:{label:"创建日期",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.createDate)+" ")]}}])}),a("el-table-column",{attrs:{label:"更新日期",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.updateDate)+" ")]}}])}),a("el-table-column",{attrs:{fixed:"right",label:"操作",width:"220"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.onUpdateFocus(t.$index,e.tableData)}}},[e._v(" 修改 ")]),a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.onRenshuDetail(t.$index,e.tableData)}}},[e._v(" 股东 ")]),a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.onCancelFocus(t.$index,e.tableData)}}},[e._v(" "+e._s(t.row.focused)+" ")]),a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.klineChart(t.$index,e.tableData)}}},[e._v(" K线图 ")])]}}])})],1),a("div",{staticClass:"gva-pagination"},[a("el-pagination",{attrs:{"current-page":e.paginationForm.pageNum,"page-size":e.paginationForm.pageSize,"page-sizes":[10,30,50,100],total:e.paginationForm.total,layout:"total, sizes, prev, pager, next, jumper"},on:{"current-change":e.handleCurrentChange,"size-change":e.handleSizeChange}})],1),a("div",[a("el-dialog",{attrs:{title:"K线查询",visible:e.lineChartForm.lineChartVisible},on:{"update:visible":function(t){return e.$set(e.lineChartForm,"lineChartVisible",t)}}},[a("el-tabs",{on:{"tab-click":e.selectTabClick},model:{value:e.lineChartForm.activeName,callback:function(t){e.$set(e.lineChartForm,"activeName",t)},expression:"lineChartForm.activeName"}},[a("el-tab-pane",{attrs:{label:"日线",name:"date"}},[a("el-image",{attrs:{src:e.lineChartForm.lineChartSrc}})],1),a("el-tab-pane",{attrs:{label:"周线",name:"contour"}},[a("el-image",{attrs:{src:e.lineChartForm.lineChartSrc}})],1),a("el-tab-pane",{attrs:{label:"分时线",name:"minute"}},[a("el-image",{attrs:{src:e.lineChartForm.lineChartSrc}})],1)],1)],1)],1),a("div",[a("el-drawer",{attrs:{title:"我是标题",visible:e.updateForm.updateFocusDrawer,"with-header":!0},on:{"update:visible":function(t){return e.$set(e.updateForm,"updateFocusDrawer",t)}}},[a("el-form",{ref:"updateForm",staticClass:"demo-ruleForm",attrs:{model:e.updateForm,"label-width":"100px"}},[a("el-form-item",{attrs:{label:"预期买入",prop:"expectPrice",rules:[{required:!0,message:"价格不能为空"},{type:"number",message:"价格必须为数字值"}]}},[a("el-input",{attrs:{type:"expectPrice",autocomplete:"off"},model:{value:e.updateForm.expectPrice,callback:function(t){e.$set(e.updateForm,"expectPrice",e._n(t))},expression:"updateForm.expectPrice"}})],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.onSubmitForm("updateForm")}}},[e._v("提交")]),a("el-button",{on:{click:function(t){return e.onResetForm("updateForm")}}},[e._v("重置")])],1)],1)],1)],1),a("div",[a("el-dialog",{attrs:{title:"股东人数变化详情",visible:e.operateForm.dialogVisible},on:{"update:visible":function(t){return e.$set(e.operateForm,"dialogVisible",t)}}},[a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:e.listLoading,expression:"listLoading"}],attrs:{data:e.operateForm.gudongHistory,"element-loading-text":"Loading",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{fixed:"left",label:"名称",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"danger",effect:"dark"}},[e._v(e._s(t.row.name))])]}}])}),a("el-table-column",{attrs:{label:"代码",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("span",[e._v(e._s(t.row.secucode))])]}}])}),a("el-table-column",{attrs:{label:"持仓人数",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"warning",effect:"dark"}},[e._v(e._s(t.row.holderTotalNum))])]}}])}),a("el-table-column",{attrs:{label:"持仓变化率",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.totalNumRatio)+" ")]}}])}),a("el-table-column",{attrs:{label:"发布时间",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.releaseDate)+" ")]}}])}),a("el-table-column",{attrs:{label:"集中度",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.holdFocus)+" ")]}}])}),a("el-table-column",{attrs:{label:"收盘价",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.price)+" ")]}}])}),a("el-table-column",{attrs:{label:"非流通股东率",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.holdRatioTotal)+" ")]}}])}),a("el-table-column",{attrs:{label:"流通股东率",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(t.row.freeholdRatioTotal)+" ")]}}])})],1)],1)],1)],1)},r=[],i=(a("b0c0"),a("a434"),a("a15b"),a("ac1f"),a("1276"),a("d43b")),o={filters:{statusFilter:function(e){var t={published:"success",draft:"gray",deleted:"danger"};return t[e]}},data:function(){return{tableData:null,timer:null,listLoading:!0,lineChartForm:{activeName:"date",lineChartVisible:!1,lineChartSrc:"",secucode:""},queryForm:{name:"",secucode:"",disabled:0,options:[{value:0,label:"全部"},{value:1,label:"已关注"},{value:2,label:"取消关注"}]},updateForm:{secucode:"",expectPrice:0,updateFocusDrawer:!1},paginationForm:{pageNum:1,pageSize:20,total:0},operateForm:{dialogVisible:!1,gudongHistory:null}}},mounted:function(){this.fetchData()},methods:{fetchData:function(){var e=this;this.listLoading=!0;var t={name:this.queryForm.name,secucode:this.queryForm.secucode,disabled:this.queryForm.disabled,limit:this.paginationForm.pageSize,offset:(this.paginationForm.pageNum-1)*this.paginationForm.pageSize};Object(i["f"])(t).then((function(t){e.tableData=t.data.items,e.listLoading=!1,e.paginationForm.total=t.data.total}))},handleCurrentChange:function(e){this.paginationForm.pageNum=e,this.fetchData()},handleSizeChange:function(e){this.paginationForm.pageNum=0,this.paginationForm.pageSize=e,this.fetchData()},onCancelFocus:function(e,t){var a=this,n=t[e],r={name:n.name,secucode:n.secucode};Object(i["a"])(r).then((function(t){a.tableData.splice(e,1),a.$message({message:"取消成功",type:"success"})}))},onQuerySubmit:function(e){this.fetchData()},klineChart:function(e,t){var a=t[e],n=a.secucode.split(".").join("").toLowerCase();this.lineChartForm.secucode=n,this.lineChartForm.lineChartVisible=!this.lineChartForm.lineChartVisible,this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/daily/n/"+n+".gif"},selectTabClick:function(e){var t=this.lineChartForm.secucode;"date"===e.name?this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/daily/n/"+t+".gif":"contour"==e.name?this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/weekly/n/"+t+".gif":this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/min/n/"+t+".gif",this.lineChartForm.activeName="date"},onUpdateFocus:function(e,t){var a=t[e];this.updateForm.secucode=a.secucode,this.updateForm.updateFocusDrawer=!0},onResetForm:function(e){this.$refs[e].resetFields()},onSubmitForm:function(e){var t=this;this.$refs[e].validate((function(e){e||t.$message({message:"提交失败",type:"warning"});var a={secucode:t.updateForm.secucode,expectPrice:t.updateForm.expectPrice};Object(i["k"])(a).then((function(e){2e4==e.code&&t.$message({message:"编辑成功",type:"success"}),t.fetchData()}))}))},onRenshuDetail:function(e,t){var a=this;this.listLoading=!0;var n=t[e];this.operateForm.dialogVisible=!0;var r={secucode:n.secucode};Object(i["c"])(r).then((function(e){a.listLoading=!1,a.operateForm.gudongHistory=e.data.items}))}}},l=o,s=(a("95d1"),a("2877")),c=Object(s["a"])(l,n,r,!1,null,"3a7aea94",null);t["default"]=c.exports},d43b:function(e,t,a){"use strict";a.d(t,"j",(function(){return r})),a.d(t,"g",(function(){return i})),a.d(t,"m",(function(){return o})),a.d(t,"e",(function(){return l})),a.d(t,"i",(function(){return s})),a.d(t,"b",(function(){return c})),a.d(t,"f",(function(){return u})),a.d(t,"a",(function(){return d})),a.d(t,"k",(function(){return m})),a.d(t,"d",(function(){return f})),a.d(t,"c",(function(){return p})),a.d(t,"l",(function(){return h})),a.d(t,"h",(function(){return b}));var n=a("b775");function r(e){return Object(n["a"])({url:"/api/backend/updateCNConfig",method:"post",data:e})}function i(e){return Object(n["a"])({url:"/api/stock/GetRecommend",method:"post",data:e})}function o(e){return Object(n["a"])({url:"/api/stock/UpdateRecommend",method:"post",data:e})}function l(e){return Object(n["a"])({url:"/api/stock/GetDailyList",method:"post",data:e})}function s(e){return Object(n["a"])({url:"/api/stock/ManualDecreaseList",method:"post",data:e})}function c(e){return Object(n["a"])({url:"/api/stock/ConfirmFocus",method:"post",data:e})}function u(e){return Object(n["a"])({url:"/api/stock/GetFocusList",method:"post",data:e})}function d(e){return Object(n["a"])({url:"/api/stock/CancelFocus",method:"post",data:e})}function m(e){return Object(n["a"])({url:"/api/stock/updateFocus",method:"post",data:e})}function f(e){return Object(n["a"])({url:"/api/stock/GDRenshuList",method:"post",data:e})}function p(e){return Object(n["a"])({url:"/api/stock/GDRenshuDetail",method:"post",data:e})}function h(e){return Object(n["a"])({url:"/api/stock/UpdateGPZhouQi",method:"post",data:e})}function b(e){return Object(n["a"])({url:"/api/stock/GPZhouQiList",method:"post",data:e})}}}]);