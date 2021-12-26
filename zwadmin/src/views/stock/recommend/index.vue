<template>
  <div>
  <div class="app-container">
    <el-form :inline="true" :model="formInline" class="demo-form-inline">
      <el-form-item label="跌幅">
        <el-input v-model.number="formInline.decrease" placeholder=0></el-input>
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
    <el-table-column fixed="right" label="操作" width="220">
      <template slot-scope="scope">
        <el-button
          @click.native.prevent="modifyRow(scope.$index, tableData)"
          type="text"
          size="small">
          修改
        </el-button>
        <el-button
          @click.native.prevent="lineChartDate(scope.$index, tableData)"
          type="text"
          size="small">
          日线
        </el-button>
        <el-button
          @click.native.prevent="lineChartContour(scope.$index, tableData)"
          type="text"
          size="small">
          周线
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
    <el-dialog title="修改推荐" :visible.sync="modifyDialogVisible">
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
  <div>
    <el-dialog title="K线查询" :visible.sync="lineChartForm.lineChartVisible">
      <el-image :src="lineChartForm.lineChartSrc"></el-image>
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
      modifyDialogVisible: false,
      formInline: {
        user: '',
        region: '',
        decrease: 0,
      },
      modifyForm: {
        name: '',
        region: '',
        secucode: '',
        priceDecrease: 0,
      },
      lineChartForm:{
        lineChartVisible: false,
        lineChartSrc: '',
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
        pDecrease:this.formInline.decrease,
      }
      getRecommendList(req).then(response => {
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
      this.modifyForm.secucode = data.secucode
      this.modifyForm.priceDecrease = data.pDecrease
      this.modifyForm.name = data.name
    },
    lineChartDate(index, rows) {//日线图
      console.log("==>>TODO 3141: ", rows[index])
      var data = rows[index]
      var secucode = data.secucode.split('.').join("").toLowerCase()
      console.log("==>>TODO 3142: ", secucode)
      this.lineChartForm.lineChartVisible = !this.lineChartForm.lineChartVisible
      this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/daily/n/'+secucode+'.gif'
    },
    lineChartContour(index, rows) {
      console.log("==>>TODO 3141: ", rows[index])
      var data = rows[index]
      var secucode = data.secucode.split('.').join("").toLowerCase()
      console.log("==>>TODO 3142: ", secucode)
      this.lineChartForm.lineChartVisible = !this.lineChartForm.lineChartVisible
      this.lineChartForm.lineChartSrc = 'http://image.sinajs.cn/newchart/weekly/n/'+secucode+'.gif'
    },
    cancelDialog(index, rows) {
      this.modifyDialogVisible = !this.modifyDialogVisible
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
        this.modifyDialogVisible = !this.modifyDialogVisible
      })
      this.fetchData()
    },
    onSubmit() {
      console.log('submit!', this.formInline.user, this.formInline.region);
      // var req = {
      //   limit: this.pageInfo.pageSize,
      //   offset: (this.pageInfo.pageNum-1)*this.pageInfo.pageSize,
      // }
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
      console.log("===>>TODO 214: ", val)
    }
  }
}
</script>
