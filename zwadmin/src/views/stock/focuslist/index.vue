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
      <el-form-item label="关注">
        <el-select v-model="queryForm.disabled" clearable placeholder="全部">
          <el-option
            v-for="item in queryForm.focusOpts"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="状态">
        <el-select v-model="queryForm.state" clearable placeholder="全部">
          <el-option
            v-for="item in queryForm.stateOpts"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="期望日期">
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
      <el-table-column class-name="status-col" label="状态" width="110" align="center">
      <template slot-scope="scope">
        <el-tag :type="scope.row.state | statusFilter" effect="dark">{{ scope.row.state }}</el-tag>
      </template>
    </el-table-column>
      <el-table-column label="价格差" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="danger" effect="light">{{ scope.row.diffPrice }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="期望价格" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" effect="light">{{ scope.row.expectPrice }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="当前价格" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.presentPrice }}
        </template>
      </el-table-column>
      <el-table-column label="期望日期" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="success" effect="light">{{ scope.row.expectDate|dateFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="最近备注" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.remark }}
        </template>
      </el-table-column>
      <el-table-column label="股东集中度" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.holdFocus }}
        </template>
      </el-table-column>
      <el-table-column prop="totalNumRatio" label="股东人数" width="110" align="center" sortable>
        <template slot-scope="scope">
          {{ scope.row.totalNumRatio }}
        </template>
      </el-table-column>
      <el-table-column label="流通市值(亿)" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.traded }}
        </template>
      </el-table-column>
      <el-table-column label="创建日期" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.createDate }}
        </template>
      </el-table-column>
      <el-table-column label="更新日期" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.updateDate }}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="220">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="onUpdateFocus(scope.$index, tableData)"
          type="text"
          size="small">
          修改
        </el-button>
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="onRenshuDetail(scope.$index, tableData)"
          type="text"
          size="small">
          股东
        </el-button>
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="onCancelFocus(scope.$index, tableData)"
          type="text"
          size="small">
          {{ scope.row.focused }}
        </el-button>
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="onKlineChart(scope.$index, tableData)"
          type="text"
          size="small">
          K线图
        </el-button>
      </template>
    </el-table-column>
    </el-table>
    <div class="gva-pagination">
      <el-pagination
        :current-page=paginationForm.pageNum
        :page-size=paginationForm.pageSize
        :page-sizes="[10, 30, 50, 100]"
        :total=paginationForm.total
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
    <div>
      <el-drawer
        title="修改关注信息"
        :visible.sync="updateForm.updateFocusDrawer"
        :with-header="true">
        <el-form :model="updateForm" ref="updateForm" label-width="100px" class="demo-ruleForm">
          <el-form-item
            label="预期买入"
            prop="expectPrice"
            :rules="[
              { required: true, message: '价格不能为空'},
              { type: 'number', message: '价格必须为数字值'}
            ]"
          >
            <!-- <el-input type="expectPrice" v-model.number="updateForm.expectPrice" autocomplete="off"></el-input> -->
            <el-input-number v-model="updateForm.expectPrice" :precision="2" :step="0.5" :max="10000"></el-input-number>
          </el-form-item>
          <el-form-item label="预期日期">
            <el-date-picker
              v-model="updateForm.expectDate"
              value-format="timestamp"
              type="date"
              placeholder="选择日期">
            </el-date-picker>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmitForm('updateForm')">提交</el-button>
            <el-button @click="onResetForm('updateForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </el-drawer>
    </div>
    <div>
      <el-dialog title="股东人数变化详情" :visible.sync="operateForm.dialogVisible">
      <el-table
      v-loading="listLoading"
      :data="operateForm.gudongHistory"
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
import { getFocusList, cancelFocus, updateFocus,gdrenshuDetail } from '@/api/stock'
import { parseTime } from '@/utils/index'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        "待定": 'success',
        "准备": 'info',
        "开始": 'warning',
        "进行中": 'danger',
      }
      return statusMap[status]
    },
    dateFilter(time) {
        return parseTime(time)
    }
  },
  data() {
    return {
      tableData: null,
      timer:null,
      listLoading: true,
      lineChartForm:{
        activeName: 'date',
        lineChartVisible: false,
        lineChartSrc: '',
        secucode: '',
      },
      queryForm: {
        name:'',
        secucode: '',
        disabled: 0,
        state: 0,
        dateRange:'',
        startDate: 0,
        endDate: 0,
        focusOpts: [{
          value: 0,
          label: '全部'
        }, {
          value: 1,
          label: '已关注'
        }, {
          value: 2,
          label: '取消关注'
        }],
        stateOpts: [{
          value: 0,
          label: '全部'
        }, {
          value: 1,
          label: '准备'
        }, {
          value: 2,
          label: '开始'
        }, {
          value: 3,
          label: '进行中'
        }],
      },
      updateForm: {
        secucode: '',
        expectPrice: 0,
        expectDate: Date.parse(new Date()),
        updateFocusDrawer: false,
      },
      paginationForm: {
        pageNum: 1,
        pageSize: 20,
        total: 0,
      },
      operateForm:{
        dialogVisible: false,
        gudongHistory: null,
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
        disabled: this.queryForm.disabled,
        expectDateStart: this.queryForm.startDate,
        expectDateEnd: this.queryForm.endDate,
        state: this.queryForm.state,
        limit: this.paginationForm.pageSize,
        offset: (this.paginationForm.pageNum-1)*this.paginationForm.pageSize,
      }

      getFocusList(req).then(response => {
        this.tableData = response.data.items
        this.listLoading = false
        this.paginationForm.total = response.data.total
      })
    },
    handleCurrentChange(val) {
      this.paginationForm.pageNum = val
      this.fetchData()
    },
    handleSizeChange(val) {
      this.paginationForm.pageNum = 0
      this.paginationForm.pageSize = val
      this.fetchData()
    },
    onCancelFocus(index, rows) {
      var data = rows[index]
      var req = {
        name: data.name,
        secucode: data.secucode
      }
      cancelFocus(req).then(response=>{
        this.tableData.splice(index, 1)
        this.$message({
          message: '取消成功',
          type: 'success'
        });
      })
    },
    onQuerySubmit(val) {
      this.fetchData()
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
    onUpdateFocus(index, rows) {
      var data = rows[index]
      this.updateForm.secucode = data.secucode
      // this.updateForm.expectDate = 
      this.updateForm.updateFocusDrawer = true
    },
    onResetForm(formName) {
      this.$refs[formName].resetFields();
    },
    onSubmitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (!valid) {
          this.$message({
            message: '提交失败',
            type: 'warning'
          });
        }

        var req = {
          secucode: this.updateForm.secucode,
          expectDate: this.updateForm.expectDate,
          expectPrice: this.updateForm.expectPrice,
        }
        updateFocus(req).then(response=>{
          if (response.code == 20000) {
            this.$message({
            message: '编辑成功',
            type: 'success'
          });
          }

          this.fetchData()
        })
      });
    },
    onRenshuDetail(index, rows) {
        this.listLoading = true
        var data = rows[index]
        this.operateForm.dialogVisible = true
        var req = {
            secucode: data.secucode
        }
        gdrenshuDetail(req).then(response => {
            this.listLoading = false
            this.operateForm.gudongHistory = response.data.items
      })
    },
    onSelectDate(val) {
      if (!!val) {
        this.queryForm.startDate = val[0]
        this.queryForm.endDate = val[1]
      } else {
        this.queryForm.startDate = 0
      this.queryForm.endDate = 0
      }
      
      console.log('===>>TODO 332: ', this.queryForm.startDate, this.queryForm.endDate)
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
