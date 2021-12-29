<template>
  <div>
  <div class="app-container">
    <el-form :inline="true" :model="queryForm" class="demo-form-inline">
      <el-form-item label="代码">
        <el-input v-model="queryForm.secucode" placeholder="SZ.000001"></el-input>
      </el-form-item>
      <el-form-item label="跌幅">
        <el-input v-model.number="queryForm.decrease" placeholder=20></el-input>
      </el-form-item>
      <el-form-item label="状态">
        <el-select v-model.number=queryForm.state placeholder="准备">
          <el-option label="准备" value=1></el-option>
          <el-option label="开始" value=2></el-option>
          <el-option label="进行中" value=3></el-option>
          <el-option label="结束" value=4></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onQuerySubmit">查询</el-button>
      </el-form-item>
    </el-form>
  </div>
  <div class="app-container">
    <el-table :data="tableData" stripe style="width: 100%" max-height="800">
    <el-table-column fixed prop="secucode" label="代码" width="150"></el-table-column>
    <el-table-column prop="name" label="名称" width="120"></el-table-column>
    <el-table-column prop="rMIndex" label="推荐指数" width="120"></el-table-column>
    <el-table-column class-name="status-col" label="Status" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.state | statusFilter" effect="dark">{{ scope.row.state }}</el-tag>
        </template>
      </el-table-column>
    <el-table-column prop="pDecrease" label="最近跌幅" width="120"></el-table-column>
    <el-table-column prop="maxPrice" label="最高价" width="120"></el-table-column>
    <el-table-column prop="rMPrice" label="推荐价格" width="220"> </el-table-column>
    <el-table-column prop="presentPrice" label="当前价" width="120"> </el-table-column>
    <el-table-column prop="gDDecrease" label="股东人数" width="120"> </el-table-column>
    <el-table-column prop="updateNum" label="更新次数" width="80"> </el-table-column>
    <el-table-column prop="updateNum" label="更新次数" width="80"> </el-table-column>
    <el-table-column prop="referDecrease" label="参考幅度" width="120"> </el-table-column>
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
          @click.native.prevent="KLineChart(scope.$index, tableData)"
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
        :page-sizes="[10, 30, 50, 100]"
        :total=pageInfo.total
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
  </div>
  <div>
    <el-dialog title="修改推荐" :visible.sync="modifyDialogVisible">
      <el-form :model="modifyForm" ref="modifyForm" label-width="100px" class="demo-ruleForm">
      <el-form-item
        label="跌幅修正"
        prop="decrease"
        :rules="[
          { type: 'number', message: '跌幅必须为数字值'}
        ]"
      >
        <el-input type="age" v-model.number="modifyForm.decrease" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="cancelDialog('modifyForm')">取消</el-button>
        <el-button type="primary" @click="confirmDialog('modifyForm')">提交</el-button>
      </el-form-item>
    </el-form>
    </el-dialog>
  </div>
  <div>
    <el-dialog title="K线查询" :visible.sync="lineChartForm.lineChartVisible">
      <!-- <el-image :src="lineChartForm.lineChartSrc"></el-image> -->
      <el-tabs v-model="lineChartForm.activeName" @tab-click="handleClick">
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
import { getRecommendList,updateRecommend } from '@/api/stock'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        "进行中": 'danger',
        "开始": 'warning',
        "准备": 'info'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      
      listLoading: true,
      modifyDialogVisible: false,
      queryForm: {
        secucode: '',
        region: 1,
        decrease: 0,
        state: 1,
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
        pageSize: 20,
        total: 0,
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
        limit: this.pageInfo.pageSize,
        offset: (this.pageInfo.pageNum-1)*this.pageInfo.pageSize,
        pDecrease:this.queryForm.decrease,
        state: this.queryForm.state,
        secucode: this.queryForm.secucode,
      }
      getRecommendList(req).then(response => {
        console.log("===>>TODO 111: ", response)

        this.tableData = response.data.items
        this.pageInfo.total = response.data.total
        this.listLoading = false
      })
      console.log("===>>TODO 211: ", this.tableData)
    },
    modifyRow(index, rows) {
      var data = rows[index]
      this.modifyDialogVisible = !this.modifyDialogVisible
      this.modifyForm.secucode = data.secucode
      this.modifyForm.priceDecrease = data.pDecrease
      this.modifyForm.name = data.name
    },
    KLineChart(index, rows) {//日线图
      console.log("==>>TODO 3141: ", rows[index])
      var data = rows[index]
      var secucode = data.secucode.split('.').join("").toLowerCase()
      this.lineChartForm.secucode = secucode
      console.log("==>>TODO 3142: ", secucode)
      this.lineChartForm.lineChartVisible = !this.lineChartForm.lineChartVisible
      this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/daily/n/'+secucode+'.gif'
    },
    cancelDialog(index, rows) {
      this.modifyDialogVisible = !this.modifyDialogVisible
    },
    confirmDialog(formName) {
      console.log("==>>TODO 3031: ", this.modifyForm)
      console.log("==>>TODO 3032: ", this.modifyForm.decrease)
      var form = this.modifyForm
      console.log("==>>TODO 3035: ", this.$refs[formName], form)

      var req = {
        name : this.modifyForm.name,
        secucode: this.modifyForm.secucode, 
        priceDecrease: this.modifyForm.decrease
      }

      updateRecommend(req).then(response=>{
        console.log("==>>TODO 3036: ", "ok")
        this.modifyDialogVisible = !this.modifyDialogVisible
        this.fetchData()
      })
      
      this.modifyForm.decrease = 0
    },
    onQuerySubmit() {
      console.log('submit!', this.queryForm.state, this.queryForm.decrease);
      this.fetchData()
      console.log("===>>TODO 212: ", this.tableData)
    },
    handleCurrentChange(val) {
      console.log("===>>TODO 2131: ", val)
      this.pageInfo.pageNum = val
      console.log("===>>TODO 2132: ", this.pageInfo.pageNum)
      var req = {
        limit: this.pageInfo.pageSize,
        offset: (this.pageInfo.pageNum-1)*this.pageInfo.pageSize,
      }
      this.fetchData()
    },
    handleSizeChange(val) {
      console.log("===>>TODO 214: ", val)
    },
    handleClick(tab) {
      console.log("===>>TODO 254: ", tab)
      var secucode = this.lineChartForm.secucode
      if (tab.name === "date") {
        this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/daily/n/'+secucode+'.gif'
      } else if (tab.name == "contour") { //contour
        this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/weekly/n/'+secucode+'.gif'
      } else { //分时线
        this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/min/n/'+secucode+'.gif'
      }
      this.lineChartForm.activeName = "date"
    }

  }
}
</script>
