(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-0bf36107"],{"101a":function(e,t,a){"use strict";a.r(t);var i=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("div",{staticClass:"app-container"},[a("el-form",{staticClass:"demo-form-inline",attrs:{inline:!0,model:e.queryForm}},[a("el-form-item",{attrs:{label:"代码"}},[a("el-input",{attrs:{placeholder:"SZ.000001"},model:{value:e.queryForm.secucode,callback:function(t){e.$set(e.queryForm,"secucode",t)},expression:"queryForm.secucode"}})],1),a("el-form-item",{attrs:{label:"跌幅"}},[a("el-input",{attrs:{placeholder:"20"},model:{value:e.queryForm.decrease,callback:function(t){e.$set(e.queryForm,"decrease",e._n(t))},expression:"queryForm.decrease"}})],1),a("el-form-item",{attrs:{label:"状态"}},[a("el-select",{attrs:{placeholder:"准备"},model:{value:e.queryForm.state,callback:function(t){e.$set(e.queryForm,"state",e._n(t))},expression:"queryForm.state"}},[a("el-option",{attrs:{label:"待定",value:"0"}}),a("el-option",{attrs:{label:"准备",value:"1"}}),a("el-option",{attrs:{label:"开始",value:"2"}}),a("el-option",{attrs:{label:"进行中",value:"3"}}),a("el-option",{attrs:{label:"结束",value:"4"}})],1)],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:e.onQuerySubmit}},[e._v("查询")])],1)],1)],1),a("div",{staticClass:"app-container"},[a("el-table",{staticStyle:{width:"100%"},attrs:{data:e.tableData,stripe:"","max-height":"800"}},[a("el-table-column",{attrs:{fixed:"",prop:"secucode",label:"代码",width:"150"}}),a("el-table-column",{attrs:{prop:"name",label:"名称",width:"120"}}),a("el-table-column",{attrs:{prop:"rMIndex",label:"推荐指数",width:"120"}}),a("el-table-column",{attrs:{"class-name":"status-col",label:"状态",width:"110",align:"center"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:e._f("statusFilter")(t.row.state),effect:"dark"}},[e._v(e._s(t.row.state))])]}}])}),a("el-table-column",{attrs:{prop:"pDecrease",label:"最近跌幅",width:"120"}}),a("el-table-column",{attrs:{prop:"maxPrice",label:"最高价",width:"120"}}),a("el-table-column",{attrs:{prop:"rMPrice",label:"推荐价格",width:"220"}}),a("el-table-column",{attrs:{prop:"presentPrice",label:"当前价",width:"120"}}),a("el-table-column",{attrs:{prop:"gDDecrease",label:"股东人数",width:"120"}}),a("el-table-column",{attrs:{prop:"updateNum",label:"更新次数",width:"80"}}),a("el-table-column",{attrs:{"class-name":"status-col",label:"参考幅度",width:"80"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-tag",{attrs:{type:"warning",effect:"dark"}},[e._v(e._s(t.row.referDecrease))])]}}])}),a("el-table-column",{attrs:{prop:"updateDate",label:"最近更新时间",width:"120"}}),a("el-table-column",{attrs:{fixed:"right",label:"操作",width:"220"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.modifyRow(t.$index,e.tableData)}}},[e._v(" 修改 ")]),a("el-divider",{attrs:{direction:"vertical"}}),a("el-button",{attrs:{type:"text",size:"small"},nativeOn:{click:function(a){return a.preventDefault(),e.KLineChart(t.$index,e.tableData)}}},[e._v(" K线图 ")])]}}])})],1),a("div",{staticClass:"gva-pagination"},[a("el-pagination",{attrs:{"current-page":e.pageInfo.pageNum,"page-size":e.pageInfo.pageSize,"page-sizes":[10,30,50,100],total:e.pageInfo.total,layout:"total, sizes, prev, pager, next, jumper"},on:{"current-change":e.handleCurrentChange,"size-change":e.handleSizeChange}})],1)],1),a("div",[a("el-dialog",{attrs:{title:"修改推荐",visible:e.modifyDialogVisible},on:{"update:visible":function(t){e.modifyDialogVisible=t}}},[a("el-form",{ref:"modifyForm",staticClass:"demo-ruleForm",attrs:{model:e.modifyForm,"label-width":"100px"}},[a("el-form-item",{attrs:{label:"跌幅修正",prop:"decrease",rules:[{type:"number",message:"跌幅必须为数字值"}]}},[a("el-input",{attrs:{type:"age",autocomplete:"off"},model:{value:e.modifyForm.decrease,callback:function(t){e.$set(e.modifyForm,"decrease",e._n(t))},expression:"modifyForm.decrease"}})],1),a("el-form-item",[a("el-button",{on:{click:function(t){return e.cancelDialog("modifyForm")}}},[e._v("取消")]),a("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.confirmDialog("modifyForm")}}},[e._v("提交")])],1)],1)],1)],1),a("div",[a("el-dialog",{attrs:{title:"K线查询",visible:e.lineChartForm.lineChartVisible},on:{"update:visible":function(t){return e.$set(e.lineChartForm,"lineChartVisible",t)}}},[a("el-tabs",{on:{"tab-click":e.handleClick},model:{value:e.lineChartForm.activeName,callback:function(t){e.$set(e.lineChartForm,"activeName",t)},expression:"lineChartForm.activeName"}},[a("el-tab-pane",{attrs:{label:"日线",name:"date"}},[a("el-image",{attrs:{src:e.lineChartForm.lineChartSrc}})],1),a("el-tab-pane",{attrs:{label:"周线",name:"contour"}},[a("el-image",{attrs:{src:e.lineChartForm.lineChartSrc}})],1),a("el-tab-pane",{attrs:{label:"分时线",name:"minute"}},[a("el-image",{attrs:{src:e.lineChartForm.lineChartSrc}})],1)],1)],1)],1)])},o=[],l=(a("b0c0"),a("a15b"),a("ac1f"),a("1276"),a("d43b")),r={filters:{statusFilter:function(e){var t={"进行中":"danger","开始":"warning","准备":"info"};return t[e]}},data:function(){return{listLoading:!0,modifyDialogVisible:!1,queryForm:{secucode:"",region:1,decrease:0,state:0},modifyForm:{name:"",region:"",secucode:"",priceDecrease:0},lineChartForm:{activeName:"date",lineChartVisible:!1,lineChartSrc:"",secucode:""},formLabelWidth:"120px",tableData:[],pageInfo:{pageNum:1,pageSize:20,total:0}}},created:function(){this.fetchData()},methods:{fetchData:function(){var e=this;this.listLoading=!0;var t={limit:this.pageInfo.pageSize,offset:(this.pageInfo.pageNum-1)*this.pageInfo.pageSize,pDecrease:this.queryForm.decrease,state:this.queryForm.state,secucode:this.queryForm.secucode};Object(l["b"])(t).then((function(t){console.log("===>>TODO 111: ",t),e.tableData=t.data.items,e.pageInfo.total=t.data.total,e.listLoading=!1})),console.log("===>>TODO 211: ",this.tableData)},modifyRow:function(e,t){var a=t[e];this.modifyDialogVisible=!this.modifyDialogVisible,this.modifyForm.secucode=a.secucode,this.modifyForm.priceDecrease=a.pDecrease,this.modifyForm.name=a.name},KLineChart:function(e,t){console.log("==>>TODO 3141: ",t[e]);var a=t[e],i=a.secucode.split(".").join("").toLowerCase();this.lineChartForm.secucode=i,console.log("==>>TODO 3142: ",i),this.lineChartForm.lineChartVisible=!this.lineChartForm.lineChartVisible,this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/daily/n/"+i+".gif"},cancelDialog:function(e,t){this.modifyDialogVisible=!this.modifyDialogVisible},confirmDialog:function(e){var t=this;console.log("==>>TODO 3031: ",this.modifyForm),console.log("==>>TODO 3032: ",this.modifyForm.decrease);var a=this.modifyForm;console.log("==>>TODO 3035: ",this.$refs[e],a);var i={name:this.modifyForm.name,secucode:this.modifyForm.secucode,priceDecrease:this.modifyForm.decrease};Object(l["d"])(i).then((function(e){console.log("==>>TODO 3036: ","ok"),t.modifyDialogVisible=!t.modifyDialogVisible,t.fetchData()})),this.modifyForm.decrease=0},onQuerySubmit:function(){console.log("submit!",this.queryForm.state,this.queryForm.decrease),this.fetchData(),console.log("===>>TODO 212: ",this.tableData)},handleCurrentChange:function(e){console.log("===>>TODO 2131: ",e),this.pageInfo.pageNum=e,console.log("===>>TODO 2132: ",this.pageInfo.pageNum);this.pageInfo.pageSize,this.pageInfo.pageNum,this.pageInfo.pageSize;this.fetchData()},handleSizeChange:function(e){console.log("===>>TODO 214: ",e)},handleClick:function(e){console.log("===>>TODO 254: ",e);var t=this.lineChartForm.secucode;"date"===e.name?this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/daily/n/"+t+".gif":"contour"==e.name?this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/weekly/n/"+t+".gif":this.lineChartForm.lineChartSrc="http://image.sinajs.cn/newchart/min/n/"+t+".gif",this.lineChartForm.activeName="date"}}},n=r,s=a("2877"),c=Object(s["a"])(n,i,o,!1,null,null,null);t["default"]=c.exports},a15b:function(e,t,a){"use strict";var i=a("23e7"),o=a("44ad"),l=a("fc6a"),r=a("a640"),n=[].join,s=o!=Object,c=r("join",",");i({target:"Array",proto:!0,forced:s||!c},{join:function(e){return n.call(l(this),void 0===e?",":e)}})},d43b:function(e,t,a){"use strict";a.d(t,"b",(function(){return o})),a.d(t,"d",(function(){return l})),a.d(t,"c",(function(){return r}));var i=a("b775");function o(e){return Object(i["a"])({url:"/api/stock/GetRecommend",method:"post",data:e})}function l(e){return Object(i["a"])({url:"/api/stock/UpdateRecommend",method:"post",data:e})}function r(e){return Object(i["a"])({url:"/api/stock/PromptBuyList",method:"post",data:e})}}}]);