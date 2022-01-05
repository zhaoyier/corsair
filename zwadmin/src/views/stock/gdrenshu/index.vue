<template>
  <div class="dashboard-container">
    <div class="app-container">
        <el-form :inline="true" :model="queryForm" class="demo-form-inline">
        <el-form-item label="名称">
            <el-input v-model="queryForm.name" placeholder=""></el-input>
        </el-form-item>
        <el-form-item label="代码">
            <el-input v-model="queryForm.secucode" placeholder="000001"></el-input>
        </el-form-item>
        <el-form-item label="增减率">
            <el-input-number v-model="queryForm.totalRatio" :step="2"  :min="-100" :max="100"  step-strictly></el-input-number>
        </el-form-item>
        <el-form-item label="排序">
            <el-select v-model="queryForm.sortTyp" clearable placeholder="全部" @change='onSelectValue'>
            <el-option
                v-for="item in queryForm.options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
            </el-option>
            </el-select>
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
      :data="tableData"
      element-loading-text="Loading"
      fit
      highlight-current-row
    >
      <el-table-column fixed="left" label="名称" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="danger" effect="dark">{{ scope.row.name }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="代码" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.secucode }}</span>
        </template>
      </el-table-column>
      <el-table-column label="持仓人数" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" effect="dark">{{ scope.row.holderTotalNum }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="totalNumRatio" label="持仓变化率" width="140" align="center" sortable>
        <template slot-scope="scope">
          {{ scope.row.totalNumRatio }}
        </template>
      </el-table-column>
      <el-table-column label="发布时间" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.releaseDate }}
        </template>
      </el-table-column>
      <el-table-column label="集中度" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.holdFocus }}
        </template>
      </el-table-column>
      <el-table-column label="当前价" width="110" align="center">
        <template slot-scope="scope">
            <el-tag type="danger" effect="light">{{ scope.row.presentPrice }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="收盘价" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" effect="light">{{ scope.row.price }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="非流通股东率" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.holdRatioTotal }}
        </template>
      </el-table-column>
      <el-table-column label="流通股东率" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.freeholdRatioTotal }}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="220">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="onRenshuDetail(scope.$index, tableData)"
          type="text"
          size="small">
          历史
        </el-button>
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="onConfirmFocus(scope.$index, tableData)"
          type="text"
          size="small">
          {{ scope.row.focused }}
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
    <el-dialog title="股东人数变化详情" :visible.sync="detailForm.dialogVisible">
      <el-table
      v-loading="listLoading"
      :data="detailData"
      element-loading-text="Loading"
      fit
      highlight-current-row
    >
      <el-table-column fixed="left" label="名称" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="danger" effect="dark">{{ scope.row.name }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="代码" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.secucode }}</span>
        </template>
      </el-table-column>
      <el-table-column label="持仓人数" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" effect="dark">{{ scope.row.holderTotalNum }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="持仓变化率" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.totalNumRatio }}
        </template>
      </el-table-column>
      <el-table-column label="发布时间" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.releaseDate }}
        </template>
      </el-table-column>
      <el-table-column label="集中度" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.holdFocus }}
        </template>
      </el-table-column>
      <el-table-column label="收盘价" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.price }}
        </template>
      </el-table-column>
      <el-table-column label="非流通股东率" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.holdRatioTotal }}
        </template>
      </el-table-column>
      <el-table-column label="流通股东率" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.freeholdRatioTotal }}
        </template>
      </el-table-column>
    </el-table>
    </el-dialog>
  </div>
  </div>
</template>

<script>
import { gdrenshuList, gdrenshuDetail,confirmFocus } from '@/api/stock'

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
      tableData: null,
      detailData: null,
      listLoading: true,
      lineChartForm:{
        activeName: 'date',
        lineChartVisible: false,
        lineChartSrc: '',
        secucode: '',
      },
      detailForm: {
          dialogVisible: false
      },
      queryForm: {
        pageNum: 1,
        pageSize: 20,
        total: 0,
        name:'',
        secucode: '',
        dateRange: '',
        startDate: 0,
        endDate: 0,
        totalRatio:0, 
        sortTyp: 0,
        options: [{
          value: 0,
          label: '全部'
        }, {
          value: 1,
          label: '减少'
        }, {
          value: 2,
          label: '增加'
        }],
      }
    }
  },
  mounted(){
    this.fetchData();
  },

  methods: {
    fetchData() {
        console.log("==>>TODO 111: ", this.queryForm.sortTyp)

      this.listLoading = true
      var req = {
        name: this.queryForm.name,
        secucode: this.queryForm.secucode,
        decrease: this.queryForm.decrease,
        releaseStart: this.queryForm.startDate,
        releaseEnd: this.queryForm.endDate,
        totalRatio: this.queryForm.totalRatio,
        limit: this.queryForm.pageSize,
        sortTyp: this.queryForm.sortTyp,
        offset: (this.queryForm.pageNum-1)*this.queryForm.pageSize,
      }
      gdrenshuList(req).then(response => {
        this.tableData = response.data.items
        this.listLoading = false
        this.queryForm.total = response.data.total
      })
    },
    handleCurrentChange(val) {
      this.queryForm.pageNum = val

      this.fetchData()
    },
    handleSizeChange() {

    },
    onSelectValue: function(){
        console.log('您选择了', this.queryForm.sortTyp)
    },
    onConfirmFocus(index, rows) {
      var data = rows[index]
      var req = {
        name : data.name,
        secucode: data.secucode,
        presentPrice: data.presentPrice,
      }

      console.log("==>>TODO fucus 01:", req)

      confirmFocus(req).then(response=>{
        console.log("==>>TODO fucus 02:", response)
        this.fetchData()
      })
    },
    onSelectDate(val) {
      console.log("==>>TODO date is: ", val[0], val[1])
      this.queryForm.startDate = val[0]
      this.queryForm.endDate = val[1]
    },
    onQuerySubmit(val) {
      console.log("==>>TODO query is: ", val)
      this.fetchData()
    },
    onRenshuDetail(index, rows) {
        this.listLoading = true

        var data = rows[index]
        this.detailForm.dialogVisible = true
        var req = {
            secucode: data.secucode
        }
        gdrenshuDetail(req).then(response => {
            this.listLoading = false
            this.detailData = response.data.items
      })
    },
    klineChart(index, rows) {//日线图
      var data = rows[index]
      var secucode = data.secucode.split('.').join("").toLowerCase()
      this.lineChartForm.secucode = secucode
      this.lineChartForm.lineChartVisible = !this.lineChartForm.lineChartVisible
      this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/daily/n/'+secucode+'.gif'
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
