<template>
  <div>
    <div class="app-container">
        <el-row type="flex">
            <el-form :inline="true" :model="queryForm" class="demo-form-inline">
            <el-form-item label="代码">
                <el-input v-model="queryForm.name" placeholder="贵州茅台"></el-input>
            </el-form-item>
            <el-form-item label="代码">
                <el-input v-model="queryForm.secucode" placeholder="SZ.000001"></el-input>
            </el-form-item>
            <el-form-item label="跌幅">
                <el-input-number v-model="queryForm.decrease" :step="2"  :min="-100" :max="100" step-strictly></el-input-number>
            </el-form-item>
            <el-form-item label="状态">
                <el-select v-model.number=queryForm.state placeholder="准备" clearable>
                <el-option label="全部" value=0></el-option>
                <el-option label="推荐" value=1></el-option>
                <el-option label="关注" value=2></el-option>
                <el-option label="忽略" value=3></el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onQuerySubmit">查询</el-button>
            </el-form-item>
            </el-form>
        </el-row>
    </div>
    <div class="app-container">
        <el-table :data="tableData" stripe style="width: 100%" max-height="800">
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column fixed="left" class-name="status-col" label="名称" width="120">
            <template slot-scope="scope">
                <el-tag type="danger" effect="dark">{{ scope.row.name }}</el-tag>
            </template>
            </el-table-column>
            <el-table-column prop="secucode" label="代码" width="120"></el-table-column>
            <el-table-column prop="decrease" label="最近跌幅" width="120"></el-table-column>
            <el-table-column prop="inflowRatioStr" label="流入占比" width="120"> </el-table-column>
            <el-table-column class-name="status-col" label="状态" width="110" align="center">
            <template slot-scope="scope">
                <el-tag :type="scope.row.state | statusFilter" effect="dark">{{ scope.row.state }}</el-tag>
            </template>
            </el-table-column>
            
            <el-table-column class-name="status-col" label="最高价格" width="160">
            <template slot-scope="scope">
                <el-tag type="success" effect="dark">{{ scope.row.maxPrice }}</el-tag>
            </template>
            </el-table-column>
            <el-table-column class-name="status-col" label="最低价" width="120">
            <template slot-scope="scope">
                <el-tag type="danger" effect="dark">{{ scope.row.minPrice }}</el-tag>
            </template>
            </el-table-column>
            <el-table-column class-name="status-col" label="收盘价" width="120">
            <template slot-scope="scope">
                <el-tag type="" effect="dark">{{ scope.row.closing }}</el-tag>
            </template>
            </el-table-column>
            <el-table-column class-name="status-col" label="当前价" width="120">
            <template slot-scope="scope">
                <el-tag type="" effect="dark">{{ scope.row.presentPrice }}</el-tag>
            </template>
            </el-table-column>
            <el-table-column prop="traded" label="流通市值" width="120"> </el-table-column>
            
            <el-table-column label="创建时间" width="110" align="center">
              <template slot-scope="scope">
                <el-tag type="success" effect="light">{{ scope.row.createDate|dateFilter }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column fixed="right" label="操作" width="220">
            <template slot-scope="scope">
                <el-button
                @click.native.prevent="modifyRow(scope.$index, tableData)"
                type="text"
                size="small">
                修改
                </el-button>
                <el-divider direction="vertical"></el-divider>
                <el-button
                @click.native.prevent="klineChart(scope.$index, tableData)"
                type="text"
                size="small">
                K线图
                </el-button>
            </template>
            </el-table-column>
        </el-table>
    </div>
    <div class="gva-pagination">
        <el-pagination
        :current-page=pageInfo.pageNum
        :page-size=pageInfo.pageSize
        :page-sizes="[30, 50, 100]"
        :total=pageInfo.total
        layout="total, sizes, prev, pager, next"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        />
    </div>
    <div>
        <el-dialog title="K线查询" :visible.sync="lineChartForm.lineChartVisible">
        <el-tabs v-model="lineChartForm.activeName" @tab-click="selectTabClick">
            <el-tab-pane label="分时线" name="minute">
            <el-image :src="lineChartForm.lineChartSrc"></el-image>
            </el-tab-pane>
            <el-tab-pane label="日线" name="date">
            <el-image :src="lineChartForm.lineChartSrc"></el-image>
            </el-tab-pane>
            <el-tab-pane label="周线" name="contour">
            <el-image :src="lineChartForm.lineChartSrc"></el-image>
            </el-tab-pane>
        </el-tabs>
        </el-dialog>
    </div>
  </div>
</template>

<script>
import echarts from 'echarts'
import { getWaterfallList } from '@/api/stock'
import { parseTime } from '@/utils/index'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        1: 'warning',
        2: 'danger',
        3: 'info',
      }
      return statusMap[status]
    },
    dateFilter(time) {
        return parseTime(time)
    }
  },
  data() {
    return {
      listLoading: true,
      modifyDialogVisible: false,
      queryForm: {
        name: '',
        secucode: '',
        decrease: 0,
        state: 0,
      },
      modifyForm: {
        name: '',
        region: '',
        secucode: '',
        priceDecrease: 0,
      },
      lineChartForm:{
        activeName: 'date',
        lineChartVisible: false,
        lineChartSrc: '',
        secucode: '',
      },
      formLabelWidth: '120px',
      tableData: [],
      pageInfo: {
        pageNum: 1,
        pageSize: 30,
        total: 0,
      }
    }
  },
  mounted() {
    // this.initChart()
  },
  // 每次销毁前清空
  beforeDestroy() {
    // if (!this.chart) {
    //   return
    // }
    // this.fundForm.chart.dispose()
    // this.fundForm.chart = null
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      console.log("==>>TODO 121:")
      this.listLoading = true
       var req = {
        name: this.queryForm.name,
        limit: this.pageInfo.pageSize,
        offset: (this.pageInfo.pageNum-1)*this.pageInfo.pageSize,
        decrease:this.queryForm.decrease,
        state: this.queryForm.state,
        secucode: this.queryForm.secucode,
      }
      getWaterfallList(req).then(response => {
        this.tableData = response.data.items
        console.log("==>>TODO 122:", this.tableData.length)
        this.pageInfo.total = response.data.total
        this.listLoading = false
      })
    },
    klineChart(index, rows) {//日线图
      var data = rows[index]
      var secucode = data.secucode.split('.').join("").toLowerCase()
      this.lineChartForm.secucode = secucode
      this.lineChartForm.lineChartVisible = !this.lineChartForm.lineChartVisible
      this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/daily/n/'+secucode+'.gif'
    },
    cancelDialog(index, rows) {
      this.modifyDialogVisible = !this.modifyDialogVisible
    },
    // confirmDialog(formName) {
    //   var form = this.modifyForm

    //   var req = {
    //     name : this.modifyForm.name,
    //     secucode: this.modifyForm.secucode, 
    //     priceDecrease: this.modifyForm.decrease
    //   }

    //   updateRecommend(req).then(response=>{
    //     this.modifyDialogVisible = !this.modifyDialogVisible
    //     this.fetchData()
    //   })
      
    //   this.modifyForm.decrease = 0
    // },
    onQuerySubmit() {
      console.log('submit!', this.queryForm.state, this.queryForm.decrease);
      this.fetchData()
    },
    handleCurrentChange(val) {
      this.pageInfo.pageNum = val
      this.fetchData()
    },
    handleSizeChange(val) {
      this.pageInfo.pageSize = val
    },
    selectTabClick(tab) {
      var secucode = this.lineChartForm.secucode
      if (tab.name === "date") {
        this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/daily/n/'+secucode+'.gif'
      } else if (tab.name == "contour") { //contour
        this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/weekly/n/'+secucode+'.gif'
      } else { //分时线
        this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/min/n/'+secucode+'.gif'
      }
      this.lineChartForm.activeName = "date"
    },
    // onSelectionChange(values) {
    //   if (!values) {
    //     return 
    //   }
    //   for (var idx in values) {
    //     var val = values[idx]
    //     this.fundForm.secucodes.push(val.secucode)
    //   }
    // },
    // onLineChart() {
    //   this.fundForm.fundFlowDrawer = true
    //   var req = {
    //     secucodes: this.fundForm.secucodes,
    //   }
    //   getFundDetailList(req).then(response=>{
    //     this.fundForm.opinionData = response.data.items
    //     this.fundForm.legendData = response.data.legendData
    //     this.fundForm.secucodes = []

    //     this.drawLine("fundLineChart")
    //   })
    // },
    // drawLine(id) {
    //   this.charts = echarts.init(document.getElementById(id))
    //   this.charts.setOption({
    //       title: {
    //         text: ''
    //       },
    //       tooltip: {
    //         trigger: 'axis'
    //       },
    //       legend: {
    //         data: this.fundForm.legendData
    //       },
    //       grid: {
    //         left: '3%',
    //         right: '4%',
    //         bottom: '3%',
    //         containLabel: true
    //       },
    //       toolbox: {
    //         feature: {
    //           saveAsImage: {}
    //         }
    //       },
    //       xAxis: {
    //         type: 'category',
    //         boundaryGap: false,
    //         data: ["1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19","20"]
    //       },
    //       yAxis: {
    //         type: 'value'
    //       },
    //       series: this.fundForm.opinionData
    //   })
    // },
  }
}
</script>
