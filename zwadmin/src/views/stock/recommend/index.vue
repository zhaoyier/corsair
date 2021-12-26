<template>
  <div>
  <div class="app-container">
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="跌幅">
        <el-input v-model="formInline.decrease" placeholder=0></el-input>
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
    <el-table :data="tableData" stripe style="width: 100%" max-height="800">
    <el-table-column fixed prop="secucode" label="代码" width="150"></el-table-column>
    <el-table-column prop="name" label="名称" width="120"></el-table-column>
    <el-table-column prop="rMIndex" label="推荐指数" width="120"></el-table-column>
    <el-table-column prop="state" label="状态" width="120"> </el-table-column>
    <el-table-column prop="pDecrease" label="最近跌幅" width="120"></el-table-column>
    <el-table-column prop="maxPrice" label="最高价" width="120"></el-table-column>
    <el-table-column prop="rMPrice" label="推荐价格" width="220"> </el-table-column>
    <el-table-column prop="presentPrice" label="当前价" width="120"> </el-table-column>
    <el-table-column prop="gDDecrease" label="股东人数" width="120"> </el-table-column>
    <el-table-column prop="updateDate" label="更新时间" width="120"> </el-table-column>
    <el-table-column fixed="right" label="操作" width="120">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="modifyRow(scope.$index, tableData)"
          type="text"
          size="small">
          修改
        </el-button>
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
  <div>
    <el-dialog title="修改推荐" :visible.sync="dialogFormVisible">
      <!-- <el-form :model="modifyForm">
        <el-form-item label="跌幅修正" :label-width="formLabelWidth">
          <el-input v-model="modifyForm.decrease" autocomplete="off"></el-input>
          <el-input v-model.number="modifyForm.decrease"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelDialog">取 消</el-button>
        <el-button type="primary" @click="confirmDialog">确 定</el-button>
      </div> -->

      <el-form :model="modifyForm" ref="modifyForm" label-width="100px" class="demo-ruleForm">
      <el-form-item
        label="跌幅修正"
        prop="decrease"
        :rules="[
          { type: 'number', message: '跌幅必须为数字值'}
        ]"
      >
        <el-input type="age" v-model.number="modifyForm.decrease" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="cancelDialog('modifyForm')">取消</el-button>
        <el-button type="primary" @click="confirmDialog('modifyForm')">提交</el-button>
      </el-form-item>
    </el-form>
    </el-dialog>
  </div>
  </div>
</template>

<script>
import { getRecommendList,updateRecommend } from '@/api/stock'

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
      listLoading: true,
      dialogFormVisible: false,
      formInline: {
        user: '',
        region: '',
      },
      modifyForm: {
        name: '',
        region: '',
        secucode: '',
        priceDecrease: 0,
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
      }
      getRecommendList(req).then(response => {
        console.log("===>>TODO 111: ", response)

        // this.tableData = response.rows
        this.tableData = response.data.items
        this.pageInfo.total = response.data.total
        this.listLoading = false
      })
      console.log("===>>TODO 211: ", this.tableData)
    },
    modifyRow(index, rows) {
     
      // rows.splice(index, 1);
      var data = rows[index]
       console.log("==>>TODO 301: ", "301", data)
      this.dialogFormVisible = !this.dialogFormVisible
      this.modifyForm.secucode = data.secucode
      this.modifyForm.priceDecrease = data.pDecrease
      this.modifyForm.name = data.name
    },
    cancelDialog(index, rows) {
      console.log("==>>TODO 302: ", "302")
      // rows.splice(index, 1);
      this.dialogFormVisible = !this.dialogFormVisible
    },
    confirmDialog(formName) {
      console.log("==>>TODO 3031: ", this.modifyForm)
      console.log("==>>TODO 3032: ", this.modifyForm.decrease)
      var form = this.modifyForm
      console.log("==>>TODO 3035: ", this.$refs[formName], form)

      var req = {
        name : this.modifyForm.name,
        secucode: this.modifyForm.secucode, 
        priceDecrease: this.modifyForm.decrease
      }

      updateRecommend(req).then(response=>{
        console.log("==>>TODO 3036: ", "ok")
        this.dialogFormVisible = !this.dialogFormVisible
      })
      this.fetchData()
    },
    onSubmit() {
      console.log('submit!', this.formInline.user, this.formInline.region);
      getRecommendList().then(response => {
        console.log("===>>TODO 112: ", response)

        this.tableData = response.data.items
        this.listLoading = false
      })
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

       getRecommendList(req).then(response => {
        console.log("===>>TODO 112: ", response)

        this.tableData = response.data.items
        this.listLoading = false
      })
      
    },
    handleSizeChange(val) {
      console.log("===>>TODO 214: ", val)
    }
  }
}
</script>
