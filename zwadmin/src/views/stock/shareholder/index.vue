<template>
  <div>
    <div class="app-container">
      <el-form :inline="true" :model="queryForm" class="demo-form-inline">
        <el-form-item label="代码">
          <el-input v-model="queryForm.secucode" placeholder="SZ.000001"></el-input>
        </el-form-item>
        <el-form-item label="股东增减">
          <el-input v-model.number="queryForm.reduceRatio" placeholder=20></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onQuerySubmit">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="app-container">
      <el-table :data="tableData" stripe style="width: 100%" max-height="800">
      <el-table-column fixed="left" class-name="status-col" label="代码" width="150">
        <template slot-scope="scope">
          <el-tag type="danger" effect="plain">{{ scope.row.secucode }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="name" label="名称" width="120"></el-table-column>
      <el-table-column class-name="status-col" label="推荐指数" width="110" align="center">
          <template slot-scope="scope">
            <el-tag type="success" effect="dark">{{ scope.row.valueIndex }}</el-tag>
          </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="股东变化" width="110" align="center">
          <template slot-scope="scope">
            <el-tag type="warning" effect="dark">{{ scope.row.gdRatio }}</el-tag>
          </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="当前价格" width="80">
        <template slot-scope="scope">
          <el-tag type="danger" effect="dark">{{ scope.row.presentPrice }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="ratioStr" label="股东比例" width="160"></el-table-column>
      <el-table-column prop="price" label="价格变动" width="160"></el-table-column>
      <el-table-column prop="date" label="日期" width="160"></el-table-column>
      <el-table-column prop="focus" label="关注度" width="200"></el-table-column>
      <el-table-column prop="updateDate" label="最近更新时间" width="120"> </el-table-column>

      <el-table-column fixed="right" label="操作" width="220">
        <template slot-scope="scope">
          <el-button
            @click.native.prevent="modifyRow(scope.$index, tableData)"
            type="text"
            size="small">
            修改
          </el-button>
          <el-divider direction="vertical"></el-divider>
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
  </div>
</template>

<script>
import { getLongLineList } from '@/api/stock'

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
      queryForm: {
        secucode: '',
        reduceRatio: 0,
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
        gdRatio: this.queryForm.reduceRatio,
        secucode: this.queryForm.secucode,
      }
      getLongLineList(req).then(response => {
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

    }
  }
}
</script>
