<template>
  <div class="dashboard-container">
    <!-- <span>系统配置</span>
    <el-divider></el-divider> -->
    <div class="config-padding">
      <el-button style="font-size: 30px;" type="warning" icon="el-icon-setting" plain @click="onConfig">参数配置</el-button>
    </div>
    <el-drawer
      title="系统参数配置"
      :visible.sync="settingForm.configDrawer"
      :with-header="true">
      <el-divider></el-divider>
      <div style="margin: 20px;">
        <el-form :model="settingForm" ref="settingForm" label-width="100px" class="demo-ruleForm">
          <el-form-item
            label="下跌幅度"
            prop="decreaseTag"
            :rules="[
              { required: true, message: '下跌幅度不能为空'},
              { type: 'number', message: '下跌幅度必须为数字值'}
            ]"
          >
            <el-input type="decreaseTag" v-model.number="settingForm.decreaseTag" autocomplete="off">
              <i slot="suffix" style="font-style:normal;margin-right: 10px;">%</i>
            </el-input>
          </el-form-item>
          <el-form-item
            label="下跌周期"
            prop="decreasePeriod"
            :rules="[
              { required: true, message: '下跌周期不能为空'},
              { type: 'number', message: '下跌周期必须为数字值'}
            ]"
          >
            <el-input type="decreasePeriod" v-model.number="settingForm.decreasePeriod" autocomplete="off">
              <i slot="suffix" style="font-style:normal;margin-right: 10px;">天</i>
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('settingForm')">提交</el-button>
            <el-button @click="resetForm('settingForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-drawer>
  </div>
</template>

<script>
import { updateCNConfig } from '@/api/stock'
  export default {
    data() {
      return {
        numberValidateForm: {
          age: ''
        },
        settingForm: {
          age: '',
          configDrawer: false,
          decreaseTag: 38,
          decreasePeriod: 30,
        }
      };
    },
    methods: {
      submitForm(formName) {
        console.log("==>>TODO 110: ", this.settingForm)
        this.$refs[formName].validate((valid) => {
          if (!valid) {
            return
          }

          var req = {
            decreaseTag: this.settingForm.decreaseTag,
            decreasePeriod: this.settingForm.decreasePeriod,
          }

          updateCNConfig(req).then(response => {
            this.$message({
              message: '配置成功',
              type: 'success'
            });
            this.resetForm(formName)
          })
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      onConfig() {
        this.settingForm.configDrawer = true
      }
    }
  }


// <script>
  

// export default {
//   data() {
//     return {
//       list: null,
//       timer:null,
//       listLoading: true,
//       settingForm: {
//         name: '',
//         pass: '',
//         age: 10,
//         configDrawer: false,
//         labelPosition: 'right',
//         decreaseTag: 10,
//         decreasePeriod: 30,

//       },
//     }
//   },
//   mounted(){
//     this.fetchData();
//     this.timer = setInterval(() => {
//         setTimeout(this.fetchData, 0)
//     }, 1000*30)
//   },
//   // created() {
//   //   this.fetchData()
//   // },
//   methods: {
//     fetchData() {
//       this.listLoading = true
//     },
//     onConfig() {
//       console.log("==>>TODO 112: ", this.list)
//       this.settingForm.configDrawer = true
//     },
//     submitConfigForm(formName) {

//     },
//     resetConfigForm(formName) {
//       console.log("==>>TODO 113: ", formName, this.$refs)
//       // this.$refs[formName].resetFields();
//     }
//   }
// }

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
.config {
  &-padding {
    padding:5px 5px 10px 20px;
  }
  &-icon {
    padding:5px 5px 10px 30px;
  }
  &-size {
    height: 30px;
    width: 30px;
  }
}

</style>
