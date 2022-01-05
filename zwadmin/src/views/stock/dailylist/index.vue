<template>
  <div class="dashboard-container">
    <div class="app-container">
    <el-form :inline="true" :model="dailyForm" class="demo-form-inline">
      <el-form-item label="名称">
        <el-input v-model="dailyForm.name" placeholder="茅台股份"></el-input>
      </el-form-item>
      <el-form-item label="代码">
        <el-input v-model="dailyForm.secucode" placeholder="000001"></el-input>
      </el-form-item>
      <el-form-item label="跌幅">
        <el-input-number v-model="dailyForm.decrease" :step="2"  :min="2" :max="20"  step-strictly></el-input-number>
      </el-form-item>
      <el-form-item label="日期">
        <el-date-picker
          @input="selectDate"
          v-model="dailyForm.dateRange"
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
        <!-- <el-button
          @click.native.prevent="modifyRow(scope.$index, tableData)"
          type="text"
          size="small">
          修改
        </el-button> -->
        <el-divider direction="vertical"></el-divider>
        <el-button
          @click.native.prevent="updateDaily(scope.$index, tableData)"
          type="text"
          size="small">
          编辑
        </el-button>
      </template>
    </el-table-column>
    </el-table>
    <div class="gva-pagination">
      <el-pagination
        :current-page=dailyForm.pageNum
        :page-size=dailyForm.pageSize
        :page-sizes="[10, 30, 50, 100]"
        :total=dailyForm.total
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
  </div>
</template>

<script>
import { getDailyList } from '@/api/stock'

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
      dailyForm: {
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
        name: this.dailyForm.name,
        secucode: this.dailyForm.secucode,
        decrease: this.dailyForm.decrease,
        startDate: this.dailyForm.startDate,
        endDate: this.dailyForm.endDate,
        limit: this.dailyForm.pageSize,
        offset: (this.dailyForm.pageNum-1)*this.dailyForm.pageSize,
      }
      getDailyList(req).then(response => {
        this.list = response.data.items
        this.listLoading = false
        this.dailyForm.total = response.data.total
        this.dailyForm.startDate = 0
        this.dailyForm.endDate = 0
      })
    },
    handleCurrentChange(val) {
      this.dailyForm.pageNum = val

      this.fetchData()
    },
    handleSizeChange() {

    },
    updateDaily(index, rows) {
      var data = rows[index]
      var secucode = data.secucode.toLowerCase()
      console.log("==>>TODO secucode is 02: ", secucode)
    },
    selectDate(val) {
      console.log("==>>TODO date is: ", val[0], val[1])
      this.dailyForm.startDate = val[0]
      this.dailyForm.endDate = val[1]
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
