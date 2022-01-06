<template>
  <div class="dashboard-container">
    <div class="app-container">
        <el-form :inline="true" :model="queryForm" class="demo-form-inline" style="float:left">
        <el-form-item label="名称">
            <el-input v-model="queryForm.name" placeholder="茅台股份"></el-input>
        </el-form-item>
        <el-form-item label="代码">
            <el-input v-model="queryForm.secucode" placeholder="SZ.000001"></el-input>
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
            v-for="item in queryForm.options"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="onQuerySubmit">查询</el-button>
        </el-form-item>
        </el-form>
        <el-button style="float:right" type="primary" @click="onCreateSubmit">新建周期</el-button>
    </div>
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
      <el-table-column label="状态" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.state | statusFilter" effect="dark">{{ scope.row.state }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="当前价格" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.presentPrice }}
        </template>
      </el-table-column>
      <el-table-column label="期望最低价" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="danger" effect="light">{{ scope.row.expectMin }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="期望最高价" width="110" align="center">
        <template slot-scope="scope">
          <el-tag effect="light">{{ scope.row.expectMax }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="期望开始时间" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" effect="light">{{ scope.row.expectStart|dateFilter }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="期望结束时间" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" effect="light">{{ scope.row.expectEnd|dateFilter }}</el-tag>
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
          @click.native.prevent="onUpdateZhouQi(scope.$index, tableData)"
          type="text"
          size="small">
          修改
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
      <!-- <el-dialog title="K线查询" :visible.sync="lineChartForm.lineChartVisible">
        <el-tabs @tab-click="onSelectTabClick">
          <el-tab-pane label="用户管理" name="first">用户管理</el-tab-pane>
          <el-tab-pane label="配置管理" name="second">配置管理</el-tab-pane>
          <el-tab-pane label="角色管理" name="third">角色管理</el-tab-pane>
          <el-tab-pane label="定时任务补偿" name="fourth">定时任务补偿</el-tab-pane>
        </el-tabs>
      </el-dialog> -->

    </div>
    <div>
      <el-drawer
        title="更新周期信息"
        :visible.sync="updateForm.updateZhouQiDrawer"
        :with-header="true">
        <el-form :model="updateForm" ref="updateForm" label-width="100px" class="demo-ruleForm">
          <el-form-item
            label="代码"
            prop="secucode"
            :rules="[
              { required: true, message: '代码不能为空'},
              { type: 'string', message: '代码必须为字符串'}
            ]"
          >
            <el-input type="secucode" v-model.number="updateForm.secucode" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item
            label="预期最低价"
            prop="expectMin"
            :rules="[
              { required: true, message: '价格不能为空'},
              { type: 'number', message: '价格必须为数字值'}
            ]"
          >
            <el-input-number type="expectMin" v-model="updateForm.expectMin" :precision="2" :step="0.5" :max="10"></el-input-number>
          </el-form-item>
          <el-form-item
            label="预期最高价"
            prop="expectMax"
            :rules="[
              { required: true, message: '价格不能为空'},
              { type: 'number', message: '价格必须为数字值'}
            ]"
          >
            <el-input-number type="expectMax" v-model="updateForm.expectMax" :precision="2" :step="0.5" :max="10"></el-input-number>
          </el-form-item>
          <el-form-item label="取消关注">
            <el-switch
             v-model="updateForm.disabled"
             active-color="#13ce66"
             :active-value=1
             :inactive-value=0>
           </el-switch>
        </el-form-item>
          <el-form-item label="预期时间">
            <el-date-picker
                @input="onSelectDate"
                v-model="updateForm.expectDate"
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
import { gpzhouQiList, updateGPZhouQi } from '@/api/stock'
import { parseTime } from '@/utils/index'


export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        "已达时间": 'danger',
        "已达价格": 'warning',
        "待定": 'success'
      }
      return statusMap[status]
    },
    dateFilter(time) {
        return parseTime(time)
    }
  },
  data() {
    return {
      activeName: 'second',
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
        state:0,
        disabled: 0,
        options: [{
          value: 0,
          label: '全部'
        }, {
          value: 1,
          label: '已达时间'
        }, {
          value: 2,
          label: '已达价格'
        }],
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
      },
      updateForm: {
        secucode: '',
        expectMin: 0,
        expectMax: 0,
        expectStart: 0,
        expectEnd: 0,
        expectDate: [],
        disabled: 0,
        updateZhouQiDrawer: false,
      },
      paginationForm: {
        pageNum: 1,
        pageSize: 20,
        total: 0,
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
        state: this.queryForm.state,
        disabled: this.queryForm.disabled,
        limit: this.paginationForm.pageSize,
        offset: (this.paginationForm.pageNum-1)*this.paginationForm.pageSize,
      }

      gpzhouQiList(req).then(response => {
        this.tableData = response.data.items
        this.listLoading = false
        this.paginationForm.total = response.data.total
      })
    },
    handleCurrentChange(val) {
      this.paginationForm.pageNum = val
    },
    handleSizeChange(val) {
      this.paginationForm.pageSize = val
    },
    onQuerySubmit(val) {
      this.fetchData()
    },
    onKlineChart(index, rows) {//日线图
      var data = rows[index]
      var secucode = data.secucode.split('.').join("").toLowerCase()
      this.lineChartForm.secucode = secucode
      this.lineChartForm.activeName = 'date'
      this.lineChartForm.lineChartVisible = !this.lineChartForm.lineChartVisible
      this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/daily/n/'+secucode+'.gif'
    },
    onSelectTabClick(tab, event) {
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
    onUpdateZhouQi(index, rows) {
      var data = rows[index]
      this.updateForm.secucode = data.secucode
      this.updateForm.expectMin = data.expectMin
      this.updateForm.expectMax = data.expectMax
      this.updateForm.disabled = data.disabled?1:0
      this.updateForm.updateZhouQiDrawer = true
      this.updateForm.expectStart = data.expectStart*1000
      this.updateForm.expectEnd = data.expectEnd*1000

      this.updateForm.expectDate =  [new Date(data.expectStart*1000), new Date(data.expectEnd*1000)]
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
          expectMin: this.updateForm.expectMin,
          expectMax: this.updateForm.expectMax,
          expectStart: this.updateForm.expectStart,
          expectEnd: this.updateForm.expectEnd,
          disabled: this.updateForm.disabled==1?true:false,
        }
        updateGPZhouQi(req).then(response=>{
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
    onCreateSubmit() {
        this.updateForm.updateZhouQiDrawer = true
    },
    onSelectDate(val) {
      this.updateForm.expectStart = val[0]
      this.updateForm.expectEnd = val[1]
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
