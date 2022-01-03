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
      <el-form-item label="状态">
        <el-select v-model="queryForm.disabled" clearable placeholder="全部">
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
      <el-table-column label="名称" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.secucode }}</span>
        </template>
      </el-table-column>
      <el-table-column label="关注价格" width="110" align="center">
        <template slot-scope="scope">
          <el-tag effect="light">{{ scope.row.focusPrice }}</el-tag>
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
      <el-table-column label="价格差" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="danger" effect="light">{{ scope.row.diffPrice }}</el-tag>
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
          @click.native.prevent="onCancelFocus(scope.$index, tableData)"
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
        <!-- <el-image :src="lineChartForm.lineChartSrc"></el-image> -->
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
      <el-drawer
        title="我是标题"
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
            <el-input type="expectPrice" v-model.number="updateForm.expectPrice" autocomplete="off"></el-input>
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
import { getFocusList, cancelFocus, updateFocus } from '@/api/stock'

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
        options: [{
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
        expectPrice: 0,
        updateFocusDrawer: false,
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
    },
    handleSizeChange(val) {
      this.paginationForm.pageSize = val
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
