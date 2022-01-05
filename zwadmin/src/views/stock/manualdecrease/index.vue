<template>
  <div class="dashboard-container">
    <div class="app-container">
    <el-form :inline="true" :model="manualForm" class="demo-form-inline">
      <el-form-item label="名称">
        <el-input v-model="manualForm.name" placeholder="茅台股份"></el-input>
      </el-form-item>
      <el-form-item label="代码">
        <el-input v-model="manualForm.secucode" placeholder="000001"></el-input>
      </el-form-item>
      <el-form-item label="跌幅">
        <el-input-number v-model="manualForm.decreaseTag" :step="2"  :min="2" :max="20"  step-strictly></el-input-number>
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
      <el-table-column label="调整幅度" width="110" align="center">
        <template slot-scope="scope">
          <el-tag type="warning" effect="dark">{{ scope.row.decreaseTag }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.createDate }}
        </template>
      </el-table-column>
      <el-table-column label="更新时间" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.updateDate }}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="220">
      <template slot-scope="scope">
        <!-- <el-button
          @click.native.prevent="modifyRow(scope.$index, tableData)"
          type="text"
          size="small">
          修改
        </el-button> -->
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="updateManualTag(scope.$index, tableData)"
          type="text"
          size="small">
          编辑
        </el-button>
      </template>
    </el-table-column>
    </el-table>
    <div class="gva-pagination">
      <el-pagination
        :current-page=manualForm.pageNum
        :page-size=manualForm.pageSize
        :page-sizes="[10, 30, 50, 100]"
        :total=manualForm.total
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
  </div>
</template>

<script>
import { manualDecreaseList } from '@/api/stock'

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
      manualForm: {
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
        decreaseTag: 38,
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
        name: this.manualForm.name,
        secucode: this.manualForm.secucode,
        decreaseTag: this.manualForm.decreaseTag,
        limit: this.manualForm.pageSize,
        offset: (this.manualForm.pageNum-1)*this.manualForm.pageSize,
      }
      manualDecreaseList(req).then(response => {
        this.list = response.data.items
        this.listLoading = false
        this.manualForm.total = response.data.total
        this.manualForm.name = ''
        this.manualForm.secucode = ''
        this.manualForm.decreaseTag = 0
        this.manualForm.pageNum = 1
      })
    },
    handleCurrentChange() {

    },
    handleSizeChange() {

    },
    updateManualTag(index, rows) {
      var data = rows[index]
      var secucode = data.secucode.toLowerCase()
      console.log("==>>TODO secucode is 01: ", secucode)
    },
    selectDate(val) {
      console.log("==>>TODO date is: ", val[0], val[1])
      this.manualForm.startDate = val[0]
      this.manualForm.endDate = val[1]
    },
    onQuerySubmit(val) {
      console.log("==>>TODO query is: ", val)
      this.fetchData()
    }


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
