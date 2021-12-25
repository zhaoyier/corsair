<template>
  <div>
  <div class="app-container">
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="审批人">
        <el-input v-model="formInline.user" placeholder="审批人"></el-input>
      </el-form-item>
      <el-form-item label="活动区域">
        <el-select v-model="formInline.region" placeholder="活动区域">
          <el-option label="区域一" value="shanghai"></el-option>
          <el-option label="区域二" value="beijing"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">查询</el-button>
      </el-form-item>
    </el-form>
  </div>
  <div class="app-container">
    <el-table :data="tableData" style="width: 100%" max-height="250">
    <el-table-column fixed prop="secucode" label="代码" width="150"></el-table-column>
    <el-table-column prop="name" label="名称" width="120"></el-table-column>
    <el-table-column prop="province" label="省份" width="120"></el-table-column>
    <el-table-column prop="city" label="市区" width="120"></el-table-column>
    <el-table-column prop="address" label="地址" width="300"></el-table-column>
    <el-table-column prop="zip" label="邮编" width="120"> </el-table-column>
    <el-table-column fixed="right" label="操作" width="120">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="deleteRow(scope.$index, tableData)"
          type="text"
          size="small">
          移除
        </el-button>
      </template>
    </el-table-column>
  </el-table>
  </div>
  </div>
</template>

<script>
import { getRecommendList } from '@/api/stock'

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
      listLoading: true,
      formInline: {
        user: '',
        region: ''
      },
      tableData: []
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getRecommendList().then(response => {
        console.log("===>>TODO 111: ", response)

        // this.tableData = response.rows
        this.tableData = response.data.items
        this.listLoading = false
      })
      console.log("===>>TODO 211: ", this.tableData)
    },
    deleteRow(index, rows) {
      rows.splice(index, 1);
    },
    onSubmit() {
      console.log('submit!', this.formInline.user, this.formInline.region);
      getRecommendList().then(response => {
        console.log("===>>TODO 112: ", response)

        this.tableData = response.data.items
        this.listLoading = false
      })
      console.log("===>>TODO 212: ", this.tableData)
    }
  }
}
</script>
