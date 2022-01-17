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
      <el-form-item label="流入金额">
        <el-input-number v-model="queryForm.inflow" :step="10"  :min="-1000" :max="1000"  step-strictly></el-input-number>
      </el-form-item>
      <el-form-item label="增减率">
        <el-input-number v-model="queryForm.inflowRatio" :step="2"  :min="-100" :max="100"  step-strictly></el-input-number>
      </el-form-item>
      <el-form-item label="流入日期">
        <el-date-picker
          @input="onSelectDate"
          v-model="queryForm.fundDate"
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
      <el-table-column class-name="status-col" label="资金日期" width="110" align="center">
      <template slot-scope="scope">
        <el-tag type="success" effect="light">{{ scope.row.fundDate|dateFilter }}</el-tag>
      </template>
    </el-table-column>
      <el-table-column label="流入金额(百万)" width="160" align="center">
        <template slot-scope="scope">
          <el-tag type="danger" effect="light">{{ scope.row.inflow }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="流入占比" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" effect="light">{{ scope.row.inflowRatio }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="流通市值" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="danger" effect="light">{{ scope.row.traded }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="当前价格" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.presentPrice }}
        </template>
      </el-table-column>
      <el-table-column label="涨跌幅" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="success" effect="light">{{ scope.row.increaseRatio }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="创建日期" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.createDate|dateFilter }}
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
          @click.native.prevent="onConfirmFocus(scope.$index, tableData)"
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
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="onDetailData(scope.$index, tableData)"
          type="text"
          size="small">
          详情
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
          <el-form-item label="主营业务">
            <el-input type="textarea" v-model="updateForm.mainBusiness"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmitForm('updateForm')">提交</el-button>
            <el-button @click="onResetForm('updateForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </el-drawer>
    </div>
    <div>
        <el-dialog title="详情" :visible.sync="detailForm.detailVisible">
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
            <el-table-column class-name="status-col" label="日期" width="110" align="center">
            <template slot-scope="scope">
                <el-tag type="success" effect="light">{{ scope.row.fundDate|dateFilter }}</el-tag>
            </template>
            </el-table-column>
            <el-table-column label="流入金额(百万)" width="160" align="center">
                <template slot-scope="scope">
                <el-tag type="danger" effect="light">{{ scope.row.inflow }}</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="流入占比" width="110" align="center">
                <template slot-scope="scope">
                <el-tag type="warning" effect="light">{{ scope.row.inflowRatio }}</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="当前价格" width="110" align="center">
                <template slot-scope="scope">
                {{ scope.row.presentPrice }}
                </template>
            </el-table-column>
            <el-table-column label="涨跌幅" width="110" align="center">
                <template slot-scope="scope">
                <el-tag type="success" effect="light">{{ scope.row.increaseRatio }}</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="创建日期" width="110" align="center">
                <template slot-scope="scope">
                {{ scope.row.createDate|dateFilter }}
                </template>
            </el-table-column>
            </el-table>
        </el-dialog>
    </div>
  </div>
</template>

<script>
import { getFocusList, cancelFocus, updateFocus,gdrenshuDetail, getFundFlowList } from '@/api/stock'
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
      detailData: null,
      queryForm: {
        name:'',
        secucode: '',
        fundDate: '',
        fundStart: 0,
        fundEnd: 0,
        inflow: 0,
        inflowRatio: 0,
        limit: 0,
        offset: 0,
      },
      lineChartForm:{
        activeName: 'date',
        lineChartVisible: false,
        lineChartSrc: '',
        secucode: '',
      },
      updateForm: {
        secucode: '',
        expectPrice: 0,
        expectDate: Date.parse(new Date()),
        updateFocusDrawer: false,
        mainBusiness: '',
      },
      paginationForm: {
        pageNum: 1,
        pageSize: 20,
        total: 0,
      },
      detailForm: {
          detailVisible: false,
      }
      
    }
  },
  mounted(){
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.listLoading = true
      var req = {
        name: this.queryForm.name,
        secucode: this.queryForm.secucode,
        fundStart: this.queryForm.fundStart,
        fundEnd: this.queryForm.fundEnd,
        inflow: this.queryForm.inflow,
        inflowRatio: this.queryForm.inflowRatio,
        limit: this.paginationForm.pageSize,
        offset: (this.paginationForm.pageNum-1)*this.paginationForm.pageSize,
      }

      getFundFlowList(req).then(response => {
        this.listLoading = false
        this.tableData = response.data.items
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
    onQuerySubmit(val) {
      this.fetchData()
    },
    onSelectDate(val) {
      if (!!val) {
        this.queryForm.startDate = val[0]
        this.queryForm.endDate = val[1]
      } else {
        this.queryForm.startDate = 0
        this.queryForm.endDate = 0
      }
    },
    onUpdateFocus(index, rows) {
      var data = rows[index]
      this.updateForm.secucode = data.secucode
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
          mainBusiness: this.updateForm.mainBusiness,
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
    onKlineChart(index, rows) {//日线图
      var data = rows[index]
      var secucode = data.secucode.split('.').join("").toLowerCase()
      this.lineChartForm.secucode = secucode
      this.lineChartForm.lineChartVisible = !this.lineChartForm.lineChartVisible
      this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/daily/n/'+secucode+'.gif'
    },
    onDetailData(index, rows) {
        var data = rows[index]
        var req = {
            secucode: data.secucode,
            limit: 20,
            offset: 0,
        }
        
        this.listLoading = true
        this.detailForm.detailVisible = true

        getFundFlowList(req).then(response=>{
            this.listLoading = false
            this.detailData = response.data.items
        })
    },
    onConfirmFocus(index, rows) {
      var data = rows[index]
      var req = {
        name : data.name,
        secucode: data.secucode,
        expectPrice: data.closing,
      }

      confirmFocus(req).then(response=>{
        this.fetchData()
      })
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
