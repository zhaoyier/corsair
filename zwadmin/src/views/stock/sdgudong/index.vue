<template>
  <div>
  <div class="app-container">
    <el-form :inline="true" :model="queryForm" class="demo-form-inline">
      <el-form-item label="关键字">
        <el-input v-model="queryForm.keyword" placeholder="社保"></el-input>
      </el-form-item>
      <el-form-item label="日期">
        <el-date-picker
          @input="onSelectDate"
          v-model="queryForm.dateRange"
          value-format="timestamp"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期">
        </el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onQuerySubmit">查询</el-button>
      </el-form-item>
    </el-form>
  </div>
  <div class="app-container">
    <!-- <el-table :data="queryForm.tableData" stripe style="width: 100%" max-height="800" @selection-change="onSelectionChange"> -->
    <el-table
      v-loading="listLoading"
      :data="queryForm.tableData"
      element-loading-text="Loading"
      fit
      highlight-current-row
    >
    <el-table-column fixed="left" class-name="status-col" label="名称" width="120">
      <template slot-scope="scope">
        <el-tag type="danger" effect="dark">{{ scope.row.name }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="secucode" label="代码" width="120"></el-table-column>
    <el-table-column prop="freeRation" label="持有比例" width="120"></el-table-column>
    <el-table-column class-name="status-col" label="状态" width="110" align="center">
      <template slot-scope="scope">
        <el-tag type="warning" effect="dark">{{ scope.row.freeRation }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column fixed="right" label="操作" width="220">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="onOpenDetail(scope.$index, queryForm.tableData)"
          type="text"
          size="small">
          详情
        </el-button>
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="onConfirmFocus(scope.$index, queryForm.tableData)"
          type="text"
          size="small">
          {{ scope.row.focused }}
        </el-button>
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="onKlineChart(scope.$index, queryForm.tableData)"
          type="text"
          size="small">
          K线图
        </el-button>
      </template>
    </el-table-column>
  </el-table>
  <div class="gva-pagination">
      <el-pagination
        :current-page=pageInfo.pageNum
        :page-size=pageInfo.pageSize
        :page-sizes="[20, 30, 50, 100]"
        :total=pageInfo.total
        layout="total, sizes, prev, pager, next"
        @current-change="onCurrentChange"
        @size-change="onSizeChange"
      />
    </div>
  </div>
    <div>
        <el-dialog title="K线查询" :visible.sync="lineChartForm.lineChartVisible">
            <el-tabs v-model="lineChartForm.activeName" @tab-click="selectTabClick">
            <el-tab-pane label="日线" name="date">
                <el-image :src="lineChartForm.lineChartSrc"></el-image>
            </el-tab-pane>
            <el-tab-pane label="周线" name="contour">
                <el-image :src="lineChartForm.lineChartSrc"></el-image>
            </el-tab-pane>
            <el-tab-pane label="分时线" name="minute">
                <el-image :src="lineChartForm.lineChartSrc"></el-image>
            </el-tab-pane>
            </el-tabs>
        </el-dialog>
    </div>
    <div>
        <el-dialog title="股东详情" :visible.sync="detailForm.dialogVisible">
            <el-table
            v-loading="listLoading"
            :data="detailForm.tableData"
            element-loading-text="Loading"
            fit
            highlight-current-row
            >
                <el-table-column fixed="left" class-name="status-col" label="代码" width="120">
                <template slot-scope="scope">
                    <el-tag type="danger" effect="dark">{{ scope.row.name }}</el-tag>
                </template>
                </el-table-column>
                <el-table-column prop="secucode" label="代码" width="120"></el-table-column>
                <el-table-column prop="holderRank" label="名次" width="100"></el-table-column>
                <el-table-column prop="holderName" label="股东名称" width="120"></el-table-column>
                <el-table-column prop="holderType" label="股东性质" width="120"></el-table-column>
                <el-table-column prop="holdNum" label="持股数" width="120"></el-table-column>
                <el-table-column prop="holdnumRation" label="持股占比" width="120"></el-table-column>
                <el-table-column prop="holdNumChange" label="持股增减" width="120"></el-table-column>
                <el-table-column class-name="status-col" label="发布时间" width="110" align="center">
                <template slot-scope="scope">
                    <el-tag type="warning" effect="dark">{{ scope.row.endDate|dateFilter }}</el-tag>
                </template>
                </el-table-column>
            </el-table>
        </el-dialog>
    </div>
  </div>
</template>

<script>
import { parseTime } from '@/utils/index'
import { getGDSDLT,getGDSDLTDetail,confirmFocus } from '@/api/stock'

export default {
  filters: {
    dateFilter(time) {
        return parseTime(time)
    }
  },
  data() {
    return {
      listLoading: true,
      modifyDialogVisible: false,
      queryForm: {
        keyword: '社保',
        dateRange: '',
        startDate: 0,
        endDate: 0,
      },
      tableData: [],
      pageInfo: {
        pageNum: 1,
        pageSize: 10,
        total: 0,
      },
      lineChartForm:{
        activeName: 'date',
        lineChartVisible: false,
        lineChartSrc: '',
        secucode: '',
      },
      detailForm: {
          secucode: '',
          tableData: [],
          dialogVisible: false,
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
        var req = {
            keyword: this.queryForm.keyword,
            releaseStart: this.queryForm.startDate,
            releaseEnd: this.queryForm.endDate,
            limit: this.pageInfo.pageSize,
            offset: (this.pageInfo.pageNum-1)*this.pageInfo.pageSize,
        }
        getGDSDLT(req).then(response => {
            this.queryForm.tableData = response.data.items
            this.pageInfo.total = response.data.total
            this.listLoading = false
        })
    },
    onConfirmFocus(index, rows) {
      var data = rows[index]
      var req = {
        name : data.name,
        secucode: data.secucode,
        expectPrice: data.presentPrice,
      }

      confirmFocus(req).then(response=>{
        this.fetchData()
      })
    },
    onKlineChart(index, rows) {//日线图
      var data = rows[index]
      var secucode = data.secucode.split('.').join("").toLowerCase()
      this.lineChartForm.secucode = secucode
      this.lineChartForm.lineChartVisible = !this.lineChartForm.lineChartVisible
      this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/daily/n/'+secucode+'.gif'
    },
    cancelDialog(index, rows) {
      this.modifyDialogVisible = !this.modifyDialogVisible
    },
    confirmDialog(formName) {
      var form = this.modifyForm

      var req = {
        name : this.modifyForm.name,
        secucode: this.modifyForm.secucode, 
        priceDecrease: this.modifyForm.decrease
      }

      updateRecommend(req).then(response=>{
        this.modifyDialogVisible = !this.modifyDialogVisible
        this.fetchData()
      })
      
      this.modifyForm.decrease = 0
    },
    onQuerySubmit() {
      console.log('submit!', this.queryForm.state, this.queryForm.decrease);
      this.fetchData()
    },
    onCurrentChange(val) {
      this.pageInfo.pageNum = val
      this.fetchData()
    },
    onSizeChange(val) {
        this.pageInfo.pageNum = 1
        this.pageInfo.pageSize = val
        this.fetchData()
    },
    onSelectDate(val) {
        if (!!val) {
            this.queryForm.startDate = val[0]
            this.queryForm.endDate = val[1]
        } else {
            this.queryForm.endDate = 0
            this.queryForm.startDate = 0
        }
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
    onSelectionChange(values) {
      if (!values) {
        return 
      }
      for (var idx in values) {
        var val = values[idx]
        this.fundForm.secucodes.push(val.secucode)
      }
    },
    onLineChart() {
      this.fundForm.fundFlowDrawer = true
      var req = {
        secucodes: this.fundForm.secucodes,
      }
      getFundDetailList(req).then(response=>{
        this.fundForm.opinionData = response.data.items
        this.fundForm.legendData = response.data.legendData
        this.fundForm.secucodes = []

        this.drawLine("fundLineChart")
      })
    },
    onOpenDetail(index, rows) {
        var data = rows[index]
        var req = {
            secucode: data.secucode,
            limit: 20,
        }
        getGDSDLTDetail(req).then(response=>{
            this.detailForm.dialogVisible = true
            this.detailForm.tableData = response.data.items
        })


    }
  }
}
</script>
