<template>
  <div class="dashboard-container">
    <div class="app-container">
    <el-form :inline="true" :model="queryForm" class="demo-form-inline">
      <el-form-item label="名称">
        <el-input v-model="queryForm.name" placeholder="茅台股份"></el-input>
      </el-form-item>
      <el-form-item label="代码">
        <el-input v-model="queryForm.secucode" placeholder="000001"></el-input>
      </el-form-item>
      <el-form-item label="跌幅">
        <el-input-number v-model="queryForm.decrease" :step="2"  :min="2" :max="20"  step-strictly></el-input-number>
      </el-form-item>
      <el-form-item label="收盘价">
        <el-input-number v-model="queryForm.closing" :precision="2" :step="0.5" :max="10000"></el-input-number>
      </el-form-item>
      <el-form-item label="市值">
        <el-input-number v-model="queryForm.market" :precision="2" :step="5" :max="10000"></el-input-number>
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
    <el-divider></el-divider>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      fit
      highlight-current-row
    >
      <el-table-column fixed="left" label="名称" width="110" align="center">
        <template slot-scope="scope">
          <!-- <span>{{ scope.row.name }}</span> -->
          <el-tag type="danger" effect="dark">{{ scope.row.name }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="名称" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.secucode }}</span>
        </template>
      </el-table-column>
      <el-table-column label="涨幅" width="110" align="center">
        <template slot-scope="scope">
          <!-- {{ scope.row.prise }} -->
          <el-tag type="warning" effect="dark">{{ scope.row.prise }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="量比" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.liangbi }}
        </template>
      </el-table-column>
      <el-table-column label="开盘价" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.opening }}
        </template>
      </el-table-column>
      <el-table-column label="收盘价" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.closing }}
        </template>
      </el-table-column>
      <el-table-column label="最高价" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.maxPrice }}
        </template>
      </el-table-column>
      <el-table-column label="最低价" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.minPrice }}
        </template>
      </el-table-column>
      <el-table-column label="市值(亿)" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.market }}
        </template>
      </el-table-column>
      <el-table-column label="流通市值(亿)" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.traded }}
        </template>
      </el-table-column>
      <el-table-column label="创建日期" width="210" align="center">
        <template slot-scope="scope">
          {{ scope.row.createDate }}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="220">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="onConfirmFocus(scope.$index, tableData)"
          type="text"
          size="small">
          {{ scope.row.focused }}
        </el-button>
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="updateDaily(scope.$index, tableData)"
          type="text"
          size="small">
          编辑
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
    <div class="gva-pagination">
      <el-pagination
        :current-page=queryForm.pageNum
        :page-size=queryForm.pageSize
        :page-sizes="[10, 30, 50, 100]"
        :total=queryForm.total
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
    <div>
      <el-dialog title="K线查询" :visible.sync="lineChartForm.lineChartVisible">
        <el-tabs v-model="lineChartForm.activeName" @tab-click="onSelectTabClick">
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
  </div>
</template>

<script>
import { getDailyList, confirmFocus } from '@/api/stock'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      timer:null,
      listLoading: true,
      lineChartForm:{
        activeName: 'date',
        lineChartVisible: false,
        lineChartSrc: '',
        secucode: '',
      },
      queryForm: {
        pageNum: 1,
        pageSize: 20,
        total: 0,
        lineChartVisible: false,
        lineChartSrc: '',
        name:'',
        secucode: '',
        dateRange: '',
        startDate: 0,
        endDate: 0,
        closing: 0,
        market: 0,
      }
    }
  },
  mounted(){
    this.fetchData();
    // this.timer = setInterval(() => {
    //     setTimeout(this.fetchData, 0)
    // }, 1000*600)
  },

  methods: {
    fetchData() {
      this.listLoading = true
      var req = {
        name: this.queryForm.name,
        secucode: this.queryForm.secucode,
        decrease: this.queryForm.decrease,
        startDate: this.queryForm.startDate,
        endDate: this.queryForm.endDate,
        closing: this.queryForm.closing,
        limit: this.queryForm.pageSize,
        market: this.queryForm.market,
        offset: (this.queryForm.pageNum-1)*this.queryForm.pageSize,
      }
      getDailyList(req).then(response => {
        this.list = response.data.items
        this.listLoading = false
        this.queryForm.total = response.data.total
      })
    },
    handleCurrentChange(val) {
      this.queryForm.pageNum = val
      this.fetchData()
    },
    handleSizeChange(val) {
      this.queryForm.pageSize = val
      this.fetchData()
    },
    updateDaily(index, rows) {
      var data = rows[index]
      var secucode = data.secucode.toLowerCase()
      console.log("==>>TODO secucode is 02: ", secucode)
    },
    onSelectDate(val) {
      this.queryForm.startDate = val[0]
      this.queryForm.endDate = val[1]
    },
    onQuerySubmit(val) {
      this.fetchData()
    },
    onConfirmFocus(index, rows) {
      var data = rows[index]
      var req = {
        name : data.name,
        secucode: data.secucode,
        presentPrice: data.closing,
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
    onSelectTabClick(tab) {
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
  }
}

</script>

<style lang="scss" scoped>
.dashboard {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
