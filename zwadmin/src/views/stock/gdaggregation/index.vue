<template>
  <div class="dashboard-container">
    <el-row type="flex">
        <el-col :span="18">
        <el-form :inline="true" :model="queryForm" class="demo-form-inline">
        <el-form-item label="名称">
            <el-input v-model="queryForm.name" placeholder=""></el-input>
        </el-form-item>
        <el-form-item label="代码">
            <el-input v-model="queryForm.secucode" placeholder="000001"></el-input>
        </el-form-item>
        <el-form-item label="股东增减">
            <el-input-number v-model="queryForm.totalNumRatio" :step="2"  :min="-100" :max="1000"  step-strictly></el-input-number>
        </el-form-item>
        <el-form-item label="价格增减">
            <el-input-number v-model="queryForm.priceRatio" :step="2"  :min="-100" :max="1000"  step-strictly></el-input-number>
        </el-form-item>
        <el-form-item label="十大股东增减">
            <el-input-number v-model="queryForm.holdRatio" :step="2"  :min="-100" :max="1000"  step-strictly></el-input-number>
        </el-form-item>
        <el-form-item label="发布日期">
            <el-date-picker
            @input="onSelectQueryDate"
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
        </el-col>
        <el-col :offset="1" :span="5" align="middle">
        <div>
            <el-button type="primary" @click="onResetSubmit">重建数据</el-button>
        </div>
        </el-col>
    </el-row>
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
      <el-table-column label="股东增减率" width="110" align="center" sortable>
        <template slot-scope="scope">
          <el-tag type="warning" effect="dark">{{ scope.row.holderRatio }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="价格变化率(%)" width="140" align="center" sortable>
        <template slot-scope="scope">
          {{ scope.row.priceRatio }}
        </template>
      </el-table-column>
      <el-table-column label="当前价" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.priceMin }}
        </template>
      </el-table-column>
      <el-table-column label="最高价" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.priceMax }}
        </template>
      </el-table-column>
      <el-table-column label="最近发布日期" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="success" effect="light">{{ scope.row.endDate|dateFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="股东人数" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.holderNum }}
        </template>
      </el-table-column>
      <el-table-column label="筹码集中度" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.holdFocus }}
        </template>
      </el-table-column>
      <el-table-column label="TOP10股东" width="110" align="center">
        <template slot-scope="scope">
            <el-tag type="danger" effect="light">{{ scope.row.holdRatio }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="TOP10流通" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" effect="light">{{ scope.row.freeholdRatio }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="更新时间" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.updateDate|dateFilter  }}
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
  <div>
      <!-- focusDrawer -->
      <el-drawer
        title="关注信息"
        :visible.sync="focusForm.focusDrawer"
        :with-header="true">
        <el-form :model="focusForm" ref="focusForm" label-width="100px" class="demo-ruleForm">
          <el-form-item label="名称">
            <el-input
                placeholder="请输入内容"
                v-model="focusForm.name"
                clearable>
            </el-input>
          </el-form-item>
          <el-form-item label="代码">
            <el-input
                placeholder="请输入内容"
                v-model="focusForm.secucode"
                clearable>
            </el-input>
          </el-form-item>
          <el-form-item label="预期价格">
            <el-input-number v-model="focusForm.expectPrice" :step="0.5"  :min="0" :max="1000"  step-strictly></el-input-number>
          </el-form-item>
          <el-form-item label="备注信息">
            <el-input placeholder="请输入备注信息" v-model="focusForm.remark" clearable> </el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmitFocusForm('focusForm')">提交</el-button>
            <el-button @click="onResetFocusForm('focusForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </el-drawer>
  </div>
  <div>
      <el-drawer
        title="重建数据"
        :visible.sync="updateForm.resetGDAggregationDrawer"
        :with-header="true">
        <el-form :model="updateForm" ref="updateForm" label-width="100px" class="demo-ruleForm">
          <el-form-item label="代码">
            <el-input
                placeholder="请输入内容"
                v-model="updateForm.secucode"
                clearable>
            </el-input>
          </el-form-item>
          <el-form-item label="预期时间">
            <el-date-picker
                @input="onSelectResetDate"
                v-model="updateForm.releaseDate"
                value-format="timestamp"
                type="daterange"
                range-separator="至"
                start-placeholder="开始日期"
                end-placeholder="结束日期">
            </el-date-picker>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmitForm('updateForm')">提交</el-button>
            <el-button @click="onResetForm('updateForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </el-drawer>
    </div>
  </div>
</template>

<script>
import { gdaggregationReset, gdaggregationList,gdrenshuDetail,confirmFocus,cancelFocus } from '@/api/stock'
import { parseTime } from '@/utils/index'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
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
      detailData: null,
      listLoading: true,
      lineChartForm:{
        activeName: 'date',
        lineChartVisible: false,
        lineChartSrc: '',
        secucode: '',
      },
      updateForm: {
          releaseDate: '',
          releaseStart:0,
          releaseEnd:0,
          secucode: '',
          resetGDAggregation: false,
          resetGDAggregationDrawer: false,
      },
      detailForm: {
          dialogVisible: false
      },
      focusForm:{
          name: '',
          secucode: '',
          focusDrawer: false,
          remark: '',
          expectPrice: 0,
      },
      queryForm: {
        pageNum: 1,
        pageSize: 20,
        total: 0,
        name:'',
        secucode: '',
        dateRange: '',
        releaseStart: 0,
        releaseEnd: 0,
        totalNumRatio:0, 
        priceRatio: 0,
        holdRatio: 0,
        holderRatio:0,
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
        totalNumRatio: this.queryForm.totalNumRatio,
        releaseStart: this.queryForm.releaseStart,
        releaseEnd: this.queryForm.releaseEnd,
        priceRatio: this.queryForm.priceRatio,
        holdRatio: this.queryForm.holdRatio,
        limit: this.queryForm.pageSize,
        offset: (this.queryForm.pageNum-1)*this.queryForm.pageSize,
      }
      gdaggregationList(req).then(response => {
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
        this.queryForm.pageSize = val

      this.fetchData()
    },
    onConfirmFocus(index, rows) {
        var data = rows[index]
        
        if (data.focused==="取消关注") {
            this.focusForm.name = ''
            this.focusForm.secucode = ''
            this.focusForm.focusDrawer = false
            var req = {
                name: data.name,
                secucode: data.secucode
            }
            cancelFocus(req).then(response=>{
                this.fetchData()
            })
        } else {
            this.focusForm.name = data.name
            this.focusForm.secucode = data.secucode
            this.focusForm.focusDrawer = true
        }


        console.log("==>>TODO fucus 0111:", data.focused)
    },
    onResetSubmit() {
        console.log("==>>TODO 351:")
        this.updateForm.resetGDAggregationDrawer = true
    },
    onSelectQueryDate(val) {
        this.updateForm.releaseStart = val[0]
        this.updateForm.releaseEnd = val[1]
    },
    onSelectResetDate(val) {
      console.log("==>>TODO date is: ", val[0], val[1])
      this.updateForm.releaseStart = val[0]
      this.updateForm.releaseEnd = val[1]
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
    onKlineChart(index, rows) {//日线图
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
    onSubmitForm(formName) {
        var req = {
            secucode: this.updateForm.secucode,
            periodEnd: this.updateForm.releaseEnd,
            periodStart: this.updateForm.releaseStart,
        }

        gdaggregationReset(req).then(response=>{

        })
    },
    onResetForm(formName) {
        
    },
    onSubmitFocusForm(formName) {
        var req = {
            name : this.focusForm.name,
            secucode: this.focusForm.secucode,
            expectPrice: this.focusForm.expectPrice,
            remark: this.focusForm.remark,
        }
        console.log("==>>TODO 3035: ", req)
        confirmFocus(req).then(response=>{
            console.log("==>>TODO fucus 02:", response)
            this.focusForm.focusDrawer = false
            this.fetchData()
        })
    
    },
    onResetFocusForm(formName) {
        console.log("==>>TODO 3036: ", this.$refs[formName])
        this.focusForm.name = ''
        this.focusForm.secucode = ''
        this.focusForm.presentPrice = ''
    }
  }
}

</script>

<style lang="scss" scoped>
.el-col {
    border-radius: 4px;
  }
.app {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}
</style>
