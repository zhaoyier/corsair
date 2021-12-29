<template>
  <div class="dashboard-container">
    <span>系统提示买入列表</span>
    <el-divider></el-divider>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column label="代码" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.secucode }}</span>
        </template>
      </el-table-column>
      <el-table-column label="名称" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="最低价" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.minPrice }}
        </template>
      </el-table-column>
      <el-table-column label="当前价" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.presentPrice }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="涨跌价" width="110" align="center">
        <template slot-scope="scope">          
          <p v-if="scope.row.priceDiff>0">
            <el-tag type="danger">{{ scope.row.priceDiff }}</el-tag>
          </p>
          <p v-else>
            <el-tag type="gray">{{ scope.row.priceDiff }}</el-tag>
          </p>
        </template>
      </el-table-column>
      <el-table-column label="买入日期" width="210" align="center">
        <template slot-scope="scope">
          {{ scope.row.createDate }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { promptBuyList } from '@/api/stock'

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
      promptForm: {
        pageNum: 1,
        pageSize: 20,
        total: 0,
      }
    }
  },
  mounted(){
    this.fetchData();
    this.timer = setInterval(() => {
        setTimeout(this.fetchData, 0)
    }, 1000*600)
  },
  // created() {
  //   this.fetchData()
  // },
  methods: {
    fetchData() {
      this.listLoading = true
      var req = {
        limit: this.promptForm.pageSize,
        offset: (this.promptForm.pageNum-1)*this.promptForm.pageSize,
      }
      promptBuyList(req).then(response => {
        this.list = response.data.items
        console.log("==>>TODO 111: ", this.list)
        this.listLoading = false
      })
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
